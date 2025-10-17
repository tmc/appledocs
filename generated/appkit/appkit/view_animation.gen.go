// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [ViewAnimation] class.
var ViewAnimationClass objc.Class

func init() {
	ViewAnimationClass = objc.GetClass("NSViewAnimation")
}

type ViewAnimation struct {
	objc.ID
}

func ViewAnimationFrom(ptr unsafe.Pointer) ViewAnimation {
	return ViewAnimation{
		ID: objc.ID(ptr),
	}
}




