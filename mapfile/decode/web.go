package decode

import (
	"fmt"
	"github.com/geo-data/mapstruct/mapfile/decode/scanner"
	"github.com/geo-data/mapstruct/types"
)

func (t *Decoder) Web() (web *types.Web, err error) {
	var token *scanner.Token
	if token, err = t.ExpectedToken(scanner.WEB); err != nil {
		return
	}
	t.Next()

	w := new(types.Web)
Loop:
	for t != nil {
		if token, err = t.Token(); err != nil {
			return
		}

		switch token.Type {
		case scanner.METADATA:
			if w.Metadata, err = t.Metadata(); err != nil {
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

	web = w
	return
}
