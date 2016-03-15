package encode

import "github.com/geo-data/mapfile/types/scalebar"

func (enc *Encoder) EncodeScalebar(s *scalebar.Scalebar) (err error) {
	if err = enc.TokenStart("SCALEBAR"); err != nil {
		return
	}

	if err = enc.TokenStringer("STATUS", s.Status); err != nil {
		return
	}
	if err = enc.TokenStringer("POSTLABELCACHE", s.PostLabelCache); err != nil {
		return
	}
	if err = enc.TokenStringer("STYLE", s.Style); err != nil {
		return
	}
	if err = enc.TokenStringer("UNITS", s.Units); err != nil {
		return
	}
	if s.Size != nil {
		if err = enc.EncodeSize(s.Size); err != nil {
			return
		}
	}
	if err = enc.TokenStringer("POSITION", s.Position); err != nil {
		return
	}
	if err = enc.TokenStringer("TRANSPARENT", s.Transparent); err != nil {
		return
	}
	if s.Color != nil {
		if err = enc.TokenStringer("COLOR", s.Color); err != nil {
			return
		}
	}
	if s.ImageColor != nil {
		if err = enc.TokenStringer("IMAGECOLOR", s.ImageColor); err != nil {
			return
		}
	}
	if s.BackgroundColor != nil {
		if err = enc.TokenStringer("BACKGROUNDCOLOR", s.BackgroundColor); err != nil {
			return
		}
	}
	if s.Label != nil {
		if err = enc.EncodeLabel(s.Label); err != nil {
			return
		}
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
