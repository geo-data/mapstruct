package encode

import "github.com/geo-data/mapfile/types/style"

func (enc *Encoder) EncodeStyle(s *style.Style) (err error) {
	if err = enc.TokenStart("STYLE"); err != nil {
		return
	}

	if s.Color != nil {
		if err = enc.TokenStringer("COLOR", s.Color); err != nil {
			return
		}
	}
	if s.OutlineColor != nil {
		if err = enc.TokenStringer("OUTLINECOLOR", s.OutlineColor); err != nil {
			return
		}
	}
	if err = enc.TokenUnion("SYMBOL", s.Symbol); err != nil {
		return
	}
	if err = enc.TokenUnion("SIZE", s.Size); err != nil {
		return
	}
	if err = enc.TokenUnion("WIDTH", s.Width); err != nil {
		return
	}
	if err = enc.TokenStringer("GEOMTRANSFORM", s.GeomTransform); err != nil {
		return
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
