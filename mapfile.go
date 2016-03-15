package main

import (
	"bytes"
	//	"encoding/json"
	"log"
	"os"

	"github.com/geo-data/mapfile/mapfile/decode/tokens"
	"github.com/geo-data/mapfile/mapfile/encode"
	"github.com/geo-data/mapfile/types/mapobj"
)

func main() {
	mapfile := os.Args[1]
	tokens, err := tokens.TokenizeMap(mapfile)
	if err != nil {
		log.Fatal(err)
	}

	var map_ *mapobj.Map
	if map_, err = tokens.Map(); err != nil {
		panic(err)
		log.Fatal(err)
	}

	/*b, err := json.Marshal(map_)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out, b, "", "  ")*/

	var out bytes.Buffer
	enc := encode.NewMapfileEncoder(&out)
	if err = enc.Encode(map_); err != nil {
		log.Fatal(err)
	}

	if _, err = out.WriteTo(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
