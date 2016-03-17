package decode

import "github.com/geo-data/mapfile/types"

func (t *Decoder) Color() (color *types.Color, err error) {
	c := new(types.Color)
	if c.R, err = t.Next().Uint8(); err != nil {
		if t.AtEnd() {
			err = EndOfTokens
		}
		return
	}

	if c.G, err = t.Next().Uint8(); err != nil {
		if t.AtEnd() {
			err = EndOfTokens
		}
		return
	}

	if c.B, err = t.Next().Uint8(); err != nil {
		if t.AtEnd() {
			err = EndOfTokens
		}
		return
	}

	c.A = 255

	color = c
	return
}
