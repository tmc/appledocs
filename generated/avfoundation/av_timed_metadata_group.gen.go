// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVTimedMetadataGroup] class.
var aVTimedMetadataGroupClass = _AVTimedMetadataGroupClass{objc.GetClass("AVTimedMetadataGroup")}

type _AVTimedMetadataGroupClass struct {
	class objc.Class
}

// A collection of metadata items that are valid for use during a specific time range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVTimedMetadataGroup

type AVTimedMetadataGroup struct {
	AVMetadataGroup
}

// AVTimedMetadataGroupFrom constructs a [AVTimedMetadataGroup] from an unsafe.Pointer.
//
// A collection of metadata items that are valid for use during a specific time range.
func AVTimedMetadataGroupFrom(ptr unsafe.Pointer) AVTimedMetadataGroup {
	return AVTimedMetadataGroup{
		AVMetadataGroup: AVMetadataGroupFrom(ptr),
	}
}



