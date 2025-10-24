// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CASpringAnimation */


/* debug [class_header]: Header for CASpringAnimation */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for SpringAnimation */
// An interface definition for the [SpringAnimation] class.
type ISpringAnimation interface {
	IBasicAnimation
	
/* debug [class_interface_properties]: Properties for SpringAnimation */
	// properties:
	AllowsOverdamping() bool
	SetAllowsOverdamping(value bool)
	Bounce() float64
	Damping() float64
	SetDamping(value float64)
	InitialVelocity() float64
	SetInitialVelocity(value float64)
	Mass() float64
	SetMass(value float64)
	PerceptualDuration() float64
	SettlingDuration() float64
	Stiffness() float64
	SetStiffness(value float64)
	ToValue() objectivec.IObject
	SetToValue(value objectivec.IObject)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for SpringAnimation */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for SpringAnimation */
// Alloc allocates a new instance without initialization.
func (sc _SpringAnimationClass) Alloc() SpringAnimation {
	rv := objc.Send[SpringAnimation](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for SpringAnimation */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for SpringAnimation */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/init(perceptualDuration:bounce:)
func NewSpringAnimationWithPerceptualDurationBounce(perceptualDuration float64, bounce float64) SpringAnimation {
	instance := getSpringAnimationClass().Alloc()
	rv := objc.Send[SpringAnimation](instance.ID, objc.Sel("initWithPerceptualDuration:bounce:"), perceptualDuration, bounce)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewSpringAnimationWithPerceptualDurationBounce */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for SpringAnimation */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for SpringAnimation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for SpringAnimation */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for SpringAnimation */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/allowsOverdamping
func (s_ SpringAnimation) AllowsOverdamping() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("allowsOverdamping"))
	return rv
}/* debug [instance_properties/getter]: allowsOverdamping */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/allowsOverdamping
func (s_ SpringAnimation) SetAllowsOverdamping(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setAllowsOverdamping:"), value)
}/* debug [instance_properties/setter]: allowsOverdamping */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/bounce
func (s_ SpringAnimation) Bounce() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("bounce"))
	return rv
}/* debug [instance_properties/getter]: bounce */


// Defines how the spring’s motion should be damped due to the forces of friction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/damping
func (s_ SpringAnimation) Damping() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("damping"))
	return rv
}/* debug [instance_properties/getter]: damping */


// Defines how the spring’s motion should be damped due to the forces of friction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/damping
func (s_ SpringAnimation) SetDamping(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDamping:"), value)
}/* debug [instance_properties/setter]: damping */


// The initial velocity of the object attached to the spring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/initialVelocity
func (s_ SpringAnimation) InitialVelocity() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("initialVelocity"))
	return rv
}/* debug [instance_properties/getter]: initialVelocity */


// The initial velocity of the object attached to the spring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/initialVelocity
func (s_ SpringAnimation) SetInitialVelocity(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setInitialVelocity:"), value)
}/* debug [instance_properties/setter]: initialVelocity */


// The mass of the object attached to the end of the spring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/mass
func (s_ SpringAnimation) Mass() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("mass"))
	return rv
}/* debug [instance_properties/getter]: mass */


// The mass of the object attached to the end of the spring.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/mass
func (s_ SpringAnimation) SetMass(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setMass:"), value)
}/* debug [instance_properties/setter]: mass */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/perceptualDuration
func (s_ SpringAnimation) PerceptualDuration() float64 {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("perceptualDuration"))
	return rv
}/* debug [instance_properties/getter]: perceptualDuration */


// The estimated duration required for the spring system to be considered at rest.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/settlingDuration
func (s_ SpringAnimation) SettlingDuration() float64 {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("settlingDuration"))
	return rv
}/* debug [instance_properties/getter]: settlingDuration */


// The spring stiffness coefficient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/stiffness
func (s_ SpringAnimation) Stiffness() float64 {
	rv := objc.Send[float64](s_.ID, objc.Sel("stiffness"))
	return rv
}/* debug [instance_properties/getter]: stiffness */


// The spring stiffness coefficient.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CASpringAnimation/stiffness
func (s_ SpringAnimation) SetStiffness(value float64) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setStiffness:"), value)
}/* debug [instance_properties/setter]: stiffness */


// Defines the value the receiver uses to end interpolation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cabasicanimation/tovalue
func (s_ SpringAnimation) ToValue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](s_.ID, objc.Sel("toValue"))
	return rv
}/* debug [instance_properties/getter]: toValue */


// Defines the value the receiver uses to end interpolation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/cabasicanimation/tovalue
func (s_ SpringAnimation) SetToValue(value objectivec.IObject) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setToValue:"), value)
}/* debug [instance_properties/setter]: toValue */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CASpringAnimation */


