// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit

import (
	"unsafe"

	"github.com/ebitengine/purego/objc"
)

// The class instance for the [TrackingArea] class.
var TrackingAreaClass objc.Class

func init() {
	TrackingAreaClass = objc.GetClass("NSTrackingArea")
}

type TrackingArea struct {
	objc.ID
}

func TrackingAreaFrom(ptr unsafe.Pointer) TrackingArea {
	return TrackingArea{
		ID: objc.ID(ptr),
	}
}



