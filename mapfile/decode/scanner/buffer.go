package scanner

import "bytes"

type buffer struct {
	s   *scanner
	buf bytes.Buffer
}

func (b *buffer) write(ch rune) bool {
	if _, err := b.buf.WriteRune(ch); err != nil {
		b.s.setErr(err)
		return false
	}
	return true
}

func (b *buffer) String() string {
	return b.buf.String()
}
