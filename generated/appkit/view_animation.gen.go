// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [ViewAnimation] class.
var viewAnimationClass = _ViewAnimationClass{objc.GetClass("NSViewAnimation")}

type _ViewAnimationClass struct {
	class objc.Class
}

// An animation of an app’s views, limited to changes in frame location and size, and to fade-in and fade-out effects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSViewAnimation

type ViewAnimation struct {
	Animation
}

// ViewAnimationFrom constructs a [ViewAnimation] from an unsafe.Pointer.
//
// An animation of an app’s views, limited to changes in frame location and size, and to fade-in and fade-out effects.
func ViewAnimationFrom(ptr unsafe.Pointer) ViewAnimation {
	return ViewAnimation{
		Animation: AnimationFrom(ptr),
	}
}



