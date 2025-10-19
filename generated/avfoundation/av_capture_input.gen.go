// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVCaptureInput] class.
var aVCaptureInputClass = _AVCaptureInputClass{objc.GetClass("AVCaptureInput")}

type _AVCaptureInputClass struct {
	class objc.Class
}

// An abstract superclass for objects that provide input data to a capture session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureInput

type AVCaptureInput struct {
	objectivec.Object
}

// AVCaptureInputFrom constructs a [AVCaptureInput] from an unsafe.Pointer.
//
// An abstract superclass for objects that provide input data to a capture session.
func AVCaptureInputFrom(ptr unsafe.Pointer) AVCaptureInput {
	return AVCaptureInput{objectivec.Object{objc.ID(ptr)}}
}



