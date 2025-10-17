// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDICIDevice] class.
var mIDICIDeviceClass = _MIDICIDeviceClass{objc.GetClass("MIDICIDevice")}

type _MIDICIDeviceClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDevice

type MIDICIDevice struct {
	objectivec.Object
}

// MIDICIDeviceFrom constructs a [MIDICIDevice] from an unsafe.Pointer.
func MIDICIDeviceFrom(ptr unsafe.Pointer) MIDICIDevice {
	return MIDICIDevice{objectivec.Object{objc.ID(ptr)}}
}



