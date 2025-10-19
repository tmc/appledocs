// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [CollectionViewFlowLayout] class.
var (
	collectionViewFlowLayoutClass     _CollectionViewFlowLayoutClass
	collectionViewFlowLayoutClassOnce sync.Once
)

func getCollectionViewFlowLayoutClass() _CollectionViewFlowLayoutClass {
	collectionViewFlowLayoutClassOnce.Do(func() {
		collectionViewFlowLayoutClass = _CollectionViewFlowLayoutClass{objc.GetClass("NSCollectionViewFlowLayout")}
	})
	return collectionViewFlowLayoutClass
}

type _CollectionViewFlowLayoutClass struct {
	class objc.Class
}

// An interface definition for the [CollectionViewFlowLayout] class.
type ICollectionViewFlowLayout interface {
	ICollectionViewLayout
}

// A layout that organizes items into a flexible and configurable arrangement.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionViewFlowLayout
type CollectionViewFlowLayout struct {
	CollectionViewLayout
}

// CollectionViewFlowLayoutFrom constructs a [CollectionViewFlowLayout] from an unsafe.Pointer.
//
// A layout that organizes items into a flexible and configurable arrangement.
func CollectionViewFlowLayoutFrom(ptr unsafe.Pointer) CollectionViewFlowLayout {
	return CollectionViewFlowLayout{
		CollectionViewLayout: CollectionViewLayoutFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CollectionViewFlowLayoutClass) Alloc() CollectionViewFlowLayout {
	rv := objc.Send[CollectionViewFlowLayout](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CollectionViewFlowLayoutClass) New() CollectionViewFlowLayout {
	rv := objc.Send[CollectionViewFlowLayout](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionViewFlowLayout) Init() CollectionViewFlowLayout {
	rv := objc.Send[CollectionViewFlowLayout](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionViewFlowLayout) Autorelease() CollectionViewFlowLayout {
	rv := objc.Send[CollectionViewFlowLayout](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionViewFlowLayout creates a new CollectionViewFlowLayout instance.
func NewCollectionViewFlowLayout() CollectionViewFlowLayout {
	return getCollectionViewFlowLayoutClass().New()
}




