package encode

import (
	"github.com/geo-data/mapfile/types"
	"github.com/geo-data/mapfile/types/metadata"
)

func (enc *Encoder) EncodeMetadata(m metadata.Metadata) (err error) {
	if err = enc.StartDirective("METADATA"); err != nil {
		return
	}

	for k, v := range m {
		if err = enc.EncodeStringers(types.String(k), types.String(v)); err != nil {
			return
		}
	}

	if err = enc.EndDirective(); err != nil {
		return
	}

	return
}
