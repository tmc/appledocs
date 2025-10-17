// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SimpleCString] class.
var SimpleCStringClass = _SimpleCStringClass{objc.GetClass("NSSimpleCString")}

type _SimpleCStringClass struct {
	class objc.Class
}

type SimpleCString struct {
	objc.ID
}

func SimpleCStringFrom(ptr unsafe.Pointer) SimpleCString {
	return SimpleCString{
		ID: objc.ID(ptr),
	}
}




