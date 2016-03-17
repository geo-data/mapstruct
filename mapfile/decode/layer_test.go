package decode_test

import (
	"errors"
	"github.com/geo-data/mapfile/mapfile/decode"
	"github.com/geo-data/mapfile/types"
	"reflect"
	"testing"
)

var layerTests = []struct {
	input    string
	expected *types.Layer // expected result
}{
	{`
LAYER
END`, &types.Layer{}},
	{`
LAYER
  NAME "Test"
END`, &types.Layer{
		Name: types.String("Test"),
	}},
	{`
LAYER
  EXTENT -0.5 50.977222 0.5 51.977222
END`, &types.Layer{
		Extent: types.NewExtent(-0.5, 50.977222, 0.5, 51.977222),
	}},
	{`
LAYER
  TYPE RASTER
END`, &types.Layer{
		Type: types.Keyword("RASTER"),
	}},
	{`
LAYER
  DEBUG OFF
END`, &types.Layer{
		Debug: types.Union(types.Keyword("OFF")),
	}},
	{`
LAYER
  DEBUG 3
END`, &types.Layer{
		Debug: types.Union(types.Integer(3)),
	}},
	{`
LAYER
  PROJECTION
  END
END`, &types.Layer{}},
	{`
LAYER
  DATA "raster.tif"
END`, &types.Layer{
		Data: types.String("raster.tif"),
	}},
	{`
LAYER
  PROCESSING "foo=bar"
END`, &types.Layer{
		Processing: types.String("foo=bar"),
	}},
	{`
LAYER
  STATUS ON
END`, &types.Layer{
		Status: types.Keyword("ON"),
	}},
	{`
LAYER
  METADATA
  END
END`, &types.Layer{
		Metadata: types.NewMetadata(),
	}},
	{`
LAYER
  CLASSITEM "foobar"
END`, &types.Layer{
		ClassItem: types.Attribute("foobar"),
	}},
	{`
LAYER
  CLASSITEM [foobar]
END`, &types.Layer{
		ClassItem: types.Attribute("foobar"),
	}},
	{`
LAYER
  LABELITEM "foobar"
END`, &types.Layer{
		LabelItem: types.Attribute("foobar"),
	}},
	{`
LAYER
  LABELITEM [foobar]
END`, &types.Layer{
		LabelItem: types.Attribute("foobar"),
	}},
	{`
LAYER
  CLASS
    NAME "one"
  END
  CLASS
    NAME "two"
  END
END`, &types.Layer{
		Classes: []*types.Class{
			&types.Class{
				Name: types.String("one"),
			},
			&types.Class{
				Name: types.String("two"),
			},
		},
	}},
	{`
LAYER
  FEATURE
    TEXT "one"
  END
  FEATURE
    TEXT "two"
  END
END`, &types.Layer{
		Features: []*types.Feature{
			&types.Feature{
				Text: types.String("one"),
			},
			&types.Feature{
				Text: types.String("two"),
			},
		},
	}},
}

var layerErrorTests = []struct {
	input    string
	expected error // expected result
}{
	{"LAYER", decode.EndOfTokens},
	{`
FOOBAR
END`, errors.New(`expected token LAYER, got: "FOOBAR"`)},
	{`
LAYER
  FOO BAR
END`, errors.New(`unhandled mapfile token: "FOO"`)},
}

func TestDecodeLayer(t *testing.T) {
	for _, tt := range layerTests {
		dec, err := decode.DecodeString(tt.input)
		if err != nil {
			t.Error("For decoding:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}
		actual, err := dec.Layer()
		if err != nil {
			t.Error("For:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}

		if !reflect.DeepEqual(tt.expected, actual) {
			t.Error("For", tt.input, "expected", tt.expected, "got", actual)
		}
	}
}

func TestDecodeLayerError(t *testing.T) {
	for _, tt := range layerErrorTests {
		dec, err := decode.DecodeString(tt.input)
		if err != nil {
			t.Error("For decoding:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}
		actual, err := dec.Layer()
		if actual != nil {
			t.Error("For:", tt.input, ", expected layer:", nil, ", got:", actual)
			continue
		}

		if err.Error() != tt.expected.Error() {
			t.Error("For:", tt.input, ", expected error:", tt.expected, ", got:", err)
		}
	}
}
