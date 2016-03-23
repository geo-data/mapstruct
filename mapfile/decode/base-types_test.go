package decode_test

import (
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
	{`STATUS foo`, types.Attribute(`foo`)},
}

func TestDecodeAttribute(t *testing.T) {
	for _, tt := range attributeTests {
		dec := decode.DecodeString(tt.input)
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
	{`STATUS on`, types.Keyword(`on`)},
}

func TestDecodeKeyword(t *testing.T) {
	for _, tt := range keywordTests {
		dec := decode.DecodeString(tt.input)
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
	{`STATUS on`, types.String("on")},
}

func TestDecodeString(t *testing.T) {
	for _, tt := range stringTests {
		dec := decode.DecodeString(tt.input)
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
