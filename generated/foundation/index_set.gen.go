// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [IndexSet] class.
var (
	indexSetClass     _IndexSetClass
	indexSetClassOnce sync.Once
)

func getIndexSetClass() _IndexSetClass {
	indexSetClassOnce.Do(func() {
		indexSetClass = _IndexSetClass{objc.GetClass("NSIndexSet")}
	})
	return indexSetClass
}

type _IndexSetClass struct {
	class objc.Class
}

// An interface definition for the [IndexSet] class.
type IIndexSet interface {
	objectivec.IObject
}

// An immutable collection of unique integer values that represent indexes in another collection. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexSet
type IndexSet struct {
	objectivec.Object
}

// IndexSetFrom constructs a [IndexSet] from an unsafe.Pointer.
//
// An immutable collection of unique integer values that represent indexes in another collection.
func IndexSetFrom(ptr unsafe.Pointer) IndexSet {
	return IndexSet{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _IndexSetClass) Alloc() IndexSet {
	rv := objc.Send[IndexSet](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IndexSetClass) New() IndexSet {
	rv := objc.Send[IndexSet](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IndexSet) Init() IndexSet {
	rv := objc.Send[IndexSet](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IndexSet) Autorelease() IndexSet {
	rv := objc.Send[IndexSet](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIndexSet creates a new IndexSet instance.
func NewIndexSet() IndexSet {
	return getIndexSetClass().New()
}




