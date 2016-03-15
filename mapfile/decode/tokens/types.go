package tokens

import "strings"

type Type int

const (
	Attribute = 1 << iota
	String
	Double
	Integer
	Keyword
)

func (t Type) Is(t2 Type) bool {
	return t&t2 != 0
}

func (t Type) String() string {
	types := []string{}
	if t.Is(Attribute) {
		types = append(types, "Attribute")
	}
	if t.Is(String) {
		types = append(types, "String")
	}
	if t.Is(Double) {
		types = append(types, "Double")
	}
	if t.Is(Integer) {
		types = append(types, "Integer")
	}
	if t.Is(Keyword) {
		types = append(types, "Keyword")
	}

	return strings.Join(types, ", ")
}
