// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDICIDeviceManager] class.
var mIDICIDeviceManagerClass = _MIDICIDeviceManagerClass{objc.GetClass("MIDICIDeviceManager")}

type _MIDICIDeviceManagerClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDeviceManager

type MIDICIDeviceManager struct {
	objectivec.Object
}

// MIDICIDeviceManagerFrom constructs a [MIDICIDeviceManager] from an unsafe.Pointer.
func MIDICIDeviceManagerFrom(ptr unsafe.Pointer) MIDICIDeviceManager {
	return MIDICIDeviceManager{objectivec.Object{objc.ID(ptr)}}
}



