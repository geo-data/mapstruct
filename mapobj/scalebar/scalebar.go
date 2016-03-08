package scalebar

import (
	"fmt"
	"github.com/geo-data/mapfile/mapobj/color"
	"github.com/geo-data/mapfile/mapobj/label"
	"github.com/geo-data/mapfile/mapobj/size"
	"github.com/geo-data/mapfile/tokens"
)

type Scalebar struct {
	Status          string `json:",omitempty"`
	PostLabelCache  string `json:",omitempty"`
	Style           uint8
	Units           string       `json:",omitempty"`
	Size            *size.Size   `json:",omitempty"`
	Position        string       `json:",omitempty"`
	Transparent     string       `json:",omitempty"`
	Color           *color.Color `json:",omitempty"`
	ImageColor      *color.Color `json:",omitempty"`
	BackgroundColor *color.Color `json:",omitempty"`
	Label           *label.Label `json:",omitempty"`
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
