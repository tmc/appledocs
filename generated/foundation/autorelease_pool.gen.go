// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AutoreleasePool] class.
var AutoreleasePoolClass = _AutoreleasePoolClass{objc.GetClass("NSAutoreleasePool")}

type _AutoreleasePoolClass struct {
	class objc.Class
}

type AutoreleasePool struct {
	objc.ID
}

func AutoreleasePoolFrom(ptr unsafe.Pointer) AutoreleasePool {
	return AutoreleasePool{
		ID: objc.ID(ptr),
	}
}




