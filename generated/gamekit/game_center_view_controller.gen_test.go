// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit_test

import (
	"github.com/tmc/appledocs/generated/gamekit"
)

// Suppress unused import errors
var _ = gamekit.NewGameCenterViewController

// ExampleNewGameCenterViewControllerWithAchievementID demonstrates how to create a GameCenterViewController instance using NewGameCenterViewControllerWithAchievementID.
// Creates a view controller that presents an achievement.
func ExampleNewGameCenterViewControllerWithAchievementID() {
	_ = gamekit.NewGameCenterViewControllerWithAchievementID(
		"achievementID", // achievementID string
	)
	// Output:
}

// ExampleNewGameCenterViewControllerWithLeaderboardSetID demonstrates how to create a GameCenterViewController instance using NewGameCenterViewControllerWithLeaderboardSetID.
// Creates a view controller that presents a leaderboard set.
func ExampleNewGameCenterViewControllerWithLeaderboardSetID() {
	_ = gamekit.NewGameCenterViewControllerWithLeaderboardSetID(
		"leaderboardSetID", // leaderboardSetID string
	)
	// Output:
}

// ExampleNewGameCenterViewControllerWithPlayer demonstrates how to create a GameCenterViewController instance using NewGameCenterViewControllerWithPlayer.
// Creates a view controller that presents a player’s Game Center profile.
func ExampleNewGameCenterViewControllerWithPlayer() {
	_ = gamekit.NewGameCenterViewControllerWithPlayer(
		gamekit.GKPlayer{}, // player GKPlayer
	)
	// Output:
}

// ExampleNewGameCenterViewControllerWithState demonstrates how to create a GameCenterViewController instance using NewGameCenterViewControllerWithState.
// Creates a view controller that presents the specified Game Center content.
func ExampleNewGameCenterViewControllerWithState() {
	_ = gamekit.NewGameCenterViewControllerWithState(
		gamekit.GameCenterViewControllerState{}, // state GameCenterViewControllerState
	)
	// Output:
}
