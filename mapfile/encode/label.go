package encode

import "github.com/geo-data/mapfile/types/label"

func (enc *MapfileEncoder) EncodeLabel(l *label.Label) (err error) {
	if err = enc.TokenStart("LABEL"); err != nil {
		return
	}

	if err = enc.TokenStringer("TYPE", l.Type); err != nil {
		return
	}
	if err = enc.TokenUnion("SIZE", l.Size); err != nil {
		return
	}
	if err = enc.TokenStringer("FONT", l.Font); err != nil {
		return
	}
	if l.Color != nil {
		if err = enc.TokenStringer("COLOR", l.Color); err != nil {
			return
		}
	}
	if err = enc.TokenStringer("POSITION", l.Position); err != nil {
		return
	}
	if uint32(l.Buffer) > uint32(0) {
		if err = enc.TokenStringer("BUFFER", l.Buffer); err != nil {
			return
		}
	}

	for _, style := range l.Styles {
		if err = enc.EncodeStyle(style); err != nil {
			return
		}
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
