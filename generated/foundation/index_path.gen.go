// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [IndexPath] class.
var (
	indexPathClass     _IndexPathClass
	indexPathClassOnce sync.Once
)

func getIndexPathClass() _IndexPathClass {
	indexPathClassOnce.Do(func() {
		indexPathClass = _IndexPathClass{objc.GetClass("NSIndexPath")}
	})
	return indexPathClass
}

type _IndexPathClass struct {
	class objc.Class
}

// An interface definition for the [IndexPath] class.
type IIndexPath interface {
	objectivec.IObject
}

// A list of indexes that together represent the path to a specific location in a tree of nested arrays.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/NSIndexPath
type IndexPath struct {
	objectivec.Object
}

// IndexPathFrom constructs a [IndexPath] from an unsafe.Pointer.
//
// A list of indexes that together represent the path to a specific location in a tree of nested arrays.
func IndexPathFrom(ptr unsafe.Pointer) IndexPath {
	return IndexPath{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ic _IndexPathClass) Alloc() IndexPath {
	rv := objc.Send[IndexPath](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _IndexPathClass) New() IndexPath {
	rv := objc.Send[IndexPath](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ IndexPath) Init() IndexPath {
	rv := objc.Send[IndexPath](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ IndexPath) Autorelease() IndexPath {
	rv := objc.Send[IndexPath](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewIndexPath creates a new IndexPath instance.
func NewIndexPath() IndexPath {
	return getIndexPathClass().New()
}




