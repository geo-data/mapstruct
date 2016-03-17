package decode

import (
	"fmt"
	"github.com/geo-data/mapfile/types"
)

func (t *Decoder) Projection() (projection types.Projection, err error) {
	token := t.Value()
	if token != "PROJECTION" {
		err = fmt.Errorf("expected token PROJECTION, got: %s", token)
		return
	}
	t.Next()

	var p types.Projection
Loop:
	for t != nil {
		token := t.Value()
		switch token {
		case "END":
			break Loop
		case "":
			if t.AtEnd() {
				err = EndOfTokens
				return
			}
			fallthrough
		default:
			var s types.String
			if s, err = t.String(); err != nil {
				return
			}
			p = append(p, string(s))
		}

		t = t.Next()
	}

	projection = p
	return
}
