// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewSpeechRecognizer

// ExampleNewSpeechRecognizer demonstrates how to create a SpeechRecognizer instance.
// Initializes and returns an instance of the   class.
func ExampleNewSpeechRecognizer() {
	_ = appkit.NewSpeechRecognizer()
	// Output:
}
// ExampleSpeechRecognizer_StartListening demonstrates using StartListening on a SpeechRecognizer instance.
// Tells the speech recognition engine to begin listening for commands.
func ExampleSpeechRecognizer_StartListening() {
	obj := appkit.NewSpeechRecognizer()
	obj.StartListening()
	// Output:
	}

// ExampleSpeechRecognizer_StopListening demonstrates using StopListening on a SpeechRecognizer instance.
// Tells the speech recognition engine to suspend listening for commands.
func ExampleSpeechRecognizer_StopListening() {
	obj := appkit.NewSpeechRecognizer()
	obj.StopListening()
	// Output:
	}

