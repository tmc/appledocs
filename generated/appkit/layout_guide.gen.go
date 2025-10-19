// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [LayoutGuide] class.
var (
	layoutGuideClass     _LayoutGuideClass
	layoutGuideClassOnce sync.Once
)

func getLayoutGuideClass() _LayoutGuideClass {
	layoutGuideClassOnce.Do(func() {
		layoutGuideClass = _LayoutGuideClass{objc.GetClass("NSLayoutGuide")}
	})
	return layoutGuideClass
}

type _LayoutGuideClass struct {
	class objc.Class
}

// An interface definition for the [LayoutGuide] class.
type ILayoutGuide interface {
	objectivec.IObject
}

// A rectangular area that can interact with Auto Layout.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSLayoutGuide
type LayoutGuide struct {
	objectivec.Object
}

// LayoutGuideFrom constructs a [LayoutGuide] from an unsafe.Pointer.
//
// A rectangular area that can interact with Auto Layout.
func LayoutGuideFrom(ptr unsafe.Pointer) LayoutGuide {
	return LayoutGuide{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (lc _LayoutGuideClass) Alloc() LayoutGuide {
	rv := objc.Send[LayoutGuide](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LayoutGuideClass) New() LayoutGuide {
	rv := objc.Send[LayoutGuide](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LayoutGuide) Init() LayoutGuide {
	rv := objc.Send[LayoutGuide](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LayoutGuide) Autorelease() LayoutGuide {
	rv := objc.Send[LayoutGuide](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLayoutGuide creates a new LayoutGuide instance.
func NewLayoutGuide() LayoutGuide {
	return getLayoutGuideClass().New()
}




