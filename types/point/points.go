package point

import (
	"encoding/json"
	"fmt"
	"github.com/geo-data/mapfile/encoding"
	"github.com/geo-data/mapfile/tokens"
	"github.com/geo-data/mapfile/types"
)

type Point struct {
	X, Y types.Double
}

func NewPoint(tokens *tokens.Tokens) (p *Point, err error) {
	p = new(Point)
	if p.X, err = tokens.Double(); err != nil {
		return
	}

	if p.Y, err = tokens.Next().Double(); err != nil {
		return
	}

	return
}

func (p *Point) String() string {
	return fmt.Sprintf("%s %s", p.X, p.Y)
}

func (p *Point) MarshalJSON() ([]byte, error) {
	return json.Marshal([]types.Double{p.X, p.Y})
}

type Points struct {
	points []*Point
}

func NewPoints(tokens *tokens.Tokens) (ps *Points, err error) {
	token := tokens.Value()
	if token != "POINTS" {
		err = fmt.Errorf("expected token POINTS, got: %s", token)
		return
	}
	tokens.Next()

	ps = new(Points)
	for tokens != nil {
		if tokens.Value() == "END" {
			break
		}

		var point *Point
		if point, err = NewPoint(tokens); err != nil {
			return
		}
		ps.points = append(ps.points, point)

		tokens = tokens.Next()
	}
	return
}

func (p *Points) MarshalJSON() ([]byte, error) {
	return json.Marshal(p.points)
}

func (p *Points) Encode(enc *encoding.MapfileEncoder) (err error) {
	if err = enc.TokenStart("POINTS"); err != nil {
		return
	}

	var points []fmt.Stringer
	for _, point := range p.points {
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
