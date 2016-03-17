package decode_test

import (
	"errors"
	"github.com/geo-data/mapfile/mapfile/decode"
	"github.com/geo-data/mapfile/types"
	"reflect"
	"testing"
)

var mapTests = []struct {
	input    string
	expected *types.Map // expected result
}{
	{`
MAP
END`, &types.Map{}},
	{`
MAP
  NAME "Testing"
END`, &types.Map{
		Name: types.String("Testing"),
	}},
	{`
MAP
  EXTENT -0.5 50.977222 0.5 51.977222
END`, &types.Map{
		Extent: types.NewExtent(-0.5, 50.977222, 0.5, 51.977222),
	}},
	{`
MAP
  IMAGETYPE PNG
END`, &types.Map{
		ImageType: types.String("PNG"),
	}},
	{`
MAP
  IMAGECOLOR 255 255 255
END`, &types.Map{
		ImageColor: types.NewColor(255, 255, 255, 255),
	}},
	{`
MAP
  STATUS ON
END`, &types.Map{
		Status: types.Keyword("ON"),
	}},
	{`
MAP
  SIZE 200 200
END`, &types.Map{
		Size: types.NewSize(200, 200),
	}},
	{`
MAP
  FONTSET "fonts.txt"
END`, &types.Map{
		Fontset: types.String("fonts.txt"),
	}},
	{`
MAP
  SYMBOLSET "symbols.txt"
END`, &types.Map{
		Symbolset: types.String("symbols.txt"),
	}},
	{`
MAP
  SCALEBAR
  END
END`, &types.Map{
		Scalebar: &types.Scalebar{},
	}},
	{`
MAP
  LEGEND
  END
END`, &types.Map{
		Legend: &types.Legend{},
	}},
	{`
MAP
  WEB
  END
END`, &types.Map{
		Web: &types.Web{},
	}},
	{`
MAP
  PROJECTION
  END
END`, &types.Map{}},
	{`
MAP
  LAYER
    NAME "first"
  END
  LAYER
    NAME "second"
  END
END`, &types.Map{
		Layers: []*types.Layer{
			&types.Layer{Name: types.String("first")},
			&types.Layer{Name: types.String("second")},
		},
	}},
}

var mapErrorTests = []struct {
	input    string
	expected error // expected result
}{
	{"MAP", decode.EndOfTokens},
	{`
FOOBAR
END`, errors.New(`expected token MAP, got: "FOOBAR"`)},
	{`
MAP
  FOO BAR
END`, errors.New(`unhandled mapfile token: "FOO"`)},
}

func TestDecodeMap(t *testing.T) {
	for _, tt := range mapTests {
		dec, err := decode.DecodeString(tt.input)
		if err != nil {
			t.Error("For decoding:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}
		actual, err := dec.Map()
		if err != nil {
			t.Error("For:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}

		if !reflect.DeepEqual(tt.expected, actual) {
			t.Error("For:", tt.input, ", expected:", tt.expected, ", got:", actual)
		}
	}
}

func TestDecodeMapError(t *testing.T) {
	for _, tt := range mapErrorTests {
		dec, err := decode.DecodeString(tt.input)
		if err != nil {
			t.Error("For decoding:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}
		actual, err := dec.Map()
		if actual != nil {
			t.Error("For:", tt.input, ", expected map:", nil, ", got:", actual)
			continue
		}

		if err.Error() != tt.expected.Error() {
			t.Error("For:", tt.input, ", expected error:", tt.expected, ", got:", err)
		}
	}
}
