package color

import (
	"github.com/geo-data/mapfile/tokens"
	"strconv"
)

type Color struct {
	R, G, B uint
}

func (c *Color) FromTokens(tokens *tokens.Tokens) error {
	i, err := strconv.ParseUint(tokens.Next().Value(), 10, 8)
	if err != nil {
		return err
	}
	c.R = uint(i)

	i, err = strconv.ParseUint(tokens.Next().Value(), 10, 8)
	if err != nil {
		return err
	}
	c.G = uint(i)

	i, err = strconv.ParseUint(tokens.Next().Value(), 10, 8)
	if err != nil {
		return err
	}
	c.B = uint(i)

	return nil
}
