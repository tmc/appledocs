// Code generated from Apple documentation for CoreMIDI. DO NOT EDIT.

package coremidi_test

import (
	"github.com/tmc/appledocs/generated/coremidi"
)

// Suppress unused import errors
var _ = coremidi.NewMIDIUMPMutableFunctionBlock

// ExampleNewMIDIUMPMutableFunctionBlockWithNameDirectionFirstGroupTotalGroupsSpannedMaxSysEx8StreamsMIDI1InfoUIHintIsEnabled demonstrates how to create a MIDIUMPMutableFunctionBlock instance using NewMIDIUMPMutableFunctionBlockWithNameDirectionFirstGroupTotalGroupsSpannedMaxSysEx8StreamsMIDI1InfoUIHintIsEnabled.
func ExampleNewMIDIUMPMutableFunctionBlockWithNameDirectionFirstGroupTotalGroupsSpannedMaxSysEx8StreamsMIDI1InfoUIHintIsEnabled() {
	_ = coremidi.NewMIDIUMPMutableFunctionBlockWithNameDirectionFirstGroupTotalGroupsSpannedMaxSysEx8StreamsMIDI1InfoUIHintIsEnabled(
		"name", // name string
		coremidi.MIDIUMPFunctionBlockDirection{}, // direction MIDIUMPFunctionBlockDirection
		coremidi.MIDIUMPGroupNumber{}, // firstGroup MIDIUMPGroupNumber
		coremidi.MIDIUInteger7{}, // totalGroupsSpanned MIDIUInteger7
		coremidi.MIDIUInteger7{}, // maxSysEx8Streams MIDIUInteger7
		coremidi.MIDIUMPFunctionBlockMIDI1Info{}, // MIDI1Info MIDIUMPFunctionBlockMIDI1Info
		coremidi.MIDIUMPFunctionBlockUIHint{}, // UIHint MIDIUMPFunctionBlockUIHint
		false, // isEnabled bool
	)
	// Output:
}

