package types

import (
	"encoding/json"
	"fmt"
)

type Point struct {
	X, Y Double
}

func NewPoint(x, y float64) *Point {
	return &Point{
		X: Double(x),
		Y: Double(y),
	}
}

func (p *Point) String() string {
	return fmt.Sprintf("%s %s", p.X, p.Y)
}

func (p *Point) MarshalJSON() ([]byte, error) {
	return json.Marshal([]Double{p.X, p.Y})
}

type Points []*Point

func NewPoints(points []float64) (p Points, err error) {
	if len(points)%2 != 0 {
		err = fmt.Errorf("number of points is not even: %d", len(points))
		return
	}

	for i := 0; i < len(points); i += 2 {
		x := points[i]
		y := points[i+1]
		p = append(p, NewPoint(x, y))
	}
	return
}
