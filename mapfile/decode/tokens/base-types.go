package tokens

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/geo-data/mapfile/types"
)

func (t *Tokens) Decode(kinds Type) (types.Union, error) {
	if kinds.Is(Integer) {
		if v, err := t.Integer(); err == nil {
			return types.Union(v), nil
		}
	}
	if kinds.Is(Double) {
		if v, err := t.Double(); err == nil {
			return types.Union(v), nil
		}
	}
	if kinds.Is(String) {
		if v, err := t.String(); err == nil {
			return types.Union(v), nil
		}
	}
	if kinds.Is(Keyword) {
		if v, err := t.Keyword(); err == nil {
			return types.Union(v), nil
		}
	}
	if kinds.Is(Attribute) {
		if v, err := t.Attribute(); err == nil {
			return types.Union(v), nil
		}
	}

	return nil, fmt.Errorf("decode failed, expected one of %s: %s", kinds, t.Value())
}

func (t *Tokens) Attribute() (attr types.Attribute, err error) {
	v := t.Value()
	if v[0] == '"' && v[len(v)-1] == '"' {
		v = v[1 : len(v)-1]
	}

	attr = types.Attribute(v)
	return
}

func (t *Tokens) Keyword() (kwd types.Keyword, err error) {
	v := t.Value()
	kwd = types.Keyword(v)
	return
}

func (t *Tokens) String() (s types.String, err error) {
	v := t.Value()
	if v[0] != '"' && v[len(v)-1] != '"' {
		err = fmt.Errorf("not a map string: %s", v)
		return
	}

	v = strings.Replace(v[1:len(v)-1], `\"`, `"`, -1)
	s = types.String(v)
	return
}

func (t *Tokens) Uint8() (i types.Uint8, err error) {
	var j uint64
	if j, err = strconv.ParseUint(t.Value(), 10, 8); err != nil {
		return
	} else {
		i = types.Uint8(j)
	}

	return
}

func (t *Tokens) Uint32() (i types.Uint32, err error) {
	var j uint64
	if j, err = strconv.ParseUint(t.Value(), 10, 32); err != nil {
		return
	} else {
		i = types.Uint32(j)
	}

	return
}

func (t *Tokens) Integer() (i types.Integer, err error) {
	var j int64
	if j, err = strconv.ParseInt(t.Value(), 10, 64); err != nil {
		return
	} else {
		i = types.Integer(j)
	}

	return
}

func (t *Tokens) Double() (f types.Double, err error) {
	var tf float64
	if tf, err = strconv.ParseFloat(t.Value(), 64); err != nil {
		return
	}

	f = types.Double(tf)
	return
}
