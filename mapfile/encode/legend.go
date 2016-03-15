package encode

import "github.com/geo-data/mapfile/types"

func (enc *Encoder) EncodeLegend(l *types.Legend) (err error) {
	if err = enc.StartDirective("LEGEND"); err != nil {
		return
	}

	if l.ImageColor != nil {
		if err = enc.EncodeDirectiveStringer("IMAGECOLOR", l.ImageColor); err != nil {
			return
		}
	}

	if err = enc.EndDirective(); err != nil {
		return
	}

	return
}
