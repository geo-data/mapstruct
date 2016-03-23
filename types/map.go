package types

type Map struct {
	Name        String     `json:",omitempty"`
	Extent      *Extent    `json:",omitempty"`
	DataPattern Regex      `json:",omitempty"`
	ImageType   String     `json:",omitempty"`
	ImageColor  *Color     `json:",omitempty"`
	Status      Keyword    `json:",omitempty"`
	Size        *Size      `json:",omitempty"`
	Fontset     String     `json:",omitempty"`
	Symbolset   String     `json:",omitempty"`
	Legend      *Legend    `json:",omitempty"`
	Scalebar    *Scalebar  `json:",omitempty"`
	Web         *Web       `json:",omitempty"`
	Projection  Projection `json:",omitempty"`
	Layers      []*Layer   `json:",omitempty"`
}
