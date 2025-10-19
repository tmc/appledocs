// Code generated from Apple documentation for CoreImage. DO NOT EDIT.

package coreimage

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [hasTrackingID] class.
var hasTrackingIDClass = _hasTrackingIDClass{objc.GetClass("hasTrackingID")}

type _hasTrackingIDClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreImage/CIFaceFeature/hasTrackingID-c.ivar

type hasTrackingID struct {
	objectivec.Object
}

// hasTrackingIDFrom constructs a [hasTrackingID] from an unsafe.Pointer.
func hasTrackingIDFrom(ptr unsafe.Pointer) hasTrackingID {
	return hasTrackingID{objectivec.Object{objc.ID(ptr)}}
}



