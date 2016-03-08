package extent

import (
	"github.com/geo-data/mapfile/tokens"
	"strconv"
)

type Point struct {
	Min, Max float64
}

func (p *Point) FromTokens(tokens *tokens.Tokens) error {
	f, err := strconv.ParseFloat(tokens.Next().Value(), 64)
	if err != nil {
		return err
	}
	p.Min = f

	f, err = strconv.ParseFloat(tokens.Next().Value(), 64)
	if err != nil {
		return err
	}
	p.Max = f

	return nil
}

type Extent struct {
	Min, Max Point
}

func (e *Extent) FromTokens(tokens *tokens.Tokens) error {
	if err := e.Min.FromTokens(tokens); err != nil {
		return err
	}
	if err := e.Max.FromTokens(tokens); err != nil {
		return err
	}
	return nil
}
