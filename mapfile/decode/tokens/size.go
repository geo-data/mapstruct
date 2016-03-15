package tokens

import (
	"fmt"
	"github.com/geo-data/mapfile/types/size"
)

func (t *Tokens) Size() (s *size.Size, err error) {
	token := t.Value()
	if token != "SIZE" {
		err = fmt.Errorf("expected token SIZE, got: %s", token)
		return
	}

	s = new(size.Size)
	if s.Width, err = t.Next().Uint32(); err != nil {
		return
	}

	if s.Height, err = t.Next().Uint32(); err != nil {
		return
	}

	return
}
