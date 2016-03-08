package size

import (
	"github.com/geo-data/mapfile/tokens"
	"strconv"
)

type Size struct {
	Width, Height uint
}

func (s *Size) FromTokens(tokens *tokens.Tokens) error {
	i, err := strconv.ParseUint(tokens.Next().Value(), 10, 32)
	if err != nil {
		return err
	}
	s.Width = uint(i)

	i, err = strconv.ParseUint(tokens.Next().Value(), 10, 32)
	if err != nil {
		return err
	}
	s.Height = uint(i)

	return nil
}
