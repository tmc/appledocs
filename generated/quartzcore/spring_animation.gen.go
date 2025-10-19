// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SpringAnimation] class.
var springAnimationClass = _SpringAnimationClass{objc.GetClass("CASpringAnimation")}

type _SpringAnimationClass struct {
	class objc.Class
}

// An interface definition for the [SpringAnimation] class.
type ISpringAnimation interface {
	IBasicAnimation
}

// An animation that applies a spring-like force to a layer’s properties. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return springAnimationClass.New()
}




