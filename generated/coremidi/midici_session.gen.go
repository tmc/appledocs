// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDICISession] class.
var mIDICISessionClass = _MIDICISessionClass{objc.GetClass("MIDICISession")}

type _MIDICISessionClass struct {
	class objc.Class
}

// An object that represents a MIDI-CI session. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICISession

type MIDICISession struct {
	objectivec.Object
}

// MIDICISessionFrom constructs a [MIDICISession] from an unsafe.Pointer.
//
// An object that represents a MIDI-CI session.
func MIDICISessionFrom(ptr unsafe.Pointer) MIDICISession {
	return MIDICISession{objectivec.Object{objc.ID(ptr)}}
}



