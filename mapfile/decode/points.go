package decode

import (
	"fmt"
	"github.com/geo-data/mapfile/types/point"
)

func (t *Decoder) Point() (p *point.Point, err error) {
	p = new(point.Point)
	if p.X, err = t.Double(); err != nil {
		return
	}

	if p.Y, err = t.Next().Double(); err != nil {
		return
	}

	return
}

func (t *Decoder) Points() (ps point.Points, err error) {
	token := t.Value()
	if token != "POINTS" {
		err = fmt.Errorf("expected token POINTS, got: %s", token)
		return
	}
	t.Next()

	for t != nil {
		if t.Value() == "END" {
			break
		}

		var point *point.Point
		if point, err = t.Point(); err != nil {
			return
		}
		ps = append(ps, point)

		t = t.Next()
	}
	return
}
