package legend

import (
	"github.com/geo-data/mapfile/mapfile/encode"
	"github.com/geo-data/mapfile/types/color"
)

type Legend struct {
	ImageColor *color.Color `json:",omitempty"`
}

func (l *Legend) Encode(enc *encode.MapfileEncoder) (err error) {
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
