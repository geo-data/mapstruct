package decode

import (
	"fmt"
	"github.com/geo-data/mapfile/types"
)

func (t *Decoder) Point() (point *types.Point, err error) {
	p := new(types.Point)
	if p.X, err = t.Double(); err != nil {
		if t.AtEnd() {
			err = EndOfTokens
		} else {
			err = fmt.Errorf("invalid syntax for X coordinate: %s", t.Value())
		}
		return
	}

	if p.Y, err = t.Next().Double(); err != nil {
		if t.AtEnd() {
			err = EndOfTokens
		} else {
			err = fmt.Errorf("invalid syntax for Y coordinate: %s", t.Value())
		}
		return
	}

	point = p
	return
}
