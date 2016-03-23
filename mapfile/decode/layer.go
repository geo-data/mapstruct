package decode

import (
	"fmt"
	"github.com/geo-data/mapstruct/mapfile/decode/scanner"
	"github.com/geo-data/mapstruct/types"
)

func (t *Decoder) Layer() (layer *types.Layer, err error) {
	var token *scanner.Token
	if token, err = t.ExpectedToken(scanner.LAYER); err != nil {
		return
	}
	t.Next()

	l := new(types.Layer)
Loop:
	for t != nil {
		if token, err = t.Token(); err != nil {
			return
		}

		switch token.Type {
		case scanner.NAME:
			if l.Name, err = t.Next().String(); err != nil {
				return
			}
		case scanner.EXTENT:
			if l.Extent, err = t.Extent(); err != nil {
				return
			}
		case scanner.TYPE:
			if l.Type, err = t.Next().Keyword(); err != nil {
				return
			}
		case scanner.DEBUG:
			if l.Debug, err = t.Next().Decode(Keyword | Integer); err != nil {
				return
			}
		case scanner.PROJECTION:
			if l.Projection, err = t.Projection(); err != nil {
				return
			}
		case scanner.DATA:
			if l.Data, err = t.Next().String(); err != nil {
				return
			}
		case scanner.PROCESSING:
			if l.Processing, err = t.Next().String(); err != nil {
				return
			}
		case scanner.STATUS:
			if l.Status, err = t.Next().Keyword(); err != nil {
				return
			}
		case scanner.METADATA:
			if l.Metadata, err = t.Metadata(); err != nil {
				return
			}
		case scanner.CLASSITEM:
			if l.ClassItem, err = t.Next().Attribute(); err != nil {
				return
			}
		case scanner.LABELITEM:
			if l.LabelItem, err = t.Next().Attribute(); err != nil {
				return
			}
		case scanner.CLASS:
			var c *types.Class
			if c, err = t.Class(); err != nil {
				return
			}
			l.Classes = append(l.Classes, c)
		case scanner.FEATURE:
			var f *types.Feature
			if f, err = t.Feature(); err != nil {
				return
			}
			l.Features = append(l.Features, f)
		case scanner.END:
			break Loop
		default:
			err = fmt.Errorf("unhandled mapfile token: %s", token)
			return
		}

		t = t.Next()
	}

	layer = l
	return
}
