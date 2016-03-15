package types

import "fmt"

type Size struct {
	Width, Height Uint32
}

func (s *Size) String() string {
	return fmt.Sprintf("%s %s", s.Width, s.Height)
}
