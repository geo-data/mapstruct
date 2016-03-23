package decode

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/geo-data/mapfile/mapfile/decode/scanner"
	"io"
)

var (
	EndOfTokens = errors.New("unexpected end of mapfile")
)

type Decoder struct {
	scanner scanner.Scanner
	current *scanner.Token
}

func (t *Decoder) Value() string {
	if tok, _ := t.Token(); tok != nil {
		return tok.Value
	}
	return ""
}

func (t *Decoder) Type() scanner.TokenType {
	if tok, _ := t.Token(); tok != nil {
		return tok.Type
	}
	return scanner.EOF
}

func (t *Decoder) Token() (token *scanner.Token, err error) {
	if t.current == nil {
		t.current = t.nextToken()
	}

	if t.current == nil || t.current.Type == scanner.EOF {
		err = EndOfTokens
	}

	token = t.current
	return
}

func (t *Decoder) ExpectedToken(expected scanner.TokenType) (token *scanner.Token, err error) {
	if token, err = t.Token(); err != nil {
		return
	}

	if token.Type != expected {
		err = fmt.Errorf("expected token %s, got: %s", expected, token)
		return
	}

	return
}

func (t *Decoder) nextToken() *scanner.Token {
	for token := t.scanner.Scan(); token != nil && token.Type != scanner.EOF; token = t.scanner.Scan() {
		switch token.Type {
		case scanner.WS:
			continue
		case scanner.MS_COMMENT:
			continue
		default:
			return token
		}
	}

	return &scanner.Token{
		scanner.EOF,
		"",
	}
}

func (t *Decoder) Next() *Decoder {
	if t.Type() == scanner.EOF {
		return t
	}

	if token := t.nextToken(); token != nil {
		t.current = token
	}

	return t
}

func NewDecoder(scanner scanner.Scanner) *Decoder {
	return &Decoder{
		scanner: scanner,
	}
}

func DecodeMapfile(mapfile io.Reader) *Decoder {
	lexer := scanner.NewScanner(mapfile)
	return NewDecoder(lexer)
}

func DecodeString(mapfile string) *Decoder {
	buf := bytes.NewBufferString(mapfile)
	lexer := scanner.NewScanner(buf)
	return NewDecoder(lexer)
}
