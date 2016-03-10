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
	Type       tokens.String          `json:",omitempty"`
	Debug      tokens.String          `json:",omitempty"`
	Projection *projection.Projection `json:",omitempty"`
	Data       tokens.String          `json:",omitempty"`
	Processing tokens.String          `json:",omitempty"`
	Status     tokens.String          `json:",omitempty"`
	Metadata   *metadata.Metadata     `json:",omitempty"`
	ClassItem  tokens.String          `json:",omitempty"`
	LabelItem  tokens.String          `json:",omitempty"`
	Classes    []*class.Class         `json:",omitempty"`
	Features   []*feature.Feature     `json:",omitempty"`
}

func New(tokens *tokens.Tokens) (l *Layer, err error) {
	token := tokens.Value()
	if token != "LAYER" {
		err = fmt.Errorf("expected token LAYER, got: %s", token)
		return
	}
	tokens.Next()

	l = new(Layer)
	for tokens != nil {
		token := tokens.Value()
		switch token {
		case "NAME":
			l.Name = tokens.Next().Value()
		case "EXTENT":
			if l.Extent, err = extent.New(tokens); err != nil {
				return
			}
		case "TYPE":
			l.Type = tokens.Next().Value()
		case "DEBUG":
			l.Debug = tokens.Next().Value()
		case "PROJECTION":
			if l.Projection, err = projection.New(tokens); err != nil {
				return
			}
		case "DATA":
			l.Data = tokens.Next().Value()
		case "PROCESSING":
			l.Processing = tokens.Next().Value()
		case "STATUS":
			l.Status = tokens.Next().Value()
		case "METADATA":
			if l.Metadata, err = metadata.New(tokens); err != nil {
				return
			}
		case "CLASSITEM":
			l.ClassItem = tokens.Next().Value()
		case "LABELITEM":
			l.LabelItem = tokens.Next().Value()
		case "CLASS":
			var c *class.Class
			if c, err = class.New(tokens); err != nil {
				return
			}
			l.Classes = append(l.Classes, c)
		case "FEATURE":
			var f *feature.Feature
			if f, err = feature.New(tokens); err != nil {
				return
			}
			l.Features = append(l.Features, f)
		case "END":
			return
		default:
			err = fmt.Errorf("unhandled mapfile token: %s", token)
			return
		}

		tokens = tokens.Next()
	}

	return
}

func (l *Layer) Encode(enc *encoding.MapfileEncoder) (err error) {
	if err = enc.TokenStart("LAYER"); err != nil {
		return
	}

	if err = enc.TokenString("NAME", l.Type); err != nil {
		return
	}
	if l.Extent != nil {
		if err = l.Extent.Encode(enc); err != nil {
			return
		}
	}
	if err = enc.TokenValue("TYPE", l.Type); err != nil {
		return
	}
	if err = enc.TokenValue("DEBUG", l.Debug); err != nil {
		return
	}
	if l.Projection != nil {
		if err = l.Projection.Encode(enc); err != nil {
			return
		}
	}
	if err = enc.TokenString("DATA", l.Data); err != nil {
		return
	}
	if err = enc.TokenString("PROCESSING", l.Processing); err != nil {
		return
	}
	if err = enc.TokenValue("STATUS", l.Status); err != nil {
		return
	}
	if l.Metadata != nil {
		if err = l.Metadata.Encode(enc); err != nil {
			return
		}
	}
	if err = enc.TokenValue("CLASSITEM", l.ClassItem); err != nil {
		return
	}
	if err = enc.TokenValue("LABELITEM", l.LabelItem); err != nil {
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
