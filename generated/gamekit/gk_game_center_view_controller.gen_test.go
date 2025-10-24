// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit_test

import (
	"github.com/tmc/appledocs/generated/gamekit"
)

// Suppress unused import errors
var _ = gamekit.NewGameCenterViewController

// ExampleNewGameCenterViewControllerWithState demonstrates how to create a GameCenterViewController instance using NewGameCenterViewControllerWithState.
// Creates a view controller that presents the specified Game Center content.
func ExampleNewGameCenterViewControllerWithState() {
	_ = gamekit.NewGameCenterViewControllerWithState(
		gamekit.GameCenterViewControllerState /* not a class type */{}, // state GameCenterViewControllerState /* not a class type */
	)
	// Output:
}
