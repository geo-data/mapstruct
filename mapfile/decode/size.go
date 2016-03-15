package decode

import (
	"fmt"
	"github.com/geo-data/mapfile/types"
)

func (t *Decoder) Size() (s *types.Size, err error) {
	token := t.Value()
	if token != "SIZE" {
		err = fmt.Errorf("expected token SIZE, got: %s", token)
		return
	}

	s = new(types.Size)
	if s.Width, err = t.Next().Uint32(); err != nil {
		return
	}

	if s.Height, err = t.Next().Uint32(); err != nil {
		return
	}

	return
}
