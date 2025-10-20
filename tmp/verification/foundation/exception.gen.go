// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var exceptionClass _ExceptionClass

func init() {
	exceptionClass = _ExceptionClass{objc.GetClass("NSException")}
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




