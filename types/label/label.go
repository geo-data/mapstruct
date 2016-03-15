package label

import (
	"fmt"
	"github.com/geo-data/mapfile/encoding"
	"github.com/geo-data/mapfile/tokens"
	"github.com/geo-data/mapfile/types"
	"github.com/geo-data/mapfile/types/color"
	"github.com/geo-data/mapfile/types/style"
)

type Label struct {
	Type     types.Keyword  `json:",omitempty"`
	Size     types.Union    `json:",omitempty"`
	Font     types.String   `json:",omitempty"`
	Color    *color.Color   `json:",omitempty"`
	Position types.Keyword  `json:",omitempty"`
	Buffer   types.Uint32   `json:",omitempty"`
	Styles   []*style.Style `json:",omitempty"`
}

func New(toks *tokens.Tokens) (l *Label, err error) {
	token := toks.Value()
	if token != "LABEL" {
		err = fmt.Errorf("expected token LABEL, got: %s", token)
		return
	}
	toks.Next()

	l = new(Label)
	for toks != nil {
		token := toks.Value()
		switch token {
		case "TYPE":
			if l.Type, err = toks.Next().Keyword(); err != nil {
				return
			}
		case "SIZE":
			if l.Size, err = toks.Next().Decode(tokens.Double | tokens.Keyword | tokens.Attribute); err != nil {
				return
			}
		case "FONT":
			if l.Font, err = toks.Next().String(); err != nil {
				return
			}
		case "BUFFER":
			if l.Buffer, err = toks.Next().Uint32(); err != nil {
				return
			}
		case "POSITION":
			if l.Position, err = toks.Next().Keyword(); err != nil {
				return
			}
		case "COLOR":
			if l.Color, err = color.New(toks); err != nil {
				return
			}
		case "STYLE":
			var s *style.Style
			if s, err = style.New(toks); err != nil {
				return
			}
			l.Styles = append(l.Styles, s)
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

func (l *Label) Encode(enc *encoding.MapfileEncoder) (err error) {
	if err = enc.TokenStart("LABEL"); err != nil {
		return
	}

	if err = enc.TokenStringer("TYPE", l.Type); err != nil {
		return
	}
	if err = enc.TokenUnion("SIZE", l.Size); err != nil {
		return
	}
	if err = enc.TokenStringer("FONT", l.Font); err != nil {
		return
	}
	if l.Color != nil {
		if err = enc.TokenStringer("COLOR", l.Color); err != nil {
			return
		}
	}
	if err = enc.TokenStringer("POSITION", l.Position); err != nil {
		return
	}
	if uint32(l.Buffer) > uint32(0) {
		if err = enc.TokenStringer("BUFFER", l.Buffer); err != nil {
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
