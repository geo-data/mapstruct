package legend

import (
	"fmt"
	"github.com/geo-data/mapfile/mapobj/color"
	"github.com/geo-data/mapfile/tokens"
)

type Legend struct {
	ImageColor color.Color
}

func (l *Legend) FromTokens(tokens *tokens.Tokens) error {
	token := tokens.Value()
	if token != "LEGEND" {
		return fmt.Errorf("expected token LEGEND, got: %s", token)
	}
	tokens.Next()

	for tokens != nil {
		token := tokens.Value()
		switch token {
		case "IMAGECOLOR":
			if err := l.ImageColor.FromTokens(tokens); err != nil {
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
