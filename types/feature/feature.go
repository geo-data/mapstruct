package feature

import (
	"fmt"
	"github.com/geo-data/mapfile/mapfile/encode"
	"github.com/geo-data/mapfile/tokens"
	"github.com/geo-data/mapfile/types"
	"github.com/geo-data/mapfile/types/point"
)

type Feature struct {
	Wkt    types.String  `json:",omitempty"`
	Items  types.String  `json:",omitempty"`
	Text   types.String  `json:",omitempty"`
	Points *point.Points `json:",omitempty"`
}

func New(toks *tokens.Tokens) (c *Feature, err error) {
	token := toks.Value()
	if token != "FEATURE" {
		err = fmt.Errorf("expected token FEATURE, got: %s", token)
		return
	}
	toks.Next()

	c = new(Feature)
	for toks != nil {
		token := toks.Value()
		switch token {
		case "WKT":
			if c.Wkt, err = toks.Next().String(); err != nil {
				return
			}
		case "ITEMS":
			if c.Items, err = toks.Next().String(); err != nil {
				return
			}
		case "TEXT":
			if c.Text, err = toks.Next().String(); err != nil {
				return
			}
		case "POINTS":
			if c.Points, err = point.NewPoints(toks); err != nil {
				return
			}
		case "END":
			return
		default:
			err = fmt.Errorf("unhandled mapfile token: %s", token)
			return
		}

		toks = toks.Next()
	}

	return
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
