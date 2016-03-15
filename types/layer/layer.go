package layer

import (
	"github.com/geo-data/mapfile/mapfile/encode"
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

func (l *Layer) Encode(enc *encode.MapfileEncoder) (err error) {
	if err = enc.TokenStart("LAYER"); err != nil {
		return
	}

	if err = enc.TokenStringer("NAME", l.Name); err != nil {
		return
	}
	if l.Extent != nil {
		if err = l.Extent.Encode(enc); err != nil {
			return
		}
	}
	if err = enc.TokenStringer("TYPE", l.Type); err != nil {
		return
	}
	if err = enc.TokenUnion("DEBUG", l.Debug); err != nil {
		return
	}
	if l.Projection != nil {
		if err = l.Projection.Encode(enc); err != nil {
			return
		}
	}
	if err = enc.TokenStringer("DATA", l.Data); err != nil {
		return
	}
	if err = enc.TokenStringer("PROCESSING", l.Processing); err != nil {
		return
	}
	if err = enc.TokenStringer("STATUS", l.Status); err != nil {
		return
	}
	if l.Metadata != nil {
		if err = l.Metadata.Encode(enc); err != nil {
			return
		}
	}
	if err = enc.TokenString("CLASSITEM", l.ClassItem.QuotedString()); err != nil {
		return
	}
	if err = enc.TokenString("LABELITEM", l.LabelItem.QuotedString()); err != nil {
		return
	}

	for _, class := range l.Classes {
		if err = class.Encode(enc); err != nil {
			return
		}
	}

	for _, feature := range l.Features {
		if err = feature.Encode(enc); err != nil {
			return
		}
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
