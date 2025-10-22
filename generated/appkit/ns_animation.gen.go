// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Animation] class.
var (
	AnimationClass     _AnimationClass
	AnimationClassOnce sync.Once
)

func getAnimationClass() _AnimationClass {
	AnimationClassOnce.Do(func() {
		AnimationClass = _AnimationClass{objc.GetClass("NSAnimation")}
	})
	return AnimationClass
}

type _AnimationClass struct {
	class objc.Class
}

// An interface definition for the [Animation] class.
type IAnimation interface {
	objectivec.IObject
	CurrentProgress() AnimationProgress
	SetCurrentProgress(value IAnimationProgress)
	AnimationBlockingMode() unsafe.Pointer
	SetAnimationBlockingMode(value unsafe.Pointer)
	AnimationCurve() unsafe.Pointer
	SetAnimationCurve(value unsafe.Pointer)
	CurrentValue() float32
	SetCurrentValue(value float32)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Duration() unsafe.Pointer
	SetDuration(value unsafe.Pointer)
	FrameRate() float32
	SetFrameRate(value float32)
	IsAnimating() bool
	SetIsAnimating(value bool)
	ProgressMarks() foundation.Number
	SetProgressMarks(value foundation.INumber)
	RunLoopModesForAnimating() unsafe.Pointer
	SetRunLoopModesForAnimating(value unsafe.Pointer)
}

// An object that manages the timing and progress of animations in the user interface.
//
// also lets you link together multiple animations so that when one animation ends another one starts. It does not provide any drawing support for animation and does not directly deal with views, targets, or actions. objects have several characteristics, including duration, frame rate, and animation curve, which describes the relative speed of the animation over its course. You can set progress marks in an animation, each of which specifies a percentage of the animation completed; when an animation reaches a progress mark, it notifies its delegate and posts a notification to any observers. Animations execute in one of three blocking modes: blocking, non-blocking on the main thread, and non-blocking on a separate thread. The non-blocking modes permit the handling of user events while the animation is running.


// An object that manages the timing and progress of animations in the user interface.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation

type Animation struct {
	objectivec.Object
}

// AnimationFrom constructs a [Animation] from an unsafe.Pointer.
//
// An object that manages the timing and progress of animations in the user interface.
func AnimationFrom(ptr unsafe.Pointer) Animation {
	return Animation{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AnimationClass) Alloc() Animation {
	rv := objc.Send[Animation](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AnimationClass) New() Animation {
	rv := objc.Send[Animation](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ Animation) Init() Animation {
	rv := objc.Send[Animation](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ Animation) Autorelease() Animation {
	rv := objc.Send[Animation](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAnimation creates a new Animation instance.
func NewAnimation() Animation {
	return getAnimationClass().New()
}



// The current progress of the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/currentProgress

func (a_ Animation) CurrentProgress() AnimationProgress {
	rv := objc.Send[AnimationProgress](a_.ID, objc.Sel("currentProgress"))
	return rv
}


// The current progress of the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/currentProgress

func (a_ Animation) SetCurrentProgress(value IAnimationProgress) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentProgress:"), value)
}


// The blocking mode of the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/animationblockingmode

func (a_ Animation) AnimationBlockingMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("animationBlockingMode"))
	return rv
}


// The blocking mode of the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/animationblockingmode

func (a_ Animation) SetAnimationBlockingMode(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAnimationBlockingMode:"), value)
}


// The timing curve for the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/animationcurve

func (a_ Animation) AnimationCurve() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("animationCurve"))
	return rv
}


// The timing curve for the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/animationcurve

func (a_ Animation) SetAnimationCurve(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAnimationCurve:"), value)
}


// The current value of the animation effect, based on the current progress
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/currentvalue

func (a_ Animation) CurrentValue() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("currentValue"))
	return rv
}


// The current value of the animation effect, based on the current progress
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/currentvalue

func (a_ Animation) SetCurrentValue(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentValue:"), value)
}


// The animation delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/delegate

func (a_ Animation) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("delegate"))
	return rv
}


// The animation delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/delegate

func (a_ Animation) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}


// The duration of the animation, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/duration

func (a_ Animation) Duration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("duration"))
	return rv
}


// The duration of the animation, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/duration

func (a_ Animation) SetDuration(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDuration:"), value)
}


// The number of frame updates per second to generate for the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/framerate

func (a_ Animation) FrameRate() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("frameRate"))
	return rv
}


// The number of frame updates per second to generate for the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/framerate

func (a_ Animation) SetFrameRate(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFrameRate:"), value)
}


// A Boolean value indicating whether the animation is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/isanimating

func (a_ Animation) IsAnimating() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isAnimating"))
	return rv
}


// A Boolean value indicating whether the animation is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/isanimating

func (a_ Animation) SetIsAnimating(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsAnimating:"), value)
}


// An array of floating-point numbers representing current progress marks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/progressmarks

func (a_ Animation) ProgressMarks() foundation.Number {
	rv := objc.Send[foundation.Number](a_.ID, objc.Sel("progressMarks"))
	return rv
}


// An array of floating-point numbers representing current progress marks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/progressmarks

func (a_ Animation) SetProgressMarks(value foundation.INumber) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setProgressMarks:"), value)
}


// An array of strings representing the run loop modes in which the animation can run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/runloopmodesforanimating

func (a_ Animation) RunLoopModesForAnimating() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("runLoopModesForAnimating"))
	return rv
}


// An array of strings representing the run loop modes in which the animation can run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/runloopmodesforanimating

func (a_ Animation) SetRunLoopModesForAnimating(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRunLoopModesForAnimating:"), value)
}



