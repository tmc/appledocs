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

// An interface definition for the [MIDIUMPCIProfile] class.
type IMIDIUMPCIProfile interface {
	objectivec.IObject
	SetProfileStateEnabledChannelCountError(isEnabled bool, enabledChannelCount unsafe.Pointer, error unsafe.Pointer) bool
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
// Alloc allocates a new instance without initialization.
func (mc _MIDIUMPCIProfileClass) Alloc() MIDIUMPCIProfile {
	rv := objc.Send[MIDIUMPCIProfile](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new instance with a +1 retain count.
func (mc _MIDIUMPCIProfileClass) New() MIDIUMPCIProfile {
	rv := objc.Send[MIDIUMPCIProfile](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDIUMPCIProfile) Init() MIDIUMPCIProfile {
	rv := objc.Send[MIDIUMPCIProfile](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDIUMPCIProfile) Autorelease() MIDIUMPCIProfile {
	rv := objc.Send[MIDIUMPCIProfile](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDIUMPCIProfile creates a new MIDIUMPCIProfile instance.
func NewMIDIUMPCIProfile() MIDIUMPCIProfile {
	return mIDIUMPCIProfileClass.New()
}


//
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDIUMPCIProfile/setProfileState(_:enabledChannelCount:)
func (m_ MIDIUMPCIProfile) SetProfileStateEnabledChannelCountError(isEnabled bool, enabledChannelCount unsafe.Pointer, error unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("setProfileState:enabledChannelCount:error:"), isEnabled, enabledChannelCount, error)
	return rv
}


