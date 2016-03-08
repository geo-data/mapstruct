package color

import (
	"github.com/geo-data/mapfile/tokens"
)

type Color struct {
	R, G, B uint8
}

func New(tokens *tokens.Tokens) (c *Color, err error) {
	c = new(Color)
	if c.R, err = tokens.Next().Uint8(); err != nil {
		return
	}

	if c.G, err = tokens.Next().Uint8(); err != nil {
		return
	}

	if c.B, err = tokens.Next().Uint8(); err != nil {
		return
	}

	return
}
