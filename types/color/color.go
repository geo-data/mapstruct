package color

import (
	"fmt"
	"github.com/geo-data/mapfile/types"
)

type Color struct {
	R, G, B types.Uint8
}

func (c *Color) String() string {
	return fmt.Sprintf("%d %d %d", c.R, c.G, c.B)
}
