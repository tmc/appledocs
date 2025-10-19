// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVMetadataItem] class.
var aVMetadataItemClass = _AVMetadataItemClass{objc.GetClass("AVMetadataItem")}

type _AVMetadataItemClass struct {
	class objc.Class
}

// A metadata item for an audiovisual asset or one of its tracks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVMetadataItem

type AVMetadataItem struct {
	objectivec.Object
}

// AVMetadataItemFrom constructs a [AVMetadataItem] from an unsafe.Pointer.
//
// A metadata item for an audiovisual asset or one of its tracks.
func AVMetadataItemFrom(ptr unsafe.Pointer) AVMetadataItem {
	return AVMetadataItem{objectivec.Object{objc.ID(ptr)}}
}



