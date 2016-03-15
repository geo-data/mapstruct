package decode

import "github.com/geo-data/mapfile/mapfile/decode/tokenize"

type Decoder struct {
	tokens []string
	idx    uint
}

func (t *Decoder) Value() string {
	v := t.tokens[t.idx]
	for v == "" && t.idx < uint(len(t.tokens)) {
		t.idx++
		v = t.tokens[t.idx]
	}

	return v
}

func (t *Decoder) Next() *Decoder {
	if t.idx+1 < uint(len(t.tokens)) {
		t.idx++
		return t
	}

	return nil
}

func DecodeMapfile(mapfile string) (dec *Decoder, err error) {
	dec = &Decoder{}
	if dec.tokens, err = tokenize.TokenizeMapfile(mapfile); err != nil {
		return
	}

	return
}
