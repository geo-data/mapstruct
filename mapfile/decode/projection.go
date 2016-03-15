package decode

import (
	"fmt"
	"github.com/geo-data/mapfile/types"
)

func (t *Decoder) Projection() (p types.Projection, err error) {
	token := t.Value()
	if token != "PROJECTION" {
		err = fmt.Errorf("expected token PROJECTION, got: %s", token)
		return
	}
	t.Next()

	for t != nil {
		token := t.Value()
		switch token {
		case "END":
			return
		default:
			var s types.String
			if s, err = t.String(); err != nil {
				return
			}
			p = append(p, string(s))
		}

		t = t.Next()
	}

	return
}
