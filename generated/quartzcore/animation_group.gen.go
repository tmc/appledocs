// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AnimationGroup] class.
var (
	animationGroupClass     _AnimationGroupClass
	animationGroupClassOnce sync.Once
)

func getAnimationGroupClass() _AnimationGroupClass {
	animationGroupClassOnce.Do(func() {
		animationGroupClass = _AnimationGroupClass{objc.GetClass("CAAnimationGroup")}
	})
	return animationGroupClass
}

type _AnimationGroupClass struct {
	class objc.Class
}

// An interface definition for the [AnimationGroup] class.
type IAnimationGroup interface {
	IAnimation
}

// An object that allows multiple animations to be grouped and run concurrently.
//
// The grouped animations run in the time space specified by the instance. The duration of the grouped animations are not scaled to the duration of their . Instead, the animations are clipped to the duration of the animation group. For example, a 10 second animation grouped within an animation group with a duration of 5 seconds displays only the first 5 seconds of the animation. The following code shows how you can create a grouped animation containing opacity and scale animations to fade out a layer while expanding it. The animation starts with an opacity of and a scale of on all axes. As the animation’s scale increases to , the opacity drops to and the animated layer vanishes.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimationGroup
type AnimationGroup struct {
	Animation
}

// AnimationGroupFrom constructs a [AnimationGroup] from an unsafe.Pointer.
//
// An object that allows multiple animations to be grouped and run concurrently.
func AnimationGroupFrom(ptr unsafe.Pointer) AnimationGroup {
	return AnimationGroup{
		Animation: AnimationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ac _AnimationGroupClass) Alloc() AnimationGroup {
	rv := objc.Send[AnimationGroup](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AnimationGroupClass) New() AnimationGroup {
	rv := objc.Send[AnimationGroup](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AnimationGroup) Init() AnimationGroup {
	rv := objc.Send[AnimationGroup](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AnimationGroup) Autorelease() AnimationGroup {
	rv := objc.Send[AnimationGroup](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAnimationGroup creates a new AnimationGroup instance.
func NewAnimationGroup() AnimationGroup {
	return getAnimationGroupClass().New()
}


// An array of objects to be evaluated in the time space of the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimationGroup/animations
func (a_ AnimationGroup) Animations() []Animation {
	rv := objc.Send[[]Animation](a_.ID, objc.Sel("animations"))
	return rv
}


// SetAnimations sets the value of the animations property.
// An array of objects to be evaluated in the time space of the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimationGroup/animations
func (a_ AnimationGroup) SetAnimations(value []Animation) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAnimations:"), value)
}


