package encode

import "github.com/geo-data/mapfile/types/web"

func (enc *Encoder) EncodeWeb(w *web.Web) (err error) {
	if err = enc.TokenStart("WEB"); err != nil {
		return
	}

	if w.Metadata != nil {
		if err = enc.EncodeMetadata(w.Metadata); err != nil {
			return
		}
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
