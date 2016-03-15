package encoding

import (
	"fmt"
	"github.com/geo-data/mapfile/types"
	"io"
	"regexp"
	"strings"
)

type Encoder interface {
	Encode(*MapfileEncoder) error
}

var isNum = regexp.MustCompile(`-?\d+(\.\d+)?`)

type MapfileEncoder struct {
	w     io.Writer
	r     *strings.Replacer
	depth int
}

func NewMapfileEncoder(w io.Writer) *MapfileEncoder {
	return &MapfileEncoder{
		w,
		strings.NewReplacer(`\`, `\\`, `"`, `\"`),
		0,
	}
}

func (e *MapfileEncoder) Encode(m Encoder) error {
	return m.Encode(e)
}

func (e *MapfileEncoder) indent() error {
	for i := 0; i < e.depth; i++ {
		if _, err := e.w.Write([]byte("  ")); err != nil {
			return err
		}
	}

	return nil
}

func (e *MapfileEncoder) TokenString(name string, value string) (err error) {
	if value == "" {
		return
	}

	if err = e.indent(); err != nil {
		return
	}

	_, err = e.w.Write([]byte(fmt.Sprintf("%s %s\n", name, value)))
	return
}

func (e *MapfileEncoder) TokenStringer(name string, value fmt.Stringer) (err error) {
	if value == nil {
		return
	}

	return e.TokenString(name, value.String())
}

func (e *MapfileEncoder) TokenUnion(name string, value types.Union) (err error) {
	var s string
	switch t := value.(type) {
	case nil:
		return // Don't encode.
	case types.Float64:
		s = fmt.Stringer(t).String()
	case types.String:
		s = fmt.Stringer(t).String()
	case types.Attribute:
		s = fmt.Stringer(t).String()
	case types.Keyword:
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
	return e.TokenString(name, s)
}

func (e *MapfileEncoder) EncodeString(value fmt.Stringer) (err error) {
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

func (e *MapfileEncoder) EncodeStringers(values ...fmt.Stringer) (err error) {
	if len(values) == 0 {
		return
	}

	var svals []string
	for _, value := range values {
		if value != nil {
			svals = append(svals, value.String())
		}
	}

	if err = e.indent(); err != nil {
		return
	}

	join := strings.Join(svals, " ")
	if _, err = e.w.Write([]byte(fmt.Sprintf("%s\n", join))); err != nil {
		return
	}

	return
}

func (e *MapfileEncoder) TokenStart(name string) (err error) {
	if err = e.indent(); err != nil {
		return
	}

	if _, err = e.w.Write([]byte(fmt.Sprintf("%s\n", name))); err != nil {
		return
	}

	e.depth++

	return
}

func (e *MapfileEncoder) TokenEnd() (err error) {
	e.depth--

	if err = e.indent(); err != nil {
		return
	}

	if _, err = e.w.Write([]byte("END\n")); err != nil {
		return
	}

	return
}
