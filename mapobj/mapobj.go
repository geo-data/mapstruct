package mapobj

import (
	"fmt"

	"github.com/geo-data/mapfile/mapobj/color"
	"github.com/geo-data/mapfile/mapobj/extent"
	"github.com/geo-data/mapfile/mapobj/legend"
	"github.com/geo-data/mapfile/mapobj/projection"
	"github.com/geo-data/mapfile/mapobj/scalebar"
	"github.com/geo-data/mapfile/mapobj/size"
	"github.com/geo-data/mapfile/mapobj/web"
	"github.com/geo-data/mapfile/tokens"
)

type Map struct {
	Name       string
	Extent     extent.Extent
	ImageType  string
	ImageColor color.Color
	Status     string
	Size       size.Size
	Fontset    string
	Symbolset  string
	Legend     legend.Legend
	Scalebar   scalebar.Scalebar
	Web        web.Web
	Projection *projection.Projection `json:",omitempty"`
}

func (m *Map) FromTokens(tokens *tokens.Tokens) error {
	token := tokens.Value()
	if token != "MAP" {
		return fmt.Errorf("expected token MAP, got: %s", token)
	}
	tokens.Next()

	for tokens != nil {
		token = tokens.Value()
		switch token {
		case "IMAGETYPE":
			m.ImageType = tokens.Next().Value()
		case "NAME":
			m.Name = tokens.Next().Value()
		case "STATUS":
			m.Status = tokens.Next().Value()
		case "FONTSET":
			m.Fontset = tokens.Next().Value()
		case "SYMBOLSET":
			m.Symbolset = tokens.Next().Value()
		case "EXTENT":
			if err := m.Extent.FromTokens(tokens); err != nil {
				return err
			}
		case "IMAGECOLOR":
			if err := m.ImageColor.FromTokens(tokens); err != nil {
				return err
			}
		case "SIZE":
			if err := m.Size.FromTokens(tokens); err != nil {
				return err
			}
		case "SCALEBAR":
			if err := m.Scalebar.FromTokens(tokens); err != nil {
				return err
			}
		case "LEGEND":
			if err := m.Legend.FromTokens(tokens); err != nil {
				return err
			}
		case "PROJECTION":
			p := new(projection.Projection)
			if err := p.FromTokens(tokens); err != nil {
				return err
			}
			m.Projection = p
		case "WEB":
			if err := m.Web.FromTokens(tokens); err != nil {
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
