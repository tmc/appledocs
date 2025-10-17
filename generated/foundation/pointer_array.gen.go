// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [PointerArray] class.
var PointerArrayClass objc.Class

func init() {
	PointerArrayClass = objc.GetClass("NSPointerArray")
}

type PointerArray struct {
	objc.ID
}

func PointerArrayFrom(ptr unsafe.Pointer) PointerArray {
	return PointerArray{
		ID: objc.ID(ptr),
	}
}



