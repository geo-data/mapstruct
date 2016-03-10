package encoding

import (
	"fmt"
	"github.com/geo-data/mapfile/tokens"
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

func (e *MapfileEncoder) TokenString(name string, value fmt.Stringer) (err error) {
	sv := value.String()
	if sv == "" {
		return
	}

	if err = e.indent(); err != nil {
		return
	}

	_, err = e.w.Write([]byte(fmt.Sprintf("%s \"%s\"\n", name, e.r.Replace(sv))))
	return
}

func (e *MapfileEncoder) TokenValue(name string, value fmt.Stringer) (err error) {
	sv := value.String()
	if sv == "" {
		return
	}

	if err = e.indent(); err != nil {
		return
	}

	_, err = e.w.Write([]byte(fmt.Sprintf("%s %s\n", name, sv)))
	return
}

func (e *MapfileEncoder) EncodeString(value tokens.String) (err error) {
	if value == "" {
		return
	}

	if err = e.indent(); err != nil {
		return
	}

	_, err = e.w.Write([]byte(fmt.Sprintf("\"%s\"\n", e.r.Replace(value.String()))))
	return
}

func (e *MapfileEncoder) EncodeValues(values ...fmt.Stringer) (err error) {
	if len(values) == 0 {
		return
	}

	var svals []string
	for _, value := range values {
		svals = append(svals, value.String())
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

func (e *MapfileEncoder) EncodeStrings(values ...fmt.Stringer) (err error) {
	if len(values) == 0 {
		return
	}

	var svals []string
	for _, value := range values {
		svals = append(svals, fmt.Sprintf("\"%s\"", e.r.Replace(value.String())))
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
