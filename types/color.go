package types

import "fmt"

type Color struct {
	R, G, B, A Uint8
}

func NewColor(R, G, B, A uint8) *Color {
	return &Color{
		Uint8(R),
		Uint8(G),
		Uint8(B),
		Uint8(A),
	}
}

func (c *Color) String() string {
	return fmt.Sprintf("%d %d %d", c.R, c.G, c.B)
}
