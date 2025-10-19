// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVCaptureDevice] class.
var aVCaptureDeviceClass = _AVCaptureDeviceClass{objc.GetClass("AVCaptureDevice")}

type _AVCaptureDeviceClass struct {
	class objc.Class
}

// An interface definition for the [AVCaptureDevice] class.
type IAVCaptureDevice interface {
	objectivec.IObject
}

// An object that represents a hardware or virtual capture device like a camera or microphone. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice

type AVCaptureDevice struct {
	objectivec.Object
}

// AVCaptureDeviceFrom constructs a [AVCaptureDevice] from an unsafe.Pointer.
//
// An object that represents a hardware or virtual capture device like a camera or microphone.
func AVCaptureDeviceFrom(ptr unsafe.Pointer) AVCaptureDevice {
	return AVCaptureDevice{objectivec.Object{objc.ID(ptr)}}
}
// Alloc allocates a new instance without initialization.
func (ac _AVCaptureDeviceClass) Alloc() AVCaptureDevice {
	rv := objc.Send[AVCaptureDevice](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AVCaptureDeviceClass) New() AVCaptureDevice {
	rv := objc.Send[AVCaptureDevice](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVCaptureDevice) Init() AVCaptureDevice {
	rv := objc.Send[AVCaptureDevice](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVCaptureDevice) Autorelease() AVCaptureDevice {
	rv := objc.Send[AVCaptureDevice](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVCaptureDevice creates a new AVCaptureDevice instance.
func NewAVCaptureDevice() AVCaptureDevice {
	return aVCaptureDeviceClass.New()
}


// Returns the default device for the specified device type, media type, and position. [Full Topic]

//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVCaptureDevice/default(_:for:position:)
func (ac _AVCaptureDeviceClass) DefaultDeviceWithDeviceTypeMediaTypePosition(deviceType unsafe.Pointer, mediaType unsafe.Pointer, position unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(ac.class), objc.Sel("defaultDeviceWithDeviceType:mediaType:position:"), deviceType, mediaType, position)
	return rv
}


