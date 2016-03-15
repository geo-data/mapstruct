package main

import (
	"bytes"
	//	"encoding/json"
	"log"
	"os"

	"github.com/geo-data/mapfile/mapfile/decode"
	"github.com/geo-data/mapfile/mapfile/encode"
	"github.com/geo-data/mapfile/types"
)

func main() {
	mapfile := os.Args[1]
	dec, err := decode.DecodeMapfile(mapfile)
	if err != nil {
		log.Fatal(err)
	}

	var map_ *types.Map
	if map_, err = dec.Map(); err != nil {
		log.Fatal(err)
	}

	/*b, err := json.Marshal(map_)
	if err != nil {
		log.Fatal(err)
	}

	var out bytes.Buffer
	json.Indent(&out, b, "", "  ")*/

	var out bytes.Buffer
	enc := encode.NewEncoder(&out)
	if err = enc.EncodeMap(map_); err != nil {
		log.Fatal(err)
	}

	if _, err = out.WriteTo(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
