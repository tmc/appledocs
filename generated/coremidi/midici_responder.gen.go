// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDICIResponder] class.
var mIDICIResponderClass = _MIDICIResponderClass{objc.GetClass("MIDICIResponder")}

type _MIDICIResponderClass struct {
	class objc.Class
}

// An object that responds to MIDI-CI inquiries from an initiator on behalf of a MIDI client, and handles profile and property exchange operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIResponder

type MIDICIResponder struct {
	objectivec.Object
}

// MIDICIResponderFrom constructs a [MIDICIResponder] from an unsafe.Pointer.
//
// An object that responds to MIDI-CI inquiries from an initiator on behalf of a MIDI client, and handles profile and property exchange operations.
func MIDICIResponderFrom(ptr unsafe.Pointer) MIDICIResponder {
	return MIDICIResponder{objectivec.Object{objc.ID(ptr)}}
}



