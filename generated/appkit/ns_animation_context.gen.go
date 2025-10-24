// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/quartzcore"
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
	AllowsImplicitAnimation() bool
	SetAllowsImplicitAnimation(value bool)
	CompletionHandler() func()
	SetCompletionHandler(value func())
	Duration() float64
	SetDuration(value float64)
	TimingFunction() quartzcore.MediaTimingFunction
	SetTimingFunction(value quartzcore.MediaTimingFunction)
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



// Creates a new animation grouping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/beginGrouping()
func (ac _AnimationContextClass) BeginGrouping() {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("beginGrouping"))
}


// Ends the current animation grouping.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/endGrouping()
func (ac _AnimationContextClass) EndGrouping() {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("endGrouping"))
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/runAnimationGroup(_:)
func (ac _AnimationContextClass) RunAnimationGroup(changes unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("runAnimationGroup:"), changes)
}


// Allows you to specify a completion block body after the set of animation actions whose completion will trigger the completion block.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/runAnimationGroup(_:completionHandler:)
func (ac _AnimationContextClass) RunAnimationGroupCompletionHandler(changes unsafe.Pointer, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("runAnimationGroup:completionHandler:"), changes, completionHandler)
}


// Returns the current animation context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/current
func (ac _AnimationContextClass) CurrentContext() AnimationContext {
	rv := objc.Send[AnimationContext](objc.ID(ac.class), objc.Sel("currentContext"))
	return rv
}

// Determine if animations are enabled or not for animations that occur as a result of another property change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/allowsImplicitAnimation
func (a_ AnimationContext) AllowsImplicitAnimation() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("allowsImplicitAnimation"))
	return rv
}


// Determine if animations are enabled or not for animations that occur as a result of another property change.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/allowsImplicitAnimation
func (a_ AnimationContext) SetAllowsImplicitAnimation(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAllowsImplicitAnimation:"), value)
}


// A completion Block that is called when the animations in the grouping are completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/completionHandler
func (a_ AnimationContext) CompletionHandler() func() {
	rv := objc.Send[func()](a_.ID, objc.Sel("completionHandler"))
	return rv
}


// A completion Block that is called when the animations in the grouping are completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/completionHandler
func (a_ AnimationContext) SetCompletionHandler(value func()) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCompletionHandler:"), value)
}


// Returns the current animation context.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/current
func (a_ AnimationContext) CurrentContext() IAnimationContext {
	rv := objc.Send[AnimationContext](a_.ID, objc.Sel("currentContext"))
	return rv
}


// The duration used by animations created as a result of setting new values for an animatable property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/duration
func (a_ AnimationContext) Duration() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("duration"))
	return rv
}


// The duration used by animations created as a result of setting new values for an animatable property.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/duration
func (a_ AnimationContext) SetDuration(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDuration:"), value)
}


// The timing function used for all animations within this animation proxy group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/timingFunction
func (a_ AnimationContext) TimingFunction() quartzcore.MediaTimingFunction {
	rv := objc.Send[quartzcore.MediaTimingFunction](a_.ID, objc.Sel("timingFunction"))
	return rv
}


// The timing function used for all animations within this animation proxy group.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext/timingFunction
func (a_ AnimationContext) SetTimingFunction(value quartzcore.MediaTimingFunction) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTimingFunction:"), value)
}



