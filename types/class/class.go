package class

import (
	"github.com/geo-data/mapfile/mapfile/encode"
	"github.com/geo-data/mapfile/types"
	"github.com/geo-data/mapfile/types/label"
	"github.com/geo-data/mapfile/types/metadata"
	"github.com/geo-data/mapfile/types/style"
)

type Class struct {
	Name       types.String      `json:",omitempty"`
	Expression types.String      `json:",omitempty"`
	Metadata   metadata.Metadata `json:",omitempty"`
	Styles     []*style.Style    `json:",omitempty"`
	Label      *label.Label      `json:",omitempty"`
	Template   types.String      `json:",omitempty"`
	Text       types.String      `json:",omitempty"`
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
