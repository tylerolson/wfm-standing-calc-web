package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	wfmplatefficiency "github.com/tylerolson/wfm-plat-efficiency"
)

//go:embed frontend/build/*
var dist embed.FS

type Server struct {
	calculator *wfmplatefficiency.Calculator
	updatedAt  time.Time
	updating   bool
}

type BasicVendor struct {
	Slug           string                 `json:"slug"`
	Name           string                 `json:"name"`
	MostProfitable wfmplatefficiency.Item `json:"mostProfitable"`
	MostVolume     wfmplatefficiency.Item `json:"mostVolume"`
	MostEfficient  wfmplatefficiency.Item `json:"mostEfficient"`
}

// a basic overview of the vendors
type BasicVendorsResponse struct {
	UpdatedAt time.Time     `json:"updatedAt"`
	Updating  bool          `json:"updating"`
	Vendors   []BasicVendor `json:"vendors"`
}

type VendorsResponse struct {
	UpdatedAt time.Time                `json:"updatedAt"`
	Updating  bool                     `json:"updating"`
	Vendor    wfmplatefficiency.Vendor `json:"vendor"`
}

func (s *Server) updateAllVendors() {
	s.updating = true
	fmt.Println("Starting all vendors at: ", time.Now())
	for _, vendor := range s.calculator.GetVendors() {
		resultChan, err := s.calculator.UpdateVendorStats(vendor.Slug)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Starting %v\n", vendor.Slug)
		for value := range resultChan { // Loop until the channel is closed
			if value.Err != nil {
				fmt.Printf("Failed to fetch %s: %v\n", value.ItemName, value.Err)
			} else {
				fmt.Printf("Fetched %v\n", value.ItemName)
			}
		}
	}
	s.updatedAt = time.Now()
	s.updating = false
	fmt.Println("Done at", time.Now())
}

func (s *Server) getVendorsOverview() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		basicVendors := make([]BasicVendor, 0)

		// TODO FIX
		// most efficient is nil here when every vendors isnt fetched already
		for _, vendor := range s.calculator.GetVendors() {
			// return most profitable if there is a nil value, will fix later
			mostEfficient := vendor.MostProfit()
			if vendor.MostEfficient() != nil {
				mostEfficient = vendor.MostEfficient()
			}
			basicVendors = append(basicVendors, BasicVendor{
				Slug:           vendor.Slug,
				Name:           vendor.Name,
				MostProfitable: *vendor.MostProfit(),
				MostVolume:     *vendor.MostVolume(),
				MostEfficient:  *mostEfficient,
			})
		}

		// maybe send date by unix?
		vendorResponse := BasicVendorsResponse{
			s.updatedAt,
			s.updating,
			basicVendors,
		}

		w.Header().Set("Content-Type", "application/json")
		err := json.NewEncoder(w).Encode(vendorResponse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func (s *Server) getVendor() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		name := r.PathValue("name")
		vendor, err := s.calculator.GetVendor(name)
		if err != nil {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		vendorResponse := VendorsResponse{
			UpdatedAt: s.updatedAt,
			Updating:  s.updating,
			Vendor:    *vendor,
		}

		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(vendorResponse)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
}

func main() {
	scraper := wfmplatefficiency.NewScraper()
	err := scraper.LoadVendors()
	if err != nil {
		log.Fatal(err)
		return
	}

	ticker := time.NewTicker(5 * time.Hour)
	defer ticker.Stop()

	server := &Server{
		scraper,
		time.Time{},
		true,
	}

	// needs a mutex blah blah blah
	go func() {
		server.updateAllVendors()
		for range ticker.C {
			server.updateAllVendors()
		}
	}()

	http.HandleFunc("GET /api/vendors", server.getVendorsOverview())
	http.HandleFunc("GET /api/vendors/{name}", server.getVendor())

	distFS, err := fs.Sub(dist, "frontend/build")
	if err != nil {
		log.Fatal(err)
	}

	//http.Handle("/", http.FileServer(http.FS(distFS)))

	hfs := http.FileServer(http.FS(distFS))
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path != "/" {
			r.URL.Path = strings.TrimPrefix(r.URL.Path, "/")

			//check if file exists, if it doesn't server index for SPA routing
			_, err := fs.Stat(distFS, r.URL.Path)
			if err != nil {
				http.ServeFileFS(w, r, distFS, "index.html")
				return
			}
		}
		hfs.ServeHTTP(w, r)
	})

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	fmt.Println("Listening on port ", PORT)
	err = http.ListenAndServe(":"+PORT, nil)
	if err != nil {
		log.Fatal(err)
		return
	}
}
