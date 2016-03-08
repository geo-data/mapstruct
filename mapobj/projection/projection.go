package projection

import (
	"encoding/json"
	"fmt"
	"github.com/geo-data/mapfile/tokens"
)

type Projection struct {
	params []string
}

func (m Projection) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.params)
}

func (p *Projection) FromTokens(tokens *tokens.Tokens) error {
	token := tokens.Value()
	if token != "PROJECTION" {
		return fmt.Errorf("expected token PROJECTION, got: %s", token)
	}
	tokens.Next()

	for tokens != nil {
		token := tokens.Value()
		switch token {
		case "END":
			return nil
		default:
			p.params = append(p.params, token)
		}

		tokens = tokens.Next()
	}

	return nil
}
