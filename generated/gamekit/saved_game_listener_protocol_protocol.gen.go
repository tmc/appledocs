// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

// PSavedGameListener is the GKSavedGameListener protocol interface.
//
// A protocol that handles events related to saving game data.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - visionOS 1.0+
//
// See: doc://com.apple.gamekit/documentation/GameKit/GKSavedGameListener
type PSavedGameListener interface {
	// Optional methods
	PlayerDidModifySavedGame(player IGKPlayer, savedGame IGKSavedGame)
	HasPlayerDidModifySavedGame() bool
	PlayerHasConflictingSavedGames(player IGKPlayer, savedGames []SavedGame)
	HasPlayerHasConflictingSavedGames() bool
}
