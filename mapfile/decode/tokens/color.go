package tokens

import "github.com/geo-data/mapfile/types/color"

func (t *Tokens) Color() (c *color.Color, err error) {
	c = new(color.Color)
	if c.R, err = t.Next().Uint8(); err != nil {
		return
	}

	if c.G, err = t.Next().Uint8(); err != nil {
		return
	}

	if c.B, err = t.Next().Uint8(); err != nil {
		return
	}

	return
}
