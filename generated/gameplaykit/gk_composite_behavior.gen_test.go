// Code generated from Apple documentation for GameplayKit. DO NOT EDIT.

package gameplaykit_test

import (
	"github.com/tmc/appledocs/generated/gameplaykit"
)

// Suppress unused import errors
var _ = gameplaykit.NewCompositeBehavior

// ExampleNewCompositeBehaviorWithBehaviors demonstrates how to create a CompositeBehavior instance using NewCompositeBehaviorWithBehaviors.
// Creates a composite behavior from the specified individual behaviors.
func ExampleNewCompositeBehaviorWithBehaviors() {
	_ = gameplaykit.NewCompositeBehaviorWithBehaviors(
		[]gameplaykit.Behavior{}, // behaviors []Behavior
	)
	// Output:
}
// ExampleCompositeBehavior_RemoveAllBehaviors demonstrates using RemoveAllBehaviors on a CompositeBehavior instance.
// Removes all individual behaviors from the composite behavior.
func ExampleCompositeBehavior_RemoveAllBehaviors() {
	obj := gameplaykit.NewCompositeBehavior()
	obj.RemoveAllBehaviors()
	// Output:
	}

