// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/coregraphics"
)

// The class instance for the [OverlayTransitionContext] class.
var (
	OverlayTransitionContextClass     _OverlayTransitionContextClass
	OverlayTransitionContextClassOnce sync.Once
)

func getOverlayTransitionContextClass() _OverlayTransitionContextClass {
	OverlayTransitionContextClassOnce.Do(func() {
		OverlayTransitionContextClass = _OverlayTransitionContextClass{objc.GetClass("SKOverlayTransitionContext")}
	})
	return OverlayTransitionContextClass
}

type _OverlayTransitionContextClass struct {
	class objc.Class
}

// An interface definition for the [OverlayTransitionContext] class.
type IOverlayTransitionContext interface {
	objectivec.IObject
	AddAnimationBlock(block unsafe.Pointer)
}

// A context object you can use to animate UI changes while the platform presents or dismisses an overlay.
//
// For more information on animating UI changes while the system presents or dismisses an overlay, see and .
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/TransitionContext
type OverlayTransitionContext struct {
	objectivec.Object
}

// OverlayTransitionContextFrom constructs a [OverlayTransitionContext] from an unsafe.Pointer.
//
// A context object you can use to animate UI changes while the platform presents or dismisses an overlay.
func OverlayTransitionContextFrom(ptr unsafe.Pointer) OverlayTransitionContext {
	return OverlayTransitionContext{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (oc _OverlayTransitionContextClass) Alloc() OverlayTransitionContext {
	rv := objc.Send[OverlayTransitionContext](objc.ID(oc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (oc _OverlayTransitionContextClass) New() OverlayTransitionContext {
	rv := objc.Send[OverlayTransitionContext](objc.ID(oc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (o_ OverlayTransitionContext) Init() OverlayTransitionContext {
	rv := objc.Send[OverlayTransitionContext](o_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (o_ OverlayTransitionContext) Autorelease() OverlayTransitionContext {
	rv := objc.Send[OverlayTransitionContext](o_.ID, objc.Sel("autorelease"))
	return rv
}

// NewOverlayTransitionContext creates a new OverlayTransitionContext instance.
func NewOverlayTransitionContext() OverlayTransitionContext {
	return getOverlayTransitionContextClass().New()
}


// Adds a closure you can use to animate view properties.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/TransitionContext/addAnimation(_:)
func (o_ OverlayTransitionContext) AddAnimationBlock(block unsafe.Pointer) {
	objc.Send[objc.ID](o_.ID, objc.Sel("addAnimationBlock:"), block)
}

// The size and location of the overlay at the end of the transition.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/TransitionContext/endFrame
func (o_ OverlayTransitionContext) EndFrame() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](o_.ID, objc.Sel("endFrame"))
	return rv
}

// The size and location of the overlay before the transition.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKOverlay/TransitionContext/startFrame
func (o_ OverlayTransitionContext) StartFrame() coregraphics.CGRect {
	rv := objc.Send[coregraphics.CGRect](o_.ID, objc.Sel("startFrame"))
	return rv
}



