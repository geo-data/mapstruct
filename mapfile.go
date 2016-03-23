package main

import (
	"bytes"
	//	"encoding/json"
	"log"
	"os"

	"github.com/geo-data/mapstruct/mapfile/decode"
	"github.com/geo-data/mapstruct/mapfile/encode"
	"github.com/geo-data/mapstruct/types"
)

func main() {
	mapfile := os.Args[1]

	fh, err := os.Open(mapfile)
	if err != nil {
		log.Fatal(err)
	}
	defer fh.Close()

	dec := decode.DecodeMapfile(fh)

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
		panic(err)
	}

	if _, err = out.WriteTo(os.Stdout); err != nil {
		log.Fatal(err)
	}
}
