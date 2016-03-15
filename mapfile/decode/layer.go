package decode

import (
	"fmt"
	"github.com/geo-data/mapfile/types/class"
	"github.com/geo-data/mapfile/types/feature"
	"github.com/geo-data/mapfile/types/layer"
)

func (t *Decoder) Layer() (l *layer.Layer, err error) {
	token := t.Value()
	if token != "LAYER" {
		err = fmt.Errorf("expected token LAYER, got: %s", token)
		return
	}
	t.Next()

	l = new(layer.Layer)
	for t != nil {
		token := t.Value()
		switch token {
		case "NAME":
			if l.Name, err = t.Next().String(); err != nil {
				return
			}
		case "EXTENT":
			if l.Extent, err = t.Extent(); err != nil {
				return
			}
		case "TYPE":
			if l.Type, err = t.Next().Keyword(); err != nil {
				return
			}
		case "DEBUG":
			if l.Debug, err = t.Next().Decode(Keyword | Integer); err != nil {
				return
			}
		case "PROJECTION":
			if l.Projection, err = t.Projection(); err != nil {
				return
			}
		case "DATA":
			if l.Data, err = t.Next().String(); err != nil {
				return
			}
		case "PROCESSING":
			if l.Processing, err = t.Next().String(); err != nil {
				return
			}
		case "STATUS":
			if l.Status, err = t.Next().Keyword(); err != nil {
				return
			}
		case "METADATA":
			if l.Metadata, err = t.Metadata(); err != nil {
				return
			}
		case "CLASSITEM":
			if l.ClassItem, err = t.Next().Attribute(); err != nil {
				return
			}
		case "LABELITEM":
			if l.LabelItem, err = t.Next().Attribute(); err != nil {
				return
			}
		case "CLASS":
			var c *class.Class
			if c, err = t.Class(); err != nil {
				return
			}
			l.Classes = append(l.Classes, c)
		case "FEATURE":
			var f *feature.Feature
			if f, err = t.Feature(); err != nil {
				return
			}
			l.Features = append(l.Features, f)
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
