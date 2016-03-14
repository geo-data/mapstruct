package size

import (
	"fmt"
	"github.com/geo-data/mapfile/encoding"
	"github.com/geo-data/mapfile/tokens"
)

type Size struct {
	Width, Height tokens.Uint32
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

func (s *Size) Encode(enc *encoding.MapfileEncoder) error {
	return enc.TokenStringer("SIZE", s)
}
