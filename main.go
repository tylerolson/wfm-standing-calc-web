package main

import (
	"context"
	"embed"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	wfmplatefficiency "github.com/tylerolson/wfm-plat-efficiency"
)

//go:embed frontend/build/*
var dist embed.FS

func main() {
	calculator := wfmplatefficiency.NewScraper()
	if err := calculator.LoadVendors(); err != nil {
		log.Fatalf("Failed to load vendors: %v", err)
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	server := NewServer(calculator)
	s := http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: server,
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)

	go server.startVendorUpdates(ctx, 4*time.Hour)

	go func() {
		log.Printf("Starting server on port %s", port)
		if err := s.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Fatal("HTTP server error", "error", err)
		}

		log.Println("HTTP server shutdown")
	}()

	sig := <-c
	log.Printf("Received signal: %v", sig)
	log.Println("Initiating graceful shutdown...")

	// cancel vendor updates on shutdown
	log.Println("Stopping vendor updates...")
	cancel()

	log.Println("Shutting down server")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	log.Println("Shutting down HTTP server...")
	if err := s.Shutdown(shutdownCtx); err != nil {
		log.Fatalf("Server shutdown error %v", err)
	}
	log.Println("Server exited gracefully")
}
