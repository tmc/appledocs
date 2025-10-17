// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Gradient] class.
var gradientClass = _GradientClass{objc.GetClass("NSGradient")}

type _GradientClass struct {
	class objc.Class
}

// An object that can draw gradient fill colors [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSGradient

type Gradient struct {
	objectivec.Object
}

// GradientFrom constructs a [Gradient] from an unsafe.Pointer.
//
// An object that can draw gradient fill colors
func GradientFrom(ptr unsafe.Pointer) Gradient {
	return Gradient{objectivec.Object{objc.ID(ptr)}}
}



