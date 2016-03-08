package projection

import (
	"encoding/json"
	"fmt"
	"github.com/geo-data/mapfile/tokens"
)

type Projection struct {
	params []string
}

func (m *Projection) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.params)
}

func New(tokens *tokens.Tokens) (p *Projection, err error) {
	token := tokens.Value()
	if token != "PROJECTION" {
		err = fmt.Errorf("expected token PROJECTION, got: %s", token)
		return
	}
	tokens.Next()

	p = new(Projection)

	for tokens != nil {
		token := tokens.Value()
		switch token {
		case "END":
			break
		default:
			p.params = append(p.params, token)
		}

		tokens = tokens.Next()
	}

	return
}
