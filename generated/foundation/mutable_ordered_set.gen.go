// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableOrderedSet] class.
var (
	mutableOrderedSetClass     _MutableOrderedSetClass
	mutableOrderedSetClassOnce sync.Once
)

func getMutableOrderedSetClass() _MutableOrderedSetClass {
	mutableOrderedSetClassOnce.Do(func() {
		mutableOrderedSetClass = _MutableOrderedSetClass{objc.GetClass("NSMutableOrderedSet")}
	})
	return mutableOrderedSetClass
}

type _MutableOrderedSetClass struct {
	class objc.Class
}

// An interface definition for the [MutableOrderedSet] class.
type IMutableOrderedSet interface {
	IOrderedSet
}

// A dynamic, ordered collection of unique objects. [Full Topic]
//
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




