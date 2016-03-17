package decode

import (
	"fmt"
	"github.com/geo-data/mapfile/types"
)

func (t *Decoder) Size() (size *types.Size, err error) {
	token := t.Value()
	if token != "SIZE" {
		err = fmt.Errorf("expected token SIZE, got: %s", token)
		return
	}

	s := new(types.Size)
	if s.Width, err = t.Next().Uint32(); err != nil {
		if t.AtEnd() {
			err = EndOfTokens
		} else {
			err = fmt.Errorf("invalid syntax for width: %s", t.Value())
		}
		return
	}

	if s.Height, err = t.Next().Uint32(); err != nil {
		if t.AtEnd() {
			err = EndOfTokens
		} else {
			err = fmt.Errorf("invalid syntax for height: %s", t.Value())
		}
		return
	}

	size = s
	return
}
