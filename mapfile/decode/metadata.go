package decode

import (
	"fmt"
	"github.com/geo-data/mapfile/types"
)

func (t *Decoder) Metadata() (metadata types.Metadata, err error) {
	token := t.Value()
	if token != "METADATA" {
		err = fmt.Errorf("expected token METADATA, got: %s", token)
		return
	}
	t.Next()

	m := types.NewMetadata()
Loop:
	for t != nil {
		var key, value types.String

		switch t.Value() {
		case "END":
			break Loop
		case "":
			if t.AtEnd() {
				err = EndOfTokens
				return
			}
			fallthrough
		default:
			if key, err = t.String(); err != nil {
				return
			}
		}

		switch t.Next().Value() {
		case "END":
			err = fmt.Errorf("expected value after key: %s", key)
			return
		case "":
			if t.AtEnd() {
				err = EndOfTokens
				return
			}
			fallthrough
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
