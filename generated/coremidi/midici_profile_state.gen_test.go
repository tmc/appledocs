// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi_test

import (
	"github.com/tmc/appledocs/generated/coremidi"
)

// Suppress unused import errors
var _ = coremidi.NewMIDICIProfileState

// ExampleNewMIDICIProfileStateWithChannelEnabledProfilesDisabledProfiles demonstrates how to create a MIDICIProfileState instance using NewMIDICIProfileStateWithChannelEnabledProfilesDisabledProfiles.
// Creates a new profile state object for the specified MIDI channel and profiles.
func ExampleNewMIDICIProfileStateWithChannelEnabledProfilesDisabledProfiles() {
	_ = coremidi.NewMIDICIProfileStateWithChannelEnabledProfilesDisabledProfiles(
		coremidi.MIDIChannelNumber /* typedef */{}, // midiChannelNum MIDIChannelNumber /* typedef */
		[]coremidi.MIDICIProfile{}, // enabled []MIDICIProfile
		[]coremidi.MIDICIProfile{}, // disabled []MIDICIProfile
	)
	// Output:
}
// ExampleNewMIDICIProfileStateWithEnabledProfilesDisabledProfiles demonstrates how to create a MIDICIProfileState instance using NewMIDICIProfileStateWithEnabledProfilesDisabledProfiles.
// Creates a new profile state object for the specified profiles.
func ExampleNewMIDICIProfileStateWithEnabledProfilesDisabledProfiles() {
	_ = coremidi.NewMIDICIProfileStateWithEnabledProfilesDisabledProfiles(
		[]coremidi.MIDICIProfile{}, // enabled []MIDICIProfile
		[]coremidi.MIDICIProfile{}, // disabled []MIDICIProfile
	)
	// Output:
}
