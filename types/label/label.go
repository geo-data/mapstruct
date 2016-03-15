package label

import (
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
