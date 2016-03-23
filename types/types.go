package types

import (
	"fmt"
	"strconv"
	"strings"
)

type Union interface{}

type String string

type Attribute string

type Regex string

type Listex string

type Expression string

type Keyword string

type Uint8 uint8

type Integer int64

type Uint32 uint32

type Double float64

// String implements fmt.Stringer()
func (a Attribute) String() string {
	s := string(a)
	if len(s) > 0 {
		return fmt.Sprintf("[%s]", string(s))
	}
	return ""
}

func (a Attribute) QuotedString() string {
	s := string(a)
	if len(s) > 0 {
		return fmt.Sprintf(`"%s"`, string(s))
	}
	return ""
}

// String implements fmt.Stringer()
func (r Regex) String() string {
	s := string(r)
	if len(s) > 0 {
		return fmt.Sprintf("/%s/", string(s))
	}
	return ""
}

// String implements fmt.Stringer()
func (l Listex) String() string {
	attr := string(l)
	if len(attr) > 0 {
		return fmt.Sprintf("{%s}", string(attr))
	}
	return ""
}

// String implements fmt.Stringer()
func (e Expression) String() string {
	s := string(e)
	if len(s) > 0 {
		return fmt.Sprintf("(%s)", string(s))
	}
	return ""
}

// String implements fmt.Stringer()
func (t String) String() string {
	s := string(t)
	if len(s) > 0 {
		return fmt.Sprintf(`"%s"`, strings.Replace(s, `"`, `\"`, -1))
	}
	return s
}

// String implements fmt.Stringer()
func (k Keyword) String() string {
	return string(k)
}

// String implements fmt.Stringer()
func (t Uint8) String() string {
	return fmt.Sprintf("%d", uint8(t))
}

// String implements fmt.Stringer()
func (t Integer) String() string {
	return strconv.FormatInt(int64(t), 10)
}

// String implements fmt.Stringer()
func (t Uint32) String() string {
	return fmt.Sprintf("%d", uint32(t))
}

// String implements fmt.Stringer()
func (t Double) String() string {
	return strconv.FormatFloat(float64(t), 'f', -1, 64)
}
