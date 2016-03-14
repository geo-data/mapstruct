package metadata

import (
	"encoding/json"
	"fmt"
	"github.com/geo-data/mapfile/encoding"
	"github.com/geo-data/mapfile/tokens"
)

type Metadata struct {
	kvmap map[tokens.String]tokens.String
}

func (m Metadata) MarshalJSON() ([]byte, error) {
	tmap := make(map[string]string)
	for k, v := range m.kvmap {
		tmap[k.String()] = v.String()
	}
	return json.Marshal(tmap)
}

func New(toks *tokens.Tokens) (m *Metadata, err error) {
	token := toks.Value()
	if token != "METADATA" {
		err = fmt.Errorf("expected token METADATA, got: %s", token)
		return
	}
	toks.Next()

	m = &Metadata{
		kvmap: make(map[tokens.String]tokens.String),
	}

	for toks != nil {
		var key, value tokens.String
		if toks.Value() == "END" {
			break
		}
		if key, err = toks.String(); err != nil {
			return
		}

		if toks.Next().Value() == "END" {
			break
		}
		if value, err = toks.String(); err != nil {
			return
		}

		m.kvmap[key] = value

		toks = toks.Next()
	}

	return
}

func (p *Metadata) Encode(enc *encoding.MapfileEncoder) (err error) {
	if err = enc.TokenStart("METADATA"); err != nil {
		return
	}

	for k, v := range p.kvmap {
		if err = enc.EncodeStrings(k, v); err != nil {
			return
		}
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
