package encode

import (
	"github.com/geo-data/mapfile/types"
	"github.com/geo-data/mapfile/types/metadata"
)

func (enc *Encoder) EncodeMetadata(m metadata.Metadata) (err error) {
	if err = enc.TokenStart("METADATA"); err != nil {
		return
	}

	for k, v := range m {
		if err = enc.EncodeStringers(types.String(k), types.String(v)); err != nil {
			return
		}
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
