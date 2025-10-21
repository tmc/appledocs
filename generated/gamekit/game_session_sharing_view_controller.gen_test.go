// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit_test

import (
	"github.com/tmc/appledocs/generated/gamekit"
)

// Suppress unused import errors
var _ = gamekit.NewGameSessionSharingViewController

// ExampleNewGameSessionSharingViewControllerWithSession demonstrates how to create a GameSessionSharingViewController instance using NewGameSessionSharingViewControllerWithSession.
// Creates a new sharing view controller for a specified session.
func ExampleNewGameSessionSharingViewControllerWithSession() {
	_ = gamekit.NewGameSessionSharingViewControllerWithSession(
		gamekit.GKGameSession{}, // session GKGameSession
	)
	// Output:
}
