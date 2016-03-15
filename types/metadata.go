package types

type Metadata map[string]string

func NewMetadata() Metadata {
	return Metadata(make(map[string]string))
}
