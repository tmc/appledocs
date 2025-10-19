// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [MutableIndexSet] class.
var (
	mutableIndexSetClass     _MutableIndexSetClass
	mutableIndexSetClassOnce sync.Once
)

func getMutableIndexSetClass() _MutableIndexSetClass {
	mutableIndexSetClassOnce.Do(func() {
		mutableIndexSetClass = _MutableIndexSetClass{objc.GetClass("NSMutableIndexSet")}
	})
	return mutableIndexSetClass
}

type _MutableIndexSetClass struct {
	class objc.Class
}

// An interface definition for the [MutableIndexSet] class.
type IMutableIndexSet interface {
	IIndexSet
}

// A mutable collection of unique integer values that represent indexes in another collection. [Full Topic]
//
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




