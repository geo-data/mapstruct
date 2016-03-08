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
	Name       string                 `json:",omitempty"`
	Extent     *extent.Extent         `json:",omitempty"`
	ImageType  string                 `json:",omitempty"`
	ImageColor *color.Color           `json:",omitempty"`
	Status     string                 `json:",omitempty"`
	Size       *size.Size             `json:",omitempty"`
	Fontset    string                 `json:",omitempty"`
	Symbolset  string                 `json:",omitempty"`
	Legend     *legend.Legend         `json:",omitempty"`
	Scalebar   *scalebar.Scalebar     `json:",omitempty"`
	Web        *web.Web               `json:",omitempty"`
	Projection *projection.Projection `json:",omitempty"`
}

func New(tokens *tokens.Tokens) (m *Map, err error) {
	token := tokens.Value()
	if token != "MAP" {
		err = fmt.Errorf("expected token MAP, got: %s", token)
		return
	}
	tokens.Next()

	m = new(Map)

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
			if m.Extent, err = extent.New(tokens); err != nil {
				return
			}
		case "IMAGECOLOR":
			if m.ImageColor, err = color.New(tokens); err != nil {
				return
			}
		case "SIZE":
			if m.Size, err = size.New(tokens); err != nil {
				return
			}
		case "SCALEBAR":
			if m.Scalebar, err = scalebar.New(tokens); err != nil {
				return
			}
		case "LEGEND":
			if m.Legend, err = legend.New(tokens); err != nil {
				return
			}
		case "PROJECTION":
			if m.Projection, err = projection.New(tokens); err != nil {
				return
			}
		case "WEB":
			if m.Web, err = web.New(tokens); err != nil {
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
