// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSCollectionLayoutEdgeSpacing */


/* debug [class_header]: Header for NSCollectionLayoutEdgeSpacing */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CollectionLayoutEdgeSpacing */
// An interface definition for the [CollectionLayoutEdgeSpacing] class.
type ICollectionLayoutEdgeSpacing interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CollectionLayoutEdgeSpacing */
	// properties:
	Bottom() ICollectionLayoutSpacing
	Leading() ICollectionLayoutSpacing
	Top() ICollectionLayoutSpacing
	Trailing() ICollectionLayoutSpacing
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CollectionLayoutEdgeSpacing */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CollectionLayoutEdgeSpacing */
// Alloc allocates a new instance without initialization.
func (cc _CollectionLayoutEdgeSpacingClass) Alloc() CollectionLayoutEdgeSpacing {
	rv := objc.Send[CollectionLayoutEdgeSpacing](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CollectionLayoutEdgeSpacing */
// An object that defines the space around the edges of items in a collection view.
//
// You use edge spacing to create additional spacing around the edges of an item to adjust the position of the item in relation to its container and other items. The leading and trailing spaces within edge spacing differ in left-to-right versus right-to-left environments. In a left-to-right environment, the leading space is on the left, and the trailing space is on the right. In a right-to-left environment, the leading space is on the right, and the trailing space is on the left. This difference ensures that your collection view layout is built with support for right-to-left languages. The following diagram shows the difference between adding 2 points of trailing edge spacing in a left-to-right versus a right-to-left environment.


// An object that defines the space around the edges of items in a collection view.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CollectionLayoutEdgeSpacing */

// Creates an edge spacing object with the specified leading, top, trailing, and bottom spacing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutEdgeSpacing/init(leading:top:trailing:bottom:)
func NewCollectionLayoutEdgeSpacingForLeadingTopTrailingBottom(leading ICollectionLayoutSpacing, top ICollectionLayoutSpacing, trailing ICollectionLayoutSpacing, bottom ICollectionLayoutSpacing) CollectionLayoutEdgeSpacing {
	rv := objc.Send[CollectionLayoutEdgeSpacing](objc.ID(getCollectionLayoutEdgeSpacingClass().class), objc.Sel("spacingForLeading:top:trailing:bottom:"), leading, top, trailing, bottom)
	return rv
}/* debug [class_init_methods/constructor]: NewCollectionLayoutEdgeSpacingForLeadingTopTrailingBottom */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CollectionLayoutEdgeSpacing */

// Creates an edge spacing object with the specified leading, top, trailing, and bottom spacing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutEdgeSpacing/init(leading:top:trailing:bottom:)
func (cc _CollectionLayoutEdgeSpacingClass) SpacingForLeadingTopTrailingBottom(leading ICollectionLayoutSpacing, top ICollectionLayoutSpacing, trailing ICollectionLayoutSpacing, bottom ICollectionLayoutSpacing) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(cc.class), objc.Sel("spacingForLeading:top:trailing:bottom:"), leading, top, trailing, bottom)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SpacingForLeadingTopTrailingBottom) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CollectionLayoutEdgeSpacing */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CollectionLayoutEdgeSpacing */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CollectionLayoutEdgeSpacing */

// The bottom edge spacing value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutEdgeSpacing/bottom
func (c_ CollectionLayoutEdgeSpacing) Bottom() ICollectionLayoutSpacing {
	rv := objc.Send[CollectionLayoutSpacing](c_.ID, objc.Sel("bottom"))
	return rv
}/* debug [instance_properties/getter]: bottom */


// The leading edge spacing value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutEdgeSpacing/leading
func (c_ CollectionLayoutEdgeSpacing) Leading() ICollectionLayoutSpacing {
	rv := objc.Send[CollectionLayoutSpacing](c_.ID, objc.Sel("leading"))
	return rv
}/* debug [instance_properties/getter]: leading */


// The top edge spacing value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutEdgeSpacing/top
func (c_ CollectionLayoutEdgeSpacing) Top() ICollectionLayoutSpacing {
	rv := objc.Send[CollectionLayoutSpacing](c_.ID, objc.Sel("top"))
	return rv
}/* debug [instance_properties/getter]: top */


// The trailing edge spacing value.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSCollectionLayoutEdgeSpacing/trailing
func (c_ CollectionLayoutEdgeSpacing) Trailing() ICollectionLayoutSpacing {
	rv := objc.Send[CollectionLayoutSpacing](c_.ID, objc.Sel("trailing"))
	return rv
}/* debug [instance_properties/getter]: trailing */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSCollectionLayoutEdgeSpacing */


