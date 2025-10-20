// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var ExceptionClass _ExceptionClass

func init() {
	ExceptionClass = _ExceptionClass{objc.GetClass("NSException")}
}

type _ExceptionClass struct {
	class objc.Class
}

type Exception struct {
	objc.ID
}

func ExceptionFrom(ptr unsafe.Pointer) Exception {
	return Exception{
		ID: objc.ID(ptr),
	}
}




