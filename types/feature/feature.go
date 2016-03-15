package feature

import (
	"github.com/geo-data/mapfile/types"
	"github.com/geo-data/mapfile/types/point"
)

type Feature struct {
	Wkt    types.String `json:",omitempty"`
	Items  types.String `json:",omitempty"`
	Text   types.String `json:",omitempty"`
	Points point.Points `json:",omitempty"`
}
