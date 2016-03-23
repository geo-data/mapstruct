package decode

import (
	"github.com/geo-data/mapfile/mapfile/decode/scanner"
	"github.com/geo-data/mapfile/types"
)

func (t *Decoder) Projection() (projection types.Projection, err error) {
	var token *scanner.Token
	if token, err = t.ExpectedToken(scanner.PROJECTION); err != nil {
		return
	}
	t.Next()

	var p types.Projection
Loop:
	for t != nil {
		if token, err = t.Token(); err != nil {
			return
		}

		switch token.Type {
		case scanner.END:
			break Loop
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
