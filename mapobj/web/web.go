package web

import (
	"fmt"
	"github.com/geo-data/mapfile/mapobj/metadata"
	"github.com/geo-data/mapfile/tokens"
)

type Web struct {
	Metadata *metadata.Metadata `json:",omitempty"`
}

func New(tokens *tokens.Tokens) (w *Web, err error) {
	token := tokens.Value()
	if token != "WEB" {
		err = fmt.Errorf("expected token WEB, got: %s", token)
		return
	}
	tokens.Next()

	w = new(Web)
	for tokens != nil {
		token := tokens.Value()
		switch token {
		case "METADATA":
			if w.Metadata, err = metadata.New(tokens); err != nil {
				return
			}
		case "END":
			return
		default:
			err = fmt.Errorf("unhandled mapfile token: %s", token)
			return
		}

		tokens = tokens.Next()
	}

	return
}
