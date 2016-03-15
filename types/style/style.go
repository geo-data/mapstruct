package style

import (
	"github.com/geo-data/mapfile/types"
	"github.com/geo-data/mapfile/types/color"
)

type Style struct {
	Color         *color.Color `json:",omitempty"`
	OutlineColor  *color.Color `json:",omitempty"`
	Symbol        types.Union  `json:",omitempty"`
	Size          types.Union  `json:",omitempty"`
	Width         types.Union  `json:",omitempty"`
	GeomTransform types.String `json:",omitempty"`
}
