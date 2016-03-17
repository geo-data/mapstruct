package decode

import (
	"fmt"
	"github.com/geo-data/mapfile/types"
)

func (t *Decoder) Legend() (legend *types.Legend, err error) {
	token := t.Value()
	if token != "LEGEND" {
		err = fmt.Errorf("expected token LEGEND, got: %s", token)
		return
	}
	t.Next()

	l := new(types.Legend)
Loop:
	for t != nil {
		token := t.Value()
		switch token {
		case "IMAGECOLOR":
			if l.ImageColor, err = t.Color(); err != nil {
				return
			}
		case "END":
			break Loop
		case "":
			if t.AtEnd() {
				err = EndOfTokens
				return
			}
			fallthrough
		default:
			err = fmt.Errorf("unhandled mapfile token: %s", token)
			return
		}

		t = t.Next()
	}

	legend = l
	return
}
