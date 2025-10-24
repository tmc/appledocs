// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CAAnimationGroup */


/* debug [class_header]: Header for CAAnimationGroup */
// The class instance for the [AnimationGroup] class.
var (
	AnimationGroupClass     _AnimationGroupClass
	AnimationGroupClassOnce sync.Once
)

func getAnimationGroupClass() _AnimationGroupClass {
	AnimationGroupClassOnce.Do(func() {
		AnimationGroupClass = _AnimationGroupClass{objc.GetClass("CAAnimationGroup")}
	})
	return AnimationGroupClass
}

type _AnimationGroupClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AnimationGroup */
// An interface definition for the [AnimationGroup] class.
type IAnimationGroup interface {
	IAnimation
	
/* debug [class_interface_properties]: Properties for AnimationGroup */
	// properties:
	Animations() []Animation
	SetAnimations(value []Animation)
	Delegate() objectivec.IObject
	SetDelegate(value objectivec.IObject)
	IsRemovedOnCompletion() bool
	SetIsRemovedOnCompletion(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AnimationGroup */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AnimationGroup */
// Alloc allocates a new instance without initialization.
func (ac _AnimationGroupClass) Alloc() AnimationGroup {
	rv := objc.Send[AnimationGroup](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
	return getAnimationGroupClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AnimationGroup */
// An object that allows multiple animations to be grouped and run concurrently.
//
// The grouped animations run in the time space specified by the instance. The duration of the grouped animations are not scaled to the duration of their . Instead, the animations are clipped to the duration of the animation group. For example, a 10 second animation grouped within an animation group with a duration of 5 seconds displays only the first 5 seconds of the animation. The following code shows how you can create a grouped animation containing opacity and scale animations to fade out a layer while expanding it. The animation starts with an opacity of and a scale of on all axes. As the animation’s scale increases to , the opacity drops to and the animated layer vanishes.


// An object that allows multiple animations to be grouped and run concurrently.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AnimationGroup *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AnimationGroup */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AnimationGroup */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AnimationGroup */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AnimationGroup */

// An array of objects to be evaluated in the time space of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimationGroup/animations
func (a_ AnimationGroup) Animations() []Animation {
	rv := objc.Send[[]Animation](a_.ID, objc.Sel("animations"))
	return rv
}/* debug [instance_properties/getter]: animations */


// An array of objects to be evaluated in the time space of the receiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAAnimationGroup/animations
func (a_ AnimationGroup) SetAnimations(value []Animation) {
	var nsArray objc.ID
	if len(value) > 0 {
		nsArray = objc.ID(objc.GetClass("NSMutableArray")).Send(objc.Sel("arrayWithCapacity:"), len(value))
		for _, item := range value {
			nsArray.Send(objc.Sel("addObject:"), item)
		}
	} else {
		nsArray = objc.ID(objc.GetClass("NSArray")).Send(objc.Sel("array"))
	}
	objc.Send[objc.ID](a_.ID, objc.Sel("setAnimations:"), nsArray)
}/* debug [instance_properties/setter]: animations */


// Specifies the receiver’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/delegate
func (a_ AnimationGroup) Delegate() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](a_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// Specifies the receiver’s delegate object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/delegate
func (a_ AnimationGroup) SetDelegate(value objectivec.IObject) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// Determines if the animation is removed from the target layer’s animations upon completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/isremovedoncompletion
func (a_ AnimationGroup) IsRemovedOnCompletion() bool {
	rv := objc.Send[bool](a_.ID, objc.Sel("isRemovedOnCompletion"))
	return rv
}/* debug [instance_properties/getter]: isRemovedOnCompletion */


// Determines if the animation is removed from the target layer’s animations upon completion.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/quartzcore/caanimation/isremovedoncompletion
func (a_ AnimationGroup) SetIsRemovedOnCompletion(value bool) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setIsRemovedOnCompletion:"), value)
}/* debug [instance_properties/setter]: isRemovedOnCompletion */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CAAnimationGroup */



