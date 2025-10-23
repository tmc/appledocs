// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CountedSet] class.
var (
	CountedSetClass     _CountedSetClass
	CountedSetClassOnce sync.Once
)

func getCountedSetClass() _CountedSetClass {
	CountedSetClassOnce.Do(func() {
		CountedSetClass = _CountedSetClass{objc.GetClass("NSCountedSet")}
	})
	return CountedSetClass
}

type _CountedSetClass struct {
	class objc.Class
}

// An interface definition for the [CountedSet] class.
type ICountedSet interface {
	IMutableSet
	// properties:
	Count() int /* primitive/slice/pointer */
	SetCount(value int /* primitive/slice/pointer */)
	// methods:
}

// A mutable, unordered collection of distinct objects that may appear more than once in the collection.
//
// Each distinct object inserted into an object has a counter associated with it. keeps track of the number of times objects are inserted and requires that objects be removed the same number of times. Thus, there is only one instance of an object in an object even if the object has been added to the set multiple times. The method defined by the superclass has special significance; it returns the number of distinct objects, not the total number of times objects are represented in the set. The and classes are provided for static and dynamic sets, respectively, whose elements are distinct. While and are not toll-free bridged, they provide similar functionality. For more information about , see the .


// A mutable, unordered collection of distinct objects that may appear more than once in the collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountedSet
type CountedSet struct {
	MutableSet
}

// CountedSetFrom constructs a [CountedSet] from an unsafe.Pointer.
//
// A mutable, unordered collection of distinct objects that may appear more than once in the collection.
func CountedSetFrom(ptr unsafe.Pointer) CountedSet {
	return CountedSet{
		MutableSet: MutableSetFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CountedSetClass) Alloc() CountedSet {
	rv := objc.Send[CountedSet](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CountedSetClass) New() CountedSet {
	rv := objc.Send[CountedSet](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CountedSet) Init() CountedSet {
	rv := objc.Send[CountedSet](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CountedSet) Autorelease() CountedSet {
	rv := objc.Send[CountedSet](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCountedSet creates a new CountedSet instance.
func NewCountedSet() CountedSet {
	return getCountedSetClass().New()
}



// The number of members in the set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/count
func (c_ CountedSet) Count() int /* primitive/slice/pointer */ {
	rv := objc.Send[int](c_.ID, objc.Sel("count"))
	return rv
}


// The number of members in the set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsset/count
func (c_ CountedSet) SetCount(value int /* primitive/slice/pointer */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCount:"), value)
}



