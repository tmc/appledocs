// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [BasicAnimation] class.
var basicAnimationClass = _BasicAnimationClass{objc.GetClass("CABasicAnimation")}

type _BasicAnimationClass struct {
	class objc.Class
}

// An object that provides basic, single-keyframe animation capabilities for a layer property. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CABasicAnimation

type BasicAnimation struct {
	PropertyAnimation
}

// BasicAnimationFrom constructs a [BasicAnimation] from an unsafe.Pointer.
//
// An object that provides basic, single-keyframe animation capabilities for a layer property.
func BasicAnimationFrom(ptr unsafe.Pointer) BasicAnimation {
	return BasicAnimation{
		PropertyAnimation: PropertyAnimationFrom(ptr),
	}
}



