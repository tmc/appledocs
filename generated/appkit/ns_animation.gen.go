// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class NSAnimation */


/* debug [class_header]: Header for NSAnimation */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Animation */
// An interface definition for the [Animation] class.
type IAnimation interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Animation */
	// properties:
	AnimationBlockingMode() AnimationBlockingMode
	SetAnimationBlockingMode(value AnimationBlockingMode)
	AnimationCurve() AnimationCurve
	SetAnimationCurve(value AnimationCurve)
	CurrentProgress() AnimationProgress /* typedef */
	SetCurrentProgress(value AnimationProgress /* typedef */)
	CurrentValue() float32
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
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
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Animation */
	// methods:
	AddProgressMark(progressMark AnimationProgress /* typedef */)
	ClearStartAnimation()
	ClearStopAnimation()
	RemoveProgressMark(progressMark AnimationProgress /* typedef */)
	StartAnimation()
	StartWhenAnimationReachesProgress(animation IAnimation, startProgress AnimationProgress /* typedef */)
	StopAnimation()
	StopWhenAnimationReachesProgress(animation IAnimation, stopProgress AnimationProgress /* typedef */)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Animation */
// Alloc allocates a new instance without initialization.
func (ac _AnimationClass) Alloc() Animation {
	rv := objc.Send[Animation](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Animation */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Animation */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/init(coder:)
func NewAnimationWithCoder(coder foundation.Coder) Animation {
	instance := getAnimationClass().Alloc()
	rv := objc.Send[Animation](instance.ID, objc.Sel("initWithCoder:"), coder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAnimationWithCoder */


// Returns an object initialized with the specified duration and animation-curve values.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/init(duration:animationCurve:)
func NewAnimationWithDurationAnimationCurve(duration float64, animationCurve AnimationCurve) Animation {
	instance := getAnimationClass().Alloc()
	rv := objc.Send[Animation](instance.ID, objc.Sel("initWithDuration:animationCurve:"), duration, animationCurve)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewAnimationWithDurationAnimationCurve */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Animation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Animation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Animation */

// Adds the progress mark to the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/addProgressMark(_:)
func (a_ Animation) AddProgressMark(progressMark AnimationProgress /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("addProgressMark:"), progressMark)
}/* debug [instance_methods/method]: AddProgressMark */


// Clears linkage to another animation that causes the receiver to start.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/clearStart()
func (a_ Animation) ClearStartAnimation() {
	objc.Send[objc.ID](a_.ID, objc.Sel("clearStartAnimation"))
}/* debug [instance_methods/method]: ClearStartAnimation */


// Clears linkage to another animation that causes the receiver to stop.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/clearStop()
func (a_ Animation) ClearStopAnimation() {
	objc.Send[objc.ID](a_.ID, objc.Sel("clearStopAnimation"))
}/* debug [instance_methods/method]: ClearStopAnimation */


// Removes progress mark from the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/removeProgressMark(_:)
func (a_ Animation) RemoveProgressMark(progressMark AnimationProgress /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("removeProgressMark:"), progressMark)
}/* debug [instance_methods/method]: RemoveProgressMark */


// Starts the animation represented by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/start()
func (a_ Animation) StartAnimation() {
	objc.Send[objc.ID](a_.ID, objc.Sel("startAnimation"))
}/* debug [instance_methods/method]: StartAnimation */


// Starts running the animation represented by the receiver when another animation reaches a specific progress mark.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/start(when:reachesProgress:)
func (a_ Animation) StartWhenAnimationReachesProgress(animation IAnimation, startProgress AnimationProgress /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("startWhenAnimation:reachesProgress:"), animation, startProgress)
}/* debug [instance_methods/method]: StartWhenAnimationReachesProgress */


// Stops the animation represented by the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/stop()
func (a_ Animation) StopAnimation() {
	objc.Send[objc.ID](a_.ID, objc.Sel("stopAnimation"))
}/* debug [instance_methods/method]: StopAnimation */


// Stops running the animation represented by the receiver when another animation reaches a specific progress mark.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/stop(when:reachesProgress:)
func (a_ Animation) StopWhenAnimationReachesProgress(animation IAnimation, stopProgress AnimationProgress /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("stopWhenAnimation:reachesProgress:"), animation, stopProgress)
}/* debug [instance_methods/method]: StopWhenAnimationReachesProgress */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Animation */

// The blocking mode of the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/animationBlockingMode
func (a_ Animation) AnimationBlockingMode() AnimationBlockingMode {
	rv := objc.Send[AnimationBlockingMode](a_.ID, objc.Sel("animationBlockingMode"))
	return rv
}/* debug [instance_properties/getter]: animationBlockingMode */


// The blocking mode of the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/animationBlockingMode
func (a_ Animation) SetAnimationBlockingMode(value AnimationBlockingMode) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAnimationBlockingMode:"), value)
}/* debug [instance_properties/setter]: animationBlockingMode */


