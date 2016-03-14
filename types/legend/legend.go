package legend

import (
	"fmt"
	"github.com/geo-data/mapfile/encoding"
	"github.com/geo-data/mapfile/types/color"
	"github.com/geo-data/mapfile/tokens"
)

type Legend struct {
	ImageColor *color.Color `json:",omitempty"`
}

func New(tokens *tokens.Tokens) (l *Legend, err error) {
	token := tokens.Value()
	if token != "LEGEND" {
		err = fmt.Errorf("expected token LEGEND, got: %s", token)
		return
	}
	tokens.Next()

	l = new(Legend)
	for tokens != nil {
		token := tokens.Value()
		switch token {
		case "IMAGECOLOR":
			if l.ImageColor, err = color.New(tokens); err != nil {
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

func (l *Legend) Encode(enc *encoding.MapfileEncoder) (err error) {
	if err = enc.TokenStart("LEGEND"); err != nil {
		return
	}

	if l.ImageColor != nil {
		if err = enc.TokenStringer("IMAGECOLOR", l.ImageColor); err != nil {
			return
		}
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
