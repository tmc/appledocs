// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVCaptureConnection] class.
var aVCaptureConnectionClass = _AVCaptureConnectionClass{objc.GetClass("AVCaptureConnection")}

type _AVCaptureConnectionClass struct {
	class objc.Class
}

// An object that represents a connection from a capture input to a capture output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureConnection

type AVCaptureConnection struct {
	objectivec.Object
}

// AVCaptureConnectionFrom constructs a [AVCaptureConnection] from an unsafe.Pointer.
//
// An object that represents a connection from a capture input to a capture output.
func AVCaptureConnectionFrom(ptr unsafe.Pointer) AVCaptureConnection {
	return AVCaptureConnection{objectivec.Object{objc.ID(ptr)}}
}



