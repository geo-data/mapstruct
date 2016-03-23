// package scanner implements a basic scanner/lexer for the mapfile format.  It
// is adapted from the code at
// <https://blog.gopheracademy.com/advent-2014/parsers-lexers/>.

package scanner

import (
	"bufio"
	"io"
	"strings"
)

// isWhitespace returns true if ch is whitespace, false otherwise.
func isWhitespace(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

// isLetter returns true if ch is a letter, false otherwise.
func isLetter(ch rune) bool {
	return (ch >= 'a' && ch <= 'z') || (ch >= 'A' && ch <= 'Z')
}

// isIdent returns true if ch is an identifier, false otherwise.
func isIdent(ch rune) bool {
	return isLetter(ch) || ch == '_'
}

// isDigit returns true if ch is a digit, false otherwise.
func isDigit(ch rune) bool {
	return (ch >= '0' && ch <= '9')
}

// isNumber returns true if ch is a number, false otherwise.
func isNumber(ch rune) bool {
	return isDigit(ch) || ch == '-' || ch == '+' || ch == '.'
}

// eof represents End Of File.
const eof = rune(0)

// Scanner represents a lexical scanner.
type Scanner interface {
	Scan() *Token
	Err() error
}

// scanner implements a lexical scanner which tokenises an io.Reader.
type scanner struct {
	r   *bufio.Reader
	err error // Sticky error.
}

// Newscanner returns a new instance of scanner.
func NewScanner(r io.Reader) Scanner {
	return &scanner{r: bufio.NewReader(r)}
}

// setErr records the first error encountered.
func (s *scanner) setErr(err error) {
	if s.err == nil || s.err == io.EOF {
		s.err = err
	}
}

// Err returns the first non-EOF error that was encountered by the scanner.
func (s *scanner) Err() error {
	if s.err == io.EOF {
		return nil
	}
	return s.err
}

// read reads the next rune from the buffered reader.
// Returns the rune(0) if an error occurs (or io.EOF is returned).
func (s *scanner) read() rune {
	if s.err != nil {
		return eof
	}

	ch, _, err := s.r.ReadRune()
	if err != nil {
		s.setErr(err)
		return eof
	}

	return ch
}

// unread places the previously read rune back on the reader.
func (s *scanner) unread() {
	if err := s.r.UnreadRune(); err != nil {
		s.setErr(err)
	}
}

// Scan returns the next token and literal value.
func (s *scanner) Scan() *Token {
	// Read the next rune.
	ch := s.read()

	// If we see whitespace then consume all contiguous whitespace.
	// If we see a letter then consume as an keyword or reserved word.
	if isWhitespace(ch) {
		s.unread()
		return s.scanWhitespace()
	} else if isIdent(ch) {
		s.unread()
		return s.scanKeyword()
	} else if isNumber(ch) {
		s.unread()
		return s.scanNumber()
	}

	// Otherwise read the individual character.
	switch ch {
	case eof:
		return &Token{EOF, ""}
	case '#':
		s.unread()
		return s.scanComment()
	case '/':
		s.unread()
		return s.scanRegex()
	case '(':
		s.unread()
		return s.scanExpression()
	case '[':
		s.unread()
		return s.scanAttribute()
	case '{':
		s.unread()
		return s.scanListex()

	// String
	case '"':
		fallthrough
	case '\'':
		s.unread()
		return s.scanString()
	}

	if s.Err() != nil {
		return nil
	}

	return &Token{ILLEGAL, string(ch)}
}

// buffer returns a new writeable buffer object associated with scanner.
func (s *scanner) buffer() *buffer {
	return &buffer{s: s}
}

// scanWhitespace consumes the current rune and all contiguous whitespace.
func (s *scanner) scanWhitespace() *Token {
	// Create a buffer and read the current character into it.
	buf := s.buffer()
	if !buf.write(s.read()) {
		return nil
	}

	// Read every subsequent whitespace character into the buffer.
	// Non-whitespace characters and EOF will cause the loop to exit.
	for s.err == nil {
		if ch := s.read(); ch == eof {
			break
		} else if !isWhitespace(ch) {
			s.unread()
			break
		} else {
			if !buf.write(ch) {
				return nil
			}
		}
	}

	if s.Err() != nil {
		return nil
	}

	return &Token{WS, buf.String()}
}

// scanComment consumes the current rune and everything until a newline or EOF.
func (s *scanner) scanComment() *Token {
	// Create a buffer and read the current character into it.
	buf := s.buffer()
	s.read() // don't add the comment identifier.

	// Read every subsequent character into the buffer, until end of line or
	// EOF.
Loop:
	for s.err == nil {
		ch := s.read()
		switch ch {
		case eof:
			break Loop
		case '\n':
			s.unread()
			break Loop
		default:
			if !buf.write(ch) {
				return nil
			}
		}
	}

	if s.Err() != nil {
		return nil
	}

	return &Token{MS_COMMENT, buf.String()}
}

// scanNumber consumes the current rune and everything until a newline or EOF.
func (s *scanner) scanNumber() *Token {
	// Create a buffer and read the current character into it.
	buf := s.buffer()
	decimals := 0
	ch := s.read()
	if ch == '.' {
		decimals++
	}
	if !buf.write(ch) {
		return nil
	}

	// Read every subsequent character into the buffer, until end of line or
	// EOF.
Loop:
	for s.err == nil {
		ch := s.read()
		if isDigit(ch) {
			if !buf.write(ch) {
				return nil
			}
		} else if ch == '.' {
			if decimals > 0 {
				s.unread()
				break Loop
			} else {
				decimals++
				if !buf.write(ch) {
					return nil
				}
			}
		} else if ch == eof {
			break Loop
		} else {
			s.unread()
			break Loop
		}
	}

	if s.Err() != nil {
		return nil
	}

	return &Token{MS_NUMBER, buf.String()}
}

// scanString consumes the current rune and everything until the end of the
// string, identified with a '"'.
func (s *scanner) scanString() *Token {
	// Create a buffer and read the current character into it.
	buf := s.buffer()
	delimiter := s.read()

	// Read every subsequent character into the buffer, until either the
	// delimiter or EOF is reached .
Loop:
	for s.err == nil {
		ch := s.read()
		switch ch {
		case eof:
			break Loop
		case delimiter:
			break Loop
		case '\\':
			// don't assign any significance to the next character, just read it:
			if ch = s.read(); ch == eof {
				break Loop
			}
			fallthrough
		default:
			if !buf.write(ch) {
				return nil
			}
		}
	}

	if s.Err() != nil {
		return nil
	}

	return &Token{MS_STRING, buf.String()}
}

// scanRegex consumes the current rune and everything until the end of the regex
// or EOF.
func (s *scanner) scanRegex() (tok *Token) {
	if tok = s.scanDelimited('/'); tok != nil {
		tok.Type = MS_REGEX
	}
	return
}

// scanListex consumes the current rune and everything until the end of the
// listex or EOF.
func (s *scanner) scanListex() (tok *Token) {
	if tok = s.scanDelimited('}'); tok != nil {
		tok.Type = MS_LISTEX
	}
	return tok
}

// scanAttribute consumes the current rune and everything the end of the
// attribute or EOF.
func (s *scanner) scanAttribute() (tok *Token) {
	if tok = s.scanDelimited(']'); tok != nil {
		tok.Type = MS_ATTRIBUTE
	}
	return tok
}

// scanExpression consumes the current rune and everything until the end of the
// expression or EOF.
func (s *scanner) scanExpression() (tok *Token) {
	if tok = s.scanDelimited(')'); tok != nil {
		tok.Type = MS_EXPRESSION
	}
	return tok
}

// scanDelimited consumes the current rune and everything until end or EOF.
func (s *scanner) scanDelimited(end rune) *Token {
	// Create a buffer and read the current character into it.
	buf := s.buffer()
	start := s.read() // ignore the opening delimiter

	depth := 0
	// Read every subsequent character into the buffer, until end delimiter or
	// EOF is reached.
Loop:
	for s.err == nil {
		ch := s.read()
		switch ch {
		case eof:
			break Loop
		case end: // evaluate before start in case end == start.
			if depth == 0 {
				break Loop
			} else {
				depth--
				if !buf.write(ch) {
					return nil
				}
			}
		case start:
			depth++
			if !buf.write(ch) {
				return nil
			}
		default:
			if !buf.write(ch) {
				return nil
			}
		}
	}

	if s.Err() != nil {
		return nil
	}

	return &Token{MS_STRING, buf.String()}
}

// scanKeyword consumes the current rune and all contiguous keyword runes.
func (s *scanner) scanKeyword() *Token {
	// Create a buffer and read the current character into it.
	buf := s.buffer()
	if !buf.write(s.read()) {
		return nil
	}

	// Read every subsequent keyword character into the buffer.
	// Non-string characters and EOF will cause the loop to exit.
	for s.err == nil {
		if ch := s.read(); ch == eof {
			break
		} else if !isIdent(ch) {
			s.unread()
			break
		} else {
			if !buf.write(ch) {
				return nil
			}
		}
	}

	if s.Err() != nil {
		return nil
	}

	// If the string matches a keyword then return that keyword.
	if tok, ok := fromString[strings.ToUpper(buf.String())]; ok {
		return &Token{tok, buf.String()}
	}

	// Otherwise return as a regular string.
	return &Token{MS_BARE_STRING, buf.String()}
}
