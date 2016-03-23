package decode

import "strings"

type Type int

const (
	Attribute = 1 << iota
	String
	Double
	Integer
	Keyword
	Expression
	Regex
	Listex
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
	if t.Is(Expression) {
		types = append(types, "Expression")
	}
	if t.Is(Regex) {
		types = append(types, "Regular expression")
	}
	if t.Is(Listex) {
		types = append(types, "List expression")
	}

	return strings.Join(types, ", ")
}
