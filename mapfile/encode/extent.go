package encode

import "github.com/geo-data/mapfile/types"

func (enc *Encoder) EncodeExtent(e *types.Extent) error {
	return enc.EncodeDirectiveStringer("EXTENT", e)
}
