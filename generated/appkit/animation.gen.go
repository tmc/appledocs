// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Animation] class.
var animationClass = _AnimationClass{objc.GetClass("NSAnimation")}

type _AnimationClass struct {
	class objc.Class
}

// An object that manages the timing and progress of animations in the user interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimation

type Animation struct {
	objectivec.Object
}

// AnimationFrom constructs a [Animation] from an unsafe.Pointer.
//
// An object that manages the timing and progress of animations in the user interface.
func AnimationFrom(ptr unsafe.Pointer) Animation {
	return Animation{objectivec.Object{objc.ID(ptr)}}
}



