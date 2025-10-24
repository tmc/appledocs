// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/corefoundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	EndFrame() objc.IObject /* cross-framework: Rect */
	SetEndFrame(value objc.IObject /* cross-framework: Rect */)
	StartFrame() objc.IObject /* cross-framework: Rect */
	SetStartFrame(value objc.IObject /* cross-framework: Rect */)
	// methods:
}

// A context object you can use to animate UI changes while the platform presents or dismisses an overlay.
//
// For more information on animating UI changes while the system presents or dismisses an overlay, see and .


// A context object you can use to animate UI changes while the platform presents or dismisses an overlay.
//
// [Full Topic]
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



// The size and location of the overlay at the end of the transition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/transitioncontext/endframe
func (o_ OverlayTransitionContext) EndFrame() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](o_.ID, objc.Sel("endFrame"))
	return rv
}


// The size and location of the overlay at the end of the transition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/transitioncontext/endframe
func (o_ OverlayTransitionContext) SetEndFrame(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setEndFrame:"), value)
}


// The size and location of the overlay before the transition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/transitioncontext/startframe
func (o_ OverlayTransitionContext) StartFrame() objc.IObject /* cross-framework: Rect */ {
	rv := objc.Send[corefoundation.Rect](o_.ID, objc.Sel("startFrame"))
	return rv
}


// The size and location of the overlay before the transition.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skoverlay/transitioncontext/startframe
func (o_ OverlayTransitionContext) SetStartFrame(value objc.IObject /* cross-framework: Rect */) {
	objc.Send[objc.ID](o_.ID, objc.Sel("setStartFrame:"), value)
}



