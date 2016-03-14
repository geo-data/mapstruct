package web

import (
	"fmt"
	"github.com/geo-data/mapfile/encoding"
	"github.com/geo-data/mapfile/types/metadata"
	"github.com/geo-data/mapfile/tokens"
)

type Web struct {
	Metadata *metadata.Metadata `json:",omitempty"`
}

func New(toks *tokens.Tokens) (w *Web, err error) {
	token := toks.Value()
	if token != "WEB" {
		err = fmt.Errorf("expected token WEB, got: %s", token)
		return
	}
	toks.Next()

	w = new(Web)
	for toks != nil {
		token := toks.Value()
		switch token {
		case "METADATA":
			if w.Metadata, err = metadata.New(toks); err != nil {
				return
			}
		case "END":
			return
		default:
			err = fmt.Errorf("unhandled mapfile token: %s", token)
			return
		}

		toks = toks.Next()
	}

	return
}

func (w *Web) Encode(enc *encoding.MapfileEncoder) (err error) {
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
