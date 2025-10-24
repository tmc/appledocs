// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit_test

import (
	"github.com/tmc/appledocs/generated/gamekit"
)

// Suppress unused import errors
var _ = gamekit.NewLeaderboard

// ExampleNewLeaderboard demonstrates how to create a Leaderboard instance.
// Initializes a default leaderboard request.
func ExampleNewLeaderboard() {
	_ = gamekit.NewLeaderboard()
	// Output:
}
// ExampleNewLeaderboardWithPlayerIDs demonstrates how to create a Leaderboard instance using NewLeaderboardWithPlayerIDs.
// Initializes a leaderboard request to retrieve the scores of a specific group of players.
func ExampleNewLeaderboardWithPlayerIDs() {
	_ = gamekit.NewLeaderboardWithPlayerIDs(
		[]gamekit.string{}, // playerIDs []string
	)
	// Output:
}
// ExampleNewLeaderboardWithPlayers demonstrates how to create a Leaderboard instance using NewLeaderboardWithPlayers.
// Initializes a leaderboard request to retrieve the scores of a specific group of players.
func ExampleNewLeaderboardWithPlayers() {
	_ = gamekit.NewLeaderboardWithPlayers(
		[]gamekit.Player{}, // players []Player
	)
	// Output:
}
