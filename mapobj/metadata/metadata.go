package metadata

import (
	"encoding/json"
	"fmt"
	"github.com/geo-data/mapfile/tokens"
)

type Metadata struct {
	kvmap map[string]string
}

func (m Metadata) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.kvmap)
}

func New(tokens *tokens.Tokens) (m *Metadata, err error) {
	token := tokens.Value()
	if token != "METADATA" {
		err = fmt.Errorf("expected token METADATA, got: %s", token)
		return
	}
	tokens.Next()

	m = &Metadata{
		kvmap: make(map[string]string),
	}

	for tokens != nil {
		key := tokens.Value()
		if key == "END" {
			break
		}

		value := tokens.Next().Value()
		if value == "END" {
			break
		}
		m.kvmap[key] = value

		tokens = tokens.Next()
	}

	return
}
