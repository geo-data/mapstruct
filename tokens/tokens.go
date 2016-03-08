package tokens

import (
	"fmt"
	"strconv"
	"strings"
)

type String string

// String implements fmt.Stringer()
func (t String) String() string {
	return string(t)
}

type Uint8 uint8

// String implements fmt.Stringer()
func (t Uint8) String() string {
	return fmt.Sprintf("%d", uint8(t))
}

type Uint32 uint32

// String implements fmt.Stringer()
func (t Uint32) String() string {
	return fmt.Sprintf("%d", uint32(t))
}

type Float64 float64

// String implements fmt.Stringer()
func (t Float64) String() string {
	return fmt.Sprintf("%f", float64(t))
}

type Tokens struct {
	tokens []string
	idx    uint
}

func (t *Tokens) Value() String {
	v := t.tokens[t.idx]
	for v == "" && t.idx < uint(len(t.tokens)) {
		t.idx++
		v = t.tokens[t.idx]
	}

	if len(v) > 0 && v[0] == '"' {
		v = strings.Replace(v[1:len(v)-1], `\"`, `"`, -1)
	}

	return String(v)
}

func (t *Tokens) Uint8() (i Uint8, err error) {
	var j uint64
	if j, err = strconv.ParseUint(t.Value().String(), 10, 8); err != nil {
		return
	} else {
		i = Uint8(j)
	}

	return
}

func (t *Tokens) Uint32() (i Uint32, err error) {
	var j uint64
	if j, err = strconv.ParseUint(t.Value().String(), 10, 32); err != nil {
		return
	} else {
		i = Uint32(j)
	}

	return
}

func (t *Tokens) Float64() (f Float64, err error) {
	var tf float64
	if tf, err = strconv.ParseFloat(t.Value().String(), 64); err != nil {
		return
	}

	f = Float64(tf)
	return
}

func (t *Tokens) Next() *Tokens {
	if t.idx+1 < uint(len(t.tokens)) {
		t.idx++
		return t
	}

	return nil
}
