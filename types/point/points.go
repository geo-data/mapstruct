package point

import (
	"encoding/json"
	"fmt"

	"github.com/geo-data/mapfile/types"
)

type Point struct {
	X, Y types.Double
}

func (p *Point) String() string {
	return fmt.Sprintf("%s %s", p.X, p.Y)
}

func (p *Point) MarshalJSON() ([]byte, error) {
	return json.Marshal([]types.Double{p.X, p.Y})
}

type Points []*Point
