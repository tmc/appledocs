// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AnimationContext] class.
var (
	AnimationContextClass     _AnimationContextClass
	AnimationContextClassOnce sync.Once
)

func getAnimationContextClass() _AnimationContextClass {
	AnimationContextClassOnce.Do(func() {
		AnimationContextClass = _AnimationContextClass{objc.GetClass("NSAnimationContext")}
	})
	return AnimationContextClass
}

type _AnimationContextClass struct {
	class objc.Class
}

// An interface definition for the [AnimationContext] class.
type IAnimationContext interface {
	objectivec.IObject
	// properties:
	AllowsImplicitAnimation() bool /* primitive/slice/pointer. */
	SetAllowsImplicitAnimation(value bool /* primitive/slice/pointer. */)
	CompletionHandler() unsafe.Pointer
	SetCompletionHandler(value unsafe.Pointer)
	Duration() unsafe.Pointer
	SetDuration(value unsafe.Pointer)
	TimingFunction() objc.IObject /* cross-framework: MediaTimingFunction */
	SetTimingFunction(value objc.IObject /* cross-framework: MediaTimingFunction */)
	// methods:
}

// An animation context, which contains information about environment and state.
//
// is analogous to and is similar in overall concept to . Each thread maintains its own stack of nestable instances, with each new instance initialized as a copy of the instance below (so, inheriting its current properties). Multiple instances can be nested, allowing a given block of code to initiate animations using its own specified duration without affecting animations initiated by surrounding code.


// An animation context, which contains information about environment and state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext
type AnimationContext struct {
	objectivec.Object
}

// AnimationContextFrom constructs a [AnimationContext] from an unsafe.Pointer.
//
// An animation context, which contains information about environment and state.
func AnimationContextFrom(ptr unsafe.Pointer) AnimationContext {
	return AnimationContext{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AnimationContextClass) Alloc() AnimationContext {
	rv := objc.Send[AnimationContext](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AnimationContextClass) New() AnimationContext {
	rv := objc.Send[AnimationContext](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AnimationContext) Init() AnimationContext {
	rv := objc.Send[AnimationContext](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AnimationContext) Autorelease() AnimationContext {
	rv := objc.Send[AnimationContext](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAnimationContext creates a new AnimationContext instance.
func NewAnimationContext() AnimationContext {
	return getAnimationContextClass().New()
}



// Determine if animations are enabled or not for animations that occur as a result of another property change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationcontext/allowsimplicitanimation
func (a_ AnimationContext) AllowsImplicitAnimation() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsImplicitAnimation"))
	return rv
}


// Determine if animations are enabled or not for animations that occur as a result of another property change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationcontext/allowsimplicitanimation
func (a_ AnimationContext) SetAllowsImplicitAnimation(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsImplicitAnimation:"), value)
}


// A completion Block that is called when the animations in the grouping are completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationcontext/completionhandler
func (a_ AnimationContext) CompletionHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("completionHandler"))
	return rv
}


// A completion Block that is called when the animations in the grouping are completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationcontext/completionhandler
func (a_ AnimationContext) SetCompletionHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCompletionHandler:"), value)
}


// The duration used by animations created as a result of setting new values for an animatable property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationcontext/duration
func (a_ AnimationContext) Duration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("duration"))
	return rv
}


// The duration used by animations created as a result of setting new values for an animatable property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationcontext/duration
func (a_ AnimationContext) SetDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDuration:"), value)
}


// The timing function used for all animations within this animation proxy group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationcontext/timingfunction
func (a_ AnimationContext) TimingFunction() objc.IObject /* cross-framework: MediaTimingFunction */ {
	rv := objc.Send[MediaTimingFunction](a_.ID, objc.Sel("timingFunction"))
	return rv
}


// The timing function used for all animations within this animation proxy group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimationcontext/timingfunction
func (a_ AnimationContext) SetTimingFunction(value objc.IObject /* cross-framework: MediaTimingFunction */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTimingFunction:"), value)
}



