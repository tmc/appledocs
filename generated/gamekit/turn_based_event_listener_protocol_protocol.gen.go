// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

// PTurnBasedEventListener is the GKTurnBasedEventListener protocol interface.
//
// The protocol that handles turn-based and data-exchange events between participants in a match.
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
// See: doc://com.apple.gamekit/documentation/GameKit/GKTurnBasedEventListener
type PTurnBasedEventListener interface {
	// Optional methods
	PlayerDidRequestMatchWithOtherPlayers(player IGKPlayer, playersToInvite []Player)
	HasPlayerDidRequestMatchWithOtherPlayers() bool
	PlayerDidRequestMatchWithPlayers(player IGKPlayer, playerIDsToInvite []string)
	HasPlayerDidRequestMatchWithPlayers() bool
	PlayerMatchEnded(player IGKPlayer, match IGKTurnBasedMatch)
	HasPlayerMatchEnded() bool
	PlayerReceivedExchangeCancellationForMatch(player IGKPlayer, exchange IGKTurnBasedExchange, match IGKTurnBasedMatch)
	HasPlayerReceivedExchangeCancellationForMatch() bool
	PlayerReceivedExchangeRepliesForCompletedExchangeForMatch(player IGKPlayer, replies []TurnBasedExchangeReply, exchange IGKTurnBasedExchange, match IGKTurnBasedMatch)
	HasPlayerReceivedExchangeRepliesForCompletedExchangeForMatch() bool
	PlayerReceivedExchangeRequestForMatch(player IGKPlayer, exchange IGKTurnBasedExchange, match IGKTurnBasedMatch)
	HasPlayerReceivedExchangeRequestForMatch() bool
	PlayerReceivedTurnEventForMatchDidBecomeActive(player IGKPlayer, match IGKTurnBasedMatch, didBecomeActive bool)
	HasPlayerReceivedTurnEventForMatchDidBecomeActive() bool
	PlayerWantsToQuitMatch(player IGKPlayer, match IGKTurnBasedMatch)
	HasPlayerWantsToQuitMatch() bool
}
