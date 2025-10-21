// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ScrollEdgeEffectStyle] class.
var (
	ScrollEdgeEffectStyleClass     _ScrollEdgeEffectStyleClass
	ScrollEdgeEffectStyleClassOnce sync.Once
)

func getScrollEdgeEffectStyleClass() _ScrollEdgeEffectStyleClass {
	ScrollEdgeEffectStyleClassOnce.Do(func() {
		ScrollEdgeEffectStyleClass = _ScrollEdgeEffectStyleClass{objc.GetClass("NSScrollEdgeEffectStyle")}
	})
	return ScrollEdgeEffectStyleClass
}

type _ScrollEdgeEffectStyleClass struct {
	class objc.Class
}

// An interface definition for the [ScrollEdgeEffectStyle] class.
type IScrollEdgeEffectStyle interface {
	objectivec.IObject
}

// Styles for a scroll view’s edge effect.
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSScrollEdgeEffectStyle
type ScrollEdgeEffectStyle struct {
	objectivec.Object
}

// ScrollEdgeEffectStyleFrom constructs a [ScrollEdgeEffectStyle] from an unsafe.Pointer.
//
// Styles for a scroll view’s edge effect.
func ScrollEdgeEffectStyleFrom(ptr unsafe.Pointer) ScrollEdgeEffectStyle {
	return ScrollEdgeEffectStyle{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _ScrollEdgeEffectStyleClass) Alloc() ScrollEdgeEffectStyle {
	rv := objc.Send[ScrollEdgeEffectStyle](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _ScrollEdgeEffectStyleClass) New() ScrollEdgeEffectStyle {
	rv := objc.Send[ScrollEdgeEffectStyle](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ ScrollEdgeEffectStyle) Init() ScrollEdgeEffectStyle {
	rv := objc.Send[ScrollEdgeEffectStyle](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ ScrollEdgeEffectStyle) Autorelease() ScrollEdgeEffectStyle {
	rv := objc.Send[ScrollEdgeEffectStyle](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewScrollEdgeEffectStyle creates a new ScrollEdgeEffectStyle instance.
func NewScrollEdgeEffectStyle() ScrollEdgeEffectStyle {
	return getScrollEdgeEffectStyleClass().New()
}




