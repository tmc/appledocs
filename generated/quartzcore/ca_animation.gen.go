// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CAAnimation */


/* debug [class_header]: Header for CAAnimation */
// The class instance for the [Animation] class.
var (
	AnimationClass     _AnimationClass
	AnimationClassOnce sync.Once
)

func getAnimationClass() _AnimationClass {
	AnimationClassOnce.Do(func() {
		AnimationClass = _AnimationClass{objc.GetClass("CAAnimation")}
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
	AnimationEvents() []NAnimationEvent /* not a class type */
	SetAnimationEvents(value []NAnimationEvent /* not a class type */)
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	FadeInDuration() float64
	SetFadeInDuration(value float64)
	FadeOutDuration() float64
	SetFadeOutDuration(value float64)
	RemovedOnCompletion() bool
	SetRemovedOnCompletion(value bool)
	PreferredFrameRateRange() objc.IObject /* cross-framework: CAFrameRateRange */
	SetPreferredFrameRateRange(value objc.IObject /* cross-framework: CAFrameRateRange */)
	TimingFunction() IMediaTimingFunction
	SetTimingFunction(value IMediaTimingFunction)
	UsesSceneTimeBase() bool
	SetUsesSceneTimeBase(value bool)
	IsRemovedOnCompletion() bool
	SetIsRemovedOnCompletion(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Animation */
	// methods:
	ShouldArchiveValueForKey(key objc.IObject /* cross-framework: NSString */) bool
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
// The abstract superclass for animations in Core Animation.
//
// provides the basic support for the and protocols. You do not create instance of : to animate Core Animation layers or SceneKit objects, create instances of the concrete subclasses , , , or .


// The abstract superclass for animations in Core Animation.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Animation */

// Creates an animation from a SceneKit animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/init(SCNAnimation:)
func NewAnimationWithSCNAnimation(animation NAnimation /* not a class type */) Animation {
	rv := objc.Send[Animation](objc.ID(getAnimationClass().class), objc.Sel("animationWithSCNAnimation:"), animation)
	return rv
}/* debug [class_init_methods/constructor]: NewAnimationWithSCNAnimation */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Animation */

// Creates and returns a new instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/animation
func (ac _AnimationClass) Animation() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("animation"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=Animation) */


// Specifies the default value of the property with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/defaultValue(forKey:)
func (ac _AnimationClass) DefaultValueForKey(key objc.IObject /* cross-framework: NSString */) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("defaultValueForKey:"), key)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultValueForKey) */


// Creates an animation from a SceneKit animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/init(SCNAnimation:)
func (ac _AnimationClass) AnimationWithSCNAnimation(animation NAnimation /* not a class type */) IAnimation {
	rv := objc.Send[Animation](objc.ID(ac.class), objc.Sel("animationWithSCNAnimation:"), animation)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AnimationWithSCNAnimation) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Animation */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Animation */

// Specifies whether the value of the property for a given key is archived.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/shouldArchiveValue(forKey:)
func (a_ Animation) ShouldArchiveValueForKey(key objc.IObject /* cross-framework: NSString */) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("shouldArchiveValueForKey:"), key)
	return rv
}/* debug [instance_methods/method]: ShouldArchiveValueForKey */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Animation */

// For animations attached to SceneKit objects, a list of events attached to an animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/animationEvents
func (a_ Animation) AnimationEvents() []NAnimationEvent /* not a class type */ {
	rv := objc.Send[[]NAnimationEvent](a_.ID, objc.Sel("animationEvents"))
	return rv
}/* debug [instance_properties/getter]: animationEvents */


// For animations attached to SceneKit objects, a list of events attached to an animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/animationEvents
func (a_ Animation) SetAnimationEvents(value []NAnimationEvent /* not a class type */) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setAnimationEvents:"), nsArray)
}/* debug [instance_properties/setter]: animationEvents */


