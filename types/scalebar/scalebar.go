package scalebar

import (
	"github.com/geo-data/mapfile/types"
	"github.com/geo-data/mapfile/types/color"
	"github.com/geo-data/mapfile/types/label"
	"github.com/geo-data/mapfile/types/size"
)

type Scalebar struct {
	Status          types.Keyword `json:",omitempty"`
	PostLabelCache  types.Keyword `json:",omitempty"`
	Style           types.Uint8
	Units           types.Keyword `json:",omitempty"`
	Size            *size.Size    `json:",omitempty"`
	Position        types.Keyword `json:",omitempty"`
	Transparent     types.Keyword `json:",omitempty"`
	Color           *color.Color  `json:",omitempty"`
	ImageColor      *color.Color  `json:",omitempty"`
	BackgroundColor *color.Color  `json:",omitempty"`
	Label           *label.Label  `json:",omitempty"`
}
