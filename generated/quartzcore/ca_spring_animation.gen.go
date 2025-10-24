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
	// properties:
	Bounce() float64
	ToValue() unsafe.Pointer
	SetToValue(value unsafe.Pointer)
	AllowsOverdamping() bool
	SetAllowsOverdamping(value bool)
	Damping() float64
	SetDamping(value float64)
	InitialVelocity() float64
	SetInitialVelocity(value float64)
	Mass() float64
	SetMass(value float64)
	PerceptualDuration() float64
	SetPerceptualDuration(value float64)
	SettlingDuration() float64
	SetSettlingDuration(value float64)
	Stiffness() float64
	SetStiffness(value float64)
	// methods:
}

// An animation that applies a spring-like force to a layer’s properties.
//
// You would typically use a spring animation to animate a layer’s position so that it appears to be pulled towards a target by a spring. The further the layer is from the target, the greater the acceleration towards it is. allows control over physically based attributes such as the spring’s damping and stiffness. You can use a spring animation to animation properties of a layer other than its position. The following code shows how to create a spring animation that bounces a layer into view by animating its scale from to . Because the spring animation can overshoot its , the animated layer may exceed its frame.


// An animation that applies a spring-like force to a layer’s properties.
//
// [Full Topic]
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



// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/bounce
func (s_ SpringAnimation) Bounce() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("bounce"))
	return rv
}


// Defines the value the receiver uses to end interpolation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cabasicanimation/tovalue
func (s_ SpringAnimation) ToValue() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("toValue"))
	return rv
}


// Defines the value the receiver uses to end interpolation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cabasicanimation/tovalue
func (s_ SpringAnimation) SetToValue(value unsafe.Pointer) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setToValue:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caspringanimation/allowsoverdamping
func (s_ SpringAnimation) AllowsOverdamping() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("allowsOverdamping"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caspringanimation/allowsoverdamping
func (s_ SpringAnimation) SetAllowsOverdamping(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAllowsOverdamping:"), value)
}


// Defines how the spring’s motion should be damped due to the forces of friction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caspringanimation/damping
func (s_ SpringAnimation) Damping() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("damping"))
	return rv
}


// Defines how the spring’s motion should be damped due to the forces of friction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caspringanimation/damping
func (s_ SpringAnimation) SetDamping(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDamping:"), value)
}


// The initial velocity of the object attached to the spring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caspringanimation/initialvelocity
func (s_ SpringAnimation) InitialVelocity() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("initialVelocity"))
	return rv
}


// The initial velocity of the object attached to the spring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caspringanimation/initialvelocity
func (s_ SpringAnimation) SetInitialVelocity(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setInitialVelocity:"), value)
}


// The mass of the object attached to the end of the spring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caspringanimation/mass
func (s_ SpringAnimation) Mass() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("mass"))
	return rv
}


// The mass of the object attached to the end of the spring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caspringanimation/mass
func (s_ SpringAnimation) SetMass(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMass:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caspringanimation/perceptualduration
func (s_ SpringAnimation) PerceptualDuration() float64 {
	rv := objc.Send[TimeInterval](s_.ID, objc.Sel("perceptualDuration"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caspringanimation/perceptualduration
func (s_ SpringAnimation) SetPerceptualDuration(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setPerceptualDuration:"), value)
}


// The estimated duration required for the spring system to be considered at rest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caspringanimation/settlingduration
func (s_ SpringAnimation) SettlingDuration() float64 {
	rv := objc.Send[TimeInterval](s_.ID, objc.Sel("settlingDuration"))
	return rv
}


// The estimated duration required for the spring system to be considered at rest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caspringanimation/settlingduration
func (s_ SpringAnimation) SetSettlingDuration(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setSettlingDuration:"), value)
}


// The spring stiffness coefficient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caspringanimation/stiffness
func (s_ SpringAnimation) Stiffness() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("stiffness"))
	return rv
}


// The spring stiffness coefficient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caspringanimation/stiffness
func (s_ SpringAnimation) SetStiffness(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStiffness:"), value)
}



