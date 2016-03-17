package decode

import (
	"fmt"
	"github.com/geo-data/mapfile/types"
)

func (t *Decoder) Points() (points types.Points, err error) {
	token := t.Value()
	if token != "POINTS" {
		err = fmt.Errorf("expected token POINTS, got: %s", token)
		return
	}
	t.Next()

	var ps types.Points
	for t != nil {
		if t.Value() == "END" {
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
