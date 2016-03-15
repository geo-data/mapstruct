package decode

import (
	"fmt"
	"github.com/geo-data/mapfile/types/label"
	"github.com/geo-data/mapfile/types/style"
)

func (t *Decoder) Label() (l *label.Label, err error) {
	token := t.Value()
	if token != "LABEL" {
		err = fmt.Errorf("expected token LABEL, got: %s", token)
		return
	}
	t.Next()

	l = new(label.Label)
	for t != nil {
		token := t.Value()
		switch token {
		case "TYPE":
			if l.Type, err = t.Next().Keyword(); err != nil {
				return
			}
		case "SIZE":
			if l.Size, err = t.Next().Decode(Double | Keyword | Attribute); err != nil {
				return
			}
		case "FONT":
			if l.Font, err = t.Next().String(); err != nil {
				return
			}
		case "BUFFER":
			if l.Buffer, err = t.Next().Uint32(); err != nil {
				return
			}
		case "POSITION":
			if l.Position, err = t.Next().Keyword(); err != nil {
				return
			}
		case "COLOR":
			if l.Color, err = t.Color(); err != nil {
				return
			}
		case "STYLE":
			var s *style.Style
			if s, err = t.Style(); err != nil {
				return
			}
			l.Styles = append(l.Styles, s)
		case "END":
			return
		default:
			err = fmt.Errorf("unhandled mapfile token: %s", token)
			return
		}

		t = t.Next()
	}

	return
}
