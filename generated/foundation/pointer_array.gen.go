// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PointerArray] class.
var pointerArrayClass = _PointerArrayClass{objc.GetClass("NSPointerArray")}

type _PointerArrayClass struct {
	class objc.Class
}

// An interface definition for the [PointerArray] class.
type IPointerArray interface {
	objectivec.IObject
}

// A collection similar to an array, but with a broader range of available memory semantics. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray

type PointerArray struct {
	objectivec.Object
}

// PointerArrayFrom constructs a [PointerArray] from an unsafe.Pointer.
//
// A collection similar to an array, but with a broader range of available memory semantics.
func PointerArrayFrom(ptr unsafe.Pointer) PointerArray {
	return PointerArray{objectivec.Object{objc.ID(ptr)}}
}



