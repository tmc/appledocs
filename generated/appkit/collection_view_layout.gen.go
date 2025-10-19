// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CollectionViewLayout] class.
var (
	collectionViewLayoutClass     _CollectionViewLayoutClass
	collectionViewLayoutClassOnce sync.Once
)

func getCollectionViewLayoutClass() _CollectionViewLayoutClass {
	collectionViewLayoutClassOnce.Do(func() {
		collectionViewLayoutClass = _CollectionViewLayoutClass{objc.GetClass("NSCollectionViewLayout")}
	})
	return collectionViewLayoutClass
}

type _CollectionViewLayoutClass struct {
	class objc.Class
}

// An interface definition for the [CollectionViewLayout] class.
type ICollectionViewLayout interface {
	objectivec.IObject
}

// An abstract base class that you subclass and use to generate layout information for a collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewLayout
type CollectionViewLayout struct {
	objectivec.Object
}

// CollectionViewLayoutFrom constructs a [CollectionViewLayout] from an unsafe.Pointer.
//
// An abstract base class that you subclass and use to generate layout information for a collection view.
func CollectionViewLayoutFrom(ptr unsafe.Pointer) CollectionViewLayout {
	return CollectionViewLayout{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CollectionViewLayoutClass) Alloc() CollectionViewLayout {
	rv := objc.Send[CollectionViewLayout](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CollectionViewLayoutClass) New() CollectionViewLayout {
	rv := objc.Send[CollectionViewLayout](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionViewLayout) Init() CollectionViewLayout {
	rv := objc.Send[CollectionViewLayout](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionViewLayout) Autorelease() CollectionViewLayout {
	rv := objc.Send[CollectionViewLayout](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionViewLayout creates a new CollectionViewLayout instance.
func NewCollectionViewLayout() CollectionViewLayout {
	return getCollectionViewLayoutClass().New()
}




