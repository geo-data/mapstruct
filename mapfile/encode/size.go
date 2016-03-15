package encode

import "github.com/geo-data/mapfile/types/size"

func (enc *MapfileEncoder) EncodeSize(s *size.Size) error {
	return enc.TokenStringer("SIZE", s)
}
