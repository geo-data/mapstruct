package encode

import "github.com/geo-data/mapfile/types/legend"

func (enc *Encoder) EncodeLegend(l *legend.Legend) (err error) {
	if err = enc.TokenStart("LEGEND"); err != nil {
		return
	}

	if l.ImageColor != nil {
		if err = enc.TokenStringer("IMAGECOLOR", l.ImageColor); err != nil {
			return
		}
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
