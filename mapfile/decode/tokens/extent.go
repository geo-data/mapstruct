package tokens

import (
	"fmt"
	"github.com/geo-data/mapfile/types/extent"
)

func (t *Tokens) Extent() (e *extent.Extent, err error) {
	token := t.Value()
	if token != "EXTENT" {
		err = fmt.Errorf("expected token EXTENT, got: %s", token)
		return
	}
	t.Next()

	e = new(extent.Extent)
	if e.Min, err = t.Point(); err != nil {
		return
	}
	t.Next()

	if e.Max, err = t.Point(); err != nil {
		return
	}
	return
}
