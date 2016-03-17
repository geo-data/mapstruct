package decode

import (
	"fmt"
	"github.com/geo-data/mapfile/types"
)

func (t *Decoder) Extent() (extent *types.Extent, err error) {
	token := t.Value()
	if token != "EXTENT" {
		err = fmt.Errorf("expected token EXTENT, got: %s", token)
		return
	}
	t.Next()

	e := new(types.Extent)
	if e.Min, err = t.Point(); err != nil {
		return
	}
	t.Next()

	if e.Max, err = t.Point(); err != nil {
		return
	}

	extent = e
	return
}
