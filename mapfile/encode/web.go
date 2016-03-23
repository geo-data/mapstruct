package encode

import "github.com/geo-data/mapstruct/types"

func (enc *Encoder) EncodeWeb(w *types.Web) (err error) {
	if err = enc.StartDirective("WEB"); err != nil {
		return
	}

	if w.Metadata != nil {
		if err = enc.EncodeMetadata(w.Metadata); err != nil {
			return
		}
	}

	if err = enc.EndDirective(); err != nil {
		return
	}

	return
}
