package metadata

type Metadata map[string]string

func New() Metadata {
	return Metadata(make(map[string]string))
}
