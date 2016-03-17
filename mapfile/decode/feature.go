package decode

import (
	"fmt"
	"github.com/geo-data/mapfile/types"
)

func (t *Decoder) Feature() (feature *types.Feature, err error) {
	token := t.Value()
	if token != "FEATURE" {
		err = fmt.Errorf("expected token FEATURE, got: %s", token)
		return
	}
	t.Next()

	f := new(types.Feature)
Loop:
	for t != nil {
		token := t.Value()
		switch token {
		case "WKT":
			if f.Wkt, err = t.Next().String(); err != nil {
				return
			}
		case "ITEMS":
			if f.Items, err = t.Next().String(); err != nil {
				return
			}
		case "TEXT":
			if f.Text, err = t.Next().String(); err != nil {
				return
			}
		case "POINTS":
			if f.Points, err = t.Points(); err != nil {
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

	feature = f
	return
}
