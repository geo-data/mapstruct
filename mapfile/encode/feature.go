package encode

import "github.com/geo-data/mapfile/types/feature"

func (enc *Encoder) EncodeFeature(f *feature.Feature) (err error) {
	if err = enc.TokenStart("FEATURE"); err != nil {
		return
	}

	if err = enc.TokenStringer("WKT", f.Wkt); err != nil {
		return
	}
	if err = enc.TokenStringer("ITEMS", f.Items); err != nil {
		return
	}
	if err = enc.TokenStringer("TEXT", f.Text); err != nil {
		return
	}
	if f.Points != nil {
		if err = enc.EncodePoints(f.Points); err != nil {
			return
		}
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
