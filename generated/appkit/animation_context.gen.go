// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [AnimationContext] class.
var AnimationContextClass objc.Class

func init() {
	AnimationContextClass = objc.GetClass("NSAnimationContext")
}

type AnimationContext struct {
	objc.ID
}

func AnimationContextFrom(ptr unsafe.Pointer) AnimationContext {
	return AnimationContext{
		ID: objc.ID(ptr),
	}
}



