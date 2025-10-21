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

// An interface definition for the [Animation] class.
type IAnimation interface {
	objectivec.IObject
	ShouldArchiveValueForKey(key string) bool
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


// Creates an animation from a SceneKit animation.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/init(SCNAnimation:)
func NewAnimationWithSCNAnimation(animation unsafe.Pointer) Animation {
	rv := objc.Send[Animation](objc.ID(getAnimationClass().class), objc.Sel("animationWithSCNAnimation:"), animation)
	return rv
}


// Creates and returns a new instance.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/animation
func (ac _AnimationClass) Animation() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("animation"))
	return rv
}

// Specifies the default value of the property with the specified key.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/defaultValue(forKey:)
func (ac _AnimationClass) DefaultValueForKey(key string) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("defaultValueForKey:"), objc.String(key))
	return rv
}

// Creates an animation from a SceneKit animation.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/init(SCNAnimation:)
func (ac _AnimationClass) AnimationWithSCNAnimation(animation unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("animationWithSCNAnimation:"), animation)
	return rv
}

// Specifies whether the value of the property for a given key is archived.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/shouldArchiveValue(forKey:)
func (a_ Animation) ShouldArchiveValueForKey(key string) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("shouldArchiveValueForKey:"), objc.String(key))
	return rv
}

// For animations attached to SceneKit objects, a list of events attached to an animation.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/animationEvents
func (a_ Animation) AnimationEvents() []quartzcore.SCNAnimationEvent {
	rv := objc.Send[[]quartzcore.SCNAnimationEvent](a_.ID, objc.Sel("animationEvents"))
	return rv
}


// SetAnimationEvents sets the value of the animationEvents property.
// For animations attached to SceneKit objects, a list of events attached to an animation.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/animationEvents
func (a_ Animation) SetAnimationEvents(value []quartzcore.SCNAnimationEvent) {
	// Convert Go slice to NSArray
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
}
// Specifies the receiver’s delegate object.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/delegate
func (a_ Animation) Delegate() objc.ID {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// Specifies the receiver’s delegate object.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/delegate
func (a_ Animation) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
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
// For animations attached to SceneKit objects, the duration for transitioning out of the animation’s effect as it ends.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/fadeOutDuration
func (a_ Animation) FadeOutDuration() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("fadeOutDuration"))
	return rv
}


// SetFadeOutDuration sets the value of the fadeOutDuration property.
// For animations attached to SceneKit objects, the duration for transitioning out of the animation’s effect as it ends.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/fadeOutDuration
func (a_ Animation) SetFadeOutDuration(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFadeOutDuration:"), value)
}
// Determines if the animation is removed from the target layer’s animations upon completion.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/isRemovedOnCompletion
func (a_ Animation) RemovedOnCompletion() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("removedOnCompletion"))
	return rv
}


// SetRemovedOnCompletion sets the value of the removedOnCompletion property.
// Determines if the animation is removed from the target layer’s animations upon completion.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/isRemovedOnCompletion
func (a_ Animation) SetRemovedOnCompletion(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRemovedOnCompletion:"), value)
}
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/preferredFrameRateRange
func (a_ Animation) PreferredFrameRateRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("preferredFrameRateRange"))
	return rv
}


// SetPreferredFrameRateRange sets the value of the preferredFrameRateRange property.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/preferredFrameRateRange
func (a_ Animation) SetPreferredFrameRateRange(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredFrameRateRange:"), value)
}
// An optional timing function defining the pacing of the animation.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/timingFunction
func (a_ Animation) TimingFunction() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("timingFunction"))
	return rv
}


// SetTimingFunction sets the value of the timingFunction property.
// An optional timing function defining the pacing of the animation.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/timingFunction
func (a_ Animation) SetTimingFunction(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTimingFunction:"), value)
}
// For animations attached to SceneKit objects, a Boolean value that determines whether the animation is evaluated using the scene time or the system time.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/usesSceneTimeBase
func (a_ Animation) UsesSceneTimeBase() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("usesSceneTimeBase"))
	return rv
}


// SetUsesSceneTimeBase sets the value of the usesSceneTimeBase property.
// For animations attached to SceneKit objects, a Boolean value that determines whether the animation is evaluated using the scene time or the system time.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/usesSceneTimeBase
func (a_ Animation) SetUsesSceneTimeBase(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUsesSceneTimeBase:"), value)
}


