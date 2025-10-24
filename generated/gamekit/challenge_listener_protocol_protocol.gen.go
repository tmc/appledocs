// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

// PChallengeListener is the GKChallengeListener protocol interface.
//
// An object that responds to challenge events.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 26.0)
//   - iOS 7.0+ (Deprecated in 26.0)
//   - iPadOS 7.0+ (Deprecated in 26.0)
//   - macOS 10.10+ (Deprecated in 26.0)
//   - tvOS 9.0+ (Deprecated in 26.0)
//   - visionOS 1.0+ (Deprecated in 26.0)
//   - watchOS 3.0+ (Deprecated in 26.0)
//
// See: doc://com.apple.gamekit/documentation/GameKit/GKChallengeListener
type PChallengeListener interface {
	// Optional methods
	PlayerDidCompleteChallengeIssuedByFriend(player IGKPlayer, challenge IGKChallenge, friendPlayer IGKPlayer)
	HasPlayerDidCompleteChallengeIssuedByFriend() bool
	PlayerDidReceiveChallenge(player IGKPlayer, challenge IGKChallenge)
	HasPlayerDidReceiveChallenge() bool
	PlayerIssuedChallengeWasCompletedByFriend(player IGKPlayer, challenge IGKChallenge, friendPlayer IGKPlayer)
	HasPlayerIssuedChallengeWasCompletedByFriend() bool
	PlayerWantsToPlayChallenge(player IGKPlayer, challenge IGKChallenge)
	HasPlayerWantsToPlayChallenge() bool
}
