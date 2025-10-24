// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase_test

import (
	"github.com/tmc/appledocs/generated/phase"
)

// Suppress unused import errors
var _ = phase.NewPHASEGroup

// ExamplePHASEGroup_Mute demonstrates using Mute on a PHASEGroup instance.
// Silences the group.
func ExamplePHASEGroup_Mute() {
	obj := phase.NewPHASEGroup()
	obj.Mute()
	// Output:
	}

// ExamplePHASEGroup_Solo demonstrates using Solo on a PHASEGroup instance.
// Silences all other groups.
func ExamplePHASEGroup_Solo() {
	obj := phase.NewPHASEGroup()
	obj.Solo()
	// Output:
	}

// ExamplePHASEGroup_Unmute demonstrates using Unmute on a PHASEGroup instance.
// Restores the group’s volume.
func ExamplePHASEGroup_Unmute() {
	obj := phase.NewPHASEGroup()
	obj.Unmute()
	// Output:
	}

// ExamplePHASEGroup_UnregisterFromEngine demonstrates using UnregisterFromEngine on a PHASEGroup instance.
// Removes the group from the engine’s dictionary.
func ExamplePHASEGroup_UnregisterFromEngine() {
	obj := phase.NewPHASEGroup()
	obj.UnregisterFromEngine()
	// Output:
	}

// ExamplePHASEGroup_Unsolo demonstrates using Unsolo on a PHASEGroup instance.
// Restores the other groups’ volume.
func ExamplePHASEGroup_Unsolo() {
	obj := phase.NewPHASEGroup()
	obj.Unsolo()
	// Output:
	}

