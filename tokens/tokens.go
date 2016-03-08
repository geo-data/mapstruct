package tokens

import (
	"strconv"
	"strings"
)

type Tokens struct {
	tokens []string
	idx    uint
}

func (t *Tokens) Value() string {
	v := t.tokens[t.idx]
	for v == "" && t.idx < uint(len(t.tokens)) {
		t.idx++
		v = t.tokens[t.idx]
	}

	if len(v) > 0 && v[0] == '"' {
		v = strings.Replace(v[1:len(v)-1], `\"`, `"`, -1)
	}

	return v
}

func (t *Tokens) Uint8() (i uint8, err error) {
	var j uint64
	if j, err = strconv.ParseUint(t.Value(), 10, 8); err != nil {
		return
	} else {
		i = uint8(j)
	}

	return
}

func (t *Tokens) Uint32() (i uint32, err error) {
	var j uint64
	if j, err = strconv.ParseUint(t.Value(), 10, 32); err != nil {
		return
	} else {
		i = uint32(j)
	}

	return
}

func (t *Tokens) Float64() (float64, error) {
	return strconv.ParseFloat(t.Value(), 64)
}

func (t *Tokens) Next() *Tokens {
	if t.idx+1 < uint(len(t.tokens)) {
		t.idx++
		return t
	}

	return nil
}
