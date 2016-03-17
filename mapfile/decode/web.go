package decode

import (
	"fmt"
	"github.com/geo-data/mapfile/types"
)

func (t *Decoder) Web() (web *types.Web, err error) {
	token := t.Value()
	if token != "WEB" {
		err = fmt.Errorf("expected token WEB, got: %s", token)
		return
	}
	t.Next()

	w := new(types.Web)
Loop:
	for t != nil {
		token := t.Value()
		switch token {
		case "METADATA":
			if w.Metadata, err = t.Metadata(); err != nil {
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

	web = w
	return
}
