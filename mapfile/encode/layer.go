package encode

import "github.com/geo-data/mapfile/types"

func (enc *Encoder) EncodeLayer(l *types.Layer) (err error) {
	if err = enc.StartDirective("LAYER"); err != nil {
		return
	}

	if err = enc.EncodeDirectiveStringer("NAME", l.Name); err != nil {
		return
	}
	if l.Extent != nil {
		if err = enc.EncodeExtent(l.Extent); err != nil {
			return
		}
	}
	if err = enc.EncodeDirectiveStringer("TYPE", l.Type); err != nil {
		return
	}
	if err = enc.EncodeDirectiveUnion("DEBUG", l.Debug); err != nil {
		return
	}
	if l.Projection != nil {
		if err = enc.EncodeProjection(l.Projection); err != nil {
			return
		}
	}
	if err = enc.EncodeDirectiveStringer("DATA", l.Data); err != nil {
		return
	}
	if err = enc.EncodeDirectiveStringer("PROCESSING", l.Processing); err != nil {
		return
	}
	if err = enc.EncodeDirectiveStringer("STATUS", l.Status); err != nil {
		return
	}
	if l.Metadata != nil {
		if err = enc.EncodeMetadata(l.Metadata); err != nil {
			return
		}
	}
	if err = enc.EncodeDirectiveString("CLASSITEM", l.ClassItem.QuotedString()); err != nil {
		return
	}
	if err = enc.EncodeDirectiveString("LABELITEM", l.LabelItem.QuotedString()); err != nil {
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

	if err = enc.EndDirective(); err != nil {
		return
	}

	return
}
