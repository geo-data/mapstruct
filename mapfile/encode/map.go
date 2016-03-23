package encode

import "github.com/geo-data/mapfile/types"

func (enc *Encoder) EncodeMap(m *types.Map) (err error) {
	if err = enc.StartDirective("MAP"); err != nil {
		return
	}

	if err = enc.EncodeDirectiveStringer("NAME", m.Name); err != nil {
		return
	}
	if err = enc.EncodeDirectiveStringer("DATAPATTERN", m.DataPattern); err != nil {
		return
	}
	if err = enc.EncodeDirectiveStringer("IMAGETYPE", m.ImageType); err != nil {
		return
	}
	if err = enc.EncodeDirectiveStringer("STATUS", m.Status); err != nil {
		return
	}
	if err = enc.EncodeDirectiveStringer("FONTSET", m.Fontset); err != nil {
		return
	}
	if err = enc.EncodeDirectiveStringer("SYMBOLSET", m.Symbolset); err != nil {
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
		if err = enc.EncodeDirectiveStringer("IMAGECOLOR", m.ImageColor); err != nil {
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

	if err = enc.EndDirective(); err != nil {
		return
	}

	return
}
