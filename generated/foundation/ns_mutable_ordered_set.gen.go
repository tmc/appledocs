// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableOrderedSet] class.
var (
	MutableOrderedSetClass     _MutableOrderedSetClass
	MutableOrderedSetClassOnce sync.Once
)

func getMutableOrderedSetClass() _MutableOrderedSetClass {
	MutableOrderedSetClassOnce.Do(func() {
		MutableOrderedSetClass = _MutableOrderedSetClass{objc.GetClass("NSMutableOrderedSet")}
	})
	return MutableOrderedSetClass
}

type _MutableOrderedSetClass struct {
	class objc.Class
}

// An interface definition for the [MutableOrderedSet] class.
type IMutableOrderedSet interface {
	IOrderedSet
	RemoveObjectsInRange(range_ Range)
	ReplaceObjectsAtIndexesWithObjects(indexes IIndexSet, objects []objc.ID)
}

// A dynamic, ordered collection of unique objects.
//
// objects are not like C arrays. That is, even though you may specify a size when you create a mutable ordered set, the specified size is regarded as a “hint”; the actual size of the set is still 0. This means that you cannot insert an object at an index greater than the current count of an set. For example, if a set contains two objects, its size is 2, so you can add objects at indices 0, 1, or 2. Index 3 is illegal and out of bounds; if you try to add an object at index 3 (when the size of the array is 2), raises an exception.


// A dynamic, ordered collection of unique objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet

type MutableOrderedSet struct {
	OrderedSet
}

// MutableOrderedSetFrom constructs a [MutableOrderedSet] from an unsafe.Pointer.
//
// A dynamic, ordered collection of unique objects.
func MutableOrderedSetFrom(ptr unsafe.Pointer) MutableOrderedSet {
	return MutableOrderedSet{
		OrderedSet: OrderedSetFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MutableOrderedSetClass) Alloc() MutableOrderedSet {
	rv := objc.Send[MutableOrderedSet](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MutableOrderedSetClass) New() MutableOrderedSet {
	rv := objc.Send[MutableOrderedSet](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableOrderedSet) Init() MutableOrderedSet {
	rv := objc.Send[MutableOrderedSet](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableOrderedSet) Autorelease() MutableOrderedSet {
	rv := objc.Send[MutableOrderedSet](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableOrderedSet creates a new MutableOrderedSet instance.
func NewMutableOrderedSet() MutableOrderedSet {
	return getMutableOrderedSetClass().New()
}




// Removes from the mutable ordered set each of the objects within a given range.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/removeObjects(in:)-9jkis

func (m_ MutableOrderedSet) RemoveObjectsInRange(range_ Range) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeObjectsInRange:"), range_)
}



// Replaces the objects at the specified indexes with the new objects.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableOrderedSet/replaceObjects(at:with:)

func (m_ MutableOrderedSet) ReplaceObjectsAtIndexesWithObjects(indexes IIndexSet, objects []objc.ID) {
	objc.Send[objc.ID](m_.ID, objc.Sel("replaceObjectsAtIndexes:withObjects:"), indexes, objects)
}



