package color

import (
	"fmt"
	"github.com/geo-data/mapfile/tokens"
	"github.com/geo-data/mapfile/types"
)

type Color struct {
	R, G, B types.Uint8
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

func (c *Color) String() string {
	return fmt.Sprintf("%d %d %d", c.R, c.G, c.B)
}
