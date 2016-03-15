package scalebar

import (
	"fmt"
	"github.com/geo-data/mapfile/mapfile/encode"
	"github.com/geo-data/mapfile/mapfile/decode/tokens"
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

func New(toks *tokens.Tokens) (s *Scalebar, err error) {
	token := toks.Value()
	if token != "SCALEBAR" {
		err = fmt.Errorf("expected token SCALEBAR, got: %s", token)
		return
	}
	toks.Next()

	s = new(Scalebar)
	for toks != nil {
		token := toks.Value()
		switch token {
		case "STATUS":
			if s.Status, err = toks.Next().Keyword(); err != nil {
				return
			}
		case "POSTLABELCACHE":
			if s.PostLabelCache, err = toks.Next().Keyword(); err != nil {
				return
			}
		case "STYLE":
			if s.Style, err = toks.Next().Uint8(); err != nil {
				return
			}
		case "UNITS":
			if s.Units, err = toks.Next().Keyword(); err != nil {
				return
			}
		case "POSITION":
			if s.Position, err = toks.Next().Keyword(); err != nil {
				return
			}
		case "TRANSPARENT":
			if s.Transparent, err = toks.Next().Keyword(); err != nil {
				return
			}
		case "SIZE":
			if s.Size, err = size.New(toks); err != nil {
				return
			}
		case "LABEL":
			if s.Label, err = label.New(toks); err != nil {
				return
			}
		case "IMAGECOLOR":
			if s.ImageColor, err = color.New(toks); err != nil {
				return
			}
		case "COLOR":
			if s.Color, err = color.New(toks); err != nil {
				return
			}
		case "BACKGROUNDCOLOR":
			if s.BackgroundColor, err = color.New(toks); err != nil {
				return
			}
		case "END":
			return
		default:
			err = fmt.Errorf("unhandled mapfile token: %s", token)
			return
		}

		toks = toks.Next()
	}

	return
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
