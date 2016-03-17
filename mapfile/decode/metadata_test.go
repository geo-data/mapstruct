package decode_test

import (
	"errors"
	"github.com/geo-data/mapfile/mapfile/decode"
	"github.com/geo-data/mapfile/types"
	"reflect"
	"testing"
)

var metadataTests = []struct {
	input    string
	expected types.Metadata // expected result
}{
	{`
METADATA
END`, types.NewMetadata()},
	{`
METADATA
  "foo" "bar"
  "cat" "dog"
  "key" "value"
END`, types.Metadata(map[string]string{
		"foo": "bar",
		"cat": "dog",
		"key": "value",
	})},
}

var metadataErrorTests = []struct {
	input    string
	expected error // expected result
}{
	{"METADATA", decode.EndOfTokens},
	{`
METADATA
  "foo"`, decode.EndOfTokens},
	{`
METADATA
  "foo" "bar"`, decode.EndOfTokens},
	{`
FOOBAR
END`, errors.New(`expected token METADATA, got: "FOOBAR"`)},
	{`
METADATA
  "foo" "bar"
  "cat" 
END`, errors.New(`expected value after key: "cat"`)},
}

func TestDecodeMetadata(t *testing.T) {
	for _, tt := range metadataTests {
		dec, err := decode.DecodeString(tt.input)
		if err != nil {
			t.Error("For decoding:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}
		actual, err := dec.Metadata()
		if err != nil {
			t.Error("For:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}

		if !reflect.DeepEqual(tt.expected, actual) {
			t.Error("For:", tt.input, ", expected:", tt.expected, ", got:", actual)
		}
	}
}

func TestDecodeMetadataError(t *testing.T) {
	for _, tt := range metadataErrorTests {
		dec, err := decode.DecodeString(tt.input)
		if err != nil {
			t.Error("For decoding:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}
		actual, err := dec.Metadata()
		if actual != nil {
			t.Error("For:", tt.input, ", expected metadata:", nil, ", got:", actual)
			continue
		}

		if err.Error() != tt.expected.Error() {
			t.Error("For:", tt.input, ", expected error:", tt.expected, ", got:", err)
		}
	}
}
