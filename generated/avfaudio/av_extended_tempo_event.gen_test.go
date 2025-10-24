// Code generated from Apple documentation for AVFAudio. DO NOT EDIT.

package avfaudio_test

import (
	"github.com/tmc/appledocs/generated/avfaudio"
)

// Suppress unused import errors
var _ = avfaudio.NewExtendedTempoEvent

// ExampleNewExtendedTempoEventWithTempo demonstrates how to create a ExtendedTempoEvent instance using NewExtendedTempoEventWithTempo.
// Creates an extended tempo event.
func ExampleNewExtendedTempoEventWithTempo() {
	_ = avfaudio.NewExtendedTempoEventWithTempo(
		0.0, // tempo float64
	)
	// Output:
}
