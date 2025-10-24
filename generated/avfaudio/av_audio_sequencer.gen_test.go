// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio_test

import (
	"github.com/tmc/appledocs/generated/avfaudio"
)

// Suppress unused import errors
var _ = avfaudio.NewAudioSequencer

// ExampleNewAudioSequencer demonstrates how to create a AudioSequencer instance.
// Creates an audio sequencer object.
func ExampleNewAudioSequencer() {
	_ = avfaudio.NewAudioSequencer()
	// Output:
}
// ExampleAudioSequencer_CreateAndAppendTrack demonstrates using CreateAndAppendTrack on a AudioSequencer instance.
// Creates a new music track and appends it to the sequencer’s list.
func ExampleAudioSequencer_CreateAndAppendTrack() {
	obj := avfaudio.NewAudioSequencer()
	_ = obj.CreateAndAppendTrack()
	// Output:
	}

// ExampleAudioSequencer_PrepareToPlay demonstrates using PrepareToPlay on a AudioSequencer instance.
// Gets ready to play the sequence by prerolling all events.
//
// Note: This example is not executed because PrepareToPlay crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleAudioSequencer_PrepareToPlay() {
	obj := avfaudio.NewAudioSequencer()
	obj.PrepareToPlay()
	}

// ExampleAudioSequencer_ReverseEvents demonstrates using ReverseEvents on a AudioSequencer instance.
// Reverses the order of all events in all music tracks, including the tempo track.
func ExampleAudioSequencer_ReverseEvents() {
	obj := avfaudio.NewAudioSequencer()
	obj.ReverseEvents()
	// Output:
	}

// ExampleAudioSequencer_Stop demonstrates using Stop on a AudioSequencer instance.
// Stops the sequencer’s player.
func ExampleAudioSequencer_Stop() {
	obj := avfaudio.NewAudioSequencer()
	obj.Stop()
	// Output:
	}

