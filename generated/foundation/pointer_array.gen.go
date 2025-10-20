// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PointerArray] class.
var (
	pointerArrayClass     _PointerArrayClass
	pointerArrayClassOnce sync.Once
)

func getPointerArrayClass() _PointerArrayClass {
	pointerArrayClassOnce.Do(func() {
		pointerArrayClass = _PointerArrayClass{objc.GetClass("NSPointerArray")}
	})
	return pointerArrayClass
}

type _PointerArrayClass struct {
	class objc.Class
}

// An interface definition for the [PointerArray] class.
type IPointerArray interface {
	objectivec.IObject
}

// A collection similar to an array, but with a broader range of available memory semantics.
//
// The pointer array class is modeled after , but can also hold values. You can insert or remove values which contribute to the array’s . A pointer array can be initialized to maintain strong or weak references to objects, or according to any of the memory or personality options defined by . The and protocols are applicable only when a pointer array is initialized to maintain strong or weak references to objects. When enumerating a pointer array with using , the loop will yield any values present in the array. See in for more information.
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

// Alloc allocates a new instance without initialization.
func (pc _PointerArrayClass) Alloc() PointerArray {
	rv := objc.Send[PointerArray](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PointerArrayClass) New() PointerArray {
	rv := objc.Send[PointerArray](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PointerArray) Init() PointerArray {
	rv := objc.Send[PointerArray](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PointerArray) Autorelease() PointerArray {
	rv := objc.Send[PointerArray](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPointerArray creates a new PointerArray instance.
func NewPointerArray() PointerArray {
	return getPointerArrayClass().New()
}




