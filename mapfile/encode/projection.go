package encode

import (
	"github.com/geo-data/mapfile/types"
	"github.com/geo-data/mapfile/types/projection"
)

func (enc *Encoder) EncodeProjection(p projection.Projection) (err error) {
	if err = enc.StartDirective("PROJECTION"); err != nil {
		return
	}

	for _, param := range p {
		if err = enc.EncodeString(types.String(param)); err != nil {
			return
		}
	}

	if err = enc.EndDirective(); err != nil {
		return
	}

	return
}
