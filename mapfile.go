package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"

	"github.com/geo-data/mapfile/mapobj"
	"github.com/geo-data/mapfile/tokens"
)

func main() {
	tokens, err := tokens.TokenizeMap("/tmp/mapserver-7.0.1/tests/test.map")
	if err != nil {
		log.Fatal(err)
	}

	var mapfile mapobj.Map
	if err = mapfile.FromTokens(tokens); err != nil {
		log.Fatal(err)
	}

	b, err := json.Marshal(mapfile)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out, b, "", " ")
	out.WriteTo(os.Stdout)
}
