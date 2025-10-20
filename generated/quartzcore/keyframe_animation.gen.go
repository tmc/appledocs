// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [KeyframeAnimation] class.
var (
	keyframeAnimationClass     _KeyframeAnimationClass
	keyframeAnimationClassOnce sync.Once
)

func getKeyframeAnimationClass() _KeyframeAnimationClass {
	keyframeAnimationClassOnce.Do(func() {
		keyframeAnimationClass = _KeyframeAnimationClass{objc.GetClass("CAKeyframeAnimation")}
	})
	return keyframeAnimationClass
}

type _KeyframeAnimationClass struct {
	class objc.Class
}

// An interface definition for the [KeyframeAnimation] class.
type IKeyframeAnimation interface {
	IPropertyAnimation
}

// An object that provides keyframe animation capabilities for a layer object.
//
// You create a object using the inherited method, specifying the key path of the property that you want to animate on the layer. You can then specify the keyframe values to use to control the timing and animation behavior. For most types of animations, you specify the keyframe values using the and properties. During the animation, Core Animation generates intermediate values by interpolating between the values you provide. When animating a value that is a coordinate point, such as the layer’s position, you can specify a for that point to follow instead of individual values. The pacing of the animation is controlled by the timing information you provide. The following code shows how to create a keyframe animation that animates a layer’s background color from red to green to blue over a two second duration.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation
type KeyframeAnimation struct {
	PropertyAnimation
}

// KeyframeAnimationFrom constructs a [KeyframeAnimation] from an unsafe.Pointer.
//
// An object that provides keyframe animation capabilities for a layer object.
func KeyframeAnimationFrom(ptr unsafe.Pointer) KeyframeAnimation {
	return KeyframeAnimation{
		PropertyAnimation: PropertyAnimationFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (kc _KeyframeAnimationClass) Alloc() KeyframeAnimation {
	rv := objc.Send[KeyframeAnimation](objc.ID(kc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (kc _KeyframeAnimationClass) New() KeyframeAnimation {
	rv := objc.Send[KeyframeAnimation](objc.ID(kc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (k_ KeyframeAnimation) Init() KeyframeAnimation {
	rv := objc.Send[KeyframeAnimation](k_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (k_ KeyframeAnimation) Autorelease() KeyframeAnimation {
	rv := objc.Send[KeyframeAnimation](k_.ID, objc.Sel("autorelease"))
	return rv
}

// NewKeyframeAnimation creates a new KeyframeAnimation instance.
func NewKeyframeAnimation() KeyframeAnimation {
	return getKeyframeAnimationClass().New()
}


// An optional array of objects that define the pacing for each keyframe segment.
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/timingFunctions
func (k_ KeyframeAnimation) TimingFunctions() []MediaTimingFunction {
	rv := objc.Send[[]MediaTimingFunction](k_.ID, objc.Sel("timingFunctions"))
	return rv
}


// SetTimingFunctions sets the value of the timingFunctions property.
// An optional array of objects that define the pacing for each keyframe segment.

//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CAKeyframeAnimation/timingFunctions
func (k_ KeyframeAnimation) SetTimingFunctions(value []MediaTimingFunction) {
	objc.Send[objc.ID](k_.ID, objc.Sel("setTimingFunctions:"), value)
}


