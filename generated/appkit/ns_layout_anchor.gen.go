// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSLayoutAnchor */


/* debug [class_header]: Header for NSLayoutAnchor */
// The class instance for the [LayoutAnchor] class.
var (
	LayoutAnchorClass     _LayoutAnchorClass
	LayoutAnchorClassOnce sync.Once
)

func getLayoutAnchorClass() _LayoutAnchorClass {
	LayoutAnchorClassOnce.Do(func() {
		LayoutAnchorClass = _LayoutAnchorClass{objc.GetClass("NSLayoutAnchor")}
	})
	return LayoutAnchorClass
}

type _LayoutAnchorClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for LayoutAnchor */
// An interface definition for the [LayoutAnchor] class.
type ILayoutAnchor interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for LayoutAnchor */
	// properties:
	ConstraintsAffectingLayout() []LayoutConstraint
	HasAmbiguousLayout() bool
	Item() objc.ID
	Name() objc.IObject /* cross-framework: NSString */
	BottomAnchor() ILayoutYAxisAnchor
	SetBottomAnchor(value ILayoutYAxisAnchor)
	LeadingAnchor() ILayoutXAxisAnchor
	SetLeadingAnchor(value ILayoutXAxisAnchor)
	LeftAnchor() ILayoutXAxisAnchor
	SetLeftAnchor(value ILayoutXAxisAnchor)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for LayoutAnchor */
	// methods:
	ConstraintEqualToAnchor(anchor unsafe.Pointer) ILayoutConstraint
	ConstraintEqualToAnchorConstant(anchor unsafe.Pointer, c float64) ILayoutConstraint
	ConstraintGreaterThanOrEqualToAnchor(anchor unsafe.Pointer) ILayoutConstraint
	ConstraintGreaterThanOrEqualToAnchorConstant(anchor unsafe.Pointer, c float64) ILayoutConstraint
	ConstraintLessThanOrEqualToAnchor(anchor unsafe.Pointer) ILayoutConstraint
	ConstraintLessThanOrEqualToAnchorConstant(anchor unsafe.Pointer, c float64) ILayoutConstraint
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for LayoutAnchor */
// Alloc allocates a new instance without initialization.
func (lc _LayoutAnchorClass) Alloc() LayoutAnchor {
	rv := objc.Send[LayoutAnchor](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _LayoutAnchorClass) New() LayoutAnchor {
	rv := objc.Send[LayoutAnchor](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LayoutAnchor) Init() LayoutAnchor {
	rv := objc.Send[LayoutAnchor](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LayoutAnchor) Autorelease() LayoutAnchor {
	rv := objc.Send[LayoutAnchor](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLayoutAnchor creates a new LayoutAnchor instance.
func NewLayoutAnchor() LayoutAnchor {
	return getLayoutAnchorClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for LayoutAnchor */
// A factory class for creating layout constraint objects using a fluent API.
//
// Use these constraints to programatically define your layout using Auto Layout. Instead of creating objects directly, start with an or object you wish to constrain, and select one of that object’s anchor properties. These properties correspond to the main values used in Auto Layout, and provide an appropriate subclass for creating constraints to that attribute. Use the anchor’s methods to construct your constraint. As you can see from these examples, the class provides several advantages over using the API directly. The code is cleaner, more concise, and easier to read. The subclasses provide additional type checking, preventing you from creating invalid constraints. For more information on the anchor properties, see in the or .


// A factory class for creating layout constraint objects using a fluent API.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutAnchor
type LayoutAnchor struct {
	objectivec.Object
}

// LayoutAnchorFrom constructs a [LayoutAnchor] from an unsafe.Pointer.
//
// A factory class for creating layout constraint objects using a fluent API.
func LayoutAnchorFrom(ptr unsafe.Pointer) LayoutAnchor {
	return LayoutAnchor{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for LayoutAnchor *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for LayoutAnchor */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for LayoutAnchor */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for LayoutAnchor */

// Returns a constraint that defines one item’s attribute as equal to another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutAnchor/constraint(equalTo:)
func (l_ LayoutAnchor) ConstraintEqualToAnchor(anchor unsafe.Pointer) ILayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("constraintEqualToAnchor:"), anchor)
	return rv
}/* debug [instance_methods/method]: ConstraintEqualToAnchor */


// Returns a constraint that defines one item’s attribute as equal to another item’s attribute plus a constant offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutAnchor/constraint(equalTo:constant:)
func (l_ LayoutAnchor) ConstraintEqualToAnchorConstant(anchor unsafe.Pointer, c float64) ILayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("constraintEqualToAnchor:constant:"), anchor, c)
	return rv
}/* debug [instance_methods/method]: ConstraintEqualToAnchorConstant */


// Returns a constraint that defines one item’s attribute as greater than or equal to another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutAnchor/constraint(greaterThanOrEqualTo:)
func (l_ LayoutAnchor) ConstraintGreaterThanOrEqualToAnchor(anchor unsafe.Pointer) ILayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("constraintGreaterThanOrEqualToAnchor:"), anchor)
	return rv
}/* debug [instance_methods/method]: ConstraintGreaterThanOrEqualToAnchor */


// Returns a constraint that defines one item’s attribute as greater than or equal to another item’s attribute plus a constant offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutAnchor/constraint(greaterThanOrEqualTo:constant:)
func (l_ LayoutAnchor) ConstraintGreaterThanOrEqualToAnchorConstant(anchor unsafe.Pointer, c float64) ILayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("constraintGreaterThanOrEqualToAnchor:constant:"), anchor, c)
	return rv
}/* debug [instance_methods/method]: ConstraintGreaterThanOrEqualToAnchorConstant */


