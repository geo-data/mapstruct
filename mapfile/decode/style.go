package decode

import (
	"fmt"
	"github.com/geo-data/mapfile/mapfile/decode/scanner"
	"github.com/geo-data/mapfile/types"
)

func (t *Decoder) Style() (style *types.Style, err error) {
	var token *scanner.Token
	if token, err = t.ExpectedToken(scanner.STYLE); err != nil {
		return
	}
	t.Next()

	s := new(types.Style)
Loop:
	for t != nil {
		if token, err = t.Token(); err != nil {
			return
		}

		switch token.Type {
		case scanner.COLOR:
			if s.Color, err = t.Color(); err != nil {
				return
			}
		case scanner.OUTLINECOLOR:
			if s.OutlineColor, err = t.Color(); err != nil {
				return
			}
		case scanner.SYMBOL:
			if s.Symbol, err = t.Next().Decode(Integer | String | Attribute); err != nil {
				return
			}
		case scanner.SIZE:
			if s.Size, err = t.Next().Decode(Double | Attribute); err != nil {
				return
			}
		case scanner.WIDTH:
			if s.Width, err = t.Next().Decode(Double | Attribute); err != nil {
				return
			}
		case scanner.GEOMTRANSFORM:
			if s.GeomTransform, err = t.Next().String(); err != nil {
				return
			}
		case scanner.END:
			break Loop
		default:
			err = fmt.Errorf("unhandled mapfile token: %s", token)
			return
		}

		t = t.Next()
	}

	style = s
	return
}
