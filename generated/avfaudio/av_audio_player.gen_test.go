// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio_test

import (
	"github.com/tmc/appledocs/generated/avfaudio"
)

// Suppress unused import errors
var _ = avfaudio.NewAudioPlayer

// ExampleAudioPlayer_Pause demonstrates using Pause on a AudioPlayer instance.
// Pauses audio playback.
func ExampleAudioPlayer_Pause() {
	obj := avfaudio.NewAudioPlayer()
	obj.Pause()
	// Output:
	}

// ExampleAudioPlayer_Play demonstrates using Play on a AudioPlayer instance.
// Plays audio asynchronously.
func ExampleAudioPlayer_Play() {
	obj := avfaudio.NewAudioPlayer()
	_ = obj.Play()
	// Output:
	}

// ExampleAudioPlayer_PrepareToPlay demonstrates using PrepareToPlay on a AudioPlayer instance.
// Prepares the player for audio playback.
//
// Note: This example is not executed because PrepareToPlay crashes when called on bare NSObject
// (it's a protocol/category method that should be overridden by subclasses).
func ExampleAudioPlayer_PrepareToPlay() {
	obj := avfaudio.NewAudioPlayer()
	_ = obj.PrepareToPlay()
	}

// ExampleAudioPlayer_Stop demonstrates using Stop on a AudioPlayer instance.
// Stops playback and undoes the setup the system requires for playback.
func ExampleAudioPlayer_Stop() {
	obj := avfaudio.NewAudioPlayer()
	obj.Stop()
	// Output:
	}

// ExampleAudioPlayer_UpdateMeters demonstrates using UpdateMeters on a AudioPlayer instance.
// Refreshes the average and peak power values for all channels of an audio player.
func ExampleAudioPlayer_UpdateMeters() {
	obj := avfaudio.NewAudioPlayer()
	obj.UpdateMeters()
	// Output:
	}

