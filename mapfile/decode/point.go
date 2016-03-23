package decode

import (
	"fmt"
	"github.com/geo-data/mapstruct/types"
)

func (t *Decoder) Point() (point *types.Point, err error) {
	p := new(types.Point)
	if p.X, err = t.Double(); err != nil {
		err = fmt.Errorf("could not decode X coordinate: %s", err)
		return
	}

	if p.Y, err = t.Next().Double(); err != nil {
		err = fmt.Errorf("could not decode Y coordinate: %s", err)
		return
	}

	point = p
	return
}
