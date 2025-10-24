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
	// properties:
	AnimationBlockingMode() AnimationBlockingMode
	SetAnimationBlockingMode(value AnimationBlockingMode)
	AnimationCurve() AnimationCurve
	SetAnimationCurve(value AnimationCurve)
	CurrentProgress() objc.IObject /* cross-framework: AnimationProgress */
	SetCurrentProgress(value objc.IObject /* cross-framework: AnimationProgress */)
	CurrentValue() float32
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	Duration() float64
	SetDuration(value float64)
	FrameRate() float32
	SetFrameRate(value float32)
	Animating() bool
	ProgressMarks() []foundation.Number
	SetProgressMarks(value []foundation.Number)
	RunLoopModesForAnimating() []string
	IsAnimating() bool
	SetIsAnimating(value bool)
	// methods:
	AddProgressMark(progressMark objc.IObject /* cross-framework: AnimationProgress */)
	ClearStartAnimation()
	ClearStopAnimation()
	RemoveProgressMark(progressMark objc.IObject /* cross-framework: AnimationProgress */)
	StartAnimation()
	StartWhenAnimationReachesProgress(animation IAnimation, startProgress objc.IObject /* cross-framework: AnimationProgress */)
	StopAnimation()
	StopWhenAnimationReachesProgress(animation IAnimation, stopProgress objc.IObject /* cross-framework: AnimationProgress */)
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/init(coder:)
func NewAnimationWithCoder(coder foundation.Coder) Animation {
	instance := getAnimationClass().Alloc()
	rv := objc.Send[Animation](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}


// Returns an object initialized with the specified duration and animation-curve values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/init(duration:animationCurve:)
func NewAnimationWithDurationAnimationCurve(duration float64, animationCurve AnimationCurve) Animation {
	instance := getAnimationClass().Alloc()
	rv := objc.Send[Animation](instance.ID, objc.Sel("initWithDuration:animationCurve:"), duration, animationCurve)
	rv.Autorelease()
	return rv
}



// Adds the progress mark to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/addProgressMark(_:)
func (a_ Animation) AddProgressMark(progressMark objc.IObject /* cross-framework: AnimationProgress */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("addProgressMark:"), progressMark)
}


// Clears linkage to another animation that causes the receiver to start.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/clearStart()
func (a_ Animation) ClearStartAnimation() {
	objc.Send[objc.ID](a_.ID, objc.Sel("clearStartAnimation"))
}


// Clears linkage to another animation that causes the receiver to stop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/clearStop()
func (a_ Animation) ClearStopAnimation() {
	objc.Send[objc.ID](a_.ID, objc.Sel("clearStopAnimation"))
}


// Removes progress mark from the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/removeProgressMark(_:)
func (a_ Animation) RemoveProgressMark(progressMark objc.IObject /* cross-framework: AnimationProgress */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeProgressMark:"), progressMark)
}


// Starts the animation represented by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/start()
func (a_ Animation) StartAnimation() {
	objc.Send[objc.ID](a_.ID, objc.Sel("startAnimation"))
}


// Starts running the animation represented by the receiver when another animation reaches a specific progress mark.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/start(when:reachesProgress:)
func (a_ Animation) StartWhenAnimationReachesProgress(animation IAnimation, startProgress objc.IObject /* cross-framework: AnimationProgress */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("startWhenAnimation:reachesProgress:"), animation, startProgress)
}


// Stops the animation represented by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/stop()
func (a_ Animation) StopAnimation() {
	objc.Send[objc.ID](a_.ID, objc.Sel("stopAnimation"))
}


// Stops running the animation represented by the receiver when another animation reaches a specific progress mark.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/stop(when:reachesProgress:)
func (a_ Animation) StopWhenAnimationReachesProgress(animation IAnimation, stopProgress objc.IObject /* cross-framework: AnimationProgress */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("stopWhenAnimation:reachesProgress:"), animation, stopProgress)
}


// The blocking mode of the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/animationBlockingMode
func (a_ Animation) AnimationBlockingMode() AnimationBlockingMode {
	rv := objc.Send[AnimationBlockingMode](a_.ID, objc.Sel("animationBlockingMode"))
	return rv
}


// The blocking mode of the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/animationBlockingMode
func (a_ Animation) SetAnimationBlockingMode(value AnimationBlockingMode) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAnimationBlockingMode:"), value)
}


// The timing curve for the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/animationCurve
func (a_ Animation) AnimationCurve() AnimationCurve {
	rv := objc.Send[AnimationCurve](a_.ID, objc.Sel("animationCurve"))
	return rv
}


// The timing curve for the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/animationCurve
func (a_ Animation) SetAnimationCurve(value AnimationCurve) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAnimationCurve:"), value)
}


// The current progress of the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/currentProgress
func (a_ Animation) CurrentProgress() objc.IObject /* cross-framework: AnimationProgress */ {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("currentProgress"))
	return rv
}


// The current progress of the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/currentProgress
func (a_ Animation) SetCurrentProgress(value objc.IObject /* cross-framework: AnimationProgress */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentProgress:"), value)
}


// The current value of the animation effect, based on the current progress
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/currentValue
func (a_ Animation) CurrentValue() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("currentValue"))
	return rv
}


// The animation delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/delegate
func (a_ Animation) Delegate() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("delegate"))
	return rv
}


// The animation delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/delegate
func (a_ Animation) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}


// The duration of the animation, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/duration
func (a_ Animation) Duration() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("duration"))
	return rv
}


// The duration of the animation, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/duration
func (a_ Animation) SetDuration(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDuration:"), value)
}


// The number of frame updates per second to generate for the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/frameRate
func (a_ Animation) FrameRate() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("frameRate"))
	return rv
}


// The number of frame updates per second to generate for the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/frameRate
func (a_ Animation) SetFrameRate(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFrameRate:"), value)
}


// A Boolean value indicating whether the animation is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/isAnimating
func (a_ Animation) Animating() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("animating"))
	return rv
}


// An array of floating-point numbers representing current progress marks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/progressMarks
func (a_ Animation) ProgressMarks() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("progressMarks"))
	return rv
}


// An array of floating-point numbers representing current progress marks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/progressMarks
func (a_ Animation) SetProgressMarks(value []foundation.Number) {
	// Convert Go slice to NSArray
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setProgressMarks:"), nsArray)
}


// An array of strings representing the run loop modes in which the animation can run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/runLoopModesForAnimating
func (a_ Animation) RunLoopModesForAnimating() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("runLoopModesForAnimating"))
	return rv
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


