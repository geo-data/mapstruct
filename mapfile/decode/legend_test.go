package decode_test

import (
	"errors"
	"github.com/geo-data/mapstruct/mapfile/decode"
	"github.com/geo-data/mapstruct/types"
	"reflect"
	"testing"
)

var legendTests = []struct {
	input    string
	expected *types.Legend // expected result
}{
	{`
LEGEND
END`, &types.Legend{}},
	{`
LEGEND
  IMAGECOLOR 123 210 255
END`, &types.Legend{
		ImageColor: types.NewColor(123, 210, 255, 255),
	}},
}

var legendErrorTests = []struct {
	input    string
	expected error // expected result
}{
	{"LEGEND", decode.EndOfTokens},
	{`
FOOBAR
END`, errors.New(`expected token LEGEND, got: FOOBAR`)},
	{`
LEGEND
  FOO BAR
END`, errors.New(`unhandled mapfile token: FOO`)},
}

func TestDecodeLegend(t *testing.T) {
	for _, tt := range legendTests {
		dec := decode.DecodeString(tt.input)
		actual, err := dec.Legend()
		if err != nil {
			t.Error("For:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}

		if !reflect.DeepEqual(tt.expected, actual) {
			t.Error("For:", tt.input, ", expected:", tt.expected, ", got:", actual)
		}
	}
}

func TestDecodeLegendError(t *testing.T) {
	for _, tt := range legendErrorTests {
		dec := decode.DecodeString(tt.input)
		actual, err := dec.Legend()
		if actual != nil {
			t.Error("For:", tt.input, ", expected legend:", nil, ", got:", actual)
			continue
		}

		if err.Error() != tt.expected.Error() {
			t.Error("For:", tt.input, ", expected error:", tt.expected, ", got:", err)
		}
	}
}
