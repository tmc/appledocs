// Code generated from Apple documentation for AppKit. DO NOT EDIT.

package appkit_test

import (
	"github.com/tmc/appledocs/generated/appkit"
)

// Suppress unused import errors
var _ = appkit.NewSound

// ExampleNewSoundNamed demonstrates how to create a Sound instance using NewSoundNamed.
// Returns the   instance associated with a given name.
func ExampleNewSoundNamed() {
	_ = appkit.NewSoundNamed(
		appkit.SoundName{}, // name SoundName
	)
	// Output:
}

// ExampleNewSoundWithContentsOfFileByReference demonstrates how to create a Sound instance using NewSoundWithContentsOfFileByReference.
// Initializes the receiver with the audio data located at a given filepath.
func ExampleNewSoundWithContentsOfFileByReference() {
	_ = appkit.NewSoundWithContentsOfFileByReference(
		"/tmp/test", // path string
		false,       // byRef bool
	)
	// Output:
}

// ExampleNewSoundWithPasteboard demonstrates how to create a Sound instance using NewSoundWithPasteboard.
// Initializes the receiver with data from a pasteboard. The pasteboard should contain a type returned by  .   expects the data to have a proper magic number, sound header, and data for the formats it supports.
func ExampleNewSoundWithPasteboard() {
	_ = appkit.NewSoundWithPasteboard(
		appkit.NSPasteboard{}, // pasteboard NSPasteboard
	)
	// Output:
}
