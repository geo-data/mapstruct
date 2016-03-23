package decode

import (
	"fmt"
	"github.com/geo-data/mapstruct/mapfile/decode/scanner"
	"github.com/geo-data/mapstruct/types"
)

func (t *Decoder) Class() (class *types.Class, err error) {
	var token *scanner.Token
	if token, err = t.ExpectedToken(scanner.CLASS); err != nil {
		return
	}
	t.Next()

	c := new(types.Class)
Loop:
	for t != nil {
		if token, err = t.Token(); err != nil {
			return
		}

		switch token.Type {
		case scanner.NAME:
			if c.Name, err = t.Next().String(); err != nil {
				return
			}
		case scanner.EXPRESSION:
			if c.Expression, err = t.Next().Decode(String | Expression | Regex | Listex); err != nil {
				return
			}
		case scanner.TEMPLATE:
			if c.Template, err = t.Next().String(); err != nil {
				return
			}
		case scanner.TEXT:
			if c.Text, err = t.Next().String(); err != nil {
				return
			}
		case scanner.METADATA:
			if c.Metadata, err = t.Metadata(); err != nil {
				return
			}
		case scanner.STYLE:
			var s *types.Style
			if s, err = t.Style(); err != nil {
				return
			}
			c.Styles = append(c.Styles, s)
		case scanner.LABEL:
			if c.Label, err = t.Label(); err != nil {
				return
			}
		case scanner.END:
			break Loop
		default:
			err = fmt.Errorf("unhandled mapfile token: %s", token)
			return
		}

		t = t.Next()
	}

	class = c
	return
}
