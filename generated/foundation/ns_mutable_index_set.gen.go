// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableIndexSet] class.
var (
	MutableIndexSetClass     _MutableIndexSetClass
	MutableIndexSetClassOnce sync.Once
)

func getMutableIndexSetClass() _MutableIndexSetClass {
	MutableIndexSetClassOnce.Do(func() {
		MutableIndexSetClass = _MutableIndexSetClass{objc.GetClass("NSMutableIndexSet")}
	})
	return MutableIndexSetClass
}

type _MutableIndexSetClass struct {
	class objc.Class
}

// An interface definition for the [MutableIndexSet] class.
type IMutableIndexSet interface {
	IIndexSet
	AddIndex(value uint)
	RemoveIndexes(indexSet IIndexSet)
	RemoveIndex(value uint)
	RemoveIndexesInRange(range_ Range)
}

// A mutable collection of unique integer values that represent indexes in another collection.
//
// In Swift, this type bridges to ; use when you need reference semantics or other Foundation-specific behavior. The class represents a mutable collection of unique unsigned integers, known as because of the way they are used. This collection is referred to as a . The inclusive range of valid indexes is ; trying to use indexes outside this range is invalid. The values in a mutable index set are always sorted, so the order in which values are added is irrelevant. Do not subclass the class.


// A mutable collection of unique integer values that represent indexes in another collection.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableIndexSet

type MutableIndexSet struct {
	IndexSet
}

// MutableIndexSetFrom constructs a [MutableIndexSet] from an unsafe.Pointer.
//
// A mutable collection of unique integer values that represent indexes in another collection.
func MutableIndexSetFrom(ptr unsafe.Pointer) MutableIndexSet {
	return MutableIndexSet{
		IndexSet: IndexSetFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (mc _MutableIndexSetClass) Alloc() MutableIndexSet {
	rv := objc.Send[MutableIndexSet](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MutableIndexSetClass) New() MutableIndexSet {
	rv := objc.Send[MutableIndexSet](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutableIndexSet) Init() MutableIndexSet {
	rv := objc.Send[MutableIndexSet](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutableIndexSet) Autorelease() MutableIndexSet {
	rv := objc.Send[MutableIndexSet](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutableIndexSet creates a new MutableIndexSet instance.
func NewMutableIndexSet() MutableIndexSet {
	return getMutableIndexSetClass().New()
}




// Adds an index to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableIndexSet/add(_:)-6dtkj

func (m_ MutableIndexSet) AddIndex(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addIndex:"), value)
}



// Removes the indexes in an index set from the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableIndexSet/remove(_:)-196u2

func (m_ MutableIndexSet) RemoveIndexes(indexSet IIndexSet) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeIndexes:"), indexSet)
}



// Removes an index from the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableIndexSet/remove(_:)-5li0r

func (m_ MutableIndexSet) RemoveIndex(value uint) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeIndex:"), value)
}



// Removes the indexes in an index range from the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSMutableIndexSet/remove(in:)

func (m_ MutableIndexSet) RemoveIndexesInRange(range_ Range) {
	objc.Send[objc.ID](m_.ID, objc.Sel("removeIndexesInRange:"), range_)
}



