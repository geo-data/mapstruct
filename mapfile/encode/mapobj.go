package encode

import "github.com/geo-data/mapfile/types/mapobj"

func (enc *MapfileEncoder) EncodeMap(m *mapobj.Map) (err error) {
	if err = enc.TokenStart("MAP"); err != nil {
		return
	}

	if err = enc.TokenStringer("NAME", m.Name); err != nil {
		return
	}
	if err = enc.TokenStringer("IMAGETYPE", m.ImageType); err != nil {
		return
	}
	if err = enc.TokenStringer("STATUS", m.Status); err != nil {
		return
	}
	if err = enc.TokenStringer("FONTSET", m.Fontset); err != nil {
		return
	}
	if err = enc.TokenStringer("SYMBOLSET", m.Symbolset); err != nil {
		return
	}

	if m.Extent != nil {
		if err = enc.EncodeExtent(m.Extent); err != nil {
			return
		}
	}
	if m.Size != nil {
		if err = enc.EncodeSize(m.Size); err != nil {
			return
		}
	}
	if m.ImageColor != nil {
		if err = enc.TokenStringer("IMAGECOLOR", m.ImageColor); err != nil {
			return
		}
	}
	if m.Legend != nil {
		if err = enc.EncodeLegend(m.Legend); err != nil {
			return
		}
	}
	if m.Projection != nil {
		if err = enc.EncodeProjection(m.Projection); err != nil {
			return
		}
	}
	if m.Web != nil {
		if err = enc.EncodeWeb(m.Web); err != nil {
			return
		}
	}
	if m.Scalebar != nil {
		if err = enc.EncodeScalebar(m.Scalebar); err != nil {
			return
		}
	}

	for _, layer := range m.Layers {
		if err = enc.EncodeLayer(layer); err != nil {
			return
		}
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
