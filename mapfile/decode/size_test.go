package decode_test

import (
	"errors"
	"github.com/geo-data/mapfile/mapfile/decode"
	"github.com/geo-data/mapfile/types"
	"reflect"
	"testing"
)

var sizeTests = []struct {
	input    string
	expected *types.Size // expected result
}{
	{`
SIZE 0 0`, types.NewSize(0, 0)},
	{`
SIZE 20 10`, types.NewSize(20, 10)},
}

var sizeErrorTests = []struct {
	input    string
	expected error // expected result
}{
	{"SIZE", errors.New("could not decode width: unexpected end of mapfile")},
	{"SIZE 1", errors.New("could not decode height: unexpected end of mapfile")},
	{`
FOOBAR 5 10`, errors.New(`expected token SIZE, got: FOOBAR`)},
	{`
SIZE foo 10`, errors.New(`could not decode width: token is not a number: foo`)},
	{`
SIZE 6 foo`, errors.New(`could not decode height: token is not a number: foo`)},
}

func TestDecodeSize(t *testing.T) {
	for _, tt := range sizeTests {
		dec := decode.DecodeString(tt.input)
		actual, err := dec.Size()
		if err != nil {
			t.Error("For:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}

		if !reflect.DeepEqual(tt.expected, actual) {
			t.Error("For:", tt.input, ", expected:", tt.expected, ", got:", actual)
		}
	}
}

func TestDecodeSizeError(t *testing.T) {
	for _, tt := range sizeErrorTests {
		dec := decode.DecodeString(tt.input)
		actual, err := dec.Size()
		if actual != nil {
			t.Error("For:", tt.input, ", expected size:", nil, ", got:", actual)
			continue
		}

		if err.Error() != tt.expected.Error() {
			t.Error("For:", tt.input, ", expected error:", tt.expected, ", got:", err)
		}
	}
}
