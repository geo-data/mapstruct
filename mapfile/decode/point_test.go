package decode_test

import (
	"errors"
	"github.com/geo-data/mapfile/mapfile/decode"
	"github.com/geo-data/mapfile/types"
	"reflect"
	"testing"
)

var pointTests = []struct {
	input    string
	expected *types.Point // expected result
}{
	{"1 2", types.NewPoint(1, 2)},
	{"0 0", types.NewPoint(0, 0)},
}

var pointErrorTests = []struct {
	input    string
	expected error // expected result
}{
	{"", decode.EndOfTokens},
	{"1", decode.EndOfTokens},
	{"foo 1", errors.New(`invalid syntax for X coordinate: "foo"`)},
	{"1 foo", errors.New(`invalid syntax for Y coordinate: "foo"`)},
}

func TestDecodePoint(t *testing.T) {
	for _, tt := range pointTests {
		dec, err := decode.DecodeString(tt.input)
		if err != nil {
			t.Error("For decoding:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}
		actual, err := dec.Point()
		if err != nil {
			t.Error("For:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}

		if !reflect.DeepEqual(tt.expected, actual) {
			t.Error("For:", tt.input, ", expected:", tt.expected, ", got:", actual)
		}
	}
}

func TestDecodePointError(t *testing.T) {
	for _, tt := range pointErrorTests {
		dec, err := decode.DecodeString(tt.input)
		if err != nil {
			t.Error("For decoding:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}
		actual, err := dec.Point()
		if actual != nil {
			t.Error("For:", tt.input, ", expected point:", nil, ", got:", actual)
			continue
		}

		if err.Error() != tt.expected.Error() {
			t.Error("For:", tt.input, ", expected error:", tt.expected, ", got:", err)
		}
	}
}
