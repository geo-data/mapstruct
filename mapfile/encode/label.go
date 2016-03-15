package encode

import "github.com/geo-data/mapfile/types"

func (enc *Encoder) EncodeLabel(l *types.Label) (err error) {
	if err = enc.StartDirective("LABEL"); err != nil {
		return
	}

	if err = enc.EncodeDirectiveStringer("TYPE", l.Type); err != nil {
		return
	}
	if err = enc.EncodeDirectiveUnion("SIZE", l.Size); err != nil {
		return
	}
	if err = enc.EncodeDirectiveStringer("FONT", l.Font); err != nil {
		return
	}
	if l.Color != nil {
		if err = enc.EncodeDirectiveStringer("COLOR", l.Color); err != nil {
			return
		}
	}
	if err = enc.EncodeDirectiveStringer("POSITION", l.Position); err != nil {
		return
	}
	if uint32(l.Buffer) > uint32(0) {
		if err = enc.EncodeDirectiveStringer("BUFFER", l.Buffer); err != nil {
			return
		}
	}

	for _, style := range l.Styles {
		if err = enc.EncodeStyle(style); err != nil {
			return
		}
	}

	if err = enc.EndDirective(); err != nil {
		return
	}

	return
}
