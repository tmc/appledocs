// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [AVCaptureMultiCamSession] class.
var aVCaptureMultiCamSessionClass = _AVCaptureMultiCamSessionClass{objc.GetClass("AVCaptureMultiCamSession")}

type _AVCaptureMultiCamSessionClass struct {
	class objc.Class
}

// A capture session that supports simultaneous capture from multiple inputs of the same media type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureMultiCamSession

type AVCaptureMultiCamSession struct {
	AVCaptureSession
}

// AVCaptureMultiCamSessionFrom constructs a [AVCaptureMultiCamSession] from an unsafe.Pointer.
//
// A capture session that supports simultaneous capture from multiple inputs of the same media type.
func AVCaptureMultiCamSessionFrom(ptr unsafe.Pointer) AVCaptureMultiCamSession {
	return AVCaptureMultiCamSession{
		AVCaptureSession: AVCaptureSessionFrom(ptr),
	}
}



