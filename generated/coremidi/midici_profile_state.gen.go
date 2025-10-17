// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDICIProfileState] class.
var mIDICIProfileStateClass = _MIDICIProfileStateClass{objc.GetClass("MIDICIProfileState")}

type _MIDICIProfileStateClass struct {
	class objc.Class
}

// An object that provides the enabled and disabled profiles for a MIDI channel or port on a device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileState

type MIDICIProfileState struct {
	objectivec.Object
}

// MIDICIProfileStateFrom constructs a [MIDICIProfileState] from an unsafe.Pointer.
//
// An object that provides the enabled and disabled profiles for a MIDI channel or port on a device.
func MIDICIProfileStateFrom(ptr unsafe.Pointer) MIDICIProfileState {
	return MIDICIProfileState{objectivec.Object{objc.ID(ptr)}}
}



