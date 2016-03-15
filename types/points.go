package types

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	X, Y Double
}

func (p *Point) String() string {
	return fmt.Sprintf("%s %s", p.X, p.Y)
}

func (p *Point) MarshalJSON() ([]byte, error) {
	return json.Marshal([]Double{p.X, p.Y})
}

type Points []*Point
