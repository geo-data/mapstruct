package extent

import (
	"encoding/json"
	"fmt"
	"github.com/geo-data/mapfile/mapfile/encode"
	"github.com/geo-data/mapfile/tokens"
	"github.com/geo-data/mapfile/types"
	"github.com/geo-data/mapfile/types/point"
)

type Extent struct {
	Min, Max *point.Point
}

func New(tokens *tokens.Tokens) (e *Extent, err error) {
	e = new(Extent)
	if e.Min, err = point.NewPoint(tokens.Next()); err != nil {
		return
	}
	if e.Max, err = point.NewPoint(tokens.Next()); err != nil {
		return
	}
	return
}

func (e *Extent) MarshalJSON() ([]byte, error) {
	a := []types.Double{
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

func (e *Extent) Encode(enc *encode.MapfileEncoder) error {
	return enc.TokenStringer("EXTENT", e)
}
