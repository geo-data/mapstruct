package encode

import (
	"fmt"
	"io"
	"strings"

	"github.com/geo-data/mapfile/types"
)

type Encoder struct {
	w      io.Writer
	r      *strings.Replacer
	depth  int
	Indent string
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		w,
		strings.NewReplacer(`\`, `\\`, `"`, `\"`),
		0,
		"  ",
	}
}

func (e *Encoder) indent() error {
	for i := 0; i < e.depth; i++ {
		if _, err := e.w.Write([]byte(e.Indent)); err != nil {
			return err
		}
	}

	return nil
}

func (e *Encoder) EncodeDirectiveString(name string, value string) (err error) {
	if value == "" {
		return
	}

	if err = e.indent(); err != nil {
		return
	}

	_, err = e.w.Write([]byte(fmt.Sprintf("%s %s\n", name, value)))
	return
}

func (e *Encoder) EncodeDirectiveStringer(name string, value fmt.Stringer) (err error) {
	if value == nil {
		return
	}

	return e.EncodeDirectiveString(name, value.String())
}

func (e *Encoder) EncodeDirectiveUnion(name string, value types.Union) (err error) {
	var s string
	switch t := value.(type) {
	case nil:
		return // Don't encode.
	case types.Double:
		s = fmt.Stringer(t).String()
	case types.String:
		s = fmt.Stringer(t).String()
	case types.Attribute:
		s = fmt.Stringer(t).String()
	case types.Keyword:
		s = fmt.Stringer(t).String()
	case types.Expression:
		s = fmt.Stringer(t).String()
	case types.Regex:
		s = fmt.Stringer(t).String()
	case types.Listex:
		s = fmt.Stringer(t).String()
	case types.Uint8:
		s = fmt.Stringer(t).String()
	case types.Integer:
		s = fmt.Stringer(t).String()
	case types.Uint32:
		s = fmt.Stringer(t).String()
	default:
		err = fmt.Errorf("unhandled type: %v", t)
		return
	}
	return e.EncodeDirectiveString(name, s)
}

func (e *Encoder) EncodeString(value fmt.Stringer) (err error) {
	s := value.String()
	if s == "" {
		return
	}

	if err = e.indent(); err != nil {
		return
	}

	if _, err = e.w.Write([]byte(s)); err != nil {
		return
	}

	_, err = e.w.Write([]byte{'\n'})
	return
}

func (e *Encoder) EncodeStringers(values ...fmt.Stringer) (err error) {
	if len(values) == 0 {
		return
	}

	var svals []string
	for _, value := range values {
		if value != nil {
			svals = append(svals, value.String())
		}
	}

	return e.EncodeStrings(svals...)
}

func (e *Encoder) EncodeStrings(values ...string) (err error) {
	if len(values) == 0 {
		return
	}

	if err = e.indent(); err != nil {
		return
	}

	join := strings.Join(values, " ")
	if _, err = e.w.Write([]byte(fmt.Sprintf("%s\n", join))); err != nil {
		return
	}

	return
}

func (e *Encoder) StartDirective(name string) (err error) {
	if err = e.indent(); err != nil {
		return
	}

	if _, err = e.w.Write([]byte(fmt.Sprintf("%s\n", name))); err != nil {
		return
	}

	e.depth++

	return
}

func (e *Encoder) EndDirective() (err error) {
	e.depth--

	if err = e.indent(); err != nil {
		return
	}

	if _, err = e.w.Write([]byte("END\n")); err != nil {
		return
	}

	return
}
