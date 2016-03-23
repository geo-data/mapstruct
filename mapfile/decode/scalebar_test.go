package decode_test

import (
	"errors"
	"github.com/geo-data/mapstruct/mapfile/decode"
	"github.com/geo-data/mapstruct/types"
	"reflect"
	"testing"
)

var scalebarTests = []struct {
	input    string
	expected *types.Scalebar // expected result
}{
	{`
SCALEBAR
END`, &types.Scalebar{}},
	{`
SCALEBAR
  STATUS EMBED
END`, &types.Scalebar{
		Status: types.Keyword("EMBED"),
	}},
	{`
SCALEBAR
  POSTLABELCACHE true
END`, &types.Scalebar{
		PostLabelCache: types.Keyword("true"),
	}},
	{`
SCALEBAR
  STYLE 1
END`, &types.Scalebar{
		Style: types.Uint8(1),
	}},
	{`
SCALEBAR
  UNITS feet
END`, &types.Scalebar{
		Units: types.Keyword("feet"),
	}},
	{`
SCALEBAR
  SIZE 20 5
END`, &types.Scalebar{
		Size: types.NewSize(20, 5),
	}},
	{`
SCALEBAR
  POSITION UL
END`, &types.Scalebar{
		Position: types.Keyword("UL"),
	}},
	{`
SCALEBAR
  TRANSPARENT ON
END`, &types.Scalebar{
		Transparent: types.Keyword("ON"),
	}},
	{`
SCALEBAR
  IMAGECOLOR 1 2 3
END`, &types.Scalebar{
		ImageColor: types.NewColor(1, 2, 3, 255),
	}},
	{`
SCALEBAR
  COLOR 1 2 3
END`, &types.Scalebar{
		Color: types.NewColor(1, 2, 3, 255),
	}},
	{`
SCALEBAR
  BACKGROUNDCOLOR 1 2 3
END`, &types.Scalebar{
		BackgroundColor: types.NewColor(1, 2, 3, 255),
	}},
	{`
SCALEBAR
  LABEL
  END
END`, &types.Scalebar{
		Label: &types.Label{},
	}},
}

var scalebarErrorTests = []struct {
	input    string
	expected error // expected result
}{
	{"SCALEBAR", decode.EndOfTokens},
	{`
FOOBAR
END`, errors.New(`expected token SCALEBAR, got: FOOBAR`)},
	{`
SCALEBAR
  FOO BAR
END`, errors.New(`unhandled mapfile token: FOO`)},
}

func TestDecodeScalebar(t *testing.T) {
	for _, tt := range scalebarTests {
		dec := decode.DecodeString(tt.input)
		actual, err := dec.Scalebar()
		if err != nil {
			t.Error("For:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}

		if !reflect.DeepEqual(tt.expected, actual) {
			t.Error("For:", tt.input, ", expected:", tt.expected, ", got:", actual)
		}
	}
}

func TestDecodeScalebarError(t *testing.T) {
	for _, tt := range scalebarErrorTests {
		dec := decode.DecodeString(tt.input)
		actual, err := dec.Scalebar()
		if actual != nil {
			t.Error("For:", tt.input, ", expected scalebar:", nil, ", got:", actual)
			continue
		}

		if err.Error() != tt.expected.Error() {
			t.Error("For:", tt.input, ", expected error:", tt.expected, ", got:", err)
		}
	}
}
