// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [hasRightEyePosition] class.
var hasRightEyePositionClass = _hasRightEyePositionClass{objc.GetClass("hasRightEyePosition")}

type _hasRightEyePositionClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasRightEyePosition-c.ivar

type hasRightEyePosition struct {
	objectivec.Object
}

// hasRightEyePositionFrom constructs a [hasRightEyePosition] from an unsafe.Pointer.
func hasRightEyePositionFrom(ptr unsafe.Pointer) hasRightEyePosition {
	return hasRightEyePosition{objectivec.Object{objc.ID(ptr)}}
}



