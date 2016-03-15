package label

import (
	"github.com/geo-data/mapfile/mapfile/encode"
	"github.com/geo-data/mapfile/types"
	"github.com/geo-data/mapfile/types/color"
	"github.com/geo-data/mapfile/types/style"
)

type Label struct {
	Type     types.Keyword  `json:",omitempty"`
	Size     types.Union    `json:",omitempty"`
	Font     types.String   `json:",omitempty"`
	Color    *color.Color   `json:",omitempty"`
	Position types.Keyword  `json:",omitempty"`
	Buffer   types.Uint32   `json:",omitempty"`
	Styles   []*style.Style `json:",omitempty"`
}

func (l *Label) Encode(enc *encode.MapfileEncoder) (err error) {
	if err = enc.TokenStart("LABEL"); err != nil {
		return
	}

	if err = enc.TokenStringer("TYPE", l.Type); err != nil {
		return
	}
	if err = enc.TokenUnion("SIZE", l.Size); err != nil {
		return
	}
	if err = enc.TokenStringer("FONT", l.Font); err != nil {
		return
	}
	if l.Color != nil {
		if err = enc.TokenStringer("COLOR", l.Color); err != nil {
			return
		}
	}
	if err = enc.TokenStringer("POSITION", l.Position); err != nil {
		return
	}
	if uint32(l.Buffer) > uint32(0) {
		if err = enc.TokenStringer("BUFFER", l.Buffer); err != nil {
			return
		}
	}

	for _, style := range l.Styles {
		if err = style.Encode(enc); err != nil {
			return
		}
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
