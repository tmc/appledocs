// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [Gradient] class.
var GradientClass objc.Class

func init() {
	GradientClass = objc.GetClass("NSGradient")
}

type Gradient struct {
	objc.ID
}

func GradientFrom(ptr unsafe.Pointer) Gradient {
	return Gradient{
		ID: objc.ID(ptr),
	}
}



