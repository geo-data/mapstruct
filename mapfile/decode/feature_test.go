package decode_test

import (
	"errors"
	"github.com/geo-data/mapfile/mapfile/decode"
	"github.com/geo-data/mapfile/types"
	"reflect"
	"testing"
)

var featureTests = []struct {
	input    string
	expected *types.Feature // expected result
}{
	{`
FEATURE
END`, &types.Feature{}},
	{`
FEATURE
  WKT "test"
END`, &types.Feature{
		Wkt: types.String("test"),
	}},
	{`
FEATURE
  ITEMS "foo;bar"
END`, &types.Feature{
		Items: types.String("foo;bar"),
	}},
	{`
FEATURE
  TEXT "foo"
END`, &types.Feature{
		Text: types.String("foo"),
	}},
	{`
FEATURE
  POINTS END
END`, &types.Feature{}},
}

var featureErrorTests = []struct {
	input    string
	expected error // expected result
}{
	{"FEATURE", decode.EndOfTokens},
	{`
FOOBAR
END`, errors.New(`expected token FEATURE, got: "FOOBAR"`)},
	{`
FEATURE
  FOO BAR
END`, errors.New(`unhandled mapfile token: "FOO"`)},
}

func TestDecodeFeature(t *testing.T) {
	for _, tt := range featureTests {
		dec, err := decode.DecodeString(tt.input)
		if err != nil {
			t.Error("For decoding:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}
		actual, err := dec.Feature()
		if err != nil {
			t.Error("For:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}

		if !reflect.DeepEqual(tt.expected, actual) {
			t.Error("For:", tt.input, ", expected:", tt.expected, ", got:", actual)
		}
	}
}

func TestDecodeFeatureError(t *testing.T) {
	for _, tt := range featureErrorTests {
		dec, err := decode.DecodeString(tt.input)
		if err != nil {
			t.Error("For decoding:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}
		actual, err := dec.Feature()
		if actual != nil {
			t.Error("For:", tt.input, ", expected feature:", nil, ", got:", actual)
			continue
		}

		if err.Error() != tt.expected.Error() {
			t.Error("For:", tt.input, ", expected error:", tt.expected, ", got:", err)
		}
	}
}
