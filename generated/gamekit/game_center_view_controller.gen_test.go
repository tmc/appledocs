// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit_test

import (
	"github.com/tmc/appledocs/generated/gamekit"
)

// Suppress unused import errors
var _ = gamekit.NewGameCenterViewController

// ExampleNewGameCenterViewControllerWithPlayer demonstrates how to create a GameCenterViewController instance using NewGameCenterViewControllerWithPlayer.
// Creates a view controller that presents a player’s Game Center profile.
func ExampleNewGameCenterViewControllerWithPlayer() {
	_ = gamekit.NewGameCenterViewControllerWithPlayer(
		gamekit.GKPlayer{}, // player GKPlayer
	)
	// Output:
}
