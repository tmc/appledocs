// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [SpringAnimation] class.
var (
	SpringAnimationClass     _SpringAnimationClass
	SpringAnimationClassOnce sync.Once
)

func getSpringAnimationClass() _SpringAnimationClass {
	SpringAnimationClassOnce.Do(func() {
		SpringAnimationClass = _SpringAnimationClass{objc.GetClass("CASpringAnimation")}
	})
	return SpringAnimationClass
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
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/init(perceptualDuration:bounce:)
func NewSpringAnimationWithPerceptualDurationBounce(perceptualDuration unsafe.Pointer, bounce float64) SpringAnimation {
	instance := getSpringAnimationClass().Alloc()
	rv := objc.Send[SpringAnimation](instance.ID, objc.Sel("initWithPerceptualDuration:bounce:"), perceptualDuration, bounce)
	rv.Autorelease()
	return rv
}


//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/allowsOverdamping
func (s_ SpringAnimation) AllowsOverdamping() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("allowsOverdamping"))
	return rv
}


// SetAllowsOverdamping sets the value of the allowsOverdamping property.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/allowsOverdamping
func (s_ SpringAnimation) SetAllowsOverdamping(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAllowsOverdamping:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/bounce
func (s_ SpringAnimation) Bounce() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("bounce"))
	return rv
}

// Defines how the spring’s motion should be damped due to the forces of friction.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/damping
func (s_ SpringAnimation) Damping() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("damping"))
	return rv
}


// SetDamping sets the value of the damping property.
// Defines how the spring’s motion should be damped due to the forces of friction.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/damping
func (s_ SpringAnimation) SetDamping(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDamping:"), value)
}

// The initial velocity of the object attached to the spring.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/initialVelocity
func (s_ SpringAnimation) InitialVelocity() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("initialVelocity"))
	return rv
}


// SetInitialVelocity sets the value of the initialVelocity property.
// The initial velocity of the object attached to the spring.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/initialVelocity
func (s_ SpringAnimation) SetInitialVelocity(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setInitialVelocity:"), value)
}

// The mass of the object attached to the end of the spring.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/mass
func (s_ SpringAnimation) Mass() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("mass"))
	return rv
}


// SetMass sets the value of the mass property.
// The mass of the object attached to the end of the spring.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/mass
func (s_ SpringAnimation) SetMass(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMass:"), value)
}

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/perceptualDuration
func (s_ SpringAnimation) PerceptualDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("perceptualDuration"))
	return rv
}

// The estimated duration required for the spring system to be considered at rest.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/settlingDuration
func (s_ SpringAnimation) SettlingDuration() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("settlingDuration"))
	return rv
}

// The spring stiffness coefficient.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/stiffness
func (s_ SpringAnimation) Stiffness() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("stiffness"))
	return rv
}


// SetStiffness sets the value of the stiffness property.
// The spring stiffness coefficient.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/stiffness
func (s_ SpringAnimation) SetStiffness(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStiffness:"), value)
}


