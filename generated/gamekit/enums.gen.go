// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

// Enum types and constants
// GKErrorCode - Error codes for the GameKit error domain.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code
type ErrorCode uint

const (
// ErrorInvalidCredentials - The system can’t complete the requested operation because the user name or password are incorrect.
//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/invalidCredentials
ErrorInvalidCredentials ErrorCode = 0
// ErrorLockdownMode - The system can’t complete the requested operation because the player enabled Lockdown Mode on the device.
//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/lockdownMode
ErrorLockdownMode ErrorCode = 0
// ErrorNotAuthenticated - The system can’t complete the requested operation because the system hasn’t authorized the player.
//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/notAuthenticated
ErrorNotAuthenticated ErrorCode = 0
// ErrorNotAuthorized - The system can’t complete the requested operation because the system hasn’t authorized the player.
//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/notAuthorized
ErrorNotAuthorized ErrorCode = 0
// ErrorParentalControlsBlocked - The system can’t complete the requested operation because the user disabled this feature in Restrictions.
//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/parentalControlsBlocked
ErrorParentalControlsBlocked ErrorCode = 0
// ErrorPlayerPhotoFailure - The system can’t complete the requested operation to retrieve a player’s photo.
//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/playerPhotoFailure
ErrorPlayerPhotoFailure ErrorCode = 0
// ErrorPlayerStatusExceedsMaximumLength - The player’s status exceeds the maximum length.
//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/playerStatusExceedsMaximumLength
ErrorPlayerStatusExceedsMaximumLength ErrorCode = 0
// ErrorPlayerStatusInvalid - The player’s status is invalid.
//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/playerStatusInvalid
ErrorPlayerStatusInvalid ErrorCode = 0
// ErrorTurnBasedInvalidTurn - The system can’t complete the requested operation because the participant doesn’t have the required turn state.
//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/turnBasedInvalidTurn
ErrorTurnBasedInvalidTurn ErrorCode = 0
// ErrorUbiquityContainerUnavailable - The system can’t complete the requested operation because the user hasn’t signed in to iCloud or hasn’t enabled iCloud Drive.
//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/ubiquityContainerUnavailable
ErrorUbiquityContainerUnavailable ErrorCode = 0
// ErrorUnderage - The system can’t complete the requested operation because this feature isn’t available to underage players.
//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/underage
ErrorUnderage ErrorCode = 0
// ErrorUserDenied - The system can’t complete the requested operation because the user denied it.
//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/userDenied
ErrorUserDenied ErrorCode = 0
)

// GKFriendsAuthorizationStatus - Constants that indicate if the local player grants access to their friends list.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKFriendsAuthorizationStatus
type FriendsAuthorizationStatus uint

// GKMatchType - The kind of match managed by Game Center.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchType
type MatchType uint

const (
// MatchTypeHosted - A match hosted on your private server.
//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchType/hosted
MatchTypeHosted MatchType = 0
// MatchTypePeerToPeer - A peer-to-peer match hosted by Game Center.
//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchType/peerToPeer
MatchTypePeerToPeer MatchType = 0
// MatchTypeTurnBased - A turn-based match hosted by Game Center.
//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchType/turnBased
MatchTypeTurnBased MatchType = 0
)

// GKPeerPickerConnectionType - Network connections available to the peer picker dialog.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPeerPickerConnectionType
type PeerPickerConnectionType uint

// GKTransportType - The mechanism used to send messages to other players in a game session.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTransportType
type TransportType uint

// GKTurnBasedExchangeStatus - The status of an exchange or reply.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchangeStatus
type TurnBasedExchangeStatus uint

const (
// TurnBasedExchangeStatusComplete - All recipients of the exchange request replied.
//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchangeStatus/complete
TurnBasedExchangeStatusComplete TurnBasedExchangeStatus = 0
)


