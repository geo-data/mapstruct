package feature

import (
	"github.com/geo-data/mapfile/mapfile/encode"
	"github.com/geo-data/mapfile/types"
	"github.com/geo-data/mapfile/types/point"
)

type Feature struct {
	Wkt    types.String `json:",omitempty"`
	Items  types.String `json:",omitempty"`
	Text   types.String `json:",omitempty"`
	Points point.Points `json:",omitempty"`
}

func (c *Feature) Encode(enc *encode.MapfileEncoder) (err error) {
	if err = enc.TokenStart("FEATURE"); err != nil {
		return
	}

	if err = enc.TokenStringer("WKT", c.Wkt); err != nil {
		return
	}
	if err = enc.TokenStringer("ITEMS", c.Items); err != nil {
		return
	}
	if err = enc.TokenStringer("TEXT", c.Text); err != nil {
		return
	}
	if c.Points != nil {
		if err = c.Points.Encode(enc); err != nil {
			return
		}
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
