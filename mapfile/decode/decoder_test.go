package decode_test

import (
	"github.com/geo-data/mapstruct/mapfile/decode"
	"github.com/geo-data/mapstruct/mapfile/decode/scanner"
	"reflect"
	"testing"
)

func TestDecoder(t *testing.T) {

	mapfile := `"one" "two" "three"`
	dec := decode.DecodeString(mapfile)
	if dec == nil {
		t.Fatal("For", mapfile, ", expected decoder, got:", nil)
	}

	expected := []*scanner.Token{
		&scanner.Token{
			scanner.MS_STRING, "one",
		},
		&scanner.Token{
			scanner.MS_STRING, "two",
		},
		&scanner.Token{
			scanner.MS_STRING, "three",
		},
	}

	for i, token := range expected {
		actual, err := dec.Token()
		if err != nil {
			t.Error("For call count", i, ", expected error:", nil, ", got:", err)
			continue
		}
		if !reflect.DeepEqual(actual, token) {
			t.Error("For call count", i, ", expected token:", token, ", got:", actual)
		}

		typ := dec.Type()
		if typ != token.Type {
			t.Error("For call count", i, ", expected Type():", token.Type, ", got:", actual.Type)
		}

		v := dec.Value()
		if v != token.Value {
			t.Error("For call count", i, ", expected Value():", token.Value, ", got:", v)
		}

		d := dec.Next()
		if d == nil {
			t.Error("For call count", i, ", expected Next() == decoder, got:", d)
		}
	}
}
