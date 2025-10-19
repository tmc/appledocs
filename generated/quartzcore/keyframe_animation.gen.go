// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [KeyframeAnimation] class.
var keyframeAnimationClass = _KeyframeAnimationClass{objc.GetClass("CAKeyframeAnimation")}

type _KeyframeAnimationClass struct {
	class objc.Class
}

// An interface definition for the [KeyframeAnimation] class.
type IKeyframeAnimation interface {
	IPropertyAnimation
}

// An object that provides keyframe animation capabilities for a layer object. [Full Topic]
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

// New creates and returns a new instance with a +1 retain count.
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
	return keyframeAnimationClass.New()
}




