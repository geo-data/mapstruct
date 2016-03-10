package style

import (
	"fmt"
	"github.com/geo-data/mapfile/encoding"
	"github.com/geo-data/mapfile/mapobj/color"
	"github.com/geo-data/mapfile/tokens"
)

type Style struct {
	Color         *color.Color  `json:",omitempty"`
	OutlineColor  *color.Color  `json:",omitempty"`
	Symbol        tokens.String `json:",omitempty"`
	Size          tokens.String `json:",omitempty"`
	Width         tokens.String `json:",omitempty"`
	GeomTransform tokens.String `json:",omitempty"`
}

func New(tokens *tokens.Tokens) (s *Style, err error) {
	token := tokens.Value()
	if token != "STYLE" {
		err = fmt.Errorf("expected token STYLE, got: %s", token)
		return
	}
	tokens.Next()

	s = new(Style)
	for tokens != nil {
		token := tokens.Value()
		switch token {
		case "COLOR":
			if s.Color, err = color.New(tokens); err != nil {
				return
			}
		case "OUTLINECOLOR":
			if s.OutlineColor, err = color.New(tokens); err != nil {
				return
			}
		case "SYMBOL":
			s.Symbol = tokens.Next().Value()
		case "SIZE":
			s.Size = tokens.Next().Value()
		case "WIDTH":
			s.Size = tokens.Next().Value()
		case "GEOMTRANSFORM":
			s.GeomTransform = tokens.Next().Value()
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

func (s *Style) Encode(enc *encoding.MapfileEncoder) (err error) {
	if err = enc.TokenStart("STYLE"); err != nil {
		return
	}

	if s.Color != nil {
		if err = enc.TokenValue("COLOR", s.Color); err != nil {
			return
		}
	}
	if s.OutlineColor != nil {
		if err = enc.TokenValue("OUTLINECOLOR", s.OutlineColor); err != nil {
			return
		}
	}
	if err = enc.TokenString("SYMBOL", s.Symbol); err != nil {
		return
	}
	if err = enc.TokenValue("SIZE", s.Size); err != nil {
		return
	}
	if err = enc.TokenValue("WIDTH", s.Width); err != nil {
		return
	}
	if err = enc.TokenValue("GEOMTRANSFORM", s.GeomTransform); err != nil {
		return
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
