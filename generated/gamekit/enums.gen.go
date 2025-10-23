// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

// Enum types and constants
// GKErrorCode - Error codes for the GameKit error domain.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code
type GKErrorCode uint

const (
	// GKErrorChallengeInvalid - The challenge request failed due to invalid challenge data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/challengeInvalid
	GKErrorChallengeInvalid GKErrorCode = 0
	// GKErrorNotAuthenticated - The system can’t complete the requested operation because the system hasn’t authorized the player.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/notAuthenticated
	GKErrorNotAuthenticated GKErrorCode = 0
	// GKErrorPlayerPhotoFailure - The system can’t complete the requested operation to retrieve a player’s photo.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/playerPhotoFailure
	GKErrorPlayerPhotoFailure GKErrorCode = 0
	// GKErrorTurnBasedInvalidTurn - The system can’t complete the requested operation because the participant doesn’t have the required turn state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/turnBasedInvalidTurn
	GKErrorTurnBasedInvalidTurn GKErrorCode = 0
	// GKErrorUnderage - The system can’t complete the requested operation because this feature isn’t available to underage players.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/underage
	GKErrorUnderage GKErrorCode = 0
	// GKErrorUserDenied - The system can’t complete the requested operation because the user denied it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/userDenied
	GKErrorUserDenied GKErrorCode = 0
)

// GKFriendsAuthorizationStatus - Constants that indicate if the local player grants access to their friends list.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKFriendsAuthorizationStatus
type GKFriendsAuthorizationStatus uint

// GKLeaderboardTimeScope - Specifies the time period for filtering data.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/TimeScope-swift.enum
type GKLeaderboardTimeScope uint

const (
	// GKLeaderboardTimeScopeAllTime - Loads a player’s best score.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/TimeScope-swift.enum/allTime
	GKLeaderboardTimeScopeAllTime GKLeaderboardTimeScope = 0
	// GKLeaderboardTimeScopeToday - Loads data for the past 24 hours.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/TimeScope-swift.enum/today
	GKLeaderboardTimeScopeToday GKLeaderboardTimeScope = 0
)

// GKTurnBasedExchangeStatus - The status of an exchange or reply.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchangeStatus
type GKTurnBasedExchangeStatus uint

const (
	// GKTurnBasedExchangeStatusComplete - All recipients of the exchange request replied.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchangeStatus/complete
	GKTurnBasedExchangeStatusComplete GKTurnBasedExchangeStatus = 0
	// GKTurnBasedExchangeStatusResolved - The current participant saved the exchange request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchangeStatus/resolved
	GKTurnBasedExchangeStatusResolved GKTurnBasedExchangeStatus = 0
	// GKTurnBasedExchangeStatusUnknown - The state of the exchange request is unknown.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchangeStatus/unknown
	GKTurnBasedExchangeStatusUnknown GKTurnBasedExchangeStatus = 0
)


