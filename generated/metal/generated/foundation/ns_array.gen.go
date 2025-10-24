// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Array] class.
var (
	ArrayClass     _ArrayClass
	ArrayClassOnce sync.Once
)

func getArrayClass() _ArrayClass {
	ArrayClassOnce.Do(func() {
		ArrayClass = _ArrayClass{objc.GetClass("NSArray")}
	})
	return ArrayClass
}

type _ArrayClass struct {
	class objc.Class
}

// An interface definition for the [Array] class.
type IArray interface {
	objectivec.IObject
	// properties:
	Count() int
	SetCount(value int)
	Description() IString
	SetDescription(value IString)
	FirstObject() unsafe.Pointer
	SetFirstObject(value unsafe.Pointer)
	LastObject() unsafe.Pointer
	SetLastObject(value unsafe.Pointer)
	SortedArrayHint() IData
	SetSortedArrayHint(value IData)
	// methods:
}

// A static ordered collection of objects.
//
// You can use this type in Swift instead of an constant in cases that require reference semantics. and its subclass manage ordered collections of objects called . creates static arrays, and creates dynamic arrays. You can use arrays when you need an ordered collection of objects. is “toll-free bridged” with its Core Foundation counterpart, . See for more information on toll-free bridging.


// A static ordered collection of objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray
type Array struct {
	objectivec.Object
}

// ArrayFrom constructs a [Array] from an unsafe.Pointer.
//
// A static ordered collection of objects.
func ArrayFrom(ptr unsafe.Pointer) Array {
	return Array{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _ArrayClass) Alloc() Array {
	rv := objc.Send[Array](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _ArrayClass) New() Array {
	rv := objc.Send[Array](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ Array) Init() Array {
	rv := objc.Send[Array](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ Array) Autorelease() Array {
	rv := objc.Send[Array](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewArray creates a new Array instance.
func NewArray() Array {
	return getArrayClass().New()
}



// The number of objects in the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/count
func (a_ Array) Count() int {
	rv := objc.Send[int](a_.ID, objc.Sel("count"))
	return rv
}


// The number of objects in the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/count
func (a_ Array) SetCount(value int) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCount:"), value)
}


// A string that represents the contents of the array, formatted as a property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/description
func (a_ Array) Description() IString {
	rv := objc.Send[String](a_.ID, objc.Sel("description"))
	return rv
}


// A string that represents the contents of the array, formatted as a property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/description
func (a_ Array) SetDescription(value IString) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDescription:"), value)
}


// The first object in the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/firstobject
func (a_ Array) FirstObject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("firstObject"))
	return rv
}


// The first object in the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/firstobject
func (a_ Array) SetFirstObject(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFirstObject:"), value)
}


// The last object in the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/lastobject
func (a_ Array) LastObject() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("lastObject"))
	return rv
}


// The last object in the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/lastobject
func (a_ Array) SetLastObject(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setLastObject:"), value)
}


// Analyzes the array and returns a “hint” that speeds the sorting of the array when the hint is supplied to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/sortedarrayhint
func (a_ Array) SortedArrayHint() IData {
	rv := objc.Send[Data](a_.ID, objc.Sel("sortedArrayHint"))
	return rv
}


// Analyzes the array and returns a “hint” that speeds the sorting of the array when the hint is supplied to
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/sortedarrayhint
func (a_ Array) SetSortedArrayHint(value IData) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSortedArrayHint:"), value)
}



