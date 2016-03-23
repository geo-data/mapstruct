package decode

import (
	"fmt"
	"github.com/geo-data/mapfile/mapfile/decode/scanner"
	"github.com/geo-data/mapfile/types"
)

func (t *Decoder) Scalebar() (scalebar *types.Scalebar, err error) {
	var token *scanner.Token
	if token, err = t.ExpectedToken(scanner.SCALEBAR); err != nil {
		return
	}
	t.Next()

	s := new(types.Scalebar)
Loop:
	for t != nil {
		if token, err = t.Token(); err != nil {
			return
		}

		switch token.Type {
		case scanner.STATUS:
			if s.Status, err = t.Next().Keyword(); err != nil {
				return
			}
		case scanner.POSTLABELCACHE:
			if s.PostLabelCache, err = t.Next().Keyword(); err != nil {
				return
			}
		case scanner.STYLE:
			if s.Style, err = t.Next().Uint8(); err != nil {
				return
			}
		case scanner.UNITS:
			if s.Units, err = t.Next().Keyword(); err != nil {
				return
			}
		case scanner.POSITION:
			if s.Position, err = t.Next().Keyword(); err != nil {
				return
			}
		case scanner.TRANSPARENT:
			if s.Transparent, err = t.Next().Keyword(); err != nil {
				return
			}
		case scanner.SIZE:
			if s.Size, err = t.Size(); err != nil {
				return
			}
		case scanner.LABEL:
			if s.Label, err = t.Label(); err != nil {
				return
			}
		case scanner.IMAGECOLOR:
			if s.ImageColor, err = t.Color(); err != nil {
				return
			}
		case scanner.COLOR:
			if s.Color, err = t.Color(); err != nil {
				return
			}
		case scanner.BACKGROUNDCOLOR:
			if s.BackgroundColor, err = t.Color(); err != nil {
				return
			}
		case scanner.END:
			break Loop
		default:
			err = fmt.Errorf("unhandled mapfile token: %s", token)
			return
		}

		t = t.Next()
	}

	scalebar = s
	return
}
