package label

import (
	"fmt"
	"github.com/geo-data/mapfile/mapobj/color"
	"github.com/geo-data/mapfile/tokens"
)

type Label struct {
	Type     string       `json:",omitempty"`
	Size     string       `json:",omitempty"`
	Color    *color.Color `json:",omitempty"`
	Position string       `json:",omitempty"`
	Buffer   uint32       `json:",omitempty"`
}

func New(tokens *tokens.Tokens) (l *Label, err error) {
	token := tokens.Value()
	if token != "LABEL" {
		err = fmt.Errorf("expected token LABEL, got: %s", token)
		return
	}
	tokens.Next()

	l = new(Label)
	for tokens != nil {
		token := tokens.Value()
		switch token {
		case "TYPE":
			l.Type = tokens.Next().Value()
		case "SIZE":
			l.Size = tokens.Next().Value()
		case "BUFFER":
			if l.Buffer, err = tokens.Next().Uint32(); err != nil {
				return
			}
		case "POSITION":
			l.Position = tokens.Next().Value()
		case "COLOR":
			if l.Color, err = color.New(tokens); err != nil {
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
