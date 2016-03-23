package decode

import (
	"fmt"
	"github.com/geo-data/mapfile/mapfile/decode/scanner"
	"github.com/geo-data/mapfile/types"
)

func (t *Decoder) Feature() (feature *types.Feature, err error) {
	var token *scanner.Token
	if token, err = t.ExpectedToken(scanner.FEATURE); err != nil {
		return
	}
	t.Next()

	f := new(types.Feature)
Loop:
	for t != nil {
		if token, err = t.Token(); err != nil {
			return
		}

		switch token.Type {
		case scanner.WKT:
			if f.Wkt, err = t.Next().String(); err != nil {
				return
			}
		case scanner.ITEMS:
			if f.Items, err = t.Next().String(); err != nil {
				return
			}
		case scanner.TEXT:
			if f.Text, err = t.Next().String(); err != nil {
				return
			}
		case scanner.POINTS:
			if f.Points, err = t.Points(); err != nil {
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

	feature = f
	return
}
