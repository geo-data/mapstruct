package encode

import "github.com/geo-data/mapstruct/types"

func (enc *Encoder) EncodeSize(s *types.Size) error {
	return enc.EncodeDirectiveStringer("SIZE", s)
}
