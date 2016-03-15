package scalebar

import (
	"github.com/geo-data/mapfile/mapfile/encode"
	"github.com/geo-data/mapfile/types"
	"github.com/geo-data/mapfile/types/color"
	"github.com/geo-data/mapfile/types/label"
	"github.com/geo-data/mapfile/types/size"
)

type Scalebar struct {
	Status          types.Keyword `json:",omitempty"`
	PostLabelCache  types.Keyword `json:",omitempty"`
	Style           types.Uint8
	Units           types.Keyword `json:",omitempty"`
	Size            *size.Size    `json:",omitempty"`
	Position        types.Keyword `json:",omitempty"`
	Transparent     types.Keyword `json:",omitempty"`
	Color           *color.Color  `json:",omitempty"`
	ImageColor      *color.Color  `json:",omitempty"`
	BackgroundColor *color.Color  `json:",omitempty"`
	Label           *label.Label  `json:",omitempty"`
}

func (s *Scalebar) Encode(enc *encode.MapfileEncoder) (err error) {
	if err = enc.TokenStart("SCALEBAR"); err != nil {
		return
	}

	if err = enc.TokenStringer("STATUS", s.Status); err != nil {
		return
	}
	if err = enc.TokenStringer("POSTLABELCACHE", s.PostLabelCache); err != nil {
		return
	}
	if err = enc.TokenStringer("STYLE", s.Style); err != nil {
		return
	}
	if err = enc.TokenStringer("UNITS", s.Units); err != nil {
		return
	}
	if s.Size != nil {
		if err = s.Size.Encode(enc); err != nil {
			return
		}
	}
	if err = enc.TokenStringer("POSITION", s.Position); err != nil {
		return
	}
	if err = enc.TokenStringer("TRANSPARENT", s.Transparent); err != nil {
		return
	}
	if s.Color != nil {
		if err = enc.TokenStringer("COLOR", s.Color); err != nil {
			return
		}
	}
	if s.ImageColor != nil {
		if err = enc.TokenStringer("IMAGECOLOR", s.ImageColor); err != nil {
			return
		}
	}
	if s.BackgroundColor != nil {
		if err = enc.TokenStringer("BACKGROUNDCOLOR", s.BackgroundColor); err != nil {
			return
		}
	}
	if s.Label != nil {
		if err = s.Label.Encode(enc); err != nil {
			return
		}
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
