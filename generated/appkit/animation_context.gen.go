// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AnimationContext] class.
var animationContextClass = _AnimationContextClass{objc.GetClass("NSAnimationContext")}

type _AnimationContextClass struct {
	class objc.Class
}

// An animation context, which contains information about environment and state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSAnimationContext

type AnimationContext struct {
	objectivec.Object
}

// AnimationContextFrom constructs a [AnimationContext] from an unsafe.Pointer.
//
// An animation context, which contains information about environment and state.
func AnimationContextFrom(ptr unsafe.Pointer) AnimationContext {
	return AnimationContext{objectivec.Object{objc.ID(ptr)}}
}



