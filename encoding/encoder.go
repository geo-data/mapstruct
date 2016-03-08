package encoding

import (
	"fmt"
	"github.com/geo-data/mapfile/tokens"
	"io"
	"strings"
)

type Encoder interface {
	Encode(*MapfileEncoder) error
}

type MapfileEncoder struct {
	w io.Writer
	r *strings.Replacer
}

func NewMapfileEncoder(w io.Writer) *MapfileEncoder {
	return &MapfileEncoder{
		w,
		strings.NewReplacer(`\`, `\\`, `"`, `\"`),
	}
}

func (e *MapfileEncoder) Encode(m Encoder) error {
	return m.Encode(e)
}

func (e *MapfileEncoder) TokenString(name string, value fmt.Stringer) (err error) {
	sv := value.String()
	if sv == "" {
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

	_, err = e.w.Write([]byte(fmt.Sprintf("%s %s\n", name, sv)))
	return
}

func (e *MapfileEncoder) EncodeString(value tokens.String) (err error) {
	if value == "" {
		return
	}

	_, err = e.w.Write([]byte(fmt.Sprintf("\"%s\"\n", e.r.Replace(value.String()))))
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

	join := strings.Join(svals, " ")
	if _, err = e.w.Write([]byte(fmt.Sprintf("%s\n", join))); err != nil {
		return
	}

	return
}

func (e *MapfileEncoder) TokenStart(name string) (err error) {
	_, err = e.w.Write([]byte(fmt.Sprintf("%s\n", name)))
	return
}

func (e *MapfileEncoder) TokenEnd() (err error) {
	_, err = e.w.Write([]byte("END\n"))
	return
}
