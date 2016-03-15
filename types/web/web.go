package web

import (
	"github.com/geo-data/mapfile/mapfile/encode"
	"github.com/geo-data/mapfile/types/metadata"
)

type Web struct {
	Metadata metadata.Metadata `json:",omitempty"`
}

func (w *Web) Encode(enc *encode.MapfileEncoder) (err error) {
	if err = enc.TokenStart("WEB"); err != nil {
		return
	}

	if w.Metadata != nil {
		if err = w.Metadata.Encode(enc); err != nil {
			return
		}
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
