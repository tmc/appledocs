// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit_test

import (
	"github.com/tmc/appledocs/generated/gamekit"
)

// Suppress unused import errors
var _ = gamekit.NewTurnBasedMatchmakerViewController

// ExampleNewTurnBasedMatchmakerViewControllerWithMatchRequest demonstrates how to create a TurnBasedMatchmakerViewController instance using NewTurnBasedMatchmakerViewControllerWithMatchRequest.
// Creates a matchmaker view controller for the local player to start inviting other players to a turn-based game.
func ExampleNewTurnBasedMatchmakerViewControllerWithMatchRequest() {
	_ = gamekit.NewTurnBasedMatchmakerViewControllerWithMatchRequest(
		gamekit.GKMatchRequest{}, // request GKMatchRequest
	)
	// Output:
}
