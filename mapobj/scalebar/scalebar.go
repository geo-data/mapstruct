package scalebar

import (
	"fmt"
	"github.com/geo-data/mapfile/mapobj/color"
	"github.com/geo-data/mapfile/mapobj/label"
	"github.com/geo-data/mapfile/mapobj/size"
	"github.com/geo-data/mapfile/tokens"
	"strconv"
)

type Scalebar struct {
	Status          string
	PostLabelCache  string
	Style           uint
	Units           string
	Size            size.Size
	Position        string
	Transparent     string
	Color           color.Color
	ImageColor      color.Color
	BackgroundColor color.Color
	Label           label.Label
}

func (s *Scalebar) FromTokens(tokens *tokens.Tokens) error {
	token := tokens.Value()
	if token != "SCALEBAR" {
		return fmt.Errorf("expected token SCALEBAR, got: %s", token)
	}
	tokens.Next()

	for tokens != nil {
		token := tokens.Value()
		switch token {
		case "STATUS":
			s.Status = tokens.Next().Value()
		case "POSTLABELCACHE":
			s.PostLabelCache = tokens.Next().Value()
		case "STYLE":
			i, err := strconv.ParseUint(tokens.Next().Value(), 10, 0)
			if err != nil {
				return err
			}
			s.Style = uint(i)
		case "UNITS":
			s.Units = tokens.Next().Value()
		case "POSITION":
			s.Position = tokens.Next().Value()
		case "TRANSPARENT":
			s.Transparent = tokens.Next().Value()
		case "SIZE":
			if err := s.Size.FromTokens(tokens); err != nil {
				return err
			}
		case "LABEL":
			if err := s.Label.FromTokens(tokens); err != nil {
				return err
			}
		case "IMAGECOLOR":
			if err := s.ImageColor.FromTokens(tokens); err != nil {
				return err
			}
		case "COLOR":
			if err := s.Color.FromTokens(tokens); err != nil {
				return err
			}
		case "BACKGROUNDCOLOR":
			if err := s.BackgroundColor.FromTokens(tokens); err != nil {
				return err
			}
		case "END":
			return nil
		default:
			return fmt.Errorf("unhandled mapfile token: %s", token)
		}

		tokens = tokens.Next()
	}

	return nil
}
