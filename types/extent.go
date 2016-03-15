package types

import (
	"encoding/json"
	"fmt"
)

type Extent struct {
	Min, Max *Point
}

func NewExtent(minx, miny, maxx, maxy float64) *Extent {
	return &Extent{
		Min: NewPoint(minx, miny),
		Max: NewPoint(maxx, maxy),
	}
}

func (e *Extent) MarshalJSON() ([]byte, error) {
	a := []Double{
		e.Min.X,
		e.Min.Y,
		e.Max.X,
		e.Max.Y,
	}
	return json.Marshal(a)
}

func (e *Extent) String() string {
	return fmt.Sprintf("%s %s %s %s", e.Min.X, e.Min.Y, e.Max.X, e.Max.Y)
}
