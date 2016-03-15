package projection

import (
	"encoding/json"
	"fmt"
	"github.com/geo-data/mapfile/mapfile/encode"
	"github.com/geo-data/mapfile/mapfile/decode/tokens"
	"github.com/geo-data/mapfile/types"
)

type Projection struct {
	params []types.String
}

func (m *Projection) MarshalJSON() ([]byte, error) {
	return json.Marshal(m.params)
}

func New(toks *tokens.Tokens) (p *Projection, err error) {
	token := toks.Value()
	if token != "PROJECTION" {
		err = fmt.Errorf("expected token PROJECTION, got: %s", token)
		return
	}
	toks.Next()

	p = new(Projection)

	for toks != nil {
		token := toks.Value()
		switch token {
		case "END":
			return
		default:
			var s types.String
			if s, err = toks.String(); err != nil {
				return
			}
			p.params = append(p.params, s)
		}

		toks = toks.Next()
	}

	return
}

func (p *Projection) Encode(enc *encode.MapfileEncoder) (err error) {
	if err = enc.TokenStart("PROJECTION"); err != nil {
		return
	}

	for _, param := range p.params {
		if err = enc.EncodeString(param); err != nil {
			return
		}
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
