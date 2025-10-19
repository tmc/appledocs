// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [hasTrackingFrameCount] class.
var hasTrackingFrameCountClass = _hasTrackingFrameCountClass{objc.GetClass("hasTrackingFrameCount")}

type _hasTrackingFrameCountClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasTrackingFrameCount-c.ivar

type hasTrackingFrameCount struct {
	objectivec.Object
}

// hasTrackingFrameCountFrom constructs a [hasTrackingFrameCount] from an unsafe.Pointer.
func hasTrackingFrameCountFrom(ptr unsafe.Pointer) hasTrackingFrameCount {
	return hasTrackingFrameCount{objectivec.Object{objc.ID(ptr)}}
}



