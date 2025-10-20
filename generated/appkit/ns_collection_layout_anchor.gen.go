// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CollectionLayoutAnchor] class.
var (
	CollectionLayoutAnchorClass     _CollectionLayoutAnchorClass
	CollectionLayoutAnchorClassOnce sync.Once
)

func getCollectionLayoutAnchorClass() _CollectionLayoutAnchorClass {
	CollectionLayoutAnchorClassOnce.Do(func() {
		CollectionLayoutAnchorClass = _CollectionLayoutAnchorClass{objc.GetClass("NSCollectionLayoutAnchor")}
	})
	return CollectionLayoutAnchorClass
}

type _CollectionLayoutAnchorClass struct {
	class objc.Class
}

// An interface definition for the [CollectionLayoutAnchor] class.
type ICollectionLayoutAnchor interface {
	objectivec.IObject
}

// An object that defines how to attach a supplementary item to an item in a collection view.
//
// You use an anchor to attach a supplementary item to a specific item. An anchor contains information about where on the item your supplementary item is attached, including: An edge or set of edges. You can attach a supplementary item to a single edge, or to a corner by specifying two adjacent edges. An offset from the item. By default, the supplementary item is anchored within the specified edges of the item it’s attached to. You can change this location by providing a custom offset when you create an anchor.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutAnchor
type CollectionLayoutAnchor struct {
	objectivec.Object
}

// CollectionLayoutAnchorFrom constructs a [CollectionLayoutAnchor] from an unsafe.Pointer.
//
// An object that defines how to attach a supplementary item to an item in a collection view.
func CollectionLayoutAnchorFrom(ptr unsafe.Pointer) CollectionLayoutAnchor {
	return CollectionLayoutAnchor{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CollectionLayoutAnchorClass) Alloc() CollectionLayoutAnchor {
	rv := objc.Send[CollectionLayoutAnchor](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CollectionLayoutAnchorClass) New() CollectionLayoutAnchor {
	rv := objc.Send[CollectionLayoutAnchor](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionLayoutAnchor) Init() CollectionLayoutAnchor {
	rv := objc.Send[CollectionLayoutAnchor](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionLayoutAnchor) Autorelease() CollectionLayoutAnchor {
	rv := objc.Send[CollectionLayoutAnchor](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionLayoutAnchor creates a new CollectionLayoutAnchor instance.
func NewCollectionLayoutAnchor() CollectionLayoutAnchor {
	return getCollectionLayoutAnchorClass().New()
}




