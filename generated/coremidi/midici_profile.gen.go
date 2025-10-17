// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDICIProfile] class.
var mIDICIProfileClass = _MIDICIProfileClass{objc.GetClass("MIDICIProfile")}

type _MIDICIProfileClass struct {
	class objc.Class
}

// A mapping of MIDI messages to specific sounds and synthesis behaviors, such as General MIDI, a drawbar organ, and so on. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfile

type MIDICIProfile struct {
	objectivec.Object
}

// MIDICIProfileFrom constructs a [MIDICIProfile] from an unsafe.Pointer.
//
// A mapping of MIDI messages to specific sounds and synthesis behaviors, such as General MIDI, a drawbar organ, and so on.
func MIDICIProfileFrom(ptr unsafe.Pointer) MIDICIProfile {
	return MIDICIProfile{objectivec.Object{objc.ID(ptr)}}
}



