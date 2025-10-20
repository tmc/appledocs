// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SpringAnimation] class.
var (
	springAnimationClass     _SpringAnimationClass
	springAnimationClassOnce sync.Once
)

func getSpringAnimationClass() _SpringAnimationClass {
	springAnimationClassOnce.Do(func() {
		springAnimationClass = _SpringAnimationClass{objc.GetClass("CASpringAnimation")}
	})
	return springAnimationClass
}

type _SpringAnimationClass struct {
	class objc.Class
}

// An interface definition for the [SpringAnimation] class.
type ISpringAnimation interface {
	IBasicAnimation
}

// An animation that applies a spring-like force to a layer’s properties.
//
// You would typically use a spring animation to animate a layer’s position so that it appears to be pulled towards a target by a spring. The further the layer is from the target, the greater the acceleration towards it is. allows control over physically based attributes such as the spring’s damping and stiffness. You can use a spring animation to animation properties of a layer other than its position. The following code shows how to create a spring animation that bounces a layer into view by animating its scale from to . Because the spring animation can overshoot its , the animated layer may exceed its frame.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation
type SpringAnimation struct {
	BasicAnimation
}

// SpringAnimationFrom constructs a [SpringAnimation] from an unsafe.Pointer.
//
// An animation that applies a spring-like force to a layer’s properties.
func SpringAnimationFrom(ptr unsafe.Pointer) SpringAnimation {
	return SpringAnimation{
		BasicAnimation: BasicAnimationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (sc _SpringAnimationClass) Alloc() SpringAnimation {
	rv := objc.Send[SpringAnimation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SpringAnimationClass) New() SpringAnimation {
	rv := objc.Send[SpringAnimation](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ SpringAnimation) Init() SpringAnimation {
	rv := objc.Send[SpringAnimation](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ SpringAnimation) Autorelease() SpringAnimation {
	rv := objc.Send[SpringAnimation](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSpringAnimation creates a new SpringAnimation instance.
func NewSpringAnimation() SpringAnimation {
	return getSpringAnimationClass().New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/bounce
func (s_ SpringAnimation) Bounce() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("bounce"))
	return rv
}



