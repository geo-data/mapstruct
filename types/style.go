package types

type Style struct {
	Color         *Color `json:",omitempty"`
	OutlineColor  *Color `json:",omitempty"`
	Symbol        Union  `json:",omitempty"`
	Size          Union  `json:",omitempty"`
	Width         Union  `json:",omitempty"`
	GeomTransform String `json:",omitempty"`
}
