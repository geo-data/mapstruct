package metadata

import (
	"encoding/json"
	"fmt"
	"github.com/geo-data/mapfile/encoding"
	"github.com/geo-data/mapfile/tokens"
)

type Metadata struct {
	kvmap map[fmt.Stringer]fmt.Stringer
}

func (m Metadata) MarshalJSON() ([]byte, error) {
	tmap := make(map[string]string)
	for k, v := range m.kvmap {
		tmap[k.String()] = v.String()
	}
	return json.Marshal(tmap)
}

func New(tokens *tokens.Tokens) (m *Metadata, err error) {
	token := tokens.Value()
	if token != "METADATA" {
		err = fmt.Errorf("expected token METADATA, got: %s", token)
		return
	}
	tokens.Next()

	m = &Metadata{
		kvmap: make(map[fmt.Stringer]fmt.Stringer),
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
