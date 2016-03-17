package decode_test

import (
	"errors"
	"github.com/geo-data/mapfile/mapfile/decode"
	"github.com/geo-data/mapfile/types"
	"reflect"
	"testing"
)

var pointsTests = []struct {
	input    string
	expected types.Points // expected result
}{
	{`
POINTS
END`, nil},
	{`
POINTS
  -0.2 51.5 0.2 51.5 0.2 52.5 -0.2 52.5 -0.2 51.5
END`, types.Points{
		types.NewPoint(-0.2, 51.5),
		types.NewPoint(0.2, 51.5),
		types.NewPoint(0.2, 52.5),
		types.NewPoint(-0.2, 52.5),
		types.NewPoint(-0.2, 51.5),
	}},
}

var pointsErrorTests = []struct {
	input    string
	expected error // expected result
}{
	{"POINTS", decode.EndOfTokens},
	{"POINTS 1 2", decode.EndOfTokens},
	{`
FOOBAR
END`, errors.New(`expected token POINTS, got: "FOOBAR"`)},
	{`
POINTS
  -0.2 51.5 0.2 51.5 0.2 52.5 -0.2 52.5 -0.2
END`, errors.New("invalid syntax for Y coordinate: END")},
	{`
POINTS
  -0.2 51.5 0.2 51.5 0.2 52.5 -0.2 52.5 FOO
END`, errors.New(`invalid syntax for X coordinate: "FOO"`)},
}

func TestDecodePoints(t *testing.T) {
	for _, tt := range pointsTests {
		dec, err := decode.DecodeString(tt.input)
		if err != nil {
			t.Error("For decoding:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}
		actual, err := dec.Points()
		if err != nil {
			t.Error("For:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}

		if !reflect.DeepEqual(tt.expected, actual) {
			t.Error("For:", tt.input, ", expected:", tt.expected, ", got:", actual)
		}
	}
}

func TestDecodePointsError(t *testing.T) {
	for _, tt := range pointsErrorTests {
		dec, err := decode.DecodeString(tt.input)
		if err != nil {
			t.Error("For decoding:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}
		actual, err := dec.Points()
		if actual != nil {
			t.Error("For:", tt.input, ", expected points:", nil, ", got:", actual)
			continue
		}

		if err.Error() != tt.expected.Error() {
			t.Error("For:", tt.input, ", expected error:", tt.expected, ", got:", err)
		}
	}
}
