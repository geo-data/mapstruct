package feature

import (
	"fmt"
	"github.com/geo-data/mapfile/encoding"
	"github.com/geo-data/mapfile/mapobj/point"
	"github.com/geo-data/mapfile/tokens"
)

type Feature struct {
	Wkt    tokens.String `json:",omitempty"`
	Items  tokens.String `json:",omitempty"`
	Text   tokens.String `json:",omitempty"`
	Points *point.Points `json:",omitempty"`
}

func New(tokens *tokens.Tokens) (c *Feature, err error) {
	token := tokens.Value()
	if token != "FEATURE" {
		err = fmt.Errorf("expected token FEATURE, got: %s", token)
		return
	}
	tokens.Next()

	c = new(Feature)
	for tokens != nil {
		token := tokens.Value()
		switch token {
		case "WKT":
			c.Wkt = tokens.Next().Value()
		case "ITEMS":
			c.Items = tokens.Next().Value()
		case "TEXT":
			c.Text = tokens.Next().Value()
		case "POINTS":
			if c.Points, err = point.NewPoints(tokens); err != nil {
				return
			}
		case "END":
			return
		default:
			err = fmt.Errorf("unhandled mapfile token: %s", token)
			return
		}

		tokens = tokens.Next()
	}

	return
}

func (c *Feature) Encode(enc *encoding.MapfileEncoder) (err error) {
	if err = enc.TokenStart("FEATURE"); err != nil {
		return
	}

	if err = enc.TokenString("WKT", c.Wkt); err != nil {
		return
	}
	if err = enc.TokenString("ITEMS", c.Items); err != nil {
		return
	}
	if err = enc.TokenString("TEXT", c.Text); err != nil {
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
