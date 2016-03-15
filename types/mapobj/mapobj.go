package mapobj

import (
	"github.com/geo-data/mapfile/mapfile/encode"
	"github.com/geo-data/mapfile/types"
	"github.com/geo-data/mapfile/types/color"
	"github.com/geo-data/mapfile/types/extent"
	"github.com/geo-data/mapfile/types/layer"
	"github.com/geo-data/mapfile/types/legend"
	"github.com/geo-data/mapfile/types/projection"
	"github.com/geo-data/mapfile/types/scalebar"
	"github.com/geo-data/mapfile/types/size"
	"github.com/geo-data/mapfile/types/web"
)

type Map struct {
	Name       types.String          `json:",omitempty"`
	Extent     *extent.Extent        `json:",omitempty"`
	ImageType  types.String          `json:",omitempty"`
	ImageColor *color.Color          `json:",omitempty"`
	Status     types.Keyword         `json:",omitempty"`
	Size       *size.Size            `json:",omitempty"`
	Fontset    types.String          `json:",omitempty"`
	Symbolset  types.String          `json:",omitempty"`
	Legend     *legend.Legend        `json:",omitempty"`
	Scalebar   *scalebar.Scalebar    `json:",omitempty"`
	Web        *web.Web              `json:",omitempty"`
	Projection projection.Projection `json:",omitempty"`
	Layers     []*layer.Layer        `json:",omitempty"`
}

func (m *Map) Encode(enc *encode.MapfileEncoder) (err error) {
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
