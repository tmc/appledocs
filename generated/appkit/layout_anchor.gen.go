// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LayoutAnchor] class.
var (
	layoutAnchorClass     _LayoutAnchorClass
	layoutAnchorClassOnce sync.Once
)

func getLayoutAnchorClass() _LayoutAnchorClass {
	layoutAnchorClassOnce.Do(func() {
		layoutAnchorClass = _LayoutAnchorClass{objc.GetClass("NSLayoutAnchor")}
	})
	return layoutAnchorClass
}

type _LayoutAnchorClass struct {
	class objc.Class
}

// An interface definition for the [LayoutAnchor] class.
type ILayoutAnchor interface {
	objectivec.IObject
}

// A factory class for creating layout constraint objects using a fluent API. [Full Topic]
//
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
// Alloc allocates a new instance without initialization.
func (lc _LayoutAnchorClass) Alloc() LayoutAnchor {
	rv := objc.Send[LayoutAnchor](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




