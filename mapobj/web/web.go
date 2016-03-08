package web

import (
	"fmt"
	"github.com/geo-data/mapfile/mapobj/metadata"
	"github.com/geo-data/mapfile/tokens"
)

type Web struct {
	Metadata metadata.Metadata
}

func (w *Web) FromTokens(tokens *tokens.Tokens) error {
	token := tokens.Value()
	if token != "WEB" {
		return fmt.Errorf("expected token WEB, got: %s", token)
	}
	tokens.Next()

	for tokens != nil {
		token := tokens.Value()
		switch token {
		case "METADATA":
			if err := w.Metadata.FromTokens(tokens); err != nil {
				return err
			}
		case "END":
			return nil
		default:
			return fmt.Errorf("unhandled mapfile token: %s", token)
		}

		tokens = tokens.Next()
	}

	return nil
}
