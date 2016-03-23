package decode_test

import (
	"errors"
	"github.com/geo-data/mapfile/mapfile/decode"
	"github.com/geo-data/mapfile/types"
	"reflect"
	"testing"
)

var styleTests = []struct {
	input    string
	expected *types.Style // expected result
}{
	{`
STYLE
END`, &types.Style{}},
	{`
STYLE
  COLOR 1 2 3
END`, &types.Style{
		Color: types.NewColor(1, 2, 3, 255),
	}},
	{`
STYLE
  OUTLINECOLOR 1 2 3
END`, &types.Style{
		OutlineColor: types.NewColor(1, 2, 3, 255),
	}},
	{`
STYLE
  SIZE 20.5
END`, &types.Style{
		Size: types.Double(20.5),
	}},
	{`
STYLE
  SIZE [size]
END`, &types.Style{
		Size: types.Attribute("size"),
	}},
	{`
STYLE
  SYMBOL 2
END`, &types.Style{
		Symbol: types.Integer(2),
	}},
	{`
STYLE
  SYMBOL "name"
END`, &types.Style{
		Symbol: types.String("name"),
	}},
	{`
STYLE
  SYMBOL [symbol]
END`, &types.Style{
		Symbol: types.Attribute("symbol"),
	}},
	{`
STYLE
  WIDTH 2.5
END`, &types.Style{
		Width: types.Double(2.5),
	}},
	{`
STYLE
  WIDTH [width]
END`, &types.Style{
		Width: types.Attribute("width"),
	}},
	{`
STYLE
  GEOMTRANSFORM bbox
END`, &types.Style{
		GeomTransform: types.String("bbox"),
	}},
}

var styleErrorTests = []struct {
	input    string
	expected error // expected result
}{
	{"STYLE", decode.EndOfTokens},
	{`
FOOBAR
END`, errors.New(`expected token STYLE, got: FOOBAR`)},
	{`
STYLE
  FOO BAR
END`, errors.New(`unhandled mapfile token: FOO`)},
}

func TestDecodeStyle(t *testing.T) {
	for _, tt := range styleTests {
		dec := decode.DecodeString(tt.input)
		actual, err := dec.Style()
		if err != nil {
			t.Error("For:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}

		if !reflect.DeepEqual(tt.expected, actual) {
			t.Error("For:", tt.input, ", expected:", tt.expected, ", got:", actual)
		}
	}
}

func TestDecodeStyleError(t *testing.T) {
	for _, tt := range styleErrorTests {
		dec := decode.DecodeString(tt.input)
		actual, err := dec.Style()
		if actual != nil {
			t.Error("For:", tt.input, ", expected style:", nil, ", got:", actual)
			continue
		}

		if err.Error() != tt.expected.Error() {
			t.Error("For:", tt.input, ", expected error:", tt.expected, ", got:", err)
		}
	}
}
