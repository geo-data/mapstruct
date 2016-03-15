package types

type Feature struct {
	Wkt    String `json:",omitempty"`
	Items  String `json:",omitempty"`
	Text   String `json:",omitempty"`
	Points Points `json:",omitempty"`
}
