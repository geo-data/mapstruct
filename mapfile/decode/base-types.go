package decode

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/geo-data/mapstruct/mapfile/decode/scanner"
	"github.com/geo-data/mapstruct/types"
)

func (t *Decoder) Decode(kinds Type) (types.Union, error) {
	tok, err := t.Token()
	if err != nil {
		return nil, err
	}

	if kinds.Is(Integer) && tok.Type == scanner.MS_NUMBER {
		if v, err := t.Integer(); err == nil {
			return types.Union(v), nil
		}
	}
	if kinds.Is(Double) && tok.Type == scanner.MS_NUMBER {
		if v, err := t.Double(); err == nil {
			return types.Union(v), nil
		}
	}
	if kinds.Is(Listex) && tok.Type == scanner.MS_LISTEX {
		if v, err := t.Listex(); err == nil {
			return types.Union(v), nil
		}
	}
	if kinds.Is(Regex) && tok.Type == scanner.MS_REGEX {
		if v, err := t.Regex(); err == nil {
			return types.Union(v), nil
		}
	}
	if kinds.Is(Expression) && tok.Type == scanner.MS_EXPRESSION {
		if v, err := t.Expression(); err == nil {
			return types.Union(v), nil
		}
	}
	if kinds.Is(Keyword) && tok.Type == scanner.MS_BARE_STRING {
		if v, err := t.Keyword(); err == nil {
			return types.Union(v), nil
		}
	}
	if kinds.Is(Attribute) && (tok.Type == scanner.MS_ATTRIBUTE) {
		if v, err := t.Attribute(); err == nil {
			return types.Union(v), nil
		}
	}
	if kinds.Is(String) {
		if v, err := t.String(); err == nil {
			return types.Union(v), nil
		}
	}

	return nil, fmt.Errorf("decode failed for %q: expected one of %s", t.Value(), kinds)
}

func (t *Decoder) Attribute() (attr types.Attribute, err error) {
	var tok *scanner.Token
	if tok, err = t.Token(); err != nil {
		return
	}

	attr = types.Attribute(tok.Value)
	return
}

func (t *Decoder) Keyword() (kwd types.Keyword, err error) {
	var tok *scanner.Token
	if tok, err = t.ExpectedToken(scanner.MS_BARE_STRING); err != nil {
		return
	}

	kwd = types.Keyword(tok.Value)
	return
}

func (t *Decoder) String() (s types.String, err error) {
	var tok *scanner.Token
	if tok, err = t.Token(); err != nil {
		return
	}

	v := strings.Replace(tok.Value, `\"`, `"`, -1)
	s = types.String(v)
	return
}

func (t *Decoder) Regex() (attr types.Regex, err error) {
	var tok *scanner.Token
	if tok, err = t.ExpectedToken(scanner.MS_REGEX); err != nil {
		return
	}

	attr = types.Regex(tok.Value)
	return
}

func (t *Decoder) Listex() (attr types.Listex, err error) {
	var tok *scanner.Token
	if tok, err = t.ExpectedToken(scanner.MS_LISTEX); err != nil {
		return
	}

	attr = types.Listex(tok.Value)
	return
}

func (t *Decoder) Expression() (attr types.Expression, err error) {
	var tok *scanner.Token
	if tok, err = t.ExpectedToken(scanner.MS_EXPRESSION); err != nil {
		return
	}

	attr = types.Expression(tok.Value)
	return
}

func (t *Decoder) Uint8() (i types.Uint8, err error) {
	var (
		tok *scanner.Token
		j   uint64
	)

	if tok, err = t.Token(); err != nil {
		return
	}
	if tok.Type != scanner.MS_NUMBER {
		err = fmt.Errorf("token is not a number: %s", tok)
		return
	}

	if j, err = strconv.ParseUint(tok.Value, 10, 8); err != nil {
		return
	} else {
		i = types.Uint8(j)
	}

	return
}

func (t *Decoder) Uint32() (i types.Uint32, err error) {
	var (
		tok *scanner.Token
		j   uint64
	)

	if tok, err = t.Token(); err != nil {
		return
	}
	if tok.Type != scanner.MS_NUMBER {
		err = fmt.Errorf("token is not a number: %s", tok)
		return
	}

	if j, err = strconv.ParseUint(t.Value(), 10, 32); err != nil {
		return
	} else {
		i = types.Uint32(j)
	}

	return
}

func (t *Decoder) Integer() (i types.Integer, err error) {
	var (
		tok *scanner.Token
		j   int64
	)

	if tok, err = t.Token(); err != nil {
		return
	}
	if tok.Type != scanner.MS_NUMBER {
		err = fmt.Errorf("token is not a number: %s", tok)
		return
	}

	if j, err = strconv.ParseInt(t.Value(), 10, 64); err != nil {
		return
	} else {
		i = types.Integer(j)
	}

	return
}

func (t *Decoder) Double() (f types.Double, err error) {
	var (
		tok *scanner.Token
		tf  float64
	)

	if tok, err = t.Token(); err != nil {
		return
	}
	if tok.Type != scanner.MS_NUMBER {
		err = fmt.Errorf("token is not a number: %s", tok)
		return
	}

	if tf, err = strconv.ParseFloat(t.Value(), 64); err != nil {
		err = fmt.Errorf("invalid syntax for double: %s", t.Value())
		return
	}

	f = types.Double(tf)
	return
}
