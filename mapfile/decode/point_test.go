package decode_test

import (
	"errors"
	"github.com/geo-data/mapstruct/mapfile/decode"
	"github.com/geo-data/mapstruct/types"
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
	{"", errors.New("could not decode X coordinate: unexpected end of mapfile")},
	{"1", errors.New("could not decode Y coordinate: unexpected end of mapfile")},
	{"foo 1", errors.New(`could not decode X coordinate: token is not a number: foo`)},
	{"1 foo", errors.New(`could not decode Y coordinate: token is not a number: foo`)},
}

func TestDecodePoint(t *testing.T) {
	for _, tt := range pointTests {
		dec := decode.DecodeString(tt.input)
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
		dec := decode.DecodeString(tt.input)
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
