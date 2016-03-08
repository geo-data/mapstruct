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

func (m *Metadata) FromTokens(tokens *tokens.Tokens) error {
	token := tokens.Value()
	if token != "METADATA" {
		return fmt.Errorf("expected token METADATA, got: %s", token)
	}
	tokens.Next()

	if m.kvmap == nil {
		m.kvmap = make(map[string]string)
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

	return nil
}
