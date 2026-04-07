package main

import (
	"github.com/danielgtaylor/huma/v2"
	wfmplatefficiency "github.com/tylerolson/wfm-plat-efficiency"
)

// VendorSlug wraps the library VendorSlug to provide Huma OpenAPI schema
type VendorSlug wfmplatefficiency.VendorSlug

func (VendorSlug) Schema(r huma.Registry) *huma.Schema {
	return &huma.Schema{
		Type: huma.TypeString,
		Enum: []any{
			string(wfmplatefficiency.VendorSlugArbitersOfHexis),
			string(wfmplatefficiency.VendorSlugCephalonSuda),
			string(wfmplatefficiency.VendorSlugNewLoka),
			string(wfmplatefficiency.VendorSlugRedVeil),
			string(wfmplatefficiency.VendorSlugSteelMeridian),
			string(wfmplatefficiency.VendorSlugThePerrinSequence),
		},
	}
}

// ItemType wraps the library ItemType to provide Huma OpenAPI schema
type ItemType wfmplatefficiency.ItemType

func (t ItemType) MarshalJSON() ([]byte, error) {
	return wfmplatefficiency.ItemType(t).MarshalJSON()
}

func (t *ItemType) UnmarshalJSON(data []byte) error {
	var libType wfmplatefficiency.ItemType
	if err := (&libType).UnmarshalJSON(data); err != nil {
		return err
	}

	*t = ItemType(libType)
	return nil
}

func (ItemType) Schema(r huma.Registry) *huma.Schema {
	return &huma.Schema{
		Type: huma.TypeString,
		Enum: []any{"Mod", "ArchPart", "Weapon"},
	}
}

// Item wraps the library Item with local types for proper schema generation
type Item struct {
	Slug     string   `json:"slug"`
	Name     string   `json:"name"`
	Type     ItemType `json:"type"`
	Standing int      `json:"standing"`
	Price    float64  `json:"price"`
	Volume   float64  `json:"volume"`
	Score    float64  `json:"score"`
}

func ItemFromLib(i wfmplatefficiency.Item) Item {
	return Item{
		Slug:     i.Slug,
		Name:     i.Name,
		Type:     ItemType(i.Type),
		Standing: i.StandingCost,
		Price:    i.Price,
		Volume:   i.Volume,
		Score:    i.Score,
	}
}

// Vendor wraps the library Vendor with local types for proper schema generation
type Vendor struct {
	Slug  VendorSlug `json:"slug"`
	Name  string     `json:"name"`
	Items []Item     `json:"items"`
}

func VendorFromLib(v wfmplatefficiency.Vendor) Vendor {
	items := make([]Item, 0, len(v.Items))
	for _, item := range v.Items {
		if item != nil {
			items = append(items, ItemFromLib(*item))
		}
	}
	return Vendor{
		Slug:  VendorSlug(v.Slug),
		Name:  v.Name,
		Items: items,
	}
}
