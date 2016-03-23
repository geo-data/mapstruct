package decode

import (
	"fmt"
	"github.com/geo-data/mapfile/mapfile/decode/scanner"
	"github.com/geo-data/mapfile/types"
)

func (t *Decoder) Metadata() (metadata types.Metadata, err error) {
	var token *scanner.Token
	if token, err = t.ExpectedToken(scanner.METADATA); err != nil {
		return
	}
	t.Next()

	m := types.NewMetadata()
Loop:
	for t != nil {
		var key, value types.String
		if token, err = t.Token(); err != nil {
			return
		}

		switch token.Type {
		case scanner.END:
			break Loop
		default:
			if key, err = t.String(); err != nil {
				return
			}
		}

		if token, err = t.Next().Token(); err != nil {
			return
		}
		switch token.Type {
		case scanner.END:
			err = fmt.Errorf("expected value after key: %s", key)
			return
		default:
			if value, err = t.String(); err != nil {
				return
			}
		}

		m[string(key)] = string(value)

		t = t.Next()
	}

	metadata = m
	return
}
