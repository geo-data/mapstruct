package decode_test

import (
	"errors"
	"github.com/geo-data/mapstruct/mapfile/decode"
	"github.com/geo-data/mapstruct/types"
	"reflect"
	"testing"
)

var webTests = []struct {
	input    string
	expected *types.Web // expected result
}{
	{`
WEB
END`, &types.Web{}},
	{`
WEB
  METADATA
  END
END`, &types.Web{
		Metadata: types.NewMetadata(),
	}},
}

var webErrorTests = []struct {
	input    string
	expected error // expected result
}{
	{"WEB", decode.EndOfTokens},
	{`
FOOBAR
END`, errors.New(`expected token WEB, got: FOOBAR`)},
	{`
WEB
  FOO BAR
END`, errors.New(`unhandled mapfile token: FOO`)},
}

func TestDecodeWeb(t *testing.T) {
	for _, tt := range webTests {
		dec := decode.DecodeString(tt.input)
		actual, err := dec.Web()
		if err != nil {
			t.Error("For:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}

		if !reflect.DeepEqual(tt.expected, actual) {
			t.Error("For:", tt.input, ", expected:", tt.expected, ", got:", actual)
		}
	}
}

func TestDecodeWebError(t *testing.T) {
	for _, tt := range webErrorTests {
		dec := decode.DecodeString(tt.input)
		actual, err := dec.Web()
		if actual != nil {
			t.Error("For:", tt.input, ", expected web:", nil, ", got:", actual)
			continue
		}

		if err.Error() != tt.expected.Error() {
			t.Error("For:", tt.input, ", expected error:", tt.expected, ", got:", err)
		}
	}
}
