// Code generated from Apple documentation for AVFoundation. DO NOT EDIT.

package avfoundation_test

import (
	"github.com/tmc/appledocs/generated/avfoundation"
)

// Suppress unused import errors
var _ = avfoundation.NewPlayer

// ExampleNewPlayer demonstrates how to create a Player instance.
// Creates a player object.
func ExampleNewPlayer() {
	_ = avfoundation.NewPlayer()
	// Output:
}
// ExamplePlayer_CancelPendingPrerolls demonstrates using CancelPendingPrerolls on a Player instance.
// Cancels any pending preroll requests and invokes the corresponding completion handlers, if present.
func ExamplePlayer_CancelPendingPrerolls() {
	obj := avfoundation.NewPlayer()
	obj.CancelPendingPrerolls()
	// Output:
	}

// ExamplePlayer_CurrentTime demonstrates using CurrentTime on a Player instance.
// Returns the current time of the current player item.
func ExamplePlayer_CurrentTime() {
	obj := avfoundation.NewPlayer()
	_ = obj.CurrentTime()
	// Output:
	}

// ExamplePlayer_Pause demonstrates using Pause on a Player instance.
// Pauses playback of the current item.
func ExamplePlayer_Pause() {
	obj := avfoundation.NewPlayer()
	obj.Pause()
	// Output:
	}

// ExamplePlayer_Play demonstrates using Play on a Player instance.
// Begins playback of the current item.
func ExamplePlayer_Play() {
	obj := avfoundation.NewPlayer()
	obj.Play()
	// Output:
	}

