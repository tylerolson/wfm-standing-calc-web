package main

import (
	"context"
	"io/fs"
	"net/http"
	"strings"
	"time"

	wfmplatefficiency "github.com/tylerolson/wfm-plat-efficiency"
)

type BasicVendor struct {
	Slug           string                 `json:"slug"`
	Name           string                 `json:"name"`
	MostProfitable wfmplatefficiency.Item `json:"mostProfitable"`
	MostVolume     wfmplatefficiency.Item `json:"mostVolume"`
	MostEfficient  wfmplatefficiency.Item `json:"mostEfficient"`
}

// a basic overview of the vendors
type BasicVendorsResponse struct {
	Body struct {
		UpdatedAt time.Time     `json:"updatedAt"`
		Updating  bool          `json:"updating"`
		Vendors   []BasicVendor `json:"vendors"`
	}
}

type GetVendorRequest struct {
	Slog string `path:"slog"`
}

type VendorsResponse struct {
	Body struct {
		UpdatedAt time.Time                `json:"updatedAt"`
		Updating  bool                     `json:"updating"`
		Vendor    wfmplatefficiency.Vendor `json:"vendor"`
	}
}

func (s *Server) getVendorsOverview(_ context.Context, _ *struct{}) (*BasicVendorsResponse, error) {
	basicVendors := make([]BasicVendor, 0)

	// TODO FIX
	// most efficient is nil here when every vendors isnt fetched already
	for _, vendor := range s.calculator.GetVendors() {
		// return most profitable if there is a nil value, will fix later
		mostEfficient := vendor.MostEfficient()
		if mostEfficient == nil {
			mostEfficient = vendor.MostProfit()
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
	vendorResponse := &BasicVendorsResponse{}
	vendorResponse.Body.UpdatedAt = s.updatedAt
	vendorResponse.Body.Updating = s.updating
	vendorResponse.Body.Vendors = basicVendors

	return vendorResponse, nil
}

func (s *Server) getVendor(_ context.Context, input *GetVendorRequest) (*VendorsResponse, error) {
	vendor, err := s.calculator.GetVendor(input.Slog)
	if err != nil {
		return nil, err
	}

	vendorResponse := &VendorsResponse{}
	vendorResponse.Body.UpdatedAt = s.updatedAt
	vendorResponse.Body.Updating = s.updating
	vendorResponse.Body.Vendor = *vendor

	return vendorResponse, nil
}

func spaHandler(distFS fs.FS) http.HandlerFunc {
	hfs := http.FileServer(http.FS(distFS))

	return func(w http.ResponseWriter, r *http.Request) {
		path := strings.TrimPrefix(r.URL.Path, "/")

		if path != "" {
			if _, err := fs.Stat(distFS, path); err != nil {
				//check if file exists, if it doesn't serve index for SPA routing
				http.ServeFileFS(w, r, distFS, "index.html")
				return
			}
		}

		hfs.ServeHTTP(w, r)
	}
}
