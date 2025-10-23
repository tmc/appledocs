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
	Count() int /* primitive/slice/pointer. */
	SetCount(value int /* primitive/slice/pointer. */)
	Description() string /* primitive/slice/pointer. */
	SetDescription(value string /* primitive/slice/pointer. */)
	FirstObject() unsafe.Pointer
	SetFirstObject(value unsafe.Pointer)
	LastObject() unsafe.Pointer
	SetLastObject(value unsafe.Pointer)
	SortedArrayHint() IData
	SetSortedArrayHint(value IData)
	// methods:
	DifferenceFromArray(other []objc.ID /* already interface */) unsafe.Pointer
	DifferenceFromArrayWithOptions(other []objc.ID /* already interface */, options OrderedCollectionDifferenceCalculationOptions) unsafe.Pointer
	DifferenceFromArrayWithOptionsUsingEquivalenceTest(other []objc.ID /* already interface */, options OrderedCollectionDifferenceCalculationOptions, block bool /* primitive/slice/pointer. */) unsafe.Pointer
	IndexOfObjectInSortedRangeOptionsUsingComparator(obj unsafe.Pointer, r objc.IObject /* cross-framework Range */, opts BinarySearchingOptions, cmp Comparator /* not a class type */) uint /* primitive/slice/pointer. */
	ReverseObjectEnumerator() unsafe.Pointer
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



// Compares two arrays to create a difference object that represents the changes between them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/differenceFromArray:
func (a_ Array) DifferenceFromArray(other []objc.ID /* already interface */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("differenceFromArray:"), other)
	return rv
}


// Compares two arrays, with options, to create a difference object that represents the changes between them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/differenceFromArray:withOptions:
func (a_ Array) DifferenceFromArrayWithOptions(other []objc.ID /* already interface */, options OrderedCollectionDifferenceCalculationOptions) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("differenceFromArray:withOptions:"), other, options)
	return rv
}


// Compares two arrays, using the provided block and with options, to create a difference object that represents the changes between them.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/differenceFromArray:withOptions:usingEquivalenceTest:
func (a_ Array) DifferenceFromArrayWithOptionsUsingEquivalenceTest(other []objc.ID /* already interface */, options OrderedCollectionDifferenceCalculationOptions, block bool /* primitive/slice/pointer. */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("differenceFromArray:withOptions:usingEquivalenceTest:"), other, options, block)
	return rv
}


// Returns the index, within a specified range, of an object compared with elements in the array using a given block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/index(of:inSortedRange:options:usingComparator:)
func (a_ Array) IndexOfObjectInSortedRangeOptionsUsingComparator(obj unsafe.Pointer, r objc.IObject /* cross-framework Range */, opts BinarySearchingOptions, cmp Comparator /* not a class type */) uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](a_.ID, objc.Sel("indexOfObject:inSortedRange:options:usingComparator:"), obj, r, opts, cmp)
	return rv
}


// Returns an enumerator object that lets you access each object in the array, in reverse order.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSArray/reverseObjectEnumerator()
func (a_ Array) ReverseObjectEnumerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("reverseObjectEnumerator"))
	return rv
}


// The number of objects in the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/count
func (a_ Array) Count() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](a_.ID, objc.Sel("count"))
	return rv
}


// The number of objects in the array.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/count
func (a_ Array) SetCount(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCount:"), value)
}


// A string that represents the contents of the array, formatted as a property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/description
func (a_ Array) Description() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](a_.ID, objc.Sel("description"))
	return rv
}


// A string that represents the contents of the array, formatted as a property list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsarray/description
func (a_ Array) SetDescription(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDescription:"), objc.String(value))
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



