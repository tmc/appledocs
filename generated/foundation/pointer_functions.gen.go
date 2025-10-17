// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PointerFunctions] class.
var pointerFunctionsClass = _PointerFunctionsClass{objc.GetClass("NSPointerFunctions")}

type _PointerFunctionsClass struct {
	class objc.Class
}

// An interface definition for the [PointerFunctions] class.
type IPointerFunctions interface {
	objectivec.IObject
}

// An instance of defines callout functions appropriate for managing a pointer reference held somewhere else. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerFunctions

type PointerFunctions struct {
	objectivec.Object
}

// PointerFunctionsFrom constructs a [PointerFunctions] from an unsafe.Pointer.
//
// An instance of defines callout functions appropriate for managing a pointer reference held somewhere else.
func PointerFunctionsFrom(ptr unsafe.Pointer) PointerFunctions {
	return PointerFunctions{objectivec.Object{objc.ID(ptr)}}
}



