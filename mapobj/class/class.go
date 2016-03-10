package class

import (
	"fmt"
	"github.com/geo-data/mapfile/encoding"
	"github.com/geo-data/mapfile/mapobj/label"
	"github.com/geo-data/mapfile/mapobj/metadata"
	"github.com/geo-data/mapfile/mapobj/style"
	"github.com/geo-data/mapfile/tokens"
)

type Class struct {
	Name       tokens.String      `json:",omitempty"`
	Expression tokens.String      `json:",omitempty"`
	Metadata   *metadata.Metadata `json:",omitempty"`
	Styles     []*style.Style     `json:",omitempty"`
	Label      *label.Label       `json:",omitempty"`
	Template   tokens.String      `json:",omitempty"`
	Text       tokens.String      `json:",omitempty"`
}

func New(tokens *tokens.Tokens) (c *Class, err error) {
	token := tokens.Value()
	if token != "CLASS" {
		err = fmt.Errorf("expected token CLASS, got: %s", token)
		return
	}
	tokens.Next()

	c = new(Class)
	for tokens != nil {
		token := tokens.Value()
		switch token {
		case "NAME":
			c.Name = tokens.Next().Value()
		case "EXPRESSION":
			c.Expression = tokens.Next().Value()
		case "TEMPLATE":
			c.Template = tokens.Next().Value()
		case "TEXT":
			c.Text = tokens.Next().Value()
		case "METADATA":
			if c.Metadata, err = metadata.New(tokens); err != nil {
				return
			}
		case "STYLE":
			var s *style.Style
			if s, err = style.New(tokens); err != nil {
				return
			}
			c.Styles = append(c.Styles, s)
		case "LABEL":
			if c.Label, err = label.New(tokens); err != nil {
				return
			}
		case "END":
			return
		default:
			err = fmt.Errorf("unhandled mapfile token: %s", token)
			return
		}

		tokens = tokens.Next()
	}

	return
}

func (c *Class) Encode(enc *encoding.MapfileEncoder) (err error) {
	if err = enc.TokenStart("CLASS"); err != nil {
		return
	}

	if err = enc.TokenString("NAME", c.Name); err != nil {
		return
	}
	if err = enc.TokenString("EXPRESSION", c.Expression); err != nil {
		return
	}
	if err = enc.TokenString("TEMPLATE", c.Template); err != nil {
		return
	}
	if err = enc.TokenString("TEXT", c.Text); err != nil {
		return
	}
	if c.Metadata != nil {
		if err = c.Metadata.Encode(enc); err != nil {
			return
		}
	}
	if c.Label != nil {
		if err = c.Label.Encode(enc); err != nil {
			return
		}
	}

	for _, style := range c.Styles {
		if err = style.Encode(enc); err != nil {
			return
		}
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
