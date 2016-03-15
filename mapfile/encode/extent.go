package encode

import "github.com/geo-data/mapfile/types/extent"

func (enc *MapfileEncoder) EncodeExtent(e *extent.Extent) error {
	return enc.TokenStringer("EXTENT", e)
}
