package tokens

import (
	"fmt"
	"strconv"
	"strings"
)

type Attribute string

func (a Attribute) String() string {
	attr := string(a)
	if len(attr) > 0 {
		return fmt.Sprintf("[%s]", string(attr))
	}
	return ""
}

func (a Attribute) QuotedString() string {
	attr := string(a)
	if len(attr) > 0 {
		return fmt.Sprintf(`"%s"`, string(attr))
	}
	return ""
}

type String string

// String implements fmt.Stringer()
func (t String) String() string {
	s := string(t)
	if len(s) > 0 {
		return fmt.Sprintf(`"%s"`, strings.Replace(s, `"`, `\"`, -1))
	}
	return s
}

type Keyword string

// String implements fmt.Stringer()
func (k Keyword) String() string {
	return string(k)
}

type Uint8 uint8

// String implements fmt.Stringer()
func (t Uint8) String() string {
	return fmt.Sprintf("%d", uint8(t))
}

type Integer int64

// String implements fmt.Stringer()
func (t Integer) String() string {
	return strconv.FormatInt(int64(t), 10)
}

type Uint32 uint32

// String implements fmt.Stringer()
func (t Uint32) String() string {
	return fmt.Sprintf("%d", uint32(t))
}

type Float64 float64

// String implements fmt.Stringer()
func (t Float64) String() string {
	return strconv.FormatFloat(float64(t), 'f', -1, 64)
}

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

	return v
}

type Type int

const (
	ATTRIBUTE = 1 << iota
	STRING
	FLOAT64
	INTEGER
	KEYWORD
)

func (t Type) Is(t2 Type) bool {
	return t&t2 != 0
}

func (t Type) String() string {
	types := []string{}
	if t.Is(ATTRIBUTE) {
		types = append(types, "Attribute")
	}
	if t.Is(STRING) {
		types = append(types, "String")
	}
	if t.Is(FLOAT64) {
		types = append(types, "Double")
	}
	if t.Is(INTEGER) {
		types = append(types, "Integer")
	}
	if t.Is(KEYWORD) {
		types = append(types, "Keyword")
	}

	return strings.Join(types, ", ")
}

func (t *Tokens) Decode(types Type) (fmt.Stringer, error) {
	if types.Is(INTEGER) {
		if v, err := t.Integer(); err == nil {
			return v, nil
		}
	}
	if types.Is(FLOAT64) {
		if v, err := t.Float64(); err == nil {
			return v, nil
		}
	}
	if types.Is(STRING) {
		if v, err := t.String(); err == nil {
			return v, nil
		}
	}
	if types.Is(KEYWORD) {
		if v, err := t.Keyword(); err == nil {
			return v, nil
		}
	}
	if types.Is(ATTRIBUTE) {
		if v, err := t.Attribute(); err == nil {
			return v, nil
		}
	}

	return nil, fmt.Errorf("decode failed, expected one of %s: %s", types, t.Value())
}

func (t *Tokens) Attribute() (attr Attribute, err error) {
	v := t.Value()
	if v[0] == '"' && v[len(v)-1] == '"' {
		v = v[1 : len(v)-1]
	}

	attr = Attribute(v)
	return
}

func (t *Tokens) Keyword() (kwd Keyword, err error) {
	v := t.Value()
	kwd = Keyword(v)
	return
}

func (t *Tokens) String() (s String, err error) {
	v := t.Value()
	if v[0] != '"' && v[len(v)-1] != '"' {
		err = fmt.Errorf("not a map string: %s", v)
		return
	}

	v = strings.Replace(v[1:len(v)-1], `\"`, `"`, -1)
	s = String(v)
	return
}

func (t *Tokens) Uint8() (i Uint8, err error) {
	var j uint64
	if j, err = strconv.ParseUint(t.Value(), 10, 8); err != nil {
		return
	} else {
		i = Uint8(j)
	}

	return
}

func (t *Tokens) Uint32() (i Uint32, err error) {
	var j uint64
	if j, err = strconv.ParseUint(t.Value(), 10, 32); err != nil {
		return
	} else {
		i = Uint32(j)
	}

	return
}

func (t *Tokens) Integer() (i Integer, err error) {
	var j int64
	if j, err = strconv.ParseInt(t.Value(), 10, 64); err != nil {
		return
	} else {
		i = Integer(j)
	}

	return
}

func (t *Tokens) Float64() (f Float64, err error) {
	var tf float64
	if tf, err = strconv.ParseFloat(t.Value(), 64); err != nil {
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
