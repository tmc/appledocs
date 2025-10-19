// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CountedSet] class.
var (
	countedSetClass     _CountedSetClass
	countedSetClassOnce sync.Once
)

func getCountedSetClass() _CountedSetClass {
	countedSetClassOnce.Do(func() {
		countedSetClass = _CountedSetClass{objc.GetClass("NSCountedSet")}
	})
	return countedSetClass
}

type _CountedSetClass struct {
	class objc.Class
}

// An interface definition for the [CountedSet] class.
type ICountedSet interface {
	IMutableSet
}

// A mutable, unordered collection of distinct objects that may appear more than once in the collection.
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




