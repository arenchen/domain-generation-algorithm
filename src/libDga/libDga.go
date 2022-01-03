package main

import "C"

import (
	"dga/src/pkgDga"
	"unsafe"
)

//export GenerateRandomKey
func GenerateRandomKey(keySize C.int) *C.char {
	if keySize < 1 {
		keySize = 10
	}

	kSize := int(keySize)

	key := pkgDga.GenerateRandomKey(kSize)

	return C.CString(key)
}

//export GenerateDomain
func GenerateDomain(token *C.char, unixSeconds C.long, formatPattern *C.char, count C.int) **C.char {
	key := C.GoString(token)
	unixtime := int64(unixSeconds)
	format := C.GoString(formatPattern)
	length := int(count)

	domains := pkgDga.GenerateDomain(key, unixtime, format, length)

	cArray := C.malloc(C.size_t(len(domains)) * C.size_t(unsafe.Sizeof(uintptr(0))))

	// convert the C array to a Go Array so we can index it
	a := (*[1<<30 - 1]*C.char)(cArray)

	for idx, substring := range domains {
		a[idx] = C.CString(substring)
	}

	return (**C.char)(cArray)
}

func main() {}
