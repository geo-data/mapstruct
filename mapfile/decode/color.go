package decode

import "github.com/geo-data/mapfile/types"

func (t *Decoder) Color() (c *types.Color, err error) {
	c = new(types.Color)
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
