package decode

import (
	"github.com/geo-data/mapstruct/mapfile/decode/scanner"
	"github.com/geo-data/mapstruct/types"
)

func (t *Decoder) Points() (points types.Points, err error) {
	var token *scanner.Token
	if token, err = t.ExpectedToken(scanner.POINTS); err != nil {
		return
	}
	t.Next()

	var ps types.Points
	for t != nil {
		if token, err = t.Token(); err != nil {
			return
		}

		if token.Type == scanner.END {
			break
		}

		var point *types.Point
		if point, err = t.Point(); err != nil {
			return
		}
		ps = append(ps, point)

		t = t.Next()
	}

	points = ps
	return
}
