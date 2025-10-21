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
	AddObject(object unsafe.Pointer)
	CountForObject(object unsafe.Pointer) uint
	ObjectEnumerator() unsafe.Pointer
	RemoveObject(object unsafe.Pointer)
}

// A mutable, unordered collection of distinct objects that may appear more than once in the collection.
//
// Each distinct object inserted into an object has a counter associated with it. keeps track of the number of times objects are inserted and requires that objects be removed the same number of times. Thus, there is only one instance of an object in an object even if the object has been added to the set multiple times. The method defined by the superclass has special significance; it returns the number of distinct objects, not the total number of times objects are represented in the set. The and classes are provided for static and dynamic sets, respectively, whose elements are distinct. While and are not toll-free bridged, they provide similar functionality. For more information about , see the .
//
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

// Returns a counted set object initialized with the contents of a given array.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountedSet/init(array:)
func NewCountedSetWithArray(array unsafe.Pointer) CountedSet {
	instance := getCountedSetClass().Alloc()
	rv := objc.Send[CountedSet](instance.ID, objc.Sel("initWithArray:"), array)
	rv.Autorelease()
	return rv
}

// Returns a counted set object initialized with enough memory to hold a given number of objects.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountedSet/init(capacity:)
func NewCountedSetWithCapacity(numItems uint) CountedSet {
	instance := getCountedSetClass().Alloc()
	rv := objc.Send[CountedSet](instance.ID, objc.Sel("initWithCapacity:"), numItems)
	rv.Autorelease()
	return rv
}

// Returns a counted set object initialized with the contents of a given set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountedSet/init(set:)
func NewCountedSetWithSet(set unsafe.Pointer) CountedSet {
	instance := getCountedSetClass().Alloc()
	rv := objc.Send[CountedSet](instance.ID, objc.Sel("initWithSet:"), set)
	rv.Autorelease()
	return rv
}

// Adds a given object to the set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountedSet/add(_:)
func (c_ CountedSet) AddObject(object unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("addObject:"), object)
}

// Returns the count associated with a given object in the set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountedSet/count(for:)
func (c_ CountedSet) CountForObject(object unsafe.Pointer) uint {
	rv := objc.Send[uint](c_.ID, objc.Sel("countForObject:"), object)
	return rv
}

// Returns an enumerator object that lets you access each object in the set once, independent of its count.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountedSet/objectEnumerator()
func (c_ CountedSet) ObjectEnumerator() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("objectEnumerator"))
	return rv
}

// Removes a given object from the set.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSCountedSet/remove(_:)
func (c_ CountedSet) RemoveObject(object unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("removeObject:"), object)
}
