// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [trackingFrameCount] class.
var trackingFrameCountClass = _trackingFrameCountClass{objc.GetClass("trackingFrameCount")}

type _trackingFrameCountClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/trackingFrameCount-c.ivar

type trackingFrameCount struct {
	objectivec.Object
}

// trackingFrameCountFrom constructs a [trackingFrameCount] from an unsafe.Pointer.
func trackingFrameCountFrom(ptr unsafe.Pointer) trackingFrameCount {
	return trackingFrameCount{objectivec.Object{objc.ID(ptr)}}
}



