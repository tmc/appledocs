// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [MIDICIProfileState] class.
var (
	MIDICIProfileStateClass     _MIDICIProfileStateClass
	MIDICIProfileStateClassOnce sync.Once
)

func getMIDICIProfileStateClass() _MIDICIProfileStateClass {
	MIDICIProfileStateClassOnce.Do(func() {
		MIDICIProfileStateClass = _MIDICIProfileStateClass{objc.GetClass("MIDICIProfileState")}
	})
	return MIDICIProfileStateClass
}

type _MIDICIProfileStateClass struct {
	class objc.Class
}

// An interface definition for the [MIDICIProfileState] class.
type IMIDICIProfileState interface {
	objectivec.IObject
	// properties:
	DisabledProfiles() IMIDICIProfile
	SetDisabledProfiles(value IMIDICIProfile)
	EnabledProfiles() IMIDICIProfile
	SetEnabledProfiles(value IMIDICIProfile)
	MidiChannel() MIDIChannelNumber /* typedef */
	SetMidiChannel(value MIDIChannelNumber /* typedef */)
	// methods:
}

// An object that provides the enabled and disabled profiles for a MIDI channel or port on a device.


// An object that provides the enabled and disabled profiles for a MIDI channel or port on a device.
//
// [Full Topic]
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

// Alloc allocates a new instance without initialization.
func (mc _MIDICIProfileStateClass) Alloc() MIDICIProfileState {
	rv := objc.Send[MIDICIProfileState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MIDICIProfileStateClass) New() MIDICIProfileState {
	rv := objc.Send[MIDICIProfileState](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MIDICIProfileState) Init() MIDICIProfileState {
	rv := objc.Send[MIDICIProfileState](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MIDICIProfileState) Autorelease() MIDICIProfileState {
	rv := objc.Send[MIDICIProfileState](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMIDICIProfileState creates a new MIDICIProfileState instance.
func NewMIDICIProfileState() MIDICIProfileState {
	return getMIDICIProfileStateClass().New()
}



// The object’s disabled profiles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofilestate/disabledprofiles
func (m_ MIDICIProfileState) DisabledProfiles() IMIDICIProfile {
	rv := objc.Send[MIDICIProfile](m_.ID, objc.Sel("disabledProfiles"))
	return rv
}


// The object’s disabled profiles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofilestate/disabledprofiles
func (m_ MIDICIProfileState) SetDisabledProfiles(value IMIDICIProfile) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDisabledProfiles:"), value)
}


// The object’s enabled profiles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofilestate/enabledprofiles
func (m_ MIDICIProfileState) EnabledProfiles() IMIDICIProfile {
	rv := objc.Send[MIDICIProfile](m_.ID, objc.Sel("enabledProfiles"))
	return rv
}


// The object’s enabled profiles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofilestate/enabledprofiles
func (m_ MIDICIProfileState) SetEnabledProfiles(value IMIDICIProfile) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setEnabledProfiles:"), value)
}


// The MIDI channel to which this state applies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofilestate/midichannel
func (m_ MIDICIProfileState) MidiChannel() MIDIChannelNumber /* typedef */ {
	rv := objc.Send[MIDIChannelNumber](m_.ID, objc.Sel("midiChannel"))
	return rv
}


// The MIDI channel to which this state applies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coremidi/midiciprofilestate/midichannel
func (m_ MIDICIProfileState) SetMidiChannel(value MIDIChannelNumber /* typedef */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setMidiChannel:"), value)
}



