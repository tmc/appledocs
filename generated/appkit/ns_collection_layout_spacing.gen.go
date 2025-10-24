// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CollectionLayoutSpacing] class.
var (
	CollectionLayoutSpacingClass     _CollectionLayoutSpacingClass
	CollectionLayoutSpacingClassOnce sync.Once
)

func getCollectionLayoutSpacingClass() _CollectionLayoutSpacingClass {
	CollectionLayoutSpacingClassOnce.Do(func() {
		CollectionLayoutSpacingClass = _CollectionLayoutSpacingClass{objc.GetClass("NSCollectionLayoutSpacing")}
	})
	return CollectionLayoutSpacingClass
}

type _CollectionLayoutSpacingClass struct {
	class objc.Class
}

// An interface definition for the [CollectionLayoutSpacing] class.
type ICollectionLayoutSpacing interface {
	objectivec.IObject
	// properties:
	IsFixedSpacing() bool
	IsFlexibleSpacing() bool
	Spacing() float64
	IsFixed() bool
	SetIsFixed(value bool)
	IsFlexible() bool
	SetIsFlexible(value bool)
	// methods:
}

// An object that defines the space between or around items in a collection view.
//
// In a collection view layout, you use a spacing object to specify both the amount of space and the way in which it’s calculated. You can express spacing using fixed or flexible spacing. Use to provide an exact amount of space. For example, the following code creates exactly 200 points of space between the items in the group. Use to provide a minimum amount of space that can grow as more space becomes available. For example, the following code creates at least 200 points of space between the items in the group. As more space becomes available, items are respaced evenly in the additional space.


// An object that defines the space between or around items in a collection view.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSpacing
type CollectionLayoutSpacing struct {
	objectivec.Object
}

// CollectionLayoutSpacingFrom constructs a [CollectionLayoutSpacing] from an unsafe.Pointer.
//
// An object that defines the space between or around items in a collection view.
func CollectionLayoutSpacingFrom(ptr unsafe.Pointer) CollectionLayoutSpacing {
	return CollectionLayoutSpacing{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CollectionLayoutSpacingClass) Alloc() CollectionLayoutSpacing {
	rv := objc.Send[CollectionLayoutSpacing](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CollectionLayoutSpacingClass) New() CollectionLayoutSpacing {
	rv := objc.Send[CollectionLayoutSpacing](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CollectionLayoutSpacing) Init() CollectionLayoutSpacing {
	rv := objc.Send[CollectionLayoutSpacing](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CollectionLayoutSpacing) Autorelease() CollectionLayoutSpacing {
	rv := objc.Send[CollectionLayoutSpacing](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCollectionLayoutSpacing creates a new CollectionLayoutSpacing instance.
func NewCollectionLayoutSpacing() CollectionLayoutSpacing {
	return getCollectionLayoutSpacingClass().New()
}



// Creates a space equivalent to the specified number of points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSpacing/fixed(_:)
func (cc _CollectionLayoutSpacingClass) FixedSpacing(fixedSpacing float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("fixedSpacing:"), fixedSpacing)
	return rv
}


// Creates a space equivalent to or greater than the specified number of points, depending on the available space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSpacing/flexible(_:)
func (cc _CollectionLayoutSpacingClass) FlexibleSpacing(flexibleSpacing float64) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(cc.class), objc.Sel("flexibleSpacing:"), flexibleSpacing)
	return rv
}


// A Boolean value that indicates whether the space is fixed to a specific number of points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSpacing/isFixed
func (c_ CollectionLayoutSpacing) IsFixedSpacing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFixedSpacing"))
	return rv
}


// A Boolean value that indicates whether the space is flexible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSpacing/isFlexible
func (c_ CollectionLayoutSpacing) IsFlexibleSpacing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFlexibleSpacing"))
	return rv
}


// The floating-point value of the space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSpacing/spacing
func (c_ CollectionLayoutSpacing) Spacing() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("spacing"))
	return rv
}


// A Boolean value that indicates whether the space is fixed to a specific number of points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionlayoutspacing/isfixed
func (c_ CollectionLayoutSpacing) IsFixed() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFixed"))
	return rv
}


// A Boolean value that indicates whether the space is fixed to a specific number of points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionlayoutspacing/isfixed
func (c_ CollectionLayoutSpacing) SetIsFixed(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsFixed:"), value)
}


// A Boolean value that indicates whether the space is flexible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionlayoutspacing/isflexible
func (c_ CollectionLayoutSpacing) IsFlexible() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFlexible"))
	return rv
}


// A Boolean value that indicates whether the space is flexible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionlayoutspacing/isflexible
func (c_ CollectionLayoutSpacing) SetIsFlexible(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsFlexible:"), value)
}



