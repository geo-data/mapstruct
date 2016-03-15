package layer

import (
	"github.com/geo-data/mapfile/types"
	"github.com/geo-data/mapfile/types/class"
	"github.com/geo-data/mapfile/types/extent"
	"github.com/geo-data/mapfile/types/feature"
	"github.com/geo-data/mapfile/types/metadata"
	"github.com/geo-data/mapfile/types/projection"
)

type Layer struct {
	Name       types.String          `json:",omitempty"`
	Extent     *extent.Extent        `json:",omitempty"`
	Type       types.Keyword         `json:",omitempty"`
	Debug      types.Union           `json:",omitempty"`
	Projection projection.Projection `json:",omitempty"`
	Data       types.String          `json:",omitempty"`
	Processing types.String          `json:",omitempty"`
	Status     types.Keyword         `json:",omitempty"`
	Metadata   metadata.Metadata     `json:",omitempty"`
	ClassItem  types.Attribute       `json:",omitempty"`
	LabelItem  types.Attribute       `json:",omitempty"`
	Classes    []*class.Class        `json:",omitempty"`
	Features   []*feature.Feature    `json:",omitempty"`
}
