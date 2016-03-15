package encode

import "github.com/geo-data/mapfile/types"

func (enc *Encoder) EncodeMetadata(m types.Metadata) (err error) {
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
