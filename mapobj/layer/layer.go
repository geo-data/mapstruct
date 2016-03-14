package layer

import (
	"fmt"
	"github.com/geo-data/mapfile/encoding"
	"github.com/geo-data/mapfile/mapobj/class"
	"github.com/geo-data/mapfile/mapobj/extent"
	"github.com/geo-data/mapfile/mapobj/feature"
	"github.com/geo-data/mapfile/mapobj/metadata"
	"github.com/geo-data/mapfile/mapobj/projection"
	"github.com/geo-data/mapfile/tokens"
)

type Layer struct {
	Name       tokens.String          `json:",omitempty"`
	Extent     *extent.Extent         `json:",omitempty"`
	Type       tokens.Keyword         `json:",omitempty"`
	Debug      fmt.Stringer           `json:",omitempty"`
	Projection *projection.Projection `json:",omitempty"`
	Data       tokens.String          `json:",omitempty"`
	Processing tokens.String          `json:",omitempty"`
	Status     tokens.Keyword         `json:",omitempty"`
	Metadata   *metadata.Metadata     `json:",omitempty"`
	ClassItem  tokens.Attribute       `json:",omitempty"`
	LabelItem  tokens.Attribute       `json:",omitempty"`
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
			if l.Debug, err = toks.Next().Decode(tokens.KEYWORD | tokens.INTEGER); err != nil {
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
	if err = enc.TokenStringer("DEBUG", l.Debug); err != nil {
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
