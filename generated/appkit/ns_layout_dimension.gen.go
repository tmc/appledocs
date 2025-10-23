// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [LayoutDimension] class.
type ILayoutDimension interface {
	ILayoutAnchor
	ConstraintLessThanOrEqualToAnchorMultiplierConstant(anchor ILayoutDimension, m float64, c float64) LayoutConstraint
}

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

// Alloc allocates a new instance without initialization.
func (lc _LayoutDimensionClass) Alloc() LayoutDimension {
	rv := objc.Send[LayoutDimension](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Returns a constraint that defines the anchor’s size attribute as greater than or equal to the specified anchor multiplied by the constant plus an offset.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutDimension/constraint(lessThanOrEqualTo:multiplier:constant:)
func (l_ LayoutDimension) ConstraintLessThanOrEqualToAnchorMultiplierConstant(anchor ILayoutDimension, m float64, c float64) LayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("constraintLessThanOrEqualToAnchor:multiplier:constant:"), anchor, m, c)
	return rv
}



