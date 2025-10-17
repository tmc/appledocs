// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Animation] class.
var AnimationClass objc.Class

func init() {
	AnimationClass = objc.GetClass("NSAnimation")
}

type Animation struct {
	objc.ID
}

func AnimationFrom(ptr unsafe.Pointer) Animation {
	return Animation{
		ID: objc.ID(ptr),
	}
}




