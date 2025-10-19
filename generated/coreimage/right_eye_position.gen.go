// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [rightEyePosition] class.
var rightEyePositionClass = _rightEyePositionClass{objc.GetClass("rightEyePosition")}

type _rightEyePositionClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/rightEyePosition-c.ivar

type rightEyePosition struct {
	objectivec.Object
}

// rightEyePositionFrom constructs a [rightEyePosition] from an unsafe.Pointer.
func rightEyePositionFrom(ptr unsafe.Pointer) rightEyePosition {
	return rightEyePosition{objectivec.Object{objc.ID(ptr)}}
}



