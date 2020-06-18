package lzma

//#include <stdlib.h>
//#include <strings.h>
//#include "lzma.h"
import "C"

import (
	"errors"
	"unsafe"
)

func Uncompress(src []byte) ([]byte, error) {
	l := len(src)
	if l == 0 {
		return nil, errors.New("src data error")
	}
	var srcptr *C.char = (*C.char)(unsafe.Pointer(&src[0]))
	var srclen C.int = C.int(C.cxLzmaGetUncompressLen(unsafe.Pointer(srcptr)))
	var ret []byte = make([]byte, int(srclen))
	var dstptr *C.char = (*C.char)(unsafe.Pointer(&ret[0]))
	if C.cxLzmaUncompress(srcptr, C.int(l), dstptr, &srclen) == nil {
		return nil, errors.New("uncompress failed")
	}
	return ret, nil
}

func Compress(src []byte, level ...int) ([]byte, error) {
	l := len(src)
	if l == 0 {
		return nil, errors.New("src data error")
	}
	lv := 5
	if len(level) > 0 {
		lv = level[0]
	}
	var srcptr *C.char = (*C.char)(unsafe.Pointer(&src[0]))
	var srclen C.int = C.int(C.cxLzmaGetCompressLen(C.int(l)))
	var ret []byte = make([]byte, int(srclen))
	var dstptr *C.char = (*C.char)(unsafe.Pointer(&ret[0]))
	if C.cxLzmaCompress(srcptr, C.int(l), dstptr, &srclen, C.int(lv)) == nil {
		return nil, errors.New("compress failed")
	}
	return ret[0:int(srclen)], nil
}
