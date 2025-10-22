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
	PointerArrayClass     _PointerArrayClass
	PointerArrayClassOnce sync.Once
)

func getPointerArrayClass() _PointerArrayClass {
	PointerArrayClassOnce.Do(func() {
		PointerArrayClass = _PointerArrayClass{objc.GetClass("NSPointerArray")}
	})
	return PointerArrayClass
}

type _PointerArrayClass struct {
	class objc.Class
}

// An interface definition for the [PointerArray] class.
type IPointerArray interface {
	objectivec.IObject
	RemovePointerAtIndex(index uint)
	AllObjects() objc.ID
	Count() uint
	SetCount(value uint)
	PointerFunctions() NSPointerFunctions
	SetPointerFunctions(value IPointerFunctions)
}

// A collection similar to an array, but with a broader range of available memory semantics.
//
// The pointer array class is modeled after , but can also hold values. You can insert or remove values which contribute to the array’s . A pointer array can be initialized to maintain strong or weak references to objects, or according to any of the memory or personality options defined by . The and protocols are applicable only when a pointer array is initialized to maintain strong or weak references to objects. When enumerating a pointer array with using , the loop will yield any values present in the array. See in for more information.
//
// [Full Topic]
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



// Returns a new pointer array initialized to use the given options.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray/pointerArrayWithOptions:

func (pc _PointerArrayClass) PointerArrayWithOptions(options PointerFunctionsOptions) PointerArray {
	rv := objc.Send[PointerArray](objc.ID(pc.class), objc.Sel("pointerArrayWithOptions:"), options)
	return rv
}


// A new pointer array initialized to use the given functions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray/pointerArrayWithPointerFunctions:

func (pc _PointerArrayClass) PointerArrayWithPointerFunctions(functions IPointerFunctions) PointerArray {
	rv := objc.Send[PointerArray](objc.ID(pc.class), objc.Sel("pointerArrayWithPointerFunctions:"), functions)
	return rv
}


// Returns a new pointer array that maintains strong references to its elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray/pointerArrayWithStrongObjects

func (pc _PointerArrayClass) PointerArrayWithStrongObjects() objc.ID {
	rv := objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("pointerArrayWithStrongObjects"))
	return rv
}


// Returns a new pointer array that maintains weak references to its elements.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray/weakObjects()

func (pc _PointerArrayClass) WeakObjectsPointerArray() PointerArray {
	rv := objc.Send[PointerArray](objc.ID(pc.class), objc.Sel("weakObjectsPointerArray"))
	return rv
}

// Removes the pointer at a given index.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray/removePointer(at:)
func (p_ PointerArray) RemovePointerAtIndex(index uint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("removePointerAtIndex:"), index)
}

// All the objects in the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray/allObjects
func (p_ PointerArray) AllObjects() objc.ID {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("allObjects"))
	return rv
}

// The number of elements in the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray/count
func (p_ PointerArray) Count() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("count"))
	return rv
}


// SetCount sets the value of the count property.
// The number of elements in the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSPointerArray/count
func (p_ PointerArray) SetCount(value uint) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setCount:"), value)
}

// The functions in use by the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerarray/pointerfunctions
func (p_ PointerArray) PointerFunctions() NSPointerFunctions {
	rv := objc.Send[NSPointerFunctions](p_.ID, objc.Sel("pointerFunctions"))
	return rv
}


// SetPointerFunctions sets the value of the pointerFunctions property.
// The functions in use by the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspointerarray/pointerfunctions
func (p_ PointerArray) SetPointerFunctions(value IPointerFunctions) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPointerFunctions:"), value)
}



