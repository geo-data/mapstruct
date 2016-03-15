package tokens

import (
	"fmt"
	"github.com/geo-data/mapfile/types/web"
)

func (t *Tokens) Web() (w *web.Web, err error) {
	token := t.Value()
	if token != "WEB" {
		err = fmt.Errorf("expected token WEB, got: %s", token)
		return
	}
	t.Next()

	w = new(web.Web)
	for t != nil {
		token := t.Value()
		switch token {
		case "METADATA":
			if w.Metadata, err = t.Metadata(); err != nil {
				return
			}
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