package tokens

import (
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

func (t *Tokens) Next() *Tokens {
	if t.idx+1 < uint(len(t.tokens)) {
		t.idx++
		return t
	}

	return nil
}
