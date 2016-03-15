package point

import (
	"encoding/json"
	"fmt"

	"github.com/geo-data/mapfile/mapfile/encode"
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

func (p Points) Encode(enc *encode.MapfileEncoder) (err error) {
	if err = enc.TokenStart("POINTS"); err != nil {
		return
	}

	var points []fmt.Stringer
	for _, point := range p {
		points = append(points, point)
	}

	if err = enc.EncodeStringers(points...); err != nil {
		return
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
