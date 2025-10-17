// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [SimpleCString] class.
var SimpleCStringClass objc.Class

func init() {
	SimpleCStringClass = objc.GetClass("NSSimpleCString")
}

type SimpleCString struct {
	objc.ID
}

func SimpleCStringFrom(ptr unsafe.Pointer) SimpleCString {
	return SimpleCString{
		ID: objc.ID(ptr),
	}
}



