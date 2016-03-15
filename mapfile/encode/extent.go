package encode

import "github.com/geo-data/mapfile/types/extent"

func (enc *Encoder) EncodeExtent(e *extent.Extent) error {
	return enc.EncodeDirectiveStringer("EXTENT", e)
}
