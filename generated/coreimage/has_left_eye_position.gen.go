// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [hasLeftEyePosition] class.
var hasLeftEyePositionClass = _hasLeftEyePositionClass{objc.GetClass("hasLeftEyePosition")}

type _hasLeftEyePositionClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasLeftEyePosition-c.ivar

type hasLeftEyePosition struct {
	objectivec.Object
}

// hasLeftEyePositionFrom constructs a [hasLeftEyePosition] from an unsafe.Pointer.
func hasLeftEyePositionFrom(ptr unsafe.Pointer) hasLeftEyePosition {
	return hasLeftEyePosition{objectivec.Object{objc.ID(ptr)}}
}



