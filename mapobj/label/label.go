package label

import (
	"fmt"
	"github.com/geo-data/mapfile/mapobj/color"
	"github.com/geo-data/mapfile/tokens"
	"strconv"
)

type Label struct {
	Type     string
	Size     string
	Color    color.Color
	Position string
	Buffer   uint
}

func (l *Label) FromTokens(tokens *tokens.Tokens) error {
	token := tokens.Value()
	if token != "LABEL" {
		return fmt.Errorf("expected token LABEL, got: %s", token)
	}
	tokens.Next()

	for tokens != nil {
		token := tokens.Value()
		switch token {
		case "TYPE":
			l.Type = tokens.Next().Value()
		case "SIZE":
			l.Size = tokens.Next().Value()
		case "BUFFER":
			i, err := strconv.ParseUint(tokens.Next().Value(), 10, 32)
			if err != nil {
				return err
			}
			l.Buffer = uint(i)
		case "POSITION":
			l.Position = tokens.Next().Value()
		case "COLOR":
			if err := l.Color.FromTokens(tokens); err != nil {
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
