package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"time"

	wfmplatefficiency "github.com/tylerolson/wfm-plat-efficiency"
)

//go:embed frontend/dist/*
var dist embed.FS

type Server struct {
	scraper *wfmplatefficiency.Scraper
}

func updateAllVendors(scraper *wfmplatefficiency.Scraper) {
	fmt.Println("Starting all vendors at: ", time.Now())
	for _, vendor := range scraper.GetVendors() {
		resultChan, err := scraper.UpdateVendorStats(vendor.Name)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Starting %v\n", vendor.Name)
		for value := range resultChan { // Loop until the channel is closed
			if value.Err != nil {
				fmt.Printf("Failed to fetch %s: %v\n", value.ItemName, value.Err)
			} else {
				fmt.Printf("Fetched %v\n", value.ItemName)
			}
		}
	}
	fmt.Println("Done at", time.Now())
}

func (s *Server) getVendors() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(s.scraper.GetVendors())
	}
}

func (s *Server) getVendor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		vendor, err := s.scraper.GetVendor(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(vendor)
	}
}

func main() {
	scraper := wfmplatefficiency.NewScraper()
	scraper.LoadVendors()

	ticker := time.NewTicker(5 * time.Hour)
	defer ticker.Stop()

	go func() {
		updateAllVendors(scraper)
		for range ticker.C {
			updateAllVendors(scraper)
		}
	}()

	// needs a mutex blah blah blah
	server := &Server{
		scraper,
	}

	http.HandleFunc("GET /api/vendors", server.getVendors())
	http.HandleFunc("GET /api/vendors/{name}", server.getVendor())

	distFS, err := fs.Sub(dist, "frontend/dist")
	if err != nil {
		log.Fatal(err)
	}

	http.Handle("/", http.FileServer(http.FS(distFS)))

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	fmt.Println("Listening on port ", PORT)
	http.ListenAndServe(":"+PORT, nil)
}
