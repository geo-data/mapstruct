package extent

import (
	"encoding/json"
	"github.com/geo-data/mapfile/tokens"
)

type Point struct {
	X, Y float64
}

func (p *Point) FromTokens(tokens *tokens.Tokens) (err error) {
	if p.X, err = tokens.Next().Float64(); err != nil {
		return
	}

	if p.Y, err = tokens.Next().Float64(); err != nil {
		return
	}

	return
}

type Extent struct {
	Min, Max Point
}

func New(tokens *tokens.Tokens) (e *Extent, err error) {
	e = new(Extent)
	if err = e.Min.FromTokens(tokens); err != nil {
		return
	}
	if err = e.Max.FromTokens(tokens); err != nil {
		return
	}
	return
}

func (e *Extent) MarshalJSON() ([]byte, error) {
	a := []float64{
		e.Min.X,
		e.Min.Y,
		e.Max.X,
		e.Max.Y,
	}
	return json.Marshal(a)
}
