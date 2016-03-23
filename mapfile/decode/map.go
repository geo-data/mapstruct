package decode

import (
	"fmt"
	"github.com/geo-data/mapfile/mapfile/decode/scanner"
	"github.com/geo-data/mapfile/types"
)

func (t *Decoder) Map() (map_ *types.Map, err error) {
	var token *scanner.Token
	if token, err = t.ExpectedToken(scanner.MAP); err != nil {
		return
	}
	t.Next()

	m := new(types.Map)
Loop:
	for t != nil {
		if token, err = t.Token(); err != nil {
			return
		}

		switch token.Type {
		case scanner.IMAGETYPE:
			if m.ImageType, err = t.Next().String(); err != nil {
				return
			}
		case scanner.NAME:
			if m.Name, err = t.Next().String(); err != nil {
				return
			}
		case scanner.STATUS:
			if m.Status, err = t.Next().Keyword(); err != nil {
				return
			}
		case scanner.DATAPATTERN:
			if m.DataPattern, err = t.Next().Regex(); err != nil {
				return
			}
		case scanner.FONTSET:
			if m.Fontset, err = t.Next().String(); err != nil {
				return
			}
		case scanner.SYMBOLSET:
			if m.Symbolset, err = t.Next().String(); err != nil {
				return
			}
		case scanner.EXTENT:
			if m.Extent, err = t.Extent(); err != nil {
				return
			}
		case scanner.IMAGECOLOR:
			if m.ImageColor, err = t.Color(); err != nil {
				return
			}
		case scanner.SIZE:
			if m.Size, err = t.Size(); err != nil {
				return
			}
		case scanner.SCALEBAR:
			if m.Scalebar, err = t.Scalebar(); err != nil {
				return
			}
		case scanner.LEGEND:
			if m.Legend, err = t.Legend(); err != nil {
				return
			}
		case scanner.PROJECTION:
			if m.Projection, err = t.Projection(); err != nil {
				return
			}
		case scanner.WEB:
			if m.Web, err = t.Web(); err != nil {
				return
			}
		case scanner.LAYER:
			var l *types.Layer
			if l, err = t.Layer(); err != nil {
				return
			}
			m.Layers = append(m.Layers, l)
		case scanner.END:
			break Loop
		default:
			err = fmt.Errorf("unhandled mapfile token: %s", token)
			return
		}

		t = t.Next()
	}

	map_ = m
	return
}
