// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
}

// An object that manages the timing and progress of animations in the user interface.
//
// also lets you link together multiple animations so that when one animation ends another one starts. It does not provide any drawing support for animation and does not directly deal with views, targets, or actions. objects have several characteristics, including duration, frame rate, and animation curve, which describes the relative speed of the animation over its course. You can set progress marks in an animation, each of which specifies a percentage of the animation completed; when an animation reaches a progress mark, it notifies its delegate and posts a notification to any observers. Animations execute in one of three blocking modes: blocking, non-blocking on the main thread, and non-blocking on a separate thread. The non-blocking modes permit the handling of user events while the animation is running.
//
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
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/currentProgress
func (a_ Animation) CurrentProgress() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("currentProgress"))
	return rv
}

// SetCurrentProgress sets the value of the currentProgress property.
// The current progress of the animation.

//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation/currentProgress
func (a_ Animation) SetCurrentProgress(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setCurrentProgress:"), value)
}
