// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Animation] class.
var (
	animationClass     _AnimationClass
	animationClassOnce sync.Once
)

func getAnimationClass() _AnimationClass {
	animationClassOnce.Do(func() {
		animationClass = _AnimationClass{objc.GetClass("CAAnimation")}
	})
	return animationClass
}

type _AnimationClass struct {
	class objc.Class
}

// An interface definition for the [Animation] class.
type IAnimation interface {
	objectivec.IObject
}

// The abstract superclass for animations in Core Animation.
//
// provides the basic support for the and protocols. You do not create instance of : to animate Core Animation layers or SceneKit objects, create instances of the concrete subclasses , , , or .
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation
type Animation struct {
	objectivec.Object
}

// AnimationFrom constructs a [Animation] from an unsafe.Pointer.
//
// The abstract superclass for animations in Core Animation.
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


// For animations attached to SceneKit objects, the duration for transitioning into the animation’s effect as it begins.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/fadeInDuration
func (a_ Animation) FadeInDuration() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("fadeInDuration"))
	return rv
}


// SetFadeInDuration sets the value of the fadeInDuration property.
// For animations attached to SceneKit objects, the duration for transitioning into the animation’s effect as it begins.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/fadeInDuration
func (a_ Animation) SetFadeInDuration(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFadeInDuration:"), value)
}


