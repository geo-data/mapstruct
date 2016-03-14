package mapobj

import (
	"fmt"

	"github.com/geo-data/mapfile/encoding"
	"github.com/geo-data/mapfile/mapobj/color"
	"github.com/geo-data/mapfile/mapobj/extent"
	"github.com/geo-data/mapfile/mapobj/layer"
	"github.com/geo-data/mapfile/mapobj/legend"
	"github.com/geo-data/mapfile/mapobj/projection"
	"github.com/geo-data/mapfile/mapobj/scalebar"
	"github.com/geo-data/mapfile/mapobj/size"
	"github.com/geo-data/mapfile/mapobj/web"
	"github.com/geo-data/mapfile/tokens"
)

type Map struct {
	Name       tokens.String          `json:",omitempty"`
	Extent     *extent.Extent         `json:",omitempty"`
	ImageType  tokens.String          `json:",omitempty"`
	ImageColor *color.Color           `json:",omitempty"`
	Status     tokens.Keyword         `json:",omitempty"`
	Size       *size.Size             `json:",omitempty"`
	Fontset    tokens.String          `json:",omitempty"`
	Symbolset  tokens.String          `json:",omitempty"`
	Legend     *legend.Legend         `json:",omitempty"`
	Scalebar   *scalebar.Scalebar     `json:",omitempty"`
	Web        *web.Web               `json:",omitempty"`
	Projection *projection.Projection `json:",omitempty"`
	Layers     []*layer.Layer         `json:",omitempty"`
}

func New(toks *tokens.Tokens) (m *Map, err error) {
	token := toks.Value()
	if token != "MAP" {
		err = fmt.Errorf("expected token MAP, got: %s", token)
		return
	}
	toks.Next()

	m = new(Map)

	for toks != nil {
		token = toks.Value()
		switch token {
		case "IMAGETYPE":
			if m.ImageType, err = toks.Next().String(); err != nil {
				return
			}
		case "NAME":
			if m.Name, err = toks.Next().String(); err != nil {
				return
			}
		case "STATUS":
			if m.Status, err = toks.Next().Keyword(); err != nil {
				return
			}
		case "FONTSET":
			if m.Fontset, err = toks.Next().String(); err != nil {
				return
			}
		case "SYMBOLSET":
			if m.Symbolset, err = toks.Next().String(); err != nil {
				return
			}
		case "EXTENT":
			if m.Extent, err = extent.New(toks); err != nil {
				return
			}
		case "IMAGECOLOR":
			if m.ImageColor, err = color.New(toks); err != nil {
				return
			}
		case "SIZE":
			if m.Size, err = size.New(toks); err != nil {
				return
			}
		case "SCALEBAR":
			if m.Scalebar, err = scalebar.New(toks); err != nil {
				return
			}
		case "LEGEND":
			if m.Legend, err = legend.New(toks); err != nil {
				return
			}
		case "PROJECTION":
			if m.Projection, err = projection.New(toks); err != nil {
				return
			}
		case "WEB":
			if m.Web, err = web.New(toks); err != nil {
				return
			}
		case "LAYER":
			var l *layer.Layer
			if l, err = layer.New(toks); err != nil {
				return
			}
			m.Layers = append(m.Layers, l)
		case "END":
			return
		default:
			err = fmt.Errorf("unhandled mapfile token: %s", token)
			return
		}

		toks = toks.Next()
	}

	return
}

func (m *Map) Encode(enc *encoding.MapfileEncoder) (err error) {
	if err = enc.TokenStart("MAP"); err != nil {
		return
	}

	if err = enc.TokenStringer("NAME", m.Name); err != nil {
		return
	}
	if err = enc.TokenStringer("IMAGETYPE", m.ImageType); err != nil {
		return
	}
	if err = enc.TokenStringer("STATUS", m.Status); err != nil {
		return
	}
	if err = enc.TokenStringer("FONTSET", m.Fontset); err != nil {
		return
	}
	if err = enc.TokenStringer("SYMBOLSET", m.Symbolset); err != nil {
		return
	}

	if m.Extent != nil {
		if err = m.Extent.Encode(enc); err != nil {
			return
		}
	}
	if m.Size != nil {
		if err = m.Size.Encode(enc); err != nil {
			return
		}
	}
	if m.ImageColor != nil {
		if err = enc.TokenStringer("IMAGECOLOR", m.ImageColor); err != nil {
			return
		}
	}
	if m.Legend != nil {
		if err = m.Legend.Encode(enc); err != nil {
			return
		}
	}
	if m.Projection != nil {
		if err = m.Projection.Encode(enc); err != nil {
			return
		}
	}
	if m.Web != nil {
		if err = m.Web.Encode(enc); err != nil {
			return
		}
	}
	if m.Scalebar != nil {
		if err = m.Scalebar.Encode(enc); err != nil {
			return
		}
	}

	for _, layer := range m.Layers {
		if err = layer.Encode(enc); err != nil {
			return
		}
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
