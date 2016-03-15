package decode

import (
	"fmt"
	"github.com/geo-data/mapfile/types"
)

func (t *Decoder) Style() (s *types.Style, err error) {
	token := t.Value()
	if token != "STYLE" {
		err = fmt.Errorf("expected token STYLE, got: %s", token)
		return
	}
	t.Next()

	s = new(types.Style)
	for t != nil {
		token := t.Value()
		switch token {
		case "COLOR":
			if s.Color, err = t.Color(); err != nil {
				return
			}
		case "OUTLINECOLOR":
			if s.OutlineColor, err = t.Color(); err != nil {
				return
			}
		case "SYMBOL":
			if s.Symbol, err = t.Next().Decode(Integer | String | Attribute); err != nil {
				return
			}
		case "SIZE":
			if s.Size, err = t.Next().Decode(Double | Attribute); err != nil {
				return
			}
		case "WIDTH":
			if s.Width, err = t.Next().Decode(String | Double); err != nil {
				err = fmt.Errorf("could not decode WIDTH: %s", err)
				return
			}
		case "GEOMTRANSFORM":
			if s.GeomTransform, err = t.Next().String(); err != nil {
				return
			}
		case "END":
			return
		default:
			err = fmt.Errorf("unhandled mapfile token: %s", token)
			return
		}

		t = t.Next()
	}

	return
}
