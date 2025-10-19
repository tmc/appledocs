// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [AVExternalStorageDevice] class.
var aVExternalStorageDeviceClass = _AVExternalStorageDeviceClass{objc.GetClass("AVExternalStorageDevice")}

type _AVExternalStorageDeviceClass struct {
	class objc.Class
}

// Represents a physical external storage device that stores media assets. [Full Topic]
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



