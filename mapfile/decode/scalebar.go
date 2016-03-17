package decode

import (
	"fmt"
	"github.com/geo-data/mapfile/types"
)

func (t *Decoder) Scalebar() (scalebar *types.Scalebar, err error) {
	token := t.Value()
	if token != "SCALEBAR" {
		err = fmt.Errorf("expected token SCALEBAR, got: %s", token)
		return
	}
	t.Next()

	s := new(types.Scalebar)
Loop:
	for t != nil {
		token := t.Value()
		switch token {
		case "STATUS":
			if s.Status, err = t.Next().Keyword(); err != nil {
				return
			}
		case "POSTLABELCACHE":
			if s.PostLabelCache, err = t.Next().Keyword(); err != nil {
				return
			}
		case "STYLE":
			if s.Style, err = t.Next().Uint8(); err != nil {
				return
			}
		case "UNITS":
			if s.Units, err = t.Next().Keyword(); err != nil {
				return
			}
		case "POSITION":
			if s.Position, err = t.Next().Keyword(); err != nil {
				return
			}
		case "TRANSPARENT":
			if s.Transparent, err = t.Next().Keyword(); err != nil {
				return
			}
		case "SIZE":
			if s.Size, err = t.Size(); err != nil {
				return
			}
		case "LABEL":
			if s.Label, err = t.Label(); err != nil {
				return
			}
		case "IMAGECOLOR":
			if s.ImageColor, err = t.Color(); err != nil {
				return
			}
		case "COLOR":
			if s.Color, err = t.Color(); err != nil {
				return
			}
		case "BACKGROUNDCOLOR":
			if s.BackgroundColor, err = t.Color(); err != nil {
				return
			}
		case "END":
			break Loop
		case "":
			if t.AtEnd() {
				err = EndOfTokens
				return
			}
			fallthrough
		default:
			err = fmt.Errorf("unhandled mapfile token: %s", token)
			return
		}

		t = t.Next()
	}

	scalebar = s
	return
}
