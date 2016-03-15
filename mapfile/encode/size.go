package encode

import "github.com/geo-data/mapfile/types/size"

func (enc *Encoder) EncodeSize(s *size.Size) error {
	return enc.EncodeDirectiveStringer("SIZE", s)
}
