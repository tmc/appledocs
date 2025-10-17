// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [OutputStream] class.
var OutputStreamClass = _OutputStreamClass{objc.GetClass("NSOutputStream")}

type _OutputStreamClass struct {
	class objc.Class
}

type OutputStream struct {
	objc.ID
}

func OutputStreamFrom(ptr unsafe.Pointer) OutputStream {
	return OutputStream{
		ID: objc.ID(ptr),
	}
}




