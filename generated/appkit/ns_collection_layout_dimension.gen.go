// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)





// The class instance for the [CollectionLayoutDimension] class.
var (
	CollectionLayoutDimensionClass     _CollectionLayoutDimensionClass
	CollectionLayoutDimensionClassOnce sync.Once
)

func getCollectionLayoutDimensionClass() _CollectionLayoutDimensionClass {
	CollectionLayoutDimensionClassOnce.Do(func() {
		CollectionLayoutDimensionClass = _CollectionLayoutDimensionClass{objc.GetClass("NSCollectionLayoutDimension")}
	})
	return CollectionLayoutDimensionClass
}

type _CollectionLayoutDimensionClass struct {
	class objc.Class
}





// An interface definition for the [CollectionLayoutDimension] class.
type ICollectionLayoutDimension interface {
	objectivec.IObject
	

	// properties:
	Dimension() float64
	IsAbsolute() bool
	IsEstimated() bool
	IsFractionalHeight() bool
	IsFractionalWidth() bool


	

	// methods:


}





// Alloc allocates a new instance without initialization.
func (cc _CollectionLayoutDimensionClass) Alloc() CollectionLayoutDimension {
	rv := objc.Send[CollectionLayoutDimension](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CollectionLayoutDimensionClass) New() CollectionLayoutDimension {
	rv := objc.Send[CollectionLayoutDimension](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionLayoutDimension) Init() CollectionLayoutDimension {
	rv := objc.Send[CollectionLayoutDimension](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionLayoutDimension) Autorelease() CollectionLayoutDimension {
	rv := objc.Send[CollectionLayoutDimension](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionLayoutDimension creates a new CollectionLayoutDimension instance.
func NewCollectionLayoutDimension() CollectionLayoutDimension {
	return getCollectionLayoutDimensionClass().New()
}





// An individual dimension representing an item’s width or height in a collection view.
//
// Each item in a collection view has an explicit width dimension and height dimension, which combine to define the item’s size ( ). You can express an item’s dimensions using an absolute, estimated, or fractional value. Use an to specify exact dimensions, like a 44 x 44 point square: Use an if the size of your content might change at runtime, such as when data is loaded or in response to a change in system font size. You provide an initial estimated size and the system computes the actual value later. Use a to define a value that’s relative to a dimension of the item’s container. This option simplifies specifying aspect ratios. For example, the following item has a width and a height that are both equal to 20% of its container’s width, creating a square that grows and shrinks as the size of its container changes.


// An individual dimension representing an item’s width or height in a collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDimension
type CollectionLayoutDimension struct {
	objectivec.Object
}

// CollectionLayoutDimensionFrom constructs a [CollectionLayoutDimension] from an unsafe.Pointer.
//
// An individual dimension representing an item’s width or height in a collection view.
func CollectionLayoutDimensionFrom(ptr unsafe.Pointer) CollectionLayoutDimension {
	return CollectionLayoutDimension{objectivec.Object{objc.ID(ptr)}}
}










// Creates a dimension with an absolute point value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDimension/absolute(_:)
func (cc _CollectionLayoutDimensionClass) AbsoluteDimension(absoluteDimension float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("absoluteDimension:"), absoluteDimension)
	return rv
}


// Creates a dimension with an estimated point value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDimension/estimated(_:)
func (cc _CollectionLayoutDimensionClass) EstimatedDimension(estimatedDimension float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("estimatedDimension:"), estimatedDimension)
	return rv
}


// Creates a dimension that is computed as a fraction of the height of the containing group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDimension/fractionalHeight(_:)
func (cc _CollectionLayoutDimensionClass) FractionalHeightDimension(fractionalHeight float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("fractionalHeightDimension:"), fractionalHeight)
	return rv
}


// Creates a dimension that is computed as a fraction of the width of the containing group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDimension/fractionalWidth(_:)
func (cc _CollectionLayoutDimensionClass) FractionalWidthDimension(fractionalWidth float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("fractionalWidthDimension:"), fractionalWidth)
	return rv
}

















// The floating-point value of the dimension.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDimension/dimension
func (c_ CollectionLayoutDimension) Dimension() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("dimension"))
	return rv
}


// A Boolean value that indicates whether the dimension is expressed as an absolute value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDimension/isAbsolute
func (c_ CollectionLayoutDimension) IsAbsolute() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isAbsolute"))
	return rv
}


// A Boolean value that indicates whether the dimension is expressed as an estimated value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDimension/isEstimated
func (c_ CollectionLayoutDimension) IsEstimated() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEstimated"))
	return rv
}


// A Boolean value that indicates whether the dimension is expressed as a fraction of its container’s height.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDimension/isFractionalHeight
func (c_ CollectionLayoutDimension) IsFractionalHeight() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFractionalHeight"))
	return rv
}


// A Boolean value that indicates whether the dimension is expressed as a fraction of its container’s width.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutDimension/isFractionalWidth
func (c_ CollectionLayoutDimension) IsFractionalWidth() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFractionalWidth"))
	return rv
}








