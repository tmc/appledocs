// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Exception] class.
var ExceptionClass objc.Class

func init() {
	ExceptionClass = objc.GetClass("NSException")
}

type Exception struct {
	objc.ID
}

func ExceptionFrom(ptr unsafe.Pointer) Exception {
	return Exception{
		ID: objc.ID(ptr),
	}
}



