package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/a-h/templ"
	wfmplatefficiency "github.com/tylerolson/wfm-plat-efficiency"
)

func updateAllVendors(scraper *wfmplatefficiency.Scraper) {
	fmt.Println("Starting all vendors at: ", time.Now())
	for name := range scraper.GetVendors() {
		resultChan, err := scraper.UpdateVendorStats(name)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Starting %v\n", name)
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

	index := index(scraper.GetVendors())
	http.Handle("/", templ.Handler(index))

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	fmt.Println("Listening on port ", PORT)
	http.ListenAndServe(":"+PORT, nil)
}
