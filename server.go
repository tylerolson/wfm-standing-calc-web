package main

import (
	"context"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humago"
	wfmplatefficiency "github.com/tylerolson/wfm-plat-efficiency"
)

type Server struct {
	serveMux   *http.ServeMux
	humaAPI    huma.API
	mu         sync.RWMutex
	calculator *wfmplatefficiency.Calculator
	updatedAt  time.Time
	updating   bool
}

func NewServer(calculator *wfmplatefficiency.Calculator) *Server {
	server := &Server{
		serveMux:   http.NewServeMux(),
		calculator: calculator,
		updatedAt:  time.Time{},
		updating:   true,
	}

	// API routes
	server.humaAPI = humago.New(server.serveMux, huma.DefaultConfig("WFM Calculator", "v0.3.4"))
	huma.Get(server.humaAPI, "/api/vendors", server.getVendorsOverview)
	huma.Get(server.humaAPI, "/api/vendors/{slog}", server.getVendor)

	// SPA route
	distFS, err := fs.Sub(dist, "frontend/build")
	if err != nil {
		log.Fatalf("Failed to load frontend: %v", err)
	}
	server.serveMux.HandleFunc("/", spaHandler(distFS))

	return server
}

func (s *Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.serveMux.ServeHTTP(w, r)
}

func (s *Server) setUpdating(val bool) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.updating = val
}

func (s *Server) isUpdating() bool {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.updating
}

func (s *Server) setUpdatedAt(t time.Time) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.updatedAt = t
}

func (s *Server) getUpdatedAt() time.Time {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.updatedAt
}

func (s *Server) updateAllVendors(ctx context.Context) error {
	s.setUpdating(true)
	defer s.setUpdating(false)

	vendors := s.calculator.GetVendors()

	startTime := time.Now()
	log.Printf("Starting vendor updates at")

	for i, vendor := range s.calculator.GetVendors() {
		select {
		case <-ctx.Done():
			return ctx.Err()
		default:
		}

		j := 1
		resultChan, err := s.calculator.UpdateVendorStats(vendor.Slug)
		if err != nil {
			return fmt.Errorf("error updating vendor %s: %w", vendor.Slug, err)
		}

		log.Printf("[%d/%d] Starting %s", i+1, len(vendors), vendor.Name)
		for value := range resultChan { // Loop until the channel is closed
			if value.Err != nil {
				log.Printf("  └─ Failed to fetch %s: %v", value.ItemName, value.Err)
			} else {
				log.Printf("[%d/%d] Updating %s", j, len(vendor.Items), value.ItemName)
			}
			j++
		}
	}

	s.setUpdatedAt(time.Now())
	duration := time.Since(startTime)
	log.Printf("Completed vendor updates (took %v)", duration.Round(time.Second))
	return nil
}

func (s *Server) startVendorUpdates(ctx context.Context, interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	if err := s.updateAllVendors(ctx); err != nil {
		log.Printf("Update error: %v", err)
	}

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := s.updateAllVendors(ctx); err != nil {
				log.Printf("Update error: %v", err)
			}
		}
	}
}
