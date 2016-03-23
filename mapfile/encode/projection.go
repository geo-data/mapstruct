package encode

import "github.com/geo-data/mapstruct/types"

func (enc *Encoder) EncodeProjection(p types.Projection) (err error) {
	if err = enc.StartDirective("PROJECTION"); err != nil {
		return
	}

	for _, param := range p {
		if err = enc.EncodeString(types.String(param)); err != nil {
			return
		}
	}

	if err = enc.EndDirective(); err != nil {
		return
	}

	return
}
