package decode

import (
	"fmt"
	"github.com/geo-data/mapstruct/mapfile/decode/scanner"
	"github.com/geo-data/mapstruct/types"
)

func (t *Decoder) Label() (label *types.Label, err error) {
	var token *scanner.Token
	if token, err = t.ExpectedToken(scanner.LABEL); err != nil {
		return
	}
	t.Next()

	l := new(types.Label)
Loop:
	for t != nil {
		if token, err = t.Token(); err != nil {
			return
		}

		switch token.Type {
		case scanner.TYPE:
			if l.Type, err = t.Next().Keyword(); err != nil {
				return
			}
		case scanner.SIZE:
			if l.Size, err = t.Next().Decode(Double | Keyword | Attribute); err != nil {
				return
			}
		case scanner.FONT:
			if l.Font, err = t.Next().String(); err != nil {
				return
			}
		case scanner.BUFFER:
			if l.Buffer, err = t.Next().Uint32(); err != nil {
				return
			}
		case scanner.POSITION:
			if l.Position, err = t.Next().Keyword(); err != nil {
				return
			}
		case scanner.COLOR:
			if l.Color, err = t.Color(); err != nil {
				return
			}
		case scanner.STYLE:
			var s *types.Style
			if s, err = t.Style(); err != nil {
				return
			}
			l.Styles = append(l.Styles, s)
		case scanner.END:
			break Loop
		default:
			err = fmt.Errorf("unhandled mapfile token: %s", token)
			return
		}

		t = t.Next()
	}

	label = l
	return
}
