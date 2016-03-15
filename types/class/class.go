package class

import (
	"fmt"
	"github.com/geo-data/mapfile/mapfile/encode"
	"github.com/geo-data/mapfile/tokens"
	"github.com/geo-data/mapfile/types"
	"github.com/geo-data/mapfile/types/label"
	"github.com/geo-data/mapfile/types/metadata"
	"github.com/geo-data/mapfile/types/style"
)

type Class struct {
	Name       types.String       `json:",omitempty"`
	Expression types.String       `json:",omitempty"`
	Metadata   *metadata.Metadata `json:",omitempty"`
	Styles     []*style.Style     `json:",omitempty"`
	Label      *label.Label       `json:",omitempty"`
	Template   types.String       `json:",omitempty"`
	Text       types.String       `json:",omitempty"`
}

func New(toks *tokens.Tokens) (c *Class, err error) {
	token := toks.Value()
	if token != "CLASS" {
		err = fmt.Errorf("expected token CLASS, got: %s", token)
		return
	}
	toks.Next()

	c = new(Class)
	for toks != nil {
		token := toks.Value()
		switch token {
		case "NAME":
			if c.Name, err = toks.Next().String(); err != nil {
				return
			}
		case "EXPRESSION":
			if c.Expression, err = toks.Next().String(); err != nil {
				return
			}
		case "TEMPLATE":
			if c.Template, err = toks.Next().String(); err != nil {
				return
			}
		case "TEXT":
			if c.Text, err = toks.Next().String(); err != nil {
				return
			}
		case "METADATA":
			if c.Metadata, err = metadata.New(toks); err != nil {
				return
			}
		case "STYLE":
			var s *style.Style
			if s, err = style.New(toks); err != nil {
				return
			}
			c.Styles = append(c.Styles, s)
		case "LABEL":
			if c.Label, err = label.New(toks); err != nil {
				return
			}
		case "END":
			return
		default:
			err = fmt.Errorf("unhandled mapfile token: %s", token)
			return
		}

		toks = toks.Next()
	}

	return
}

func (c *Class) Encode(enc *encode.MapfileEncoder) (err error) {
	if err = enc.TokenStart("CLASS"); err != nil {
		return
	}

	if err = enc.TokenStringer("NAME", c.Name); err != nil {
		return
	}
	if err = enc.TokenStringer("EXPRESSION", c.Expression); err != nil {
		return
	}
	if err = enc.TokenStringer("TEMPLATE", c.Template); err != nil {
		return
	}
	if err = enc.TokenStringer("TEXT", c.Text); err != nil {
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
