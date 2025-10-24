// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class NSLayoutDimension */


/* debug [class_header]: Header for NSLayoutDimension */
// The class instance for the [LayoutDimension] class.
var (
	LayoutDimensionClass     _LayoutDimensionClass
	LayoutDimensionClassOnce sync.Once
)

func getLayoutDimensionClass() _LayoutDimensionClass {
	LayoutDimensionClassOnce.Do(func() {
		LayoutDimensionClass = _LayoutDimensionClass{objc.GetClass("NSLayoutDimension")}
	})
	return LayoutDimensionClass
}

type _LayoutDimensionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for LayoutDimension */
// An interface definition for the [LayoutDimension] class.
type ILayoutDimension interface {
	ILayoutAnchor
	
/* debug [class_interface_properties]: Properties for LayoutDimension */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for LayoutDimension */
	// methods:
	ConstraintEqualToAnchorMultiplier(anchor ILayoutDimension, m float64) ILayoutConstraint
	ConstraintEqualToAnchorMultiplierConstant(anchor ILayoutDimension, m float64, c float64) ILayoutConstraint
	ConstraintEqualToConstant(c float64) ILayoutConstraint
	ConstraintGreaterThanOrEqualToAnchorMultiplier(anchor ILayoutDimension, m float64) ILayoutConstraint
	ConstraintGreaterThanOrEqualToAnchorMultiplierConstant(anchor ILayoutDimension, m float64, c float64) ILayoutConstraint
	ConstraintGreaterThanOrEqualToConstant(c float64) ILayoutConstraint
	ConstraintLessThanOrEqualToAnchorMultiplier(anchor ILayoutDimension, m float64) ILayoutConstraint
	ConstraintLessThanOrEqualToAnchorMultiplierConstant(anchor ILayoutDimension, m float64, c float64) ILayoutConstraint
	ConstraintLessThanOrEqualToConstant(c float64) ILayoutConstraint
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for LayoutDimension */
// Alloc allocates a new instance without initialization.
func (lc _LayoutDimensionClass) Alloc() LayoutDimension {
	rv := objc.Send[LayoutDimension](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _LayoutDimensionClass) New() LayoutDimension {
	rv := objc.Send[LayoutDimension](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LayoutDimension) Init() LayoutDimension {
	rv := objc.Send[LayoutDimension](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LayoutDimension) Autorelease() LayoutDimension {
	rv := objc.Send[LayoutDimension](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLayoutDimension creates a new LayoutDimension instance.
func NewLayoutDimension() LayoutDimension {
	return getLayoutDimensionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for LayoutDimension */
// A factory class for creating size-based layout constraint objects using a fluent API.
//
// Use these constraints to programmatically define your layout using Auto Layout. All sizes are measured in points. In addition to providing size-specific methods for creating constraints, this class adds type information to the methods inherited from . Specifically, the generic methods declared by must now take a matching object. For more information on using layout anchors, see .


// A factory class for creating size-based layout constraint objects using a fluent API.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutDimension
type LayoutDimension struct {
	LayoutAnchor
}

// LayoutDimensionFrom constructs a [LayoutDimension] from an unsafe.Pointer.
//
// A factory class for creating size-based layout constraint objects using a fluent API.
func LayoutDimensionFrom(ptr unsafe.Pointer) LayoutDimension {
	return LayoutDimension{
		LayoutAnchor: LayoutAnchorFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for LayoutDimension *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for LayoutDimension */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for LayoutDimension */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for LayoutDimension */

// Returns a constraint that defines the anchor’s size attribute as equal to the specified anchor multiplied by the constant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutDimension/constraint(equalTo:multiplier:)
func (l_ LayoutDimension) ConstraintEqualToAnchorMultiplier(anchor ILayoutDimension, m float64) ILayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("constraintEqualToAnchor:multiplier:"), anchor, m)
	return rv
}/* debug [instance_methods/method]: ConstraintEqualToAnchorMultiplier */


// Returns a constraint that defines the anchor’s size attribute as equal to the specified size attribute multiplied by a constant plus an offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutDimension/constraint(equalTo:multiplier:constant:)
func (l_ LayoutDimension) ConstraintEqualToAnchorMultiplierConstant(anchor ILayoutDimension, m float64, c float64) ILayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("constraintEqualToAnchor:multiplier:constant:"), anchor, m, c)
	return rv
}/* debug [instance_methods/method]: ConstraintEqualToAnchorMultiplierConstant */


// Returns a constraint that defines a constant size for the anchor’s size attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutDimension/constraint(equalToConstant:)
func (l_ LayoutDimension) ConstraintEqualToConstant(c float64) ILayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("constraintEqualToConstant:"), c)
	return rv
}/* debug [instance_methods/method]: ConstraintEqualToConstant */


// Returns a constraint that defines the anchor’s size attribute as greater than or equal to the specified anchor multiplied by the constant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutDimension/constraint(greaterThanOrEqualTo:multiplier:)
func (l_ LayoutDimension) ConstraintGreaterThanOrEqualToAnchorMultiplier(anchor ILayoutDimension, m float64) ILayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("constraintGreaterThanOrEqualToAnchor:multiplier:"), anchor, m)
	return rv
}/* debug [instance_methods/method]: ConstraintGreaterThanOrEqualToAnchorMultiplier */


// Returns a constraint that defines the anchor’s size attribute as greater than or equal to the specified anchor multiplied by the constant plus an offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutDimension/constraint(greaterThanOrEqualTo:multiplier:constant:)
func (l_ LayoutDimension) ConstraintGreaterThanOrEqualToAnchorMultiplierConstant(anchor ILayoutDimension, m float64, c float64) ILayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("constraintGreaterThanOrEqualToAnchor:multiplier:constant:"), anchor, m, c)
	return rv
}/* debug [instance_methods/method]: ConstraintGreaterThanOrEqualToAnchorMultiplierConstant */


