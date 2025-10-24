// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSCollectionLayoutSpacing */


/* debug [class_header]: Header for NSCollectionLayoutSpacing */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CollectionLayoutSpacing */
// An interface definition for the [CollectionLayoutSpacing] class.
type ICollectionLayoutSpacing interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CollectionLayoutSpacing */
	// properties:
	IsFixedSpacing() bool
	IsFlexibleSpacing() bool
	Spacing() float64
	IsFixed() bool
	SetIsFixed(value bool)
	IsFlexible() bool
	SetIsFlexible(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CollectionLayoutSpacing */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CollectionLayoutSpacing */
// Alloc allocates a new instance without initialization.
func (cc _CollectionLayoutSpacingClass) Alloc() CollectionLayoutSpacing {
	rv := objc.Send[CollectionLayoutSpacing](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CollectionLayoutSpacing */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CollectionLayoutSpacing *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CollectionLayoutSpacing */

// Creates a space equivalent to the specified number of points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSpacing/fixed(_:)
func (cc _CollectionLayoutSpacingClass) FixedSpacing(fixedSpacing float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("fixedSpacing:"), fixedSpacing)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FixedSpacing) */


// Creates a space equivalent to or greater than the specified number of points, depending on the available space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSpacing/flexible(_:)
func (cc _CollectionLayoutSpacingClass) FlexibleSpacing(flexibleSpacing float64) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("flexibleSpacing:"), flexibleSpacing)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FlexibleSpacing) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CollectionLayoutSpacing */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CollectionLayoutSpacing */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CollectionLayoutSpacing */

// A Boolean value that indicates whether the space is fixed to a specific number of points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSpacing/isFixed
func (c_ CollectionLayoutSpacing) IsFixedSpacing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFixedSpacing"))
	return rv
}/* debug [instance_properties/getter]: isFixedSpacing */


// A Boolean value that indicates whether the space is flexible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSpacing/isFlexible
func (c_ CollectionLayoutSpacing) IsFlexibleSpacing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFlexibleSpacing"))
	return rv
}/* debug [instance_properties/getter]: isFlexibleSpacing */


// The floating-point value of the space.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutSpacing/spacing
func (c_ CollectionLayoutSpacing) Spacing() float64 {
	rv := objc.Send[float64](c_.ID, objc.Sel("spacing"))
	return rv
}/* debug [instance_properties/getter]: spacing */


// A Boolean value that indicates whether the space is fixed to a specific number of points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionlayoutspacing/isfixed
func (c_ CollectionLayoutSpacing) IsFixed() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFixed"))
	return rv
}/* debug [instance_properties/getter]: isFixed */


// A Boolean value that indicates whether the space is fixed to a specific number of points.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionlayoutspacing/isfixed
func (c_ CollectionLayoutSpacing) SetIsFixed(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsFixed:"), value)
}/* debug [instance_properties/setter]: isFixed */


// A Boolean value that indicates whether the space is flexible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionlayoutspacing/isflexible
func (c_ CollectionLayoutSpacing) IsFlexible() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isFlexible"))
	return rv
}/* debug [instance_properties/getter]: isFlexible */


// A Boolean value that indicates whether the space is flexible.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionlayoutspacing/isflexible
func (c_ CollectionLayoutSpacing) SetIsFlexible(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsFlexible:"), value)
}/* debug [instance_properties/setter]: isFlexible */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCollectionLayoutSpacing */



