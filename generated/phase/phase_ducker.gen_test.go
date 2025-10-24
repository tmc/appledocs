// Code generated from Apple documentation for PHASE. DO NOT EDIT.

package phase_test

import (
	"github.com/tmc/appledocs/generated/phase"
)

// Suppress unused import errors
var _ = phase.NewPHASEDucker

// ExamplePHASEDucker_Activate demonstrates using Activate on a PHASEDucker instance.
// Instructs the ducker to begin altering sound.
func ExamplePHASEDucker_Activate() {
	obj := phase.NewPHASEDucker()
	obj.Activate()
	// Output:
	}

// ExamplePHASEDucker_Deactivate demonstrates using Deactivate on a PHASEDucker instance.
// Stops the ducker from altering sound.
func ExamplePHASEDucker_Deactivate() {
	obj := phase.NewPHASEDucker()
	obj.Deactivate()
	// Output:
	}

