// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [AutoreleasePool] class.
var AutoreleasePoolClass objc.Class

func init() {
	AutoreleasePoolClass = objc.GetClass("NSAutoreleasePool")
}

type AutoreleasePool struct {
	objc.ID
}

func AutoreleasePoolFrom(ptr unsafe.Pointer) AutoreleasePool {
	return AutoreleasePool{
		ID: objc.ID(ptr),
	}
}



