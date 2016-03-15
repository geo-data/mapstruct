package layer

import (
	"fmt"
	"github.com/geo-data/mapfile/encoding"
	"github.com/geo-data/mapfile/tokens"
	"github.com/geo-data/mapfile/types"
	"github.com/geo-data/mapfile/types/class"
	"github.com/geo-data/mapfile/types/extent"
	"github.com/geo-data/mapfile/types/feature"
	"github.com/geo-data/mapfile/types/metadata"
	"github.com/geo-data/mapfile/types/projection"
)

type Layer struct {
	Name       types.String           `json:",omitempty"`
	Extent     *extent.Extent         `json:",omitempty"`
	Type       types.Keyword          `json:",omitempty"`
	Debug      types.Union            `json:",omitempty"`
	Projection *projection.Projection `json:",omitempty"`
	Data       types.String           `json:",omitempty"`
	Processing types.String           `json:",omitempty"`
	Status     types.Keyword          `json:",omitempty"`
	Metadata   *metadata.Metadata     `json:",omitempty"`
	ClassItem  types.Attribute        `json:",omitempty"`
	LabelItem  types.Attribute        `json:",omitempty"`
	Classes    []*class.Class         `json:",omitempty"`
	Features   []*feature.Feature     `json:",omitempty"`
}

func New(toks *tokens.Tokens) (l *Layer, err error) {
	token := toks.Value()
	if token != "LAYER" {
		err = fmt.Errorf("expected token LAYER, got: %s", token)
		return
	}
	toks.Next()

	l = new(Layer)
	for toks != nil {
		token := toks.Value()
		switch token {
		case "NAME":
			if l.Name, err = toks.Next().String(); err != nil {
				return
			}
		case "EXTENT":
			if l.Extent, err = extent.New(toks); err != nil {
				return
			}
		case "TYPE":
			if l.Type, err = toks.Next().Keyword(); err != nil {
				return
			}
		case "DEBUG":
			if l.Debug, err = toks.Next().Decode(tokens.Keyword | tokens.Integer); err != nil {
				return
			}
		case "PROJECTION":
			if l.Projection, err = projection.New(toks); err != nil {
				return
			}
		case "DATA":
			if l.Data, err = toks.Next().String(); err != nil {
				return
			}
		case "PROCESSING":
			if l.Processing, err = toks.Next().String(); err != nil {
				return
			}
		case "STATUS":
			if l.Status, err = toks.Next().Keyword(); err != nil {
				return
			}
		case "METADATA":
			if l.Metadata, err = metadata.New(toks); err != nil {
				return
			}
		case "CLASSITEM":
			if l.ClassItem, err = toks.Next().Attribute(); err != nil {
				return
			}
		case "LABELITEM":
			if l.LabelItem, err = toks.Next().Attribute(); err != nil {
				return
			}
		case "CLASS":
			var c *class.Class
			if c, err = class.New(toks); err != nil {
				return
			}
			l.Classes = append(l.Classes, c)
		case "FEATURE":
			var f *feature.Feature
			if f, err = feature.New(toks); err != nil {
				return
			}
			l.Features = append(l.Features, f)
		case "END":
			return
		default:
			err = fmt.Errorf("unhandled mapfile token: %s", token)
			return
		}

		toks = toks.Next()
	}

	return
}

func (l *Layer) Encode(enc *encoding.MapfileEncoder) (err error) {
	if err = enc.TokenStart("LAYER"); err != nil {
		return
	}

	if err = enc.TokenStringer("NAME", l.Name); err != nil {
		return
	}
	if l.Extent != nil {
		if err = l.Extent.Encode(enc); err != nil {
			return
		}
	}
	if err = enc.TokenStringer("TYPE", l.Type); err != nil {
		return
	}
	if err = enc.TokenUnion("DEBUG", l.Debug); err != nil {
		return
	}
	if l.Projection != nil {
		if err = l.Projection.Encode(enc); err != nil {
			return
		}
	}
	if err = enc.TokenStringer("DATA", l.Data); err != nil {
		return
	}
	if err = enc.TokenStringer("PROCESSING", l.Processing); err != nil {
		return
	}
	if err = enc.TokenStringer("STATUS", l.Status); err != nil {
		return
	}
	if l.Metadata != nil {
		if err = l.Metadata.Encode(enc); err != nil {
			return
		}
	}
	if err = enc.TokenString("CLASSITEM", l.ClassItem.QuotedString()); err != nil {
		return
	}
	if err = enc.TokenString("LABELITEM", l.LabelItem.QuotedString()); err != nil {
		return
	}

	for _, class := range l.Classes {
		if err = class.Encode(enc); err != nil {
			return
		}
	}

	for _, feature := range l.Features {
		if err = feature.Encode(enc); err != nil {
			return
		}
	}

	if err = enc.TokenEnd(); err != nil {
		return
	}

	return
}
