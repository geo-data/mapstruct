package size

import (
	"github.com/geo-data/mapfile/tokens"
)

type Size struct {
	Width, Height uint32
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
