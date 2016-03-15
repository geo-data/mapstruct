package size

import (
	"fmt"
	"github.com/geo-data/mapfile/mapfile/encode"
	"github.com/geo-data/mapfile/tokens"
	"github.com/geo-data/mapfile/types"
)

type Size struct {
	Width, Height types.Uint32
}

func New(tokens *tokens.Tokens) (s *Size, err error) {
	s = new(Size)
	if s.Width, err = tokens.Next().Uint32(); err != nil {
		return
	}

	if s.Height, err = tokens.Next().Uint32(); err != nil {
		return
	}

	return
}

func (s *Size) String() string {
	return fmt.Sprintf("%s %s", s.Width, s.Height)
}

func (s *Size) Encode(enc *encode.MapfileEncoder) error {
	return enc.TokenStringer("SIZE", s)
}
