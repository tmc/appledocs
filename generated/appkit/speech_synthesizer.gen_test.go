// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewSpeechSynthesizer

// ExampleNewSpeechSynthesizerWithVoice demonstrates how to create a SpeechSynthesizer instance using NewSpeechSynthesizerWithVoice.
// Initializes the receiver with a voice.
func ExampleNewSpeechSynthesizerWithVoice() {
	_ = appkit.NewSpeechSynthesizerWithVoice(
		appkit.SpeechSynthesizerVoiceName{}, // voice SpeechSynthesizerVoiceName
	)
	// Output:
}
