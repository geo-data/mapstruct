package style

import (
	"fmt"
	"github.com/geo-data/mapfile/encoding"
	"github.com/geo-data/mapfile/tokens"
	"github.com/geo-data/mapfile/types"
	"github.com/geo-data/mapfile/types/color"
)

type Style struct {
	Color         *color.Color `json:",omitempty"`
	OutlineColor  *color.Color `json:",omitempty"`
	Symbol        fmt.Stringer `json:",omitempty"`
	Size          fmt.Stringer `json:",omitempty"`
	Width         fmt.Stringer `json:",omitempty"`
	GeomTransform types.String `json:",omitempty"`
}

func New(toks *tokens.Tokens) (s *Style, err error) {
	token := toks.Value()
	if token != "STYLE" {
		err = fmt.Errorf("expected token STYLE, got: %s", token)
		return
	}
	toks.Next()

	s = new(Style)
	for toks != nil {
		token := toks.Value()
		switch token {
		case "COLOR":
			if s.Color, err = color.New(toks); err != nil {
				return
			}
		case "OUTLINECOLOR":
			if s.OutlineColor, err = color.New(toks); err != nil {
				return
			}
		case "SYMBOL":
			if s.Symbol, err = toks.Next().Decode(tokens.INTEGER | tokens.STRING | tokens.ATTRIBUTE); err != nil {
				return
			}
		case "SIZE":
			if s.Size, err = toks.Next().Decode(tokens.FLOAT64 | tokens.ATTRIBUTE); err != nil {
				return
			}
		case "WIDTH":
			if s.Width, err = toks.Next().Decode(tokens.STRING | tokens.FLOAT64); err != nil {
				err = fmt.Errorf("could not decode WIDTH: %s", err)
				return
			}
		case "GEOMTRANSFORM":
			if s.GeomTransform, err = toks.Next().String(); err != nil {
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

func (s *Style) Encode(enc *encoding.MapfileEncoder) (err error) {
	if err = enc.TokenStart("STYLE"); err != nil {
		return
	}

	if s.Color != nil {
		if err = enc.TokenStringer("COLOR", s.Color); err != nil {
			return
		}
	}
	if s.OutlineColor != nil {
		if err = enc.TokenStringer("OUTLINECOLOR", s.OutlineColor); err != nil {
			return
		}
	}
	if err = enc.TokenStringer("SYMBOL", s.Symbol); err != nil {
		return
	}
	if err = enc.TokenStringer("SIZE", s.Size); err != nil {
		return
	}
	if err = enc.TokenStringer("WIDTH", s.Width); err != nil {
		return
	}
	if err = enc.TokenStringer("GEOMTRANSFORM", s.GeomTransform); err != nil {
		return
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
