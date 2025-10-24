// Code generated from Apple documentation for Speech. DO NOT EDIT.

package speech_test

import (
	"github.com/tmc/appledocs/generated/speech"
)

// Suppress unused import errors
var _ = speech.NewSFSpeechRecognizer

// ExampleNewSFSpeechRecognizer demonstrates how to create a SFSpeechRecognizer instance.
// Creates a speech recognizer associated with the user’s default language settings.
func ExampleNewSFSpeechRecognizer() {
	_ = speech.NewSFSpeechRecognizer()
	// Output:
}
