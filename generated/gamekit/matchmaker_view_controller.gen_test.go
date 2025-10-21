// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit_test

import (
	"github.com/tmc/appledocs/generated/gamekit"
)

// Suppress unused import errors
var _ = gamekit.NewMatchmakerViewController

// ExampleNewMatchmakerViewControllerWithMatchRequest demonstrates how to create a MatchmakerViewController instance using NewMatchmakerViewControllerWithMatchRequest.
// Creates a matchmaker view controller for the local player to start inviting other players.
func ExampleNewMatchmakerViewControllerWithMatchRequest() {
	_ = gamekit.NewMatchmakerViewControllerWithMatchRequest(
		gamekit.GKMatchRequest{}, // request GKMatchRequest
	)
	// Output:
}
