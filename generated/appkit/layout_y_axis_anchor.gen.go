// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [LayoutYAxisAnchor] class.
var layoutYAxisAnchorClass = _LayoutYAxisAnchorClass{objc.GetClass("NSLayoutYAxisAnchor")}

type _LayoutYAxisAnchorClass struct {
	class objc.Class
}

// An interface definition for the [LayoutYAxisAnchor] class.
type ILayoutYAxisAnchor interface {
	ILayoutAnchor
}

// A factory class for creating vertical layout constraint objects using a fluent API. [Full Topic]
//
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
// Alloc allocates a new instance without initialization.
func (lc _LayoutYAxisAnchorClass) Alloc() LayoutYAxisAnchor {
	rv := objc.Send[LayoutYAxisAnchor](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
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
	return layoutYAxisAnchorClass.New()
}




