// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CollectionLayoutSize] class.
var (
	CollectionLayoutSizeClass     _CollectionLayoutSizeClass
	CollectionLayoutSizeClassOnce sync.Once
)

func getCollectionLayoutSizeClass() _CollectionLayoutSizeClass {
	CollectionLayoutSizeClassOnce.Do(func() {
		CollectionLayoutSizeClass = _CollectionLayoutSizeClass{objc.GetClass("NSCollectionLayoutSize")}
	})
	return CollectionLayoutSizeClass
}

type _CollectionLayoutSizeClass struct {
	class objc.Class
}

// An interface definition for the [CollectionLayoutSize] class.
type ICollectionLayoutSize interface {
	objectivec.IObject
	// properties:
	HeightDimension() ICollectionLayoutDimension
	WidthDimension() ICollectionLayoutDimension
	// methods:
}

// The width and the height of an item in a collection view.
//
// A size is a pair of dimensions ( ): a width dimension and a height dimension. Every component of a collection view layout has an explicit size.


// The width and the height of an item in a collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSize
type CollectionLayoutSize struct {
	objectivec.Object
}

// CollectionLayoutSizeFrom constructs a [CollectionLayoutSize] from an unsafe.Pointer.
//
// The width and the height of an item in a collection view.
func CollectionLayoutSizeFrom(ptr unsafe.Pointer) CollectionLayoutSize {
	return CollectionLayoutSize{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CollectionLayoutSizeClass) Alloc() CollectionLayoutSize {
	rv := objc.Send[CollectionLayoutSize](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CollectionLayoutSizeClass) New() CollectionLayoutSize {
	rv := objc.Send[CollectionLayoutSize](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionLayoutSize) Init() CollectionLayoutSize {
	rv := objc.Send[CollectionLayoutSize](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionLayoutSize) Autorelease() CollectionLayoutSize {
	rv := objc.Send[CollectionLayoutSize](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionLayoutSize creates a new CollectionLayoutSize instance.
func NewCollectionLayoutSize() CollectionLayoutSize {
	return getCollectionLayoutSizeClass().New()
}



// Creates a size with the specified width and height dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSize/init(widthDimension:heightDimension:)
func NewCollectionLayoutSizeWithWidthDimensionHeightDimension(width ICollectionLayoutDimension, height ICollectionLayoutDimension) CollectionLayoutSize {
	rv := objc.Send[CollectionLayoutSize](objc.ID(getCollectionLayoutSizeClass().class), objc.Sel("sizeWithWidthDimension:heightDimension:"), width, height)
	return rv
}



// Creates a size with the specified width and height dimensions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSize/init(widthDimension:heightDimension:)
func (cc _CollectionLayoutSizeClass) SizeWithWidthDimensionHeightDimension(width ICollectionLayoutDimension, height ICollectionLayoutDimension) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("sizeWithWidthDimension:heightDimension:"), width, height)
	return rv
}


// The height dimension of an item in a collection view layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSize/heightDimension
func (c_ CollectionLayoutSize) HeightDimension() ICollectionLayoutDimension {
	rv := objc.Send[CollectionLayoutDimension](c_.ID, objc.Sel("heightDimension"))
	return rv
}


// The width dimension of an item in a collection view layout.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSize/widthDimension
func (c_ CollectionLayoutSize) WidthDimension() ICollectionLayoutDimension {
	rv := objc.Send[CollectionLayoutDimension](c_.ID, objc.Sel("widthDimension"))
	return rv
}