// Specifies the receiver’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/delegate
func (a_ Animation) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// Specifies the receiver’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/delegate
func (a_ Animation) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// For animations attached to SceneKit objects, the duration for transitioning into the animation’s effect as it begins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/fadeInDuration
func (a_ Animation) FadeInDuration() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("fadeInDuration"))
	return rv
}/* debug [instance_properties/getter]: fadeInDuration */


// For animations attached to SceneKit objects, the duration for transitioning into the animation’s effect as it begins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/fadeInDuration
func (a_ Animation) SetFadeInDuration(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFadeInDuration:"), value)
}/* debug [instance_properties/setter]: fadeInDuration */


// For animations attached to SceneKit objects, the duration for transitioning out of the animation’s effect as it ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/fadeOutDuration
func (a_ Animation) FadeOutDuration() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("fadeOutDuration"))
	return rv
}/* debug [instance_properties/getter]: fadeOutDuration */


// For animations attached to SceneKit objects, the duration for transitioning out of the animation’s effect as it ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/fadeOutDuration
func (a_ Animation) SetFadeOutDuration(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFadeOutDuration:"), value)
}/* debug [instance_properties/setter]: fadeOutDuration */


// Determines if the animation is removed from the target layer’s animations upon completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/isRemovedOnCompletion
func (a_ Animation) RemovedOnCompletion() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("removedOnCompletion"))
	return rv
}/* debug [instance_properties/getter]: removedOnCompletion */


// Determines if the animation is removed from the target layer’s animations upon completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/isRemovedOnCompletion
func (a_ Animation) SetRemovedOnCompletion(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRemovedOnCompletion:"), value)
}/* debug [instance_properties/setter]: removedOnCompletion */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/preferredFrameRateRange
func (a_ Animation) PreferredFrameRateRange() objc.IObject /* cross-framework: CAFrameRateRange */ {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("preferredFrameRateRange"))
	return rv
}/* debug [instance_properties/getter]: preferredFrameRateRange */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/preferredFrameRateRange
func (a_ Animation) SetPreferredFrameRateRange(value objc.IObject /* cross-framework: CAFrameRateRange */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredFrameRateRange:"), value)
}/* debug [instance_properties/setter]: preferredFrameRateRange */


// An optional timing function defining the pacing of the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/timingFunction
func (a_ Animation) TimingFunction() IMediaTimingFunction {
	rv := objc.Send[MediaTimingFunction](a_.ID, objc.Sel("timingFunction"))
	return rv
}/* debug [instance_properties/getter]: timingFunction */


// An optional timing function defining the pacing of the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/timingFunction
func (a_ Animation) SetTimingFunction(value IMediaTimingFunction) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTimingFunction:"), value)
}/* debug [instance_properties/setter]: timingFunction */


// For animations attached to SceneKit objects, a Boolean value that determines whether the animation is evaluated using the scene time or the system time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/usesSceneTimeBase
func (a_ Animation) UsesSceneTimeBase() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("usesSceneTimeBase"))
	return rv
}/* debug [instance_properties/getter]: usesSceneTimeBase */


// For animations attached to SceneKit objects, a Boolean value that determines whether the animation is evaluated using the scene time or the system time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/usesSceneTimeBase
func (a_ Animation) SetUsesSceneTimeBase(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUsesSceneTimeBase:"), value)
}/* debug [instance_properties/setter]: usesSceneTimeBase */


// Determines if the animation is removed from the target layer’s animations upon completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/isremovedoncompletion
func (a_ Animation) IsRemovedOnCompletion() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRemovedOnCompletion"))
	return rv
}/* debug [instance_properties/getter]: isRemovedOnCompletion */


// Determines if the animation is removed from the target layer’s animations upon completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/isremovedoncompletion
func (a_ Animation) SetIsRemovedOnCompletion(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsRemovedOnCompletion:"), value)
}/* debug [instance_properties/setter]: isRemovedOnCompletion */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CAAnimation */



