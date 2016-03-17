package decode

import (
	"errors"
	"github.com/geo-data/mapfile/mapfile/decode/tokenize"
	"os"
)

var (
	EndOfTokens = errors.New("decode failed: unexpected end of mapfile")
)

type Decoder struct {
	tokens []string
	idx    uint
}

func (t *Decoder) Value() string {
	if t.AtEnd() {
		return ""
	}

	v := t.tokens[t.idx]
	for v == "" && t.idx < uint(len(t.tokens)) {
		t.idx++
		v = t.tokens[t.idx]
	}

	return v
}

func (t *Decoder) AtEnd() bool {
	return t.idx >= uint(len(t.tokens))
}

func (t *Decoder) Next() *Decoder {
	if !t.AtEnd() {
		t.idx++
		return t
	}

	return nil
}

func NewDecoder(tokens []string) *Decoder {
	return &Decoder{
		tokens: tokens,
	}
}

func DecodeMapfile(mapfile string) (dec *Decoder, err error) {
	dec = &Decoder{}
	if dec.tokens, err = tokenize.TokenizeMapfile(mapfile); err != nil {
		return
	}

	return
}

func DecodeString(mapfile string) (dec *Decoder, err error) {
	var tmpfile *os.File
	if tmpfile, err = tempFile("", "example", ".map"); err != nil {
		return
	}
	filename := tmpfile.Name()
	defer os.Remove(filename) // clean up

	if _, err = tmpfile.WriteString(mapfile); err != nil {
		return
	}
	if err = tmpfile.Close(); err != nil {
		return
	}

	dec = &Decoder{}
	if dec.tokens, err = tokenize.TokenizeMapfile(filename); err != nil {
		return
	}

	return
}
