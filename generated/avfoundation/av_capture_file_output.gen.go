// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVCaptureFileOutput] class.
var aVCaptureFileOutputClass = _AVCaptureFileOutputClass{objc.GetClass("AVCaptureFileOutput")}

type _AVCaptureFileOutputClass struct {
	class objc.Class
}

// The abstract superclass for capture outputs that can record captured data to a file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureFileOutput

type AVCaptureFileOutput struct {
	AVCaptureOutput
}

// AVCaptureFileOutputFrom constructs a [AVCaptureFileOutput] from an unsafe.Pointer.
//
// The abstract superclass for capture outputs that can record captured data to a file.
func AVCaptureFileOutputFrom(ptr unsafe.Pointer) AVCaptureFileOutput {
	return AVCaptureFileOutput{
		AVCaptureOutput: AVCaptureOutputFrom(ptr),
	}
}



