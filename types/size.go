package types

import "fmt"

type Size struct {
	Width, Height Uint32
}

func NewSize(width, height uint) *Size {
	return &Size{
		Uint32(width),
		Uint32(height),
	}
}

func (s *Size) String() string {
	return fmt.Sprintf("%s %s", s.Width, s.Height)
}
