// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase_test

import (
	"github.com/tmc/appledocs/generated/phase"
)

// Suppress unused import errors
var _ = phase.NewPHASESoundEvent

// ExamplePHASESoundEvent_Pause demonstrates using Pause on a PHASESoundEvent instance.
// Pauses the sound event.
func ExamplePHASESoundEvent_Pause() {
	obj := phase.NewPHASESoundEvent()
	obj.Pause()
	// Output:
	}

// ExamplePHASESoundEvent_Resume demonstrates using Resume on a PHASESoundEvent instance.
// Resumes the sound event.
func ExamplePHASESoundEvent_Resume() {
	obj := phase.NewPHASESoundEvent()
	obj.Resume()
	// Output:
	}

// ExamplePHASESoundEvent_StopAndInvalidate demonstrates using StopAndInvalidate on a PHASESoundEvent instance.
// Stops a sound event and prevents it from resuming.
func ExamplePHASESoundEvent_StopAndInvalidate() {
	obj := phase.NewPHASESoundEvent()
	obj.StopAndInvalidate()
	// Output:
	}

