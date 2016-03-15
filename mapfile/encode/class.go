package encode

import "github.com/geo-data/mapfile/types"

func (enc *Encoder) EncodeClass(c *types.Class) (err error) {
	if err = enc.StartDirective("CLASS"); err != nil {
		return
	}

	if err = enc.EncodeDirectiveStringer("NAME", c.Name); err != nil {
		return
	}
	if err = enc.EncodeDirectiveStringer("EXPRESSION", c.Expression); err != nil {
		return
	}
	if err = enc.EncodeDirectiveStringer("TEMPLATE", c.Template); err != nil {
		return
	}
	if err = enc.EncodeDirectiveStringer("TEXT", c.Text); err != nil {
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

	if err = enc.EndDirective(); err != nil {
		return
	}

	return
}