// Returns a constraint that defines the minimum size for the anchor’s size attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutDimension/constraint(greaterThanOrEqualToConstant:)
func (l_ LayoutDimension) ConstraintGreaterThanOrEqualToConstant(c float64) ILayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("constraintGreaterThanOrEqualToConstant:"), c)
	return rv
}/* debug [instance_methods/method]: ConstraintGreaterThanOrEqualToConstant */


// Returns a constraint that defines the anchor’s size attribute as less than or equal to the specified anchor multiplied by the constant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutDimension/constraint(lessThanOrEqualTo:multiplier:)
func (l_ LayoutDimension) ConstraintLessThanOrEqualToAnchorMultiplier(anchor ILayoutDimension, m float64) ILayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("constraintLessThanOrEqualToAnchor:multiplier:"), anchor, m)
	return rv
}/* debug [instance_methods/method]: ConstraintLessThanOrEqualToAnchorMultiplier */


// Returns a constraint that defines the anchor’s size attribute as greater than or equal to the specified anchor multiplied by the constant plus an offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutDimension/constraint(lessThanOrEqualTo:multiplier:constant:)
func (l_ LayoutDimension) ConstraintLessThanOrEqualToAnchorMultiplierConstant(anchor ILayoutDimension, m float64, c float64) ILayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("constraintLessThanOrEqualToAnchor:multiplier:constant:"), anchor, m, c)
	return rv
}/* debug [instance_methods/method]: ConstraintLessThanOrEqualToAnchorMultiplierConstant */


// Returns a constraint that defines the maximum size for the anchor’s size attribute.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutDimension/constraint(lessThanOrEqualToConstant:)
func (l_ LayoutDimension) ConstraintLessThanOrEqualToConstant(c float64) ILayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("constraintLessThanOrEqualToConstant:"), c)
	return rv
}/* debug [instance_methods/method]: ConstraintLessThanOrEqualToConstant */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for LayoutDimension */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSLayoutDimension */



