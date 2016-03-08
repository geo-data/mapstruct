package scalebar

import (
	"fmt"
	"github.com/geo-data/mapfile/encoding"
	"github.com/geo-data/mapfile/mapobj/color"
	"github.com/geo-data/mapfile/mapobj/label"
	"github.com/geo-data/mapfile/mapobj/size"
	"github.com/geo-data/mapfile/tokens"
)

type Scalebar struct {
	Status          tokens.String `json:",omitempty"`
	PostLabelCache  tokens.String `json:",omitempty"`
	Style           tokens.Uint8
	Units           tokens.String `json:",omitempty"`
	Size            *size.Size    `json:",omitempty"`
	Position        tokens.String `json:",omitempty"`
	Transparent     tokens.String `json:",omitempty"`
	Color           *color.Color  `json:",omitempty"`
	ImageColor      *color.Color  `json:",omitempty"`
	BackgroundColor *color.Color  `json:",omitempty"`
	Label           *label.Label  `json:",omitempty"`
}

func New(tokens *tokens.Tokens) (s *Scalebar, err error) {
	token := tokens.Value()
	if token != "SCALEBAR" {
		err = fmt.Errorf("expected token SCALEBAR, got: %s", token)
		return
	}
	tokens.Next()

	s = new(Scalebar)
	for tokens != nil {
		token := tokens.Value()
		switch token {
		case "STATUS":
			s.Status = tokens.Next().Value()
		case "POSTLABELCACHE":
			s.PostLabelCache = tokens.Next().Value()
		case "STYLE":
			if s.Style, err = tokens.Next().Uint8(); err != nil {
				return
			}
		case "UNITS":
			s.Units = tokens.Next().Value()
		case "POSITION":
			s.Position = tokens.Next().Value()
		case "TRANSPARENT":
			s.Transparent = tokens.Next().Value()
		case "SIZE":
			if s.Size, err = size.New(tokens); err != nil {
				return
			}
		case "LABEL":
			if s.Label, err = label.New(tokens); err != nil {
				return
			}
		case "IMAGECOLOR":
			if s.ImageColor, err = color.New(tokens); err != nil {
				return
			}
		case "COLOR":
			if s.Color, err = color.New(tokens); err != nil {
				return
			}
		case "BACKGROUNDCOLOR":
			if s.BackgroundColor, err = color.New(tokens); err != nil {
				return
			}
		case "END":
			return
		default:
			err = fmt.Errorf("unhandled mapfile token: %s", token)
			return
		}

		tokens = tokens.Next()
	}

	return
}

func (s *Scalebar) Encode(enc *encoding.MapfileEncoder) (err error) {
	if err = enc.TokenStart("SCALEBAR"); err != nil {
		return
	}

	if err = enc.TokenString("STATUS", s.Status); err != nil {
		return
	}
	if err = enc.TokenString("POSTLABELCACHE", s.PostLabelCache); err != nil {
		return
	}
	if err = enc.TokenValue("STYLE", s.Style); err != nil {
		return
	}
	if err = enc.TokenString("UNITS", s.Units); err != nil {
		return
	}
	if s.Size != nil {
		if err = s.Size.Encode(enc); err != nil {
			return
		}
	}
	if err = enc.TokenString("POSITION", s.Position); err != nil {
		return
	}
	if err = enc.TokenString("TRANSPARENT", s.Transparent); err != nil {
		return
	}
	if s.Color != nil {
		if err = enc.TokenValue("COLOR", s.Color); err != nil {
			return
		}
	}
	if s.ImageColor != nil {
		if err = enc.TokenValue("IMAGECOLOR", s.ImageColor); err != nil {
			return
		}
	}
	if s.BackgroundColor != nil {
		if err = enc.TokenValue("BACKGROUNDCOLOR", s.BackgroundColor); err != nil {
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
