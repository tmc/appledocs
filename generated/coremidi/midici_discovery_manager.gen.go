// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDICIDiscoveryManager] class.
var mIDICIDiscoveryManagerClass = _MIDICIDiscoveryManagerClass{objc.GetClass("MIDICIDiscoveryManager")}

type _MIDICIDiscoveryManagerClass struct {
	class objc.Class
}

// A singleton object that performs systemwide MIDI-CI discovery. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIDiscoveryManager

type MIDICIDiscoveryManager struct {
	objectivec.Object
}

// MIDICIDiscoveryManagerFrom constructs a [MIDICIDiscoveryManager] from an unsafe.Pointer.
//
// A singleton object that performs systemwide MIDI-CI discovery.
func MIDICIDiscoveryManagerFrom(ptr unsafe.Pointer) MIDICIDiscoveryManager {
	return MIDICIDiscoveryManager{objectivec.Object{objc.ID(ptr)}}
}



