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

// An interface definition for the [AVContinuityDevice] class.
type IAVContinuityDevice interface {
	objectivec.IObject
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
// Alloc allocates a new instance without initialization.
func (ac _AVContinuityDeviceClass) Alloc() AVContinuityDevice {
	rv := objc.Send[AVContinuityDevice](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (ac _AVContinuityDeviceClass) New() AVContinuityDevice {
	rv := objc.Send[AVContinuityDevice](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVContinuityDevice) Init() AVContinuityDevice {
	rv := objc.Send[AVContinuityDevice](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVContinuityDevice) Autorelease() AVContinuityDevice {
	rv := objc.Send[AVContinuityDevice](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVContinuityDevice creates a new AVContinuityDevice instance.
func NewAVContinuityDevice() AVContinuityDevice {
	return aVContinuityDeviceClass.New()
}




