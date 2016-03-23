package decode_test

import (
	"errors"
	"github.com/geo-data/mapfile/mapfile/decode"
	"github.com/geo-data/mapfile/types"
	"reflect"
	"testing"
)

var labelTests = []struct {
	input    string
	expected *types.Label // expected result
}{
	{`
LABEL
END`, &types.Label{}},
	{`
LABEL
  TYPE TRUETYPE
END`, &types.Label{
		Type: types.Keyword("TRUETYPE"),
	}},
	{`
LABEL
  FONT "Vera"
END`, &types.Label{
		Font: types.String("Vera"),
	}},
	{`
LABEL
  COLOR 1 2 3
END`, &types.Label{
		Color: types.NewColor(1, 2, 3, 255),
	}},
	{`
LABEL
  POSITION UR
END`, &types.Label{
		Position: types.Keyword("UR"),
	}},
	{`
LABEL
  BUFFER 5
END`, &types.Label{
		Buffer: types.Uint32(5),
	}},
	{`
LABEL
  BUFFER 0
END`, &types.Label{}},
	{`
LABEL
  SIZE 2.5
END`, &types.Label{
		Size: types.Double(2.5),
	}},
	{`
LABEL
  SIZE MEDIUM
END`, &types.Label{
		Size: types.Keyword("MEDIUM"),
	}},
	{`
LABEL
  SIZE [foobar]
END`, &types.Label{
		Size: types.Attribute("foobar"),
	}},
	{`
LABEL
  STYLE
    SYMBOL 1
  END
  STYLE
    SYMBOL 2
  END
END`, &types.Label{
		Styles: []*types.Style{
			&types.Style{
				Symbol: types.Union(types.Integer(1)),
			},
			&types.Style{
				Symbol: types.Union(types.Integer(2)),
			},
		},
	}},
}

var labelErrorTests = []struct {
	input    string
	expected error // expected result
}{
	{"LABEL", decode.EndOfTokens},
	{`
FOOBAR
END`, errors.New(`expected token LABEL, got: FOOBAR`)},
	{`
LABEL
  FOO BAR
END`, errors.New(`unhandled mapfile token: FOO`)},
}

func TestDecodeLabel(t *testing.T) {
	for _, tt := range labelTests {
		dec := decode.DecodeString(tt.input)
		actual, err := dec.Label()
		if err != nil {
			t.Error("For:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}

		if !reflect.DeepEqual(tt.expected, actual) {
			t.Error("For:", tt.input, ", expected:", tt.expected, ", got:", actual)
		}
	}
}

func TestDecodeLabelError(t *testing.T) {
	for _, tt := range labelErrorTests {
		dec := decode.DecodeString(tt.input)
		actual, err := dec.Label()
		if actual != nil {
			t.Error("For:", tt.input, ", expected label:", nil, ", got:", actual)
			continue
		}

		if err.Error() != tt.expected.Error() {
			t.Error("For:", tt.input, ", expected error:", tt.expected, ", got:", err)
		}
	}
}
