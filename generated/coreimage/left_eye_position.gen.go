// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [leftEyePosition] class.
var leftEyePositionClass = _leftEyePositionClass{objc.GetClass("leftEyePosition")}

type _leftEyePositionClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/leftEyePosition-c.ivar

type leftEyePosition struct {
	objectivec.Object
}

// leftEyePositionFrom constructs a [leftEyePosition] from an unsafe.Pointer.
func leftEyePositionFrom(ptr unsafe.Pointer) leftEyePosition {
	return leftEyePosition{objectivec.Object{objc.ID(ptr)}}
}



