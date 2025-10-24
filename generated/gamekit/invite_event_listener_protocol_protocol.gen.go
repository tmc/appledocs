// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

// PInviteEventListener is the GKInviteEventListener protocol interface.
//
// A protocol that handles invite events from Game Center.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 3.0+
//
// See: doc://com.apple.gamekit/documentation/GameKit/GKInviteEventListener
type PInviteEventListener interface {
	// Optional methods
	PlayerDidAcceptInvite(player IGKPlayer, invite IGKInvite)
	HasPlayerDidAcceptInvite() bool
	PlayerDidRequestMatchWithPlayers(player IGKPlayer, playerIDsToInvite []string)
	HasPlayerDidRequestMatchWithPlayers() bool
	PlayerDidRequestMatchWithRecipients(player IGKPlayer, recipientPlayers []Player)
	HasPlayerDidRequestMatchWithRecipients() bool
}
