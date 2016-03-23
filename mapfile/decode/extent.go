package decode

import (
	"github.com/geo-data/mapstruct/mapfile/decode/scanner"
	"github.com/geo-data/mapstruct/types"
)

func (t *Decoder) Extent() (extent *types.Extent, err error) {
	if _, err = t.ExpectedToken(scanner.EXTENT); err != nil {
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
