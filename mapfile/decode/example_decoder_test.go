package decode_test

import (
	"fmt"
	"github.com/geo-data/mapstruct/mapfile/decode"
	"log"
)

func ExampleDecoder() {
	mapfile := `
MAP
  NAME hello
  STATUS ON
  EXTENT 0 0 4000 3000
  SIZE 400 300
  IMAGECOLOR 200 255 255
  LAYER
    NAME 'credits'
    STATUS DEFAULT
    TYPE ANNOTATION
    FEATURE
      POINTS
        200 150
      END
      TEXT 'Hello world.  Mapserver rocks.'
    END
    CLASS
      LABEL
        TYPE BITMAP
        COLOR 0 0 0
      END
    END
  END
END`

	dec := decode.DecodeString(mapfile)

	map_, err := dec.Map()
	if err != nil {
		log.Fatalf("Decoder.Map() failed: %s", err)
	}

	fmt.Println(string(map_.Name))
	// Output: hello
}
