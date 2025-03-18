package util

import (
	"unsafe"
)

// emptyInterface is the header for an interface{} value.
type emptyInterface struct {
	typ  uintptr
	word uintptr
}

// IsNil returns true if v is a null pointer.
func IsNil(v interface{}) bool {
	e := (*emptyInterface)(unsafe.Pointer(&v))
	return e.word == 0
}
