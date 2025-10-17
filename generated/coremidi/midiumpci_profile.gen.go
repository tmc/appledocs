// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDIUMPCIProfile] class.
var mIDIUMPCIProfileClass = _MIDIUMPCIProfileClass{objc.GetClass("MIDIUMPCIProfile")}

type _MIDIUMPCIProfileClass struct {
	class objc.Class
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile

type MIDIUMPCIProfile struct {
	objectivec.Object
}

// MIDIUMPCIProfileFrom constructs a [MIDIUMPCIProfile] from an unsafe.Pointer.
func MIDIUMPCIProfileFrom(ptr unsafe.Pointer) MIDIUMPCIProfile {
	return MIDIUMPCIProfile{objectivec.Object{objc.ID(ptr)}}
}

//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/setProfileState(_:enabledChannelCount:)
func (m_ MIDIUMPCIProfile) SetProfileStateEnabledChannelCountError(isEnabled bool, enabledChannelCount unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("setProfileState:enabledChannelCount:error:"), isEnabled, enabledChannelCount, error)
	return rv
}


