// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVExternalStorageDevice] class.
var (
	aVExternalStorageDeviceClass     _AVExternalStorageDeviceClass
	aVExternalStorageDeviceClassOnce sync.Once
)

func getAVExternalStorageDeviceClass() _AVExternalStorageDeviceClass {
	aVExternalStorageDeviceClassOnce.Do(func() {
		aVExternalStorageDeviceClass = _AVExternalStorageDeviceClass{objc.GetClass("AVExternalStorageDevice")}
	})
	return aVExternalStorageDeviceClass
}

type _AVExternalStorageDeviceClass struct {
	class objc.Class
}

// An interface definition for the [AVExternalStorageDevice] class.
type IAVExternalStorageDevice interface {
	objectivec.IObject
}

// Represents a physical external storage device that stores media assets.
//
// [Full Topic]: https://developer.apple.com/documentation/AVFoundation/AVExternalStorageDevice
type AVExternalStorageDevice struct {
	objectivec.Object
}

// AVExternalStorageDeviceFrom constructs a [AVExternalStorageDevice] from an unsafe.Pointer.
//
// Represents a physical external storage device that stores media assets.
func AVExternalStorageDeviceFrom(ptr unsafe.Pointer) AVExternalStorageDevice {
	return AVExternalStorageDevice{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (ac _AVExternalStorageDeviceClass) Alloc() AVExternalStorageDevice {
	rv := objc.Send[AVExternalStorageDevice](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ac _AVExternalStorageDeviceClass) New() AVExternalStorageDevice {
	rv := objc.Send[AVExternalStorageDevice](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AVExternalStorageDevice) Init() AVExternalStorageDevice {
	rv := objc.Send[AVExternalStorageDevice](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AVExternalStorageDevice) Autorelease() AVExternalStorageDevice {
	rv := objc.Send[AVExternalStorageDevice](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAVExternalStorageDevice creates a new AVExternalStorageDevice instance.
func NewAVExternalStorageDevice() AVExternalStorageDevice {
	return getAVExternalStorageDeviceClass().New()
}




