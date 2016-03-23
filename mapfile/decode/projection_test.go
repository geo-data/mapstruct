package decode_test

import (
	"errors"
	"github.com/geo-data/mapstruct/mapfile/decode"
	"github.com/geo-data/mapstruct/types"
	"reflect"
	"testing"
)

var projectionTests = []struct {
	input    string
	expected types.Projection // expected result
}{
	{`
PROJECTION
END`, nil},
	{`
PROJECTION
  "proj=utm"
  "ellps=GRS80"
  "datum=NAD83"
  "zone=15"
  "units=m"
  "north"
  "no_defs"
END`, types.Projection([]string{
		"proj=utm",
		"ellps=GRS80",
		"datum=NAD83",
		"zone=15",
		"units=m",
		"north",
		"no_defs",
	})},
}

var projectionErrorTests = []struct {
	input    string
	expected error // expected result
}{
	{"PROJECTION", decode.EndOfTokens},
	{`
FOOBAR
END`, errors.New(`expected token PROJECTION, got: FOOBAR`)},
}

func TestDecodeProjection(t *testing.T) {
	for _, tt := range projectionTests {
		dec := decode.DecodeString(tt.input)
		actual, err := dec.Projection()
		if err != nil {
			t.Error("For:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}

		if !reflect.DeepEqual(tt.expected, actual) {
			t.Error("For:", tt.input, ", expected:", tt.expected, ", got:", actual)
		}
	}
}

func TestDecodeProjectionError(t *testing.T) {
	for _, tt := range projectionErrorTests {
		dec := decode.DecodeString(tt.input)
		actual, err := dec.Projection()
		if actual != nil {
			t.Error("For:", tt.input, ", expected projection:", nil, ", got:", actual)
			continue
		}

		if err.Error() != tt.expected.Error() {
			t.Error("For:", tt.input, ", expected error:", tt.expected, ", got:", err)
		}
	}
}
