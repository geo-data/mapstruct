package extent

import (
	"encoding/json"
	"fmt"
	"github.com/geo-data/mapfile/encoding"
	"github.com/geo-data/mapfile/tokens"
)

type Point struct {
	X, Y tokens.Float64
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
	a := []tokens.Float64{
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

func (e *Extent) Encode(enc *encoding.MapfileEncoder) error {
	return enc.TokenValue("EXTENT", e)
}
