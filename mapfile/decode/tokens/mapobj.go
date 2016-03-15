package tokens

import (
	"fmt"
	"github.com/geo-data/mapfile/types/layer"
	"github.com/geo-data/mapfile/types/mapobj"
)

func (t *Tokens) Map() (m *mapobj.Map, err error) {
	token := t.Value()
	if token != "MAP" {
		err = fmt.Errorf("expected token MAP, got: %s", token)
		return
	}
	t.Next()

	m = new(mapobj.Map)

	for t != nil {
		token = t.Value()
		switch token {
		case "IMAGETYPE":
			if m.ImageType, err = t.Next().String(); err != nil {
				return
			}
		case "NAME":
			if m.Name, err = t.Next().String(); err != nil {
				return
			}
		case "STATUS":
			if m.Status, err = t.Next().Keyword(); err != nil {
				return
			}
		case "FONTSET":
			if m.Fontset, err = t.Next().String(); err != nil {
				return
			}
		case "SYMBOLSET":
			if m.Symbolset, err = t.Next().String(); err != nil {
				return
			}
		case "EXTENT":
			if m.Extent, err = t.Extent(); err != nil {
				return
			}
		case "IMAGECOLOR":
			if m.ImageColor, err = t.Color(); err != nil {
				return
			}
		case "SIZE":
			if m.Size, err = t.Size(); err != nil {
				return
			}
		case "SCALEBAR":
			if m.Scalebar, err = t.Scalebar(); err != nil {
				return
			}
		case "LEGEND":
			if m.Legend, err = t.Legend(); err != nil {
				return
			}
		case "PROJECTION":
			if m.Projection, err = t.Projection(); err != nil {
				return
			}
		case "WEB":
			if m.Web, err = t.Web(); err != nil {
				return
			}
		case "LAYER":
			var l *layer.Layer
			if l, err = t.Layer(); err != nil {
				return
			}
			m.Layers = append(m.Layers, l)
		case "END":
			return
		default:
			err = fmt.Errorf("unhandled mapfile token: %s", token)
			return
		}

		t = t.Next()
	}

	return
}
