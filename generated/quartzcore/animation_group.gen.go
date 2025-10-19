// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AnimationGroup] class.
var animationGroupClass = _AnimationGroupClass{objc.GetClass("CAAnimationGroup")}

type _AnimationGroupClass struct {
	class objc.Class
}

// An interface definition for the [AnimationGroup] class.
type IAnimationGroup interface {
	IAnimation
}

// An object that allows multiple animations to be grouped and run concurrently. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return animationGroupClass.New()
}




