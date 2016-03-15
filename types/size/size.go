package size

import (
	"fmt"
	"github.com/geo-data/mapfile/mapfile/encode"
	"github.com/geo-data/mapfile/types"
)

type Size struct {
	Width, Height types.Uint32
}

func (s *Size) String() string {
	return fmt.Sprintf("%s %s", s.Width, s.Height)
}

func (s *Size) Encode(enc *encode.MapfileEncoder) error {
	return enc.TokenStringer("SIZE", s)
}
