package decode_test

import (
	"errors"
	"github.com/geo-data/mapfile/mapfile/decode"
	"github.com/geo-data/mapfile/types"
	"reflect"
	"testing"
)

var attributeTests = []struct {
	input    string
	expected types.Attribute // expected result
}{
	{`STATUS [foo]`, types.Attribute("foo")},
	{`STATUS "foo"`, types.Attribute("foo")},
	//{`"foo`, types.Attribute(`"foo`)}, // hangs the mapserver tokeniser.
	{`STATUS foo`, types.Attribute(`foo`)},
}

func TestDecodeAttribute(t *testing.T) {
	for _, tt := range attributeTests {
		dec, err := decode.DecodeString(tt.input)
		if err != nil {
			t.Error("For decoding:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}
		actual, err := dec.Next().Attribute()
		if err != nil {
			t.Error("For:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}

		if !reflect.DeepEqual(tt.expected, actual) {
			t.Error("For:", tt.input, ", expected:", tt.expected, ", got:", actual)
		}
	}
}

var keywordTests = []struct {
	input    string
	expected types.Keyword // expected result
}{
	{`STATUS "on"`, types.Keyword(`"on"`)},
	//{`STATUS "on`, types.Keyword(`"on`)}, // hangs the mapserver tokeniser.
	{`STATUS on`, types.Keyword(`on`)},
}

func TestDecodeKeyword(t *testing.T) {
	for _, tt := range keywordTests {
		dec, err := decode.DecodeString(tt.input)
		if err != nil {
			t.Error("For decoding:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}
		actual, err := dec.Next().Keyword()
		if err != nil {
			t.Error("For:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}

		if !reflect.DeepEqual(tt.expected, actual) {
			t.Error("For:", tt.input, ", expected:", tt.expected, ", got:", actual)
		}
	}
}

var stringTests = []struct {
	input    string
	expected types.String // expected result
}{
	{`STATUS "on"`, types.String("on")},
	//{`STATUS "on`, types.String("on")}, // hangs the mapserver tokeniser.
}

var stringErrorTests = []struct {
	input    string
	expected error // expected result
}{
	{`STATUS on`, errors.New(`not a map string: on`)},
}

func TestDecodeString(t *testing.T) {
	for _, tt := range stringTests {
		dec, err := decode.DecodeString(tt.input)
		if err != nil {
			t.Error("For decoding:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}
		actual, err := dec.Next().String()
		if err != nil {
			t.Error("For:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}

		if !reflect.DeepEqual(tt.expected, actual) {
			t.Error("For:", tt.input, ", expected:", tt.expected, ", got:", actual)
		}
	}
}
func TestDecodeStringError(t *testing.T) {
	for _, tt := range stringErrorTests {
		dec, err := decode.DecodeString(tt.input)
		if err != nil {
			t.Error("For decoding:", tt.input, ", expected error:", nil, ", got:", err)
			continue
		}
		var s types.String
		actual, err := dec.Next().String()
		if actual != s {
			t.Error("For:", tt.input, ", expected string:", nil, ", got:", actual)
			continue
		}

		if err.Error() != tt.expected.Error() {
			t.Error("For:", tt.input, ", expected error:", tt.expected, ", got:", err)
		}
	}
}
