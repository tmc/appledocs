// Code generated from Apple documentation for QuartzCore. DO NOT EDIT.

package quartzcore

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [Transition] class.
var transitionClass = _TransitionClass{objc.GetClass("CATransition")}

type _TransitionClass struct {
	class objc.Class
}

// An object that provides an animated transition between a layer’s states. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/QuartzCore/CATransition

type Transition struct {
	Animation
}

// TransitionFrom constructs a [Transition] from an unsafe.Pointer.
//
// An object that provides an animated transition between a layer’s states.
func TransitionFrom(ptr unsafe.Pointer) Transition {
	return Transition{
		Animation: AnimationFrom(ptr),
	}
}



