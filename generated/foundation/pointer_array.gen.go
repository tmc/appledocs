// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [PointerArray] class.
var PointerArrayClass = _PointerArrayClass{objc.GetClass("NSPointerArray")}

type _PointerArrayClass struct {
	class objc.Class
}

type PointerArray struct {
	objc.ID
}

func PointerArrayFrom(ptr unsafe.Pointer) PointerArray {
	return PointerArray{
		ID: objc.ID(ptr),
	}
}




