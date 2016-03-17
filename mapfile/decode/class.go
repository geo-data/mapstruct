package decode

import (
	"fmt"
	"github.com/geo-data/mapfile/types"
)

func (t *Decoder) Class() (class *types.Class, err error) {
	token := t.Value()
	if token != "CLASS" {
		err = fmt.Errorf("expected token CLASS, got: %s", token)
		return
	}
	t.Next()

	c := new(types.Class)
Loop:
	for t != nil {
		token := t.Value()
		switch token {
		case "NAME":
			if c.Name, err = t.Next().String(); err != nil {
				return
			}
		case "EXPRESSION":
			if c.Expression, err = t.Next().String(); err != nil {
				return
			}
		case "TEMPLATE":
			if c.Template, err = t.Next().String(); err != nil {
				return
			}
		case "TEXT":
			if c.Text, err = t.Next().String(); err != nil {
				return
			}
		case "METADATA":
			if c.Metadata, err = t.Metadata(); err != nil {
				return
			}
		case "STYLE":
			var s *types.Style
			if s, err = t.Style(); err != nil {
				return
			}
			c.Styles = append(c.Styles, s)
		case "LABEL":
			if c.Label, err = t.Label(); err != nil {
				return
			}
		case "END":
			break Loop
		case "":
			if t.AtEnd() {
				err = EndOfTokens
				return
			}
			fallthrough
		default:
			err = fmt.Errorf("unhandled mapfile token: %s", token)
			return
		}

		t = t.Next()
	}

	class = c
	return
}
