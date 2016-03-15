package encode

import "github.com/geo-data/mapfile/types/feature"

func (enc *Encoder) EncodeFeature(f *feature.Feature) (err error) {
	if err = enc.StartDirective("FEATURE"); err != nil {
		return
	}

	if err = enc.EncodeDirectiveStringer("WKT", f.Wkt); err != nil {
		return
	}
	if err = enc.EncodeDirectiveStringer("ITEMS", f.Items); err != nil {
		return
	}
	if err = enc.EncodeDirectiveStringer("TEXT", f.Text); err != nil {
		return
	}
	if f.Points != nil {
		if err = enc.EncodePoints(f.Points); err != nil {
			return
		}
	}

	if err = enc.EndDirective(); err != nil {
		return
	}

	return
}
