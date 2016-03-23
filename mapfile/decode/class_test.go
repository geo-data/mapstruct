package decode_test

import (
	"errors"
	"github.com/geo-data/mapstruct/mapfile/decode"
	"github.com/geo-data/mapstruct/types"
	"reflect"
	"testing"
)

var classTests = []struct {
	input    string
	expected *types.Class // expected result
}{
	{`
CLASS
END`, &types.Class{}},
	{`
CLASS
  NAME "test"
END`, &types.Class{
		Name: types.String("test"),
	}},
	{`
CLASS
  EXPRESSION "foo"
END`, &types.Class{
		Expression: types.String("foo"),
	}},
	{`
CLASS
  METADATA
  END
END`, &types.Class{
		Metadata: types.NewMetadata(),
	}},
	{`
CLASS
  STYLE
    SYMBOL 1
  END
  STYLE
    SYMBOL 2
  END
END`, &types.Class{
		Styles: []*types.Style{
			&types.Style{
				Symbol: types.Union(types.Integer(1)),
			},
			&types.Style{
				Symbol: types.Union(types.Integer(2)),
			},
		},
	}},
	{`
CLASS
  LABEL
  END
END`, &types.Class{
		Label: &types.Label{},
	}},
	{`
CLASS
  TEMPLATE "foo"
END`, &types.Class{
		Template: types.String("foo"),
	}},
	{`
CLASS
  TEXT "foo"
END`, &types.Class{
		Text: types.String("foo"),
	}},
}

var classErrorTests = []struct {
	input    string
	expected error // expected result
}{
	{"CLASS", decode.EndOfTokens},
	{`
FOOBAR
END`, errors.New(`expected token CLASS, got: FOOBAR`)},
	{`
CLASS
  FOO BAR
END`, errors.New(`unhandled mapfile token: FOO`)},
}

func TestDecodeClass(t *testing.T) {
	for _, tt := range classTests {
		dec := decode.DecodeString(tt.input)
		actual, err := dec.Class()
		if err != nil {
			t.Error("For:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}

		if !reflect.DeepEqual(tt.expected, actual) {
			t.Error("For:", tt.input, ", expected:", tt.expected, ", got:", actual)
		}
	}
}

func TestDecodeClassError(t *testing.T) {
	for _, tt := range classErrorTests {
		dec := decode.DecodeString(tt.input)
		actual, err := dec.Class()
		if actual != nil {
			t.Error("For:", tt.input, ", expected class:", nil, ", got:", actual)
			continue
		}

		if err.Error() != tt.expected.Error() {
			t.Error("For:", tt.input, ", expected error:", tt.expected, ", got:", err)
		}
	}
}
