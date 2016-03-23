package types

// Class represents a MapServer CLASS object (http://mapserver.org/mapfile/class.html#class).
type Class struct {
	// Name to use in legends for this class.  If not set class won’t show up in
	// legend.
	Name String `json:",omitempty"`
	// Expression defines which class a feature belongs to.  If no expression is
	// given, then all features are said to belong to this class.
	Expression Union    `json:",omitempty"`
	Metadata   Metadata `json:",omitempty"`
	// Styles contains styles applied to the class. A class can contain multiple
	// styles. Multiple styles can be used create complex symbols (by
	// overlay/stacking). See
	// <http://mapserver.org/mapfile/symbology/construction.html#sym-construction>
	// for more info.
	Styles []*Style `json:",omitempty"`
	// Label defines the labelling applied to the class. A class can contain
	// multiple labels (since MapServer 6.2).
	Label *Label `json:",omitempty"`
	// Template file or URL to use in presenting query results to the user. See
	// <http://mapserver.org/mapfile/template.html#template> for more info.
	Template String `json:",omitempty"`
	// Text to label features in this class with. This overrides values obtained
	// from the LAYER LABELITEM. The string can contain references to feature
	// attributes. This allows you to concatenate multiple attributes into a
	// single label. You can for example concatenate the attributes FIRSTNAME
	// and LASTNAME like this:
	//
	//   TEXT '[FIRSTNAME] [LASTNAME]'
	//
	// More advanced Expressions can be used to specify the labels. Since
	// version 6.0, there are functions available for formatting numbers:
	//
	//   TEXT ("Area is: " + tostring([area],"%.2f"))
	Text String `json:",omitempty"`
}
