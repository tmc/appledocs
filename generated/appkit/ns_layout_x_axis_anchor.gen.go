// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSLayoutXAxisAnchor */


/* debug [class_header]: Header for NSLayoutXAxisAnchor */
// The class instance for the [LayoutXAxisAnchor] class.
var (
	LayoutXAxisAnchorClass     _LayoutXAxisAnchorClass
	LayoutXAxisAnchorClassOnce sync.Once
)

func getLayoutXAxisAnchorClass() _LayoutXAxisAnchorClass {
	LayoutXAxisAnchorClassOnce.Do(func() {
		LayoutXAxisAnchorClass = _LayoutXAxisAnchorClass{objc.GetClass("NSLayoutXAxisAnchor")}
	})
	return LayoutXAxisAnchorClass
}

type _LayoutXAxisAnchorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for LayoutXAxisAnchor */
// An interface definition for the [LayoutXAxisAnchor] class.
type ILayoutXAxisAnchor interface {
	ILayoutAnchor
	
/* debug [class_interface_properties]: Properties for LayoutXAxisAnchor */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for LayoutXAxisAnchor */
	// methods:
	AnchorWithOffsetToAnchor(otherAnchor ILayoutXAxisAnchor) ILayoutDimension
	ConstraintEqualToSystemSpacingAfterAnchorMultiplier(anchor ILayoutXAxisAnchor, multiplier float64) ILayoutConstraint
	ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchorMultiplier(anchor ILayoutXAxisAnchor, multiplier float64) ILayoutConstraint
	ConstraintLessThanOrEqualToSystemSpacingAfterAnchorMultiplier(anchor ILayoutXAxisAnchor, multiplier float64) ILayoutConstraint
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for LayoutXAxisAnchor */
// Alloc allocates a new instance without initialization.
func (lc _LayoutXAxisAnchorClass) Alloc() LayoutXAxisAnchor {
	rv := objc.Send[LayoutXAxisAnchor](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _LayoutXAxisAnchorClass) New() LayoutXAxisAnchor {
	rv := objc.Send[LayoutXAxisAnchor](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LayoutXAxisAnchor) Init() LayoutXAxisAnchor {
	rv := objc.Send[LayoutXAxisAnchor](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LayoutXAxisAnchor) Autorelease() LayoutXAxisAnchor {
	rv := objc.Send[LayoutXAxisAnchor](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLayoutXAxisAnchor creates a new LayoutXAxisAnchor instance.
func NewLayoutXAxisAnchor() LayoutXAxisAnchor {
	return getLayoutXAxisAnchorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for LayoutXAxisAnchor */
// A factory class for creating horizontal layout constraint objects using a fluent API.
//
// adds type information to the methods inherited from . Specifically, the generic methods declared by must now take a matching object. For more information on using layout anchors, see .


// A factory class for creating horizontal layout constraint objects using a fluent API.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutXAxisAnchor
type LayoutXAxisAnchor struct {
	LayoutAnchor
}

// LayoutXAxisAnchorFrom constructs a [LayoutXAxisAnchor] from an unsafe.Pointer.
//
// A factory class for creating horizontal layout constraint objects using a fluent API.
func LayoutXAxisAnchorFrom(ptr unsafe.Pointer) LayoutXAxisAnchor {
	return LayoutXAxisAnchor{
		LayoutAnchor: LayoutAnchorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for LayoutXAxisAnchor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for LayoutXAxisAnchor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for LayoutXAxisAnchor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for LayoutXAxisAnchor */

// Creates a layout dimension object from two anchors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutXAxisAnchor/anchorWithOffset(to:)
func (l_ LayoutXAxisAnchor) AnchorWithOffsetToAnchor(otherAnchor ILayoutXAxisAnchor) ILayoutDimension {
	rv := objc.Send[LayoutDimension](l_.ID, objc.Sel("anchorWithOffsetToAnchor:"), otherAnchor)
	return rv
}/* debug [instance_methods/method]: AnchorWithOffsetToAnchor */


// Returns a constraint that defines by how much the current anchor trails the specified anchor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutXAxisAnchor/constraint(equalToSystemSpacingAfter:multiplier:)
func (l_ LayoutXAxisAnchor) ConstraintEqualToSystemSpacingAfterAnchorMultiplier(anchor ILayoutXAxisAnchor, multiplier float64) ILayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("constraintEqualToSystemSpacingAfterAnchor:multiplier:"), anchor, multiplier)
	return rv
}/* debug [instance_methods/method]: ConstraintEqualToSystemSpacingAfterAnchorMultiplier */


// Returns a constraint that defines the minimum amount by which the current anchor trails the specified anchor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutXAxisAnchor/constraint(greaterThanOrEqualToSystemSpacingAfter:multiplier:)
func (l_ LayoutXAxisAnchor) ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchorMultiplier(anchor ILayoutXAxisAnchor, multiplier float64) ILayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("constraintGreaterThanOrEqualToSystemSpacingAfterAnchor:multiplier:"), anchor, multiplier)
	return rv
}/* debug [instance_methods/method]: ConstraintGreaterThanOrEqualToSystemSpacingAfterAnchorMultiplier */


// Returns a constraint that defines the maximum amount by which the current anchor trails the specified anchor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutXAxisAnchor/constraint(lessThanOrEqualToSystemSpacingAfter:multiplier:)
func (l_ LayoutXAxisAnchor) ConstraintLessThanOrEqualToSystemSpacingAfterAnchorMultiplier(anchor ILayoutXAxisAnchor, multiplier float64) ILayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("constraintLessThanOrEqualToSystemSpacingAfterAnchor:multiplier:"), anchor, multiplier)
	return rv
}/* debug [instance_methods/method]: ConstraintLessThanOrEqualToSystemSpacingAfterAnchorMultiplier */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for LayoutXAxisAnchor */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSLayoutXAxisAnchor */



