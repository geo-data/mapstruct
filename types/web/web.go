package web

import "github.com/geo-data/mapfile/types/metadata"

type Web struct {
	Metadata metadata.Metadata `json:",omitempty"`
}
