package types

type Layer struct {
	Name       String     `json:",omitempty"`
	Extent     *Extent    `json:",omitempty"`
	Type       Keyword    `json:",omitempty"`
	Debug      Union      `json:",omitempty"`
	Projection Projection `json:",omitempty"`
	Data       String     `json:",omitempty"`
	Processing String     `json:",omitempty"`
	Status     Keyword    `json:",omitempty"`
	Metadata   Metadata   `json:",omitempty"`
	ClassItem  Attribute  `json:",omitempty"`
	LabelItem  Attribute  `json:",omitempty"`
	Classes    []*Class   `json:",omitempty"`
	Features   []*Feature `json:",omitempty"`
}
