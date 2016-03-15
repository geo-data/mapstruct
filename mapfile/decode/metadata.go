package decode

import (
	"fmt"
	"github.com/geo-data/mapfile/types"
	"github.com/geo-data/mapfile/types/metadata"
)

func (t *Decoder) Metadata() (m metadata.Metadata, err error) {
	token := t.Value()
	if token != "METADATA" {
		err = fmt.Errorf("expected token METADATA, got: %s", token)
		return
	}
	t.Next()

	m = metadata.New()

	for t != nil {
		var key, value types.String
		if t.Value() == "END" {
			break
		}
		if key, err = t.String(); err != nil {
			return
		}

		if t.Next().Value() == "END" {
			break
		}
		if value, err = t.String(); err != nil {
			return
		}

		m[string(key)] = string(value)

		t = t.Next()
	}

	return
}
