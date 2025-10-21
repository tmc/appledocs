// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewStateMachine

// ExampleNewStateMachineWithStates demonstrates how to create a StateMachine instance using NewStateMachineWithStates.
// Initializes a state machine with the specified states.
func ExampleNewStateMachineWithStates() {
	_ = gameplaykit.NewStateMachineWithStates(
		[]gameplaykit.State{}, // states []State
	)
	// Output:
}
