package metadata

import (
	"github.com/geo-data/mapfile/mapfile/encode"
	"github.com/geo-data/mapfile/types"
)

type Metadata map[string]string

func New() Metadata {
	return Metadata(make(map[string]string))
}

func (p Metadata) Encode(enc *encode.MapfileEncoder) (err error) {
	if err = enc.TokenStart("METADATA"); err != nil {
		return
	}

	for k, v := range p {
		if err = enc.EncodeStringers(types.String(k), types.String(v)); err != nil {
			return
		}
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
