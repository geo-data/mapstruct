package encode

import "github.com/geo-data/mapfile/types/class"

func (enc *MapfileEncoder) EncodeClass(c *class.Class) (err error) {
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
		if err = enc.EncodeMetadata(c.Metadata); err != nil {
			return
		}
	}
	if c.Label != nil {
		if err = enc.EncodeLabel(c.Label); err != nil {
			return
		}
	}

	for _, style := range c.Styles {
		if err = enc.EncodeStyle(style); err != nil {
			return
		}
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