// The timing curve for the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/animationCurve
func (a_ Animation) AnimationCurve() AnimationCurve {
	rv := objc.Send[AnimationCurve](a_.ID, objc.Sel("animationCurve"))
	return rv
}/* debug [instance_properties/getter]: animationCurve */


// The timing curve for the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/animationCurve
func (a_ Animation) SetAnimationCurve(value AnimationCurve) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAnimationCurve:"), value)
}/* debug [instance_properties/setter]: animationCurve */


// The current progress of the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/currentProgress
func (a_ Animation) CurrentProgress() AnimationProgress /* typedef */ {
	rv := objc.Send[float32](a_.ID, objc.Sel("currentProgress"))
	return rv
}/* debug [instance_properties/getter]: currentProgress */


// The current progress of the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/currentProgress
func (a_ Animation) SetCurrentProgress(value AnimationProgress /* typedef */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentProgress:"), value)
}/* debug [instance_properties/setter]: currentProgress */


// The current value of the animation effect, based on the current progress
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/currentValue
func (a_ Animation) CurrentValue() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("currentValue"))
	return rv
}/* debug [instance_properties/getter]: currentValue */


// The animation delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/delegate
func (a_ Animation) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The animation delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/delegate
func (a_ Animation) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The duration of the animation, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/duration
func (a_ Animation) Duration() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("duration"))
	return rv
}/* debug [instance_properties/getter]: duration */


// The duration of the animation, in seconds.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/duration
func (a_ Animation) SetDuration(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDuration:"), value)
}/* debug [instance_properties/setter]: duration */


// The number of frame updates per second to generate for the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/frameRate
func (a_ Animation) FrameRate() float32 {
	rv := objc.Send[float32](a_.ID, objc.Sel("frameRate"))
	return rv
}/* debug [instance_properties/getter]: frameRate */


// The number of frame updates per second to generate for the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/frameRate
func (a_ Animation) SetFrameRate(value float32) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFrameRate:"), value)
}/* debug [instance_properties/setter]: frameRate */


// A Boolean value indicating whether the animation is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/isAnimating
func (a_ Animation) Animating() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("animating"))
	return rv
}/* debug [instance_properties/getter]: animating */


// An array of floating-point numbers representing current progress marks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/progressMarks
func (a_ Animation) ProgressMarks() []foundation.Number {
	rv := objc.Send[[]foundation.Number](a_.ID, objc.Sel("progressMarks"))
	return rv
}/* debug [instance_properties/getter]: progressMarks */


// An array of floating-point numbers representing current progress marks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/progressMarks
func (a_ Animation) SetProgressMarks(value []foundation.Number) {
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
}/* debug [instance_properties/setter]: progressMarks */


// An array of strings representing the run loop modes in which the animation can run.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/runLoopModesForAnimating
func (a_ Animation) RunLoopModesForAnimating() []string {
	rv := objc.Send[[]string](a_.ID, objc.Sel("runLoopModesForAnimating"))
	return rv
}/* debug [instance_properties/getter]: runLoopModesForAnimating */


// A Boolean value indicating whether the animation is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/isanimating
func (a_ Animation) IsAnimating() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isAnimating"))
	return rv
}/* debug [instance_properties/getter]: isAnimating */


// A Boolean value indicating whether the animation is in progress.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/appkit/nsanimation/isanimating
func (a_ Animation) SetIsAnimating(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsAnimating:"), value)
}/* debug [instance_properties/setter]: isAnimating */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class NSAnimation */


