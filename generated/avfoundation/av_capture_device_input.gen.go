// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVCaptureDeviceInput] class.
var aVCaptureDeviceInputClass = _AVCaptureDeviceInputClass{objc.GetClass("AVCaptureDeviceInput")}

type _AVCaptureDeviceInputClass struct {
	class objc.Class
}

// An object that provides media input from a capture device to a capture session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDeviceInput

type AVCaptureDeviceInput struct {
	AVCaptureInput
}

// AVCaptureDeviceInputFrom constructs a [AVCaptureDeviceInput] from an unsafe.Pointer.
//
// An object that provides media input from a capture device to a capture session.
func AVCaptureDeviceInputFrom(ptr unsafe.Pointer) AVCaptureDeviceInput {
	return AVCaptureDeviceInput{
		AVCaptureInput: AVCaptureInputFrom(ptr),
	}
}



