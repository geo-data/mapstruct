package mapobj

import (
	"github.com/geo-data/mapfile/types"
	"github.com/geo-data/mapfile/types/color"
	"github.com/geo-data/mapfile/types/extent"
	"github.com/geo-data/mapfile/types/layer"
	"github.com/geo-data/mapfile/types/legend"
	"github.com/geo-data/mapfile/types/projection"
	"github.com/geo-data/mapfile/types/scalebar"
	"github.com/geo-data/mapfile/types/size"
	"github.com/geo-data/mapfile/types/web"
)

type Map struct {
	Name       types.String          `json:",omitempty"`
	Extent     *extent.Extent        `json:",omitempty"`
	ImageType  types.String          `json:",omitempty"`
	ImageColor *color.Color          `json:",omitempty"`
	Status     types.Keyword         `json:",omitempty"`
	Size       *size.Size            `json:",omitempty"`
	Fontset    types.String          `json:",omitempty"`
	Symbolset  types.String          `json:",omitempty"`
	Legend     *legend.Legend        `json:",omitempty"`
	Scalebar   *scalebar.Scalebar    `json:",omitempty"`
	Web        *web.Web              `json:",omitempty"`
	Projection projection.Projection `json:",omitempty"`
	Layers     []*layer.Layer        `json:",omitempty"`
}
