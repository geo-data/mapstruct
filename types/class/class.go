package class

import (
	"github.com/geo-data/mapfile/types"
	"github.com/geo-data/mapfile/types/label"
	"github.com/geo-data/mapfile/types/metadata"
	"github.com/geo-data/mapfile/types/style"
)

type Class struct {
	Name       types.String      `json:",omitempty"`
	Expression types.String      `json:",omitempty"`
	Metadata   metadata.Metadata `json:",omitempty"`
	Styles     []*style.Style    `json:",omitempty"`
	Label      *label.Label      `json:",omitempty"`
	Template   types.String      `json:",omitempty"`
	Text       types.String      `json:",omitempty"`
}
