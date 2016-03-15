package tokenize

/*
#cgo CFLAGS: -I/home/hrz/tmp/mapserver/include -I/usr/include/gdal/
#cgo LDFLAGS: -L/home/hrz/tmp/mapserver/lib -lmapserver
#include "mapserver/mapserver.h"
*/
import "C"

import (
	"fmt"
	"unsafe"
)

func TokenizeMapfile(filename string) (tokens []string, err error) {
	cfilename := C.CString(filename)
	defer C.free(unsafe.Pointer(cfilename))

	cnumtokens := C.int(0)
	ctokens := C.msTokenizeMap(cfilename, &cnumtokens)
	if ctokens == nil {
		err = fmt.Errorf("failed to tokenize mapfile %s", filename)
		return
	}
	defer C.msFreeCharArray(ctokens, cnumtokens)

	numtokens := int(cnumtokens)
	for i := 0; i < numtokens; i++ {
		ptr := uintptr(unsafe.Pointer(ctokens)) + (uintptr(i) * unsafe.Sizeof(*ctokens))
		ctoken := *(**C.char)(unsafe.Pointer(ptr))
		token := C.GoString(ctoken)
		if token != "" {
			tokens = append(tokens, token)
		}
	}

	return
}
