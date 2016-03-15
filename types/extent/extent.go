package extent

import (
	"encoding/json"
	"fmt"

	"github.com/geo-data/mapfile/types"
	"github.com/geo-data/mapfile/types/point"
)

type Extent struct {
	Min, Max *point.Point
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
