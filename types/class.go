package types

type Class struct {
	Name       String   `json:",omitempty"`
	Expression String   `json:",omitempty"`
	Metadata   Metadata `json:",omitempty"`
	Styles     []*Style `json:",omitempty"`
	Label      *Label   `json:",omitempty"`
	Template   String   `json:",omitempty"`
	Text       String   `json:",omitempty"`
}
