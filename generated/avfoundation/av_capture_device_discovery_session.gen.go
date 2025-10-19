// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVCaptureDeviceDiscoverySession] class.
var aVCaptureDeviceDiscoverySessionClass = _AVCaptureDeviceDiscoverySessionClass{objc.GetClass("AVCaptureDeviceDiscoverySession")}

type _AVCaptureDeviceDiscoverySessionClass struct {
	class objc.Class
}

// An object that finds capture devices that match specific search criteria. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/DiscoverySession

type AVCaptureDeviceDiscoverySession struct {
	objectivec.Object
}

// AVCaptureDeviceDiscoverySessionFrom constructs a [AVCaptureDeviceDiscoverySession] from an unsafe.Pointer.
//
// An object that finds capture devices that match specific search criteria.
func AVCaptureDeviceDiscoverySessionFrom(ptr unsafe.Pointer) AVCaptureDeviceDiscoverySession {
	return AVCaptureDeviceDiscoverySession{objectivec.Object{objc.ID(ptr)}}
}



