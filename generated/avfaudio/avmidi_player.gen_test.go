// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio_test

import (
	"github.com/tmc/appledocs/generated/avfaudio"
)

// Suppress unused import errors
var _ = avfaudio.NewMIDIPlayer

// ExampleMIDIPlayer_PrepareToPlay demonstrates using PrepareToPlay on a MIDIPlayer instance.
// Prepares the player to play the sequence by prerolling all events.
//
// Note: This example is not executed because PrepareToPlay crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleMIDIPlayer_PrepareToPlay() {
	obj := avfaudio.NewMIDIPlayer()
	obj.PrepareToPlay()
	}

// ExampleMIDIPlayer_Stop demonstrates using Stop on a MIDIPlayer instance.
// Stops playing the sequence.
func ExampleMIDIPlayer_Stop() {
	obj := avfaudio.NewMIDIPlayer()
	obj.Stop()
	// Output:
	}

