// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PointerFunctions] class.
var (
	pointerFunctionsClass     _PointerFunctionsClass
	pointerFunctionsClassOnce sync.Once
)

func getPointerFunctionsClass() _PointerFunctionsClass {
	pointerFunctionsClassOnce.Do(func() {
		pointerFunctionsClass = _PointerFunctionsClass{objc.GetClass("NSPointerFunctions")}
	})
	return pointerFunctionsClass
}

type _PointerFunctionsClass struct {
	class objc.Class
}

// An interface definition for the [PointerFunctions] class.
type IPointerFunctions interface {
	objectivec.IObject
}

// An instance of defines callout functions appropriate for managing a pointer reference held somewhere else.
//
// The functions specified by an instance of are separated into two clusters—those that define “personality” such as “object” or “C-string”, and those that describe memory management issues such as a memory deallocation function. There are constants for common personalities and memory manager selections (see ). , , and use an object to define the acquisition and retention behavior for the pointers they manage. Note, however, that not all combinations of personality and memory management behavior are valid for these collections. The pointer collection objects copy the object on input and output, so you cannot usefully subclass .
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

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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
	return getPointerFunctionsClass().New()
}




