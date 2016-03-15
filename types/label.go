package types

type Label struct {
	Type     Keyword  `json:",omitempty"`
	Size     Union    `json:",omitempty"`
	Font     String   `json:",omitempty"`
	Color    *Color   `json:",omitempty"`
	Position Keyword  `json:",omitempty"`
	Buffer   Uint32   `json:",omitempty"`
	Styles   []*Style `json:",omitempty"`
}
