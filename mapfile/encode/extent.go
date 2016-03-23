package encode

import "github.com/geo-data/mapstruct/types"

func (enc *Encoder) EncodeExtent(e *types.Extent) error {
	return enc.EncodeDirectiveStringer("EXTENT", e)
}
