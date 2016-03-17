package decode_test

import (
	"github.com/geo-data/mapfile/mapfile/decode"
	"testing"
)

func TestDecoder(t *testing.T) {
	tokens := []string{
		"one",
		"two",
		"three",
	}

	dec := decode.NewDecoder(tokens)
	if dec == nil {
		t.Fatal("For", tokens, ", expected decoder, got:", nil)
	}

	for i, token := range tokens {
		if dec.AtEnd() {
			t.Error("For call count", i, ", expected AtEnd():", false, ", got:", true)
		}

		v := dec.Value()
		if v != token {
			t.Error("For call count", i, ", expected Value():", token, ", got:", v)
		}

		d := dec.Next()
		if d == nil {
			t.Error("For call count", i, ", expected Next() == decoder, got:", d)
		}
	}

	if !dec.AtEnd() {
		t.Error("For call count", len(tokens), ", expected AtEnd():", true, ", got:", false)
	}
}