// Returns a constraint that defines one item’s attribute as less than or equal to another.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutAnchor/constraint(lessThanOrEqualTo:)
func (l_ LayoutAnchor) ConstraintLessThanOrEqualToAnchor(anchor unsafe.Pointer) ILayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("constraintLessThanOrEqualToAnchor:"), anchor)
	return rv
}/* debug [instance_methods/method]: ConstraintLessThanOrEqualToAnchor */


// Returns a constraint that defines one item’s attribute as less than or equal to another item’s attribute plus a constant offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutAnchor/constraint(lessThanOrEqualTo:constant:)
func (l_ LayoutAnchor) ConstraintLessThanOrEqualToAnchorConstant(anchor unsafe.Pointer, c float64) ILayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("constraintLessThanOrEqualToAnchor:constant:"), anchor, c)
	return rv
}/* debug [instance_methods/method]: ConstraintLessThanOrEqualToAnchorConstant */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for LayoutAnchor */

// The constraints that impact the layout of the anchor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutAnchor/constraintsAffectingLayout
func (l_ LayoutAnchor) ConstraintsAffectingLayout() []LayoutConstraint {
	rv := objc.Send[[]LayoutConstraint](l_.ID, objc.Sel("constraintsAffectingLayout"))
	return rv
}/* debug [instance_properties/getter]: constraintsAffectingLayout */


// A Boolean value indicating whether the constraints impacting the anchor specify its location ambiguously.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutAnchor/hasAmbiguousLayout
func (l_ LayoutAnchor) HasAmbiguousLayout() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("hasAmbiguousLayout"))
	return rv
}/* debug [instance_properties/getter]: hasAmbiguousLayout */


// The layout item used to calculate the anchor’s position.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutAnchor/item
func (l_ LayoutAnchor) Item() objc.ID {
	rv := objc.Send[objc.ID](l_.ID, objc.Sel("item"))
	return rv
}/* debug [instance_properties/getter]: item */


// The name assigned to the anchor for debugging purposes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutAnchor/name
func (l_ LayoutAnchor) Name() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("name"))
	return rv
}/* debug [instance_properties/getter]: name */


// A layout anchor representing the bottom edge of the view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/bottomanchor
func (l_ LayoutAnchor) BottomAnchor() ILayoutYAxisAnchor {
	rv := objc.Send[LayoutYAxisAnchor](l_.ID, objc.Sel("bottomAnchor"))
	return rv
}/* debug [instance_properties/getter]: bottomAnchor */


// A layout anchor representing the bottom edge of the view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/bottomanchor
func (l_ LayoutAnchor) SetBottomAnchor(value ILayoutYAxisAnchor) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setBottomAnchor:"), value)
}/* debug [instance_properties/setter]: bottomAnchor */


// A layout anchor representing the leading edge of the view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/leadinganchor
func (l_ LayoutAnchor) LeadingAnchor() ILayoutXAxisAnchor {
	rv := objc.Send[LayoutXAxisAnchor](l_.ID, objc.Sel("leadingAnchor"))
	return rv
}/* debug [instance_properties/getter]: leadingAnchor */


// A layout anchor representing the leading edge of the view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/leadinganchor
func (l_ LayoutAnchor) SetLeadingAnchor(value ILayoutXAxisAnchor) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLeadingAnchor:"), value)
}/* debug [instance_properties/setter]: leadingAnchor */


// A layout anchor representing the left edge of the view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/leftanchor
func (l_ LayoutAnchor) LeftAnchor() ILayoutXAxisAnchor {
	rv := objc.Send[LayoutXAxisAnchor](l_.ID, objc.Sel("leftAnchor"))
	return rv
}/* debug [instance_properties/getter]: leftAnchor */


// A layout anchor representing the left edge of the view’s frame.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsview/leftanchor
func (l_ LayoutAnchor) SetLeftAnchor(value ILayoutXAxisAnchor) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLeftAnchor:"), value)
}/* debug [instance_properties/setter]: leftAnchor */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSLayoutAnchor */



