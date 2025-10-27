// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)





// The class instance for the [LayoutYAxisAnchor] class.
var (
	LayoutYAxisAnchorClass     _LayoutYAxisAnchorClass
	LayoutYAxisAnchorClassOnce sync.Once
)

func getLayoutYAxisAnchorClass() _LayoutYAxisAnchorClass {
	LayoutYAxisAnchorClassOnce.Do(func() {
		LayoutYAxisAnchorClass = _LayoutYAxisAnchorClass{objc.GetClass("NSLayoutYAxisAnchor")}
	})
	return LayoutYAxisAnchorClass
}

type _LayoutYAxisAnchorClass struct {
	class objc.Class
}





// An interface definition for the [LayoutYAxisAnchor] class.
type ILayoutYAxisAnchor interface {
	ILayoutAnchor
	

	// properties:


	

	// methods:
	AnchorWithOffsetToAnchor(otherAnchor ILayoutYAxisAnchor) ILayoutDimension
	ConstraintEqualToSystemSpacingBelowAnchorMultiplier(anchor ILayoutYAxisAnchor, multiplier float64) ILayoutConstraint
	ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchorMultiplier(anchor ILayoutYAxisAnchor, multiplier float64) ILayoutConstraint
	ConstraintLessThanOrEqualToSystemSpacingBelowAnchorMultiplier(anchor ILayoutYAxisAnchor, multiplier float64) ILayoutConstraint


}





// Alloc allocates a new instance without initialization.
func (lc _LayoutYAxisAnchorClass) Alloc() LayoutYAxisAnchor {
	rv := objc.Send[LayoutYAxisAnchor](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _LayoutYAxisAnchorClass) New() LayoutYAxisAnchor {
	rv := objc.Send[LayoutYAxisAnchor](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LayoutYAxisAnchor) Init() LayoutYAxisAnchor {
	rv := objc.Send[LayoutYAxisAnchor](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LayoutYAxisAnchor) Autorelease() LayoutYAxisAnchor {
	rv := objc.Send[LayoutYAxisAnchor](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLayoutYAxisAnchor creates a new LayoutYAxisAnchor instance.
func NewLayoutYAxisAnchor() LayoutYAxisAnchor {
	return getLayoutYAxisAnchorClass().New()
}





// A factory class for creating vertical layout constraint objects using a fluent API.
//
// adds type information to the methods inherited from . Specifically, the generic methods declared by must now take a matching object. For more information on using layout anchors, see .


// A factory class for creating vertical layout constraint objects using a fluent API.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutYAxisAnchor
type LayoutYAxisAnchor struct {
	LayoutAnchor
}

// LayoutYAxisAnchorFrom constructs a [LayoutYAxisAnchor] from an unsafe.Pointer.
//
// A factory class for creating vertical layout constraint objects using a fluent API.
func LayoutYAxisAnchorFrom(ptr unsafe.Pointer) LayoutYAxisAnchor {
	return LayoutYAxisAnchor{
		LayoutAnchor: LayoutAnchorFrom(ptr),
	}
}




















// Creates a layout dimension object from two anchors.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutYAxisAnchor/anchorWithOffset(to:)
func (l_ LayoutYAxisAnchor) AnchorWithOffsetToAnchor(otherAnchor ILayoutYAxisAnchor) ILayoutDimension {
	rv := objc.Send[LayoutDimension](l_.ID, objc.Sel("anchorWithOffsetToAnchor:"), otherAnchor)
	return rv
}


// Returns a constraint that defines the specific distance at which the current anchor is positioned below the specified anchor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutYAxisAnchor/constraint(equalToSystemSpacingBelow:multiplier:)
func (l_ LayoutYAxisAnchor) ConstraintEqualToSystemSpacingBelowAnchorMultiplier(anchor ILayoutYAxisAnchor, multiplier float64) ILayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("constraintEqualToSystemSpacingBelowAnchor:multiplier:"), anchor, multiplier)
	return rv
}


// Returns a constraint that defines the minimum distance by which the current anchor is positioned below the specified anchor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutYAxisAnchor/constraint(greaterThanOrEqualToSystemSpacingBelow:multiplier:)
func (l_ LayoutYAxisAnchor) ConstraintGreaterThanOrEqualToSystemSpacingBelowAnchorMultiplier(anchor ILayoutYAxisAnchor, multiplier float64) ILayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("constraintGreaterThanOrEqualToSystemSpacingBelowAnchor:multiplier:"), anchor, multiplier)
	return rv
}


// Returns a constraint that defines the maximum distance by which the current anchor is positioned below the specified anchor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutYAxisAnchor/constraint(lessThanOrEqualToSystemSpacingBelow:multiplier:)
func (l_ LayoutYAxisAnchor) ConstraintLessThanOrEqualToSystemSpacingBelowAnchorMultiplier(anchor ILayoutYAxisAnchor, multiplier float64) ILayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("constraintLessThanOrEqualToSystemSpacingBelowAnchor:multiplier:"), anchor, multiplier)
	return rv
}













