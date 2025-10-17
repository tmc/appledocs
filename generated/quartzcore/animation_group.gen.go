// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AnimationGroup] class.
var animationGroupClass = _AnimationGroupClass{objc.GetClass("CAAnimationGroup")}

type _AnimationGroupClass struct {
	class objc.Class
}

// An object that allows multiple animations to be grouped and run concurrently. [Full Topic]
//
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



