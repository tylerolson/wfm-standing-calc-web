package main

import (
	"context"

	wfmplatefficiency "github.com/tylerolson/wfm-plat-efficiency"
)

type BasicVendor struct {
	Slug           string                 `json:"slug"`
	Name           string                 `json:"name"`
	MostProfitable wfmplatefficiency.Item `json:"mostProfitable"`
	MostVolume     wfmplatefficiency.Item `json:"mostVolume"`
	MostEfficient  wfmplatefficiency.Item `json:"mostEfficient"`
}

type BasicVendorsResponse struct {
	Body struct {
		UpdatedAt int64         `json:"updatedAt"`
		Updating  bool          `json:"updating"`
		Vendors   []BasicVendor `json:"vendors"`
	}
}

type GetVendorRequest struct {
	Slug string `path:"slug"`
}

type VendorsResponse struct {
	Body struct {
		UpdatedAt int64                    `json:"updatedAt"`
		Updating  bool                     `json:"updating"`
		Vendor    wfmplatefficiency.Vendor `json:"vendor"`
	}
}

func (s *Server) getVendorsOverview(_ context.Context, _ *struct{}) (*BasicVendorsResponse, error) {
	basicVendors := make([]BasicVendor, 0)

	for _, vendor := range s.calculator.GetVendors() {
		basicVendor := BasicVendor{
			Slug: vendor.Slug,
			Name: vendor.Name,
		}

		if mostProfit := vendor.MostProfit(); mostProfit != nil {
			basicVendor.MostProfitable = *mostProfit
		}

		if mostVolume := vendor.MostVolume(); mostVolume != nil {
			basicVendor.MostVolume = *mostVolume
		}

		if mostEfficient := vendor.MostEfficient(); mostEfficient != nil {
			basicVendor.MostEfficient = *mostEfficient
		} else {
			basicVendor.MostEfficient = basicVendor.MostProfitable
		}

		basicVendors = append(basicVendors, basicVendor)
	}

	vendorResponse := &BasicVendorsResponse{}
	vendorResponse.Body.UpdatedAt = s.getUpdatedAt().Unix()
	vendorResponse.Body.Updating = s.isUpdating()
	vendorResponse.Body.Vendors = basicVendors

	return vendorResponse, nil
}

func (s *Server) getVendor(_ context.Context, input *GetVendorRequest) (*VendorsResponse, error) {
	vendor, err := s.calculator.GetVendor(input.Slug)
	if err != nil {
		return nil, err
	}

	vendorResponse := &VendorsResponse{}
	vendorResponse.Body.UpdatedAt = s.getUpdatedAt().Unix()
	vendorResponse.Body.Updating = s.isUpdating()
	vendorResponse.Body.Vendor = *vendor

	return vendorResponse, nil
}
