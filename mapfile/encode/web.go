package encode

import "github.com/geo-data/mapfile/types/web"

func (enc *Encoder) EncodeWeb(w *web.Web) (err error) {
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
