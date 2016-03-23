package scanner

import "fmt"

// TokenType represents a lexical token.
type TokenType int

const (
	// Special tokens
	ILLEGAL TokenType = iota // Unexpected token
	EOF                      // End of file
	WS                       // Whitespace

	// Literals
	MS_ATTRIBUTE   // Attribute
	MS_COMMENT     // Mapfile comment
	MS_EXPRESSION  // Mapfile expression
	MS_LISTEX      // Mapfile list expression
	MS_NUMBER      // Integer or double
	MS_REGEX       // Regular expression
	MS_STRING      // Simple string
	MS_BARE_STRING // Unquoted Mapfile string

	// Mapfile directives
	BACKGROUNDCOLOR
	BUFFER
	CLASS
	CLASSITEM
	COLOR
	DATA
	DATAPATTERN
	DEBUG
	END
	EXPRESSION
	EXTENT
	FEATURE
	FONT
	FONTSET
	GEOMTRANSFORM
	IMAGECOLOR
	IMAGETYPE
	ITEMS
	LABEL
	LABELITEM
	LAYER
	LEGEND
	MAP
	METADATA
	NAME
	OUTLINECOLOR
	POINTS
	POSITION
	POSTLABELCACHE
	PROCESSING
	PROJECTION
	SCALEBAR
	SIZE
	STATUS
	STYLE
	SYMBOL
	SYMBOLSET
	TEMPLATE
	TEXT
	TRANSPARENT
	TYPE
	UNITS
	WEB
	WIDTH
	WKT
)

// toString maps tokens to their string representations.
var toString = map[TokenType]string{
	// Special tokens
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	WS:      "WS",

	// Literals
	MS_ATTRIBUTE:  "MS_ATTRIBUTE",
	MS_COMMENT:    "MS_COMMENT",
	MS_EXPRESSION: "MS_EXPRESSION",
	MS_LISTEX:     "MS_LISTEX",
	MS_NUMBER:     "MS_NUMBER",
	MS_REGEX:      "MS_REGEX",
	MS_STRING:     "MS_STRING",

	// Mapfile directives
	BACKGROUNDCOLOR: "BACKGROUNDCOLOR",
	BUFFER:          "BUFFER",
	CLASS:           "CLASS",
	CLASSITEM:       "CLASSITEM",
	COLOR:           "COLOR",
	DATA:            "DATA",
	DATAPATTERN:     "DATAPATTERN",
	DEBUG:           "DEBUG",
	END:             "END",
	EXPRESSION:      "EXPRESSION",
	EXTENT:          "EXTENT",
	FEATURE:         "FEATURE",
	FONT:            "FONT",
	FONTSET:         "FONTSET",
	GEOMTRANSFORM:   "GEOMTRANSFORM",
	IMAGECOLOR:      "IMAGECOLOR",
	IMAGETYPE:       "IMAGETYPE",
	ITEMS:           "ITEMS",
	LABEL:           "LABEL",
	LABELITEM:       "LABELITEM",
	LAYER:           "LAYER",
	LEGEND:          "LEGEND",
	MAP:             "MAP",
	METADATA:        "METADATA",
	NAME:            "NAME",
	OUTLINECOLOR:    "OUTLINECOLOR",
	POINTS:          "POINTS",
	POSITION:        "POSITION",
	POSTLABELCACHE:  "POSTLABELCACHE",
	PROCESSING:      "PROCESSING",
	PROJECTION:      "PROJECTION",
	SCALEBAR:        "SCALEBAR",
	SIZE:            "SIZE",
	STATUS:          "STATUS",
	STYLE:           "STYLE",
	SYMBOL:          "SYMBOL",
	SYMBOLSET:       "SYMBOLSET",
	TEMPLATE:        "TEMPLATE",
	TEXT:            "TEXT",
	TRANSPARENT:     "TRANSPARENT",
	TYPE:            "TYPE",
	UNITS:           "UNITS",
	WEB:             "WEB",
	WIDTH:           "WIDTH",
	WKT:             "WKT",
}

// String implements the fmt.Stringer interface.
func (t TokenType) String() string {
	return toString[t]
}

// fromString maps a string representation to a token.
var fromString map[string]TokenType

// init initialises token data structures.
func init() {
	// Seed the reverse lookup of strings to tokens.
	fromString = make(map[string]TokenType)
	for tok, str := range toString {
		fromString[str] = tok
	}
}

// Token represents a token with its type and value.
type Token struct {
	Type  TokenType
	Value string
}

// String implements the fmt.String interface.
func (tok Token) String() string {
	t := tok.Type.String()

	if t == "" {
		return tok.Value
	}

	return fmt.Sprintf("%s %s", tok.Type, tok.Value)
}
