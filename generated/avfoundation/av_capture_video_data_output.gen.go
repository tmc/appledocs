// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVCaptureVideoDataOutput] class.
var aVCaptureVideoDataOutputClass = _AVCaptureVideoDataOutputClass{objc.GetClass("AVCaptureVideoDataOutput")}

type _AVCaptureVideoDataOutputClass struct {
	class objc.Class
}

// A capture output that records video and provides access to video frames for processing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput

type AVCaptureVideoDataOutput struct {
	AVCaptureOutput
}

// AVCaptureVideoDataOutputFrom constructs a [AVCaptureVideoDataOutput] from an unsafe.Pointer.
//
// A capture output that records video and provides access to video frames for processing.
func AVCaptureVideoDataOutputFrom(ptr unsafe.Pointer) AVCaptureVideoDataOutput {
	return AVCaptureVideoDataOutput{
		AVCaptureOutput: AVCaptureOutputFrom(ptr),
	}
}

// Sets the sample buffer delegate and the queue for invoking callbacks. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureVideoDataOutput/setSampleBufferDelegate(_:queue:)
func (a_ AVCaptureVideoDataOutput) SetSampleBufferDelegateQueue(sampleBufferDelegate unsafe.Pointer, sampleBufferCallbackQueue unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setSampleBufferDelegate:queue:"), sampleBufferDelegate, sampleBufferCallbackQueue)
}


