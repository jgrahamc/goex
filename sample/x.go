package main

import (
	"flag"
)

// #cgo LDFLAGS: -lfreeimage
// #include "FreeImage.h"
//
// int degrade_jpeg(const char *in, const char *out, int quality) {
//    FIBITMAP *bitmap = FreeImage_Load(FIF_JPEG, in, JPEG_ACCURATE);
//    if (bitmap == NULL) {
//        return 0;
//    }
//    BOOL result = FreeImage_Save(FIF_JPEG, bitmap, out, quality);
//    FreeImage_Unload(bitmap);
//    return result?1:0;
// }
import "C"

func degradeJPEG(in, out string, quality int) bool {
	return C.degrade_jpeg(C.CString(in), C.CString(out), C.int(quality)) == 1
}

func main() {
	flag.Parse()
	degradeJPEG(flag.Arg(0), flag.Arg(1), 10)
}
