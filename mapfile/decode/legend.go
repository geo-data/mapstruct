package decode

import (
	"fmt"
	"github.com/geo-data/mapfile/mapfile/decode/scanner"
	"github.com/geo-data/mapfile/types"
)

func (t *Decoder) Legend() (legend *types.Legend, err error) {
	var token *scanner.Token
	if token, err = t.ExpectedToken(scanner.LEGEND); err != nil {
		return
	}
	t.Next()

	l := new(types.Legend)
Loop:
	for t != nil {
		if token, err = t.Token(); err != nil {
			return
		}

		switch token.Type {
		case scanner.IMAGECOLOR:
			if l.ImageColor, err = t.Color(); err != nil {
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

	legend = l
	return
}
