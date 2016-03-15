package encode

import "github.com/geo-data/mapfile/types/scalebar"

func (enc *Encoder) EncodeScalebar(s *scalebar.Scalebar) (err error) {
	if err = enc.StartDirective("SCALEBAR"); err != nil {
		return
	}

	if err = enc.EncodeDirectiveStringer("STATUS", s.Status); err != nil {
		return
	}
	if err = enc.EncodeDirectiveStringer("POSTLABELCACHE", s.PostLabelCache); err != nil {
		return
	}
	if err = enc.EncodeDirectiveStringer("STYLE", s.Style); err != nil {
		return
	}
	if err = enc.EncodeDirectiveStringer("UNITS", s.Units); err != nil {
		return
	}
	if s.Size != nil {
		if err = enc.EncodeSize(s.Size); err != nil {
			return
		}
	}
	if err = enc.EncodeDirectiveStringer("POSITION", s.Position); err != nil {
		return
	}
	if err = enc.EncodeDirectiveStringer("TRANSPARENT", s.Transparent); err != nil {
		return
	}
	if s.Color != nil {
		if err = enc.EncodeDirectiveStringer("COLOR", s.Color); err != nil {
			return
		}
	}
	if s.ImageColor != nil {
		if err = enc.EncodeDirectiveStringer("IMAGECOLOR", s.ImageColor); err != nil {
			return
		}
	}
	if s.BackgroundColor != nil {
		if err = enc.EncodeDirectiveStringer("BACKGROUNDCOLOR", s.BackgroundColor); err != nil {
			return
		}
	}
	if s.Label != nil {
		if err = enc.EncodeLabel(s.Label); err != nil {
			return
		}
	}

	if err = enc.EndDirective(); err != nil {
		return
	}

	return
}
