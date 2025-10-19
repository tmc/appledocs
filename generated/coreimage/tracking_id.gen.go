// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [trackingID] class.
var trackingIDClass = _trackingIDClass{objc.GetClass("trackingID")}

type _trackingIDClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/trackingID-c.ivar

type trackingID struct {
	objectivec.Object
}

// trackingIDFrom constructs a [trackingID] from an unsafe.Pointer.
func trackingIDFrom(ptr unsafe.Pointer) trackingID {
	return trackingID{objectivec.Object{objc.ID(ptr)}}
}



