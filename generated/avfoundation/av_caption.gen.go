// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVCaption] class.
var aVCaptionClass = _AVCaptionClass{objc.GetClass("AVCaption")}

type _AVCaptionClass struct {
	class objc.Class
}

// An object that represents text to present over a time range. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaption

type AVCaption struct {
	objectivec.Object
}

// AVCaptionFrom constructs a [AVCaption] from an unsafe.Pointer.
//
// An object that represents text to present over a time range.
func AVCaptionFrom(ptr unsafe.Pointer) AVCaption {
	return AVCaption{objectivec.Object{objc.ID(ptr)}}
}



