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



