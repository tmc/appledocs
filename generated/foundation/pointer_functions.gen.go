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
// Alloc allocates a new instance without initialization.
func (pc _PointerFunctionsClass) Alloc() PointerFunctions {
	rv := objc.Send[PointerFunctions](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (pc _PointerFunctionsClass) New() PointerFunctions {
	rv := objc.Send[PointerFunctions](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PointerFunctions) Init() PointerFunctions {
	rv := objc.Send[PointerFunctions](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PointerFunctions) Autorelease() PointerFunctions {
	rv := objc.Send[PointerFunctions](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPointerFunctions creates a new PointerFunctions instance.
func NewPointerFunctions() PointerFunctions {
	return pointerFunctionsClass.New()
}




