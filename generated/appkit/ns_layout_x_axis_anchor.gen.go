// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [LayoutXAxisAnchor] class.
type ILayoutXAxisAnchor interface {
	ILayoutAnchor
	// properties:
	// methods:
	ConstraintEqualToSystemSpacingAfterAnchorMultiplier(anchor ILayoutXAxisAnchor, multiplier float64 /* primitive/slice/pointer. */) ILayoutConstraint
}

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

// Alloc allocates a new instance without initialization.
func (lc _LayoutXAxisAnchorClass) Alloc() LayoutXAxisAnchor {
	rv := objc.Send[LayoutXAxisAnchor](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Returns a constraint that defines by how much the current anchor trails the specified anchor.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutXAxisAnchor/constraint(equalToSystemSpacingAfter:multiplier:)
func (l_ LayoutXAxisAnchor) ConstraintEqualToSystemSpacingAfterAnchorMultiplier(anchor ILayoutXAxisAnchor, multiplier float64 /* primitive/slice/pointer. */) ILayoutConstraint {
	rv := objc.Send[LayoutConstraint](l_.ID, objc.Sel("constraintEqualToSystemSpacingAfterAnchor:multiplier:"), anchor, multiplier)
	return rv
}



