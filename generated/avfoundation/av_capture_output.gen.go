// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVCaptureOutput] class.
var aVCaptureOutputClass = _AVCaptureOutputClass{objc.GetClass("AVCaptureOutput")}

type _AVCaptureOutputClass struct {
	class objc.Class
}

// An abstract superclass for objects that provide media output destinations for a capture session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureOutput

type AVCaptureOutput struct {
	objectivec.Object
}

// AVCaptureOutputFrom constructs a [AVCaptureOutput] from an unsafe.Pointer.
//
// An abstract superclass for objects that provide media output destinations for a capture session.
func AVCaptureOutputFrom(ptr unsafe.Pointer) AVCaptureOutput {
	return AVCaptureOutput{objectivec.Object{objc.ID(ptr)}}
}



