// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class MIDICIProfileState */


/* debug [class_header]: Header for MIDICIProfileState */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MIDICIProfileState */
// An interface definition for the [MIDICIProfileState] class.
type IMIDICIProfileState interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MIDICIProfileState */
	// properties:
	DisabledProfiles() []MIDICIProfile
	EnabledProfiles() []MIDICIProfile
	MidiChannel() MIDIChannelNumber /* typedef */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MIDICIProfileState */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MIDICIProfileState */
// Alloc allocates a new instance without initialization.
func (mc _MIDICIProfileStateClass) Alloc() MIDICIProfileState {
	rv := objc.Send[MIDICIProfileState](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MIDICIProfileState */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MIDICIProfileState */

// Creates a new profile state object for the specified MIDI channel and profiles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileState/init(channel:enabledProfiles:disabledProfiles:)
func NewMIDICIProfileStateWithChannelEnabledProfilesDisabledProfiles(midiChannelNum MIDIChannelNumber /* typedef */, enabled []MIDICIProfile, disabled []MIDICIProfile) MIDICIProfileState {
	instance := getMIDICIProfileStateClass().Alloc()
	rv := objc.Send[MIDICIProfileState](instance.ID, objc.Sel("initWithChannel:enabledProfiles:disabledProfiles:"), midiChannelNum, enabled, disabled)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMIDICIProfileStateWithChannelEnabledProfilesDisabledProfiles */


// Creates a new profile state object for the specified profiles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileState/init(enabledProfiles:disabledProfiles:)
func NewMIDICIProfileStateWithEnabledProfilesDisabledProfiles(enabled []MIDICIProfile, disabled []MIDICIProfile) MIDICIProfileState {
	instance := getMIDICIProfileStateClass().Alloc()
	rv := objc.Send[MIDICIProfileState](instance.ID, objc.Sel("initWithEnabledProfiles:disabledProfiles:"), enabled, disabled)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewMIDICIProfileStateWithEnabledProfilesDisabledProfiles */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MIDICIProfileState */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MIDICIProfileState */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MIDICIProfileState */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MIDICIProfileState */

// The object’s disabled profiles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileState/disabledProfiles
func (m_ MIDICIProfileState) DisabledProfiles() []MIDICIProfile {
	rv := objc.Send[[]MIDICIProfile](m_.ID, objc.Sel("disabledProfiles"))
	return rv
}/* debug [instance_properties/getter]: disabledProfiles */


// The object’s enabled profiles.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileState/enabledProfiles
func (m_ MIDICIProfileState) EnabledProfiles() []MIDICIProfile {
	rv := objc.Send[[]MIDICIProfile](m_.ID, objc.Sel("enabledProfiles"))
	return rv
}/* debug [instance_properties/getter]: enabledProfiles */


// The MIDI channel to which this state applies.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreMIDI/MIDICIProfileState/midiChannel
func (m_ MIDICIProfileState) MidiChannel() MIDIChannelNumber /* typedef */ {
	rv := objc.Send[objectivec.IObject](m_.ID, objc.Sel("midiChannel"))
	return rv
}/* debug [instance_properties/getter]: midiChannel */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class MIDICIProfileState */


