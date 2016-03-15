package tokens

import (
	"fmt"
	"github.com/geo-data/mapfile/types"
	"github.com/geo-data/mapfile/types/projection"
)

func (t *Tokens) Projection() (p projection.Projection, err error) {
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
