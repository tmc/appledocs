// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Null] class.
var NullClass objc.Class

func init() {
	NullClass = objc.GetClass("NSNull")
}

type Null struct {
	objc.ID
}

func NullFrom(ptr unsafe.Pointer) Null {
	return Null{
		ID: objc.ID(ptr),
	}
}




