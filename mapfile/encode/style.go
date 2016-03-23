package encode

import "github.com/geo-data/mapstruct/types"

func (enc *Encoder) EncodeStyle(s *types.Style) (err error) {
	if err = enc.StartDirective("STYLE"); err != nil {
		return
	}

	if s.Color != nil {
		if err = enc.EncodeDirectiveStringer("COLOR", s.Color); err != nil {
			return
		}
	}
	if s.OutlineColor != nil {
		if err = enc.EncodeDirectiveStringer("OUTLINECOLOR", s.OutlineColor); err != nil {
			return
		}
	}
	if err = enc.EncodeDirectiveUnion("SYMBOL", s.Symbol); err != nil {
		return
	}
	if err = enc.EncodeDirectiveUnion("SIZE", s.Size); err != nil {
		return
	}
	if err = enc.EncodeDirectiveUnion("WIDTH", s.Width); err != nil {
		return
	}
	if err = enc.EncodeDirectiveStringer("GEOMTRANSFORM", s.GeomTransform); err != nil {
		return
	}

	if err = enc.EndDirective(); err != nil {
		return
	}

	return
}
