package decode

import (
	"fmt"
	"github.com/geo-data/mapstruct/mapfile/decode/scanner"
	"github.com/geo-data/mapstruct/types"
)

func (t *Decoder) Size() (size *types.Size, err error) {
	if _, err = t.ExpectedToken(scanner.SIZE); err != nil {
		return
	}

	s := new(types.Size)
	if s.Width, err = t.Next().Uint32(); err != nil {
		err = fmt.Errorf("could not decode width: %s", err)
		return
	}

	if s.Height, err = t.Next().Uint32(); err != nil {
		err = fmt.Errorf("could not decode height: %s", err)
		return
	}

	size = s
	return
}
