package label

import (
	"fmt"
	"github.com/geo-data/mapfile/encoding"
	"github.com/geo-data/mapfile/mapobj/color"
	"github.com/geo-data/mapfile/mapobj/style"
	"github.com/geo-data/mapfile/tokens"
)

type Label struct {
	Type     tokens.String  `json:",omitempty"`
	Size     tokens.String  `json:",omitempty"`
	Font     tokens.String  `json:",omitempty"`
	Color    *color.Color   `json:",omitempty"`
	Position tokens.String  `json:",omitempty"`
	Buffer   tokens.Uint32  `json:",omitempty"`
	Styles   []*style.Style `json:",omitempty"`
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
		case "FONT":
			l.Font = tokens.Next().Value()
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
		case "STYLE":
			var s *style.Style
			if s, err = style.New(tokens); err != nil {
				return
			}
			l.Styles = append(l.Styles, s)
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

func (l *Label) Encode(enc *encoding.MapfileEncoder) (err error) {
	if err = enc.TokenStart("LABEL"); err != nil {
		return
	}

	if err = enc.TokenValue("TYPE", l.Type); err != nil {
		return
	}
	if err = enc.TokenValue("SIZE", l.Size); err != nil {
		return
	}
	if err = enc.TokenString("FONT", l.Font); err != nil {
		return
	}
	if l.Color != nil {
		if err = enc.TokenValue("COLOR", l.Color); err != nil {
			return
		}
	}
	if err = enc.TokenValue("POSITION", l.Position); err != nil {
		return
	}
	if uint32(l.Buffer) > uint32(0) {
		if err = enc.TokenValue("BUFFER", l.Buffer); err != nil {
			return
		}
	}

	for _, style := range l.Styles {
		if err = style.Encode(enc); err != nil {
			return
		}
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
