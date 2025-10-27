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
	

	// properties:
	AnimationEvents() []NAnimationEvent /* not a class type */
	SetAnimationEvents(value []NAnimationEvent /* not a class type */)
	FadeInDuration() float64
	SetFadeInDuration(value float64)
	FadeOutDuration() float64
	SetFadeOutDuration(value float64)
	RemovedOnCompletion() bool
	SetRemovedOnCompletion(value bool)
	PreferredFrameRateRange() CAFrameRateRange
	SetPreferredFrameRateRange(value CAFrameRateRange)
	TimingFunction() IMediaTimingFunction
	SetTimingFunction(value IMediaTimingFunction)
	UsesSceneTimeBase() bool
	SetUsesSceneTimeBase(value bool)
	IsRemovedOnCompletion() bool
	SetIsRemovedOnCompletion(value bool)


	

	// methods:
	ShouldArchiveValueForKey(key foundation.foundation.INSString) bool


}





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






// Creates an animation from a SceneKit animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/init(SCNAnimation:)
func NewAnimationWithSCNAnimation(animation NAnimation /* not a class type */) Animation {
	rv := objc.Send[Animation](objc.ID(getAnimationClass().class), objc.Sel("animationWithSCNAnimation:"), animation)
	return rv
}







// Creates and returns a new instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/animation
func (ac _AnimationClass) Animation() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(ac.class), objc.Sel("animation"))
	return rv
}


// Specifies the default value of the property with the specified key.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/defaultValue(forKey:)
func (ac _AnimationClass) DefaultValueForKey(key foundation.foundation.INSString) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(ac.class), objc.Sel("defaultValueForKey:"), key)
	return rv
}


// Creates an animation from a SceneKit animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/init(SCNAnimation:)
func (ac _AnimationClass) AnimationWithSCNAnimation(animation NAnimation /* not a class type */) IAnimation {
	rv := objc.Send[Animation](objc.ID(ac.class), objc.Sel("animationWithSCNAnimation:"), animation)
	return rv
}












// Specifies whether the value of the property for a given key is archived.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/shouldArchiveValue(forKey:)
func (a_ Animation) ShouldArchiveValueForKey(key foundation.foundation.INSString) bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("shouldArchiveValueForKey:"), key)
	return rv
}







// For animations attached to SceneKit objects, a list of events attached to an animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/animationEvents
func (a_ Animation) AnimationEvents() []NAnimationEvent /* not a class type */ {
	rv := objc.Send[[]NAnimationEvent](a_.ID, objc.Sel("animationEvents"))
	// Slice of concrete type to slice of interface - needs conversion
	result := make([]NAnimationEvent /* not a class type */, len(rv))
	for i, v := range rv {
		result[i] = v
	}
	return result
}


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
}


// For animations attached to SceneKit objects, the duration for transitioning into the animation’s effect as it begins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/fadeInDuration
func (a_ Animation) FadeInDuration() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("fadeInDuration"))
	return rv
}


// For animations attached to SceneKit objects, the duration for transitioning into the animation’s effect as it begins.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/fadeInDuration
func (a_ Animation) SetFadeInDuration(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFadeInDuration:"), value)
}


// For animations attached to SceneKit objects, the duration for transitioning out of the animation’s effect as it ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/fadeOutDuration
func (a_ Animation) FadeOutDuration() float64 {
	rv := objc.Send[float64](a_.ID, objc.Sel("fadeOutDuration"))
	return rv
}


// For animations attached to SceneKit objects, the duration for transitioning out of the animation’s effect as it ends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/fadeOutDuration
func (a_ Animation) SetFadeOutDuration(value float64) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setFadeOutDuration:"), value)
}


// Determines if the animation is removed from the target layer’s animations upon completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/isRemovedOnCompletion
func (a_ Animation) RemovedOnCompletion() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("removedOnCompletion"))
	return rv
}


// Determines if the animation is removed from the target layer’s animations upon completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/isRemovedOnCompletion
func (a_ Animation) SetRemovedOnCompletion(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setRemovedOnCompletion:"), value)
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/preferredFrameRateRange
func (a_ Animation) PreferredFrameRateRange() CAFrameRateRange {
	rv := objc.Send[objc.ID](a_.ID, objc.Sel("preferredFrameRateRange"))
	return rv
}


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/preferredFrameRateRange
func (a_ Animation) SetPreferredFrameRateRange(value CAFrameRateRange) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setPreferredFrameRateRange:"), value)
}


// An optional timing function defining the pacing of the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/timingFunction
func (a_ Animation) TimingFunction() IMediaTimingFunction {
	rv := objc.Send[MediaTimingFunction](a_.ID, objc.Sel("timingFunction"))
	return rv
}


// An optional timing function defining the pacing of the animation.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/timingFunction
func (a_ Animation) SetTimingFunction(value IMediaTimingFunction) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setTimingFunction:"), value)
}


// For animations attached to SceneKit objects, a Boolean value that determines whether the animation is evaluated using the scene time or the system time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/usesSceneTimeBase
func (a_ Animation) UsesSceneTimeBase() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("usesSceneTimeBase"))
	return rv
}


// For animations attached to SceneKit objects, a Boolean value that determines whether the animation is evaluated using the scene time or the system time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimation/usesSceneTimeBase
func (a_ Animation) SetUsesSceneTimeBase(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setUsesSceneTimeBase:"), value)
}


// Determines if the animation is removed from the target layer’s animations upon completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/isremovedoncompletion
func (a_ Animation) IsRemovedOnCompletion() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRemovedOnCompletion"))
	return rv
}


// Determines if the animation is removed from the target layer’s animations upon completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/isremovedoncompletion
func (a_ Animation) SetIsRemovedOnCompletion(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsRemovedOnCompletion:"), value)
}








