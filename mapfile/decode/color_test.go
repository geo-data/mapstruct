package decode_test

import (
	"errors"
	"github.com/geo-data/mapstruct/mapfile/decode"
	"github.com/geo-data/mapstruct/types"
	"reflect"
	"testing"
)

var colorTests = []struct {
	input    string
	expected *types.Color // expected result
}{
	{`
COLOR 0 0 0`, types.NewColor(0, 0, 0, 255)},
	{`
COLOR 1 2 3`, types.NewColor(1, 2, 3, 255)},
}

var colorErrorTests = []struct {
	input    string
	expected error // expected result
}{
	{`
FOOBAR 5 10`, decode.EndOfTokens},
	{`
COLOR 256 2 3`, errors.New(`strconv.ParseUint: parsing "256": value out of range`)},
	{`
COLOR foo 2 3`, errors.New(`token is not a number: foo`)},
	{`
COLOR 1 foo 3`, errors.New(`token is not a number: foo`)},
	{`
COLOR 1 2 foo`, errors.New(`token is not a number: foo`)},
}

func TestDecodeColor(t *testing.T) {
	for _, tt := range colorTests {
		dec := decode.DecodeString(tt.input)
		actual, err := dec.Color()
		if err != nil {
			t.Error("For:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}

		if !reflect.DeepEqual(tt.expected, actual) {
			t.Error("For:", tt.input, ", expected:", tt.expected, ", got:", actual)
		}
	}
}

func TestDecodeColorError(t *testing.T) {
	for _, tt := range colorErrorTests {
		dec := decode.DecodeString(tt.input)
		actual, err := dec.Color()
		if actual != nil {
			t.Error("For:", tt.input, ", expected color:", nil, ", got:", actual)
			continue
		}

		if err.Error() != tt.expected.Error() {
			t.Error("For:", tt.input, ", expected error:", tt.expected, ", got:", err)
		}
	}
}
