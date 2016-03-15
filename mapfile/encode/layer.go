package encode

import "github.com/geo-data/mapfile/types/layer"

func (enc *Encoder) EncodeLayer(l *layer.Layer) (err error) {
	if err = enc.TokenStart("LAYER"); err != nil {
		return
	}

	if err = enc.TokenStringer("NAME", l.Name); err != nil {
		return
	}
	if l.Extent != nil {
		if err = enc.EncodeExtent(l.Extent); err != nil {
			return
		}
	}
	if err = enc.TokenStringer("TYPE", l.Type); err != nil {
		return
	}
	if err = enc.TokenUnion("DEBUG", l.Debug); err != nil {
		return
	}
	if l.Projection != nil {
		if err = enc.EncodeProjection(l.Projection); err != nil {
			return
		}
	}
	if err = enc.TokenStringer("DATA", l.Data); err != nil {
		return
	}
	if err = enc.TokenStringer("PROCESSING", l.Processing); err != nil {
		return
	}
	if err = enc.TokenStringer("STATUS", l.Status); err != nil {
		return
	}
	if l.Metadata != nil {
		if err = enc.EncodeMetadata(l.Metadata); err != nil {
			return
		}
	}
	if err = enc.TokenString("CLASSITEM", l.ClassItem.QuotedString()); err != nil {
		return
	}
	if err = enc.TokenString("LABELITEM", l.LabelItem.QuotedString()); err != nil {
		return
	}

	for _, class := range l.Classes {
		if err = enc.EncodeClass(class); err != nil {
			return
		}
	}

	for _, feature := range l.Features {
		if err = enc.EncodeFeature(feature); err != nil {
			return
		}
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
