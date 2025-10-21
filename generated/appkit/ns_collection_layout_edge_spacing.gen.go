// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CollectionLayoutEdgeSpacing] class.
var (
	CollectionLayoutEdgeSpacingClass     _CollectionLayoutEdgeSpacingClass
	CollectionLayoutEdgeSpacingClassOnce sync.Once
)

func getCollectionLayoutEdgeSpacingClass() _CollectionLayoutEdgeSpacingClass {
	CollectionLayoutEdgeSpacingClassOnce.Do(func() {
		CollectionLayoutEdgeSpacingClass = _CollectionLayoutEdgeSpacingClass{objc.GetClass("NSCollectionLayoutEdgeSpacing")}
	})
	return CollectionLayoutEdgeSpacingClass
}

type _CollectionLayoutEdgeSpacingClass struct {
	class objc.Class
}

// An interface definition for the [CollectionLayoutEdgeSpacing] class.
type ICollectionLayoutEdgeSpacing interface {
	objectivec.IObject
}

// An object that defines the space around the edges of items in a collection view.
//
// You use edge spacing to create additional spacing around the edges of an item to adjust the position of the item in relation to its container and other items. The leading and trailing spaces within edge spacing differ in left-to-right versus right-to-left environments. In a left-to-right environment, the leading space is on the left, and the trailing space is on the right. In a right-to-left environment, the leading space is on the right, and the trailing space is on the left. This difference ensures that your collection view layout is built with support for right-to-left languages. The following diagram shows the difference between adding 2 points of trailing edge spacing in a left-to-right versus a right-to-left environment.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutEdgeSpacing
type CollectionLayoutEdgeSpacing struct {
	objectivec.Object
}

// CollectionLayoutEdgeSpacingFrom constructs a [CollectionLayoutEdgeSpacing] from an unsafe.Pointer.
//
// An object that defines the space around the edges of items in a collection view.
func CollectionLayoutEdgeSpacingFrom(ptr unsafe.Pointer) CollectionLayoutEdgeSpacing {
	return CollectionLayoutEdgeSpacing{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CollectionLayoutEdgeSpacingClass) Alloc() CollectionLayoutEdgeSpacing {
	rv := objc.Send[CollectionLayoutEdgeSpacing](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CollectionLayoutEdgeSpacingClass) New() CollectionLayoutEdgeSpacing {
	rv := objc.Send[CollectionLayoutEdgeSpacing](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionLayoutEdgeSpacing) Init() CollectionLayoutEdgeSpacing {
	rv := objc.Send[CollectionLayoutEdgeSpacing](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionLayoutEdgeSpacing) Autorelease() CollectionLayoutEdgeSpacing {
	rv := objc.Send[CollectionLayoutEdgeSpacing](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionLayoutEdgeSpacing creates a new CollectionLayoutEdgeSpacing instance.
func NewCollectionLayoutEdgeSpacing() CollectionLayoutEdgeSpacing {
	return getCollectionLayoutEdgeSpacingClass().New()
}




