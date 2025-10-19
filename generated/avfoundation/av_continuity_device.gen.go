// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVContinuityDevice] class.
var aVContinuityDeviceClass = _AVContinuityDeviceClass{objc.GetClass("AVContinuityDevice")}

type _AVContinuityDeviceClass struct {
	class objc.Class
}

// A class that represents a physical iOS device that’s nearby and can provide access to its cameras and microphones. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVContinuityDevice

type AVContinuityDevice struct {
	objectivec.Object
}

// AVContinuityDeviceFrom constructs a [AVContinuityDevice] from an unsafe.Pointer.
//
// A class that represents a physical iOS device that’s nearby and can provide access to its cameras and microphones.
func AVContinuityDeviceFrom(ptr unsafe.Pointer) AVContinuityDevice {
	return AVContinuityDevice{objectivec.Object{objc.ID(ptr)}}
}



