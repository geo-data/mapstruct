package tokens

import (
	"fmt"
	"github.com/geo-data/mapfile/types/feature"
)

func (t *Tokens) Feature() (c *feature.Feature, err error) {
	token := t.Value()
	if token != "FEATURE" {
		err = fmt.Errorf("expected token FEATURE, got: %s", token)
		return
	}
	t.Next()

	c = new(feature.Feature)
	for t != nil {
		token := t.Value()
		switch token {
		case "WKT":
			if c.Wkt, err = t.Next().String(); err != nil {
				return
			}
		case "ITEMS":
			if c.Items, err = t.Next().String(); err != nil {
				return
			}
		case "TEXT":
			if c.Text, err = t.Next().String(); err != nil {
				return
			}
		case "POINTS":
			if c.Points, err = t.Points(); err != nil {
				return
			}
		case "END":
			return
		default:
			err = fmt.Errorf("unhandled mapfile token: %s", token)
			return
		}

		t = t.Next()
	}

	return
}
