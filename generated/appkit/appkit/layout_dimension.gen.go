// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [LayoutDimension] class.
var (
	layoutDimensionClass     _LayoutDimensionClass
	layoutDimensionClassOnce sync.Once
)

func getLayoutDimensionClass() _LayoutDimensionClass {
	layoutDimensionClassOnce.Do(func() {
		layoutDimensionClass = _LayoutDimensionClass{objc.GetClass("NSLayoutDimension")}
	})
	return layoutDimensionClass
}

type _LayoutDimensionClass struct {
	class objc.Class
}

// An interface definition for the [LayoutDimension] class.
type ILayoutDimension interface {
	ILayoutAnchor
}

// A factory class for creating size-based layout constraint objects using a fluent API. [Full Topic]
//
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

// New creates and returns a new instance with a +1 retain count.
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




