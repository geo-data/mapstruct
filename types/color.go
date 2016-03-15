package types

import "fmt"

type Color struct {
	R, G, B Uint8
}

func (c *Color) String() string {
	return fmt.Sprintf("%d %d %d", c.R, c.G, c.B)
}
