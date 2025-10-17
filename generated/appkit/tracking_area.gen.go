// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [TrackingArea] class.
var trackingAreaClass = _TrackingAreaClass{objc.GetClass("NSTrackingArea")}

type _TrackingAreaClass struct {
	class objc.Class
}

// A region of a view that generates mouse-tracking and cursor-update events when the pointer is over that region. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AppKit/NSTrackingArea

type TrackingArea struct {
	objectivec.Object
}

// TrackingAreaFrom constructs a [TrackingArea] from an unsafe.Pointer.
//
// A region of a view that generates mouse-tracking and cursor-update events when the pointer is over that region.
func TrackingAreaFrom(ptr unsafe.Pointer) TrackingArea {
	return TrackingArea{objectivec.Object{objc.ID(ptr)}}
}



