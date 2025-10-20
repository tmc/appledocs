// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

var PointerArrayClass _PointerArrayClass

func init() {
	PointerArrayClass = _PointerArrayClass{objc.GetClass("NSPointerArray")}
}

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




