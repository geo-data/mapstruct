package projection

import (
	"github.com/geo-data/mapfile/mapfile/encode"
	"github.com/geo-data/mapfile/types"
)

type Projection []string

func (p Projection) Encode(enc *encode.MapfileEncoder) (err error) {
	if err = enc.TokenStart("PROJECTION"); err != nil {
		return
	}

	for _, param := range p {
		if err = enc.EncodeString(types.String(param)); err != nil {
			return
		}
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
