package types

type Scalebar struct {
	Status          Keyword `json:",omitempty"`
	PostLabelCache  Keyword `json:",omitempty"`
	Style           Uint8
	Units           Keyword `json:",omitempty"`
	Size            *Size   `json:",omitempty"`
	Position        Keyword `json:",omitempty"`
	Transparent     Keyword `json:",omitempty"`
	Color           *Color  `json:",omitempty"`
	ImageColor      *Color  `json:",omitempty"`
	BackgroundColor *Color  `json:",omitempty"`
	Label           *Label  `json:",omitempty"`
}
