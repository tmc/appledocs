// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

/* debug [enums.gen.go]: Generating 31 enums for GameKit */
// Enum types and constants
/* debug [enums.gen.go]: Processing enum GKAccessPointLocation (4 cases) */
// GKAccessPointLocation - Specifies the corner of the screen to display the access point.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/Location-swift.enum
type GKAccessPointLocation uint

const (
	// GKAccessPointLocationBottomLeading - The lower-left corner of the screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/Location-swift.enum/bottomLeading
	GKAccessPointLocationBottomLeading GKAccessPointLocation = 0
	// GKAccessPointLocationBottomTrailing - The lower-right corner of the screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/Location-swift.enum/bottomTrailing
	GKAccessPointLocationBottomTrailing GKAccessPointLocation = 0
	// GKAccessPointLocationTopLeading - The upper-left corner of the screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/Location-swift.enum/topLeading
	GKAccessPointLocationTopLeading GKAccessPointLocation = 0
	// GKAccessPointLocationTopTrailing - The upper-right corner of the screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAccessPoint/Location-swift.enum/topTrailing
	GKAccessPointLocationTopTrailing GKAccessPointLocation = 0
)

/* debug [enums.gen.go]: Processing enum GKErrorCode (42 cases) */
// GKErrorCode - Error codes for the GameKit error domain.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code
type GKErrorCode uint

const (
	// GKErrorAPINotAvailable - The system can’t complete the requested operation because the API isn’t available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/apiNotAvailable
	GKErrorAPINotAvailable GKErrorCode = 0
	// GKErrorAPIObsolete - The system can’t complete the requested operation because Apple deprecated the API.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/apiObsolete
	GKErrorAPIObsolete GKErrorCode = 0
	// GKErrorAppUnlisted - The system can’t complete the requested operation because the game isn’t available on the App Store.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/appUnlisted
	GKErrorAppUnlisted GKErrorCode = 0
	// GKErrorAuthenticationInProgress - The system can’t complete the requested operation because the local player is already authenticating.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/authenticationInProgress
	GKErrorAuthenticationInProgress GKErrorCode = 0
	// GKErrorCancelled - The system canceled the requested operation or the user disabled it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/cancelled
	GKErrorCancelled GKErrorCode = 0
	// GKErrorChallengeInvalid - The challenge request failed due to invalid challenge data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/challengeInvalid
	GKErrorChallengeInvalid GKErrorCode = 0
	// GKErrorCommunicationsFailure - The system can’t complete the requested operation due to an error communicating with the server.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/communicationsFailure
	GKErrorCommunicationsFailure GKErrorCode = 0
	// GKErrorConnectionTimeout - The system can’t complete the requested operation because the connection timed out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/connectionTimeout
	GKErrorConnectionTimeout GKErrorCode = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/debugMode
	GKErrorDebugMode GKErrorCode = 0
	// GKErrorFriendListDenied - Access to the local player’s list of friends denied.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/friendListDenied
	GKErrorFriendListDenied GKErrorCode = 0
	// GKErrorFriendListDescriptionMissing - Access to the local player’s list of friends denied for lack of a reason.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/friendListDescriptionMissing
	GKErrorFriendListDescriptionMissing GKErrorCode = 0
	// GKErrorFriendListRestricted - Access to the local player’s list of friends restricted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/friendListRestricted
	GKErrorFriendListRestricted GKErrorCode = 0
	// GKErrorFriendRequestNotAvailable - The player can’t send a friend request at this time from this device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/friendRequestNotAvailable
	GKErrorFriendRequestNotAvailable GKErrorCode = 0
	// GKErrorGameSessionRequestInvalid - The properties of the game session request are impossible to fulfill.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/gameSessionRequestInvalid
	GKErrorGameSessionRequestInvalid GKErrorCode = 0
	// GKErrorGameUnrecognized - The system can’t complete the requested operation because Game Center doesn’t recognize the app.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/gameUnrecognized
	GKErrorGameUnrecognized GKErrorCode = 0
	// GKErrorICloudUnavailable - The system can’t complete the requested operation because it can’t access the player’s iCloud account.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/iCloudUnavailable
	GKErrorICloudUnavailable GKErrorCode = 0
	// GKErrorInvalidCredentials - The system can’t complete the requested operation because the user name or password are incorrect.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/invalidCredentials
	GKErrorInvalidCredentials GKErrorCode = 0
	// GKErrorInvalidParameter - The system can’t complete the requested operation because one or more parameters are invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/invalidParameter
	GKErrorInvalidParameter GKErrorCode = 0
	// GKErrorInvalidPlayer - The system can’t complete the requested operation because the player is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/invalidPlayer
	GKErrorInvalidPlayer GKErrorCode = 0
	// GKErrorInvitationsDisabled - The system can’t complete the requested operation because the receiving player has disabled invitations.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/invitationsDisabled
	GKErrorInvitationsDisabled GKErrorCode = 0
	// GKErrorLockdownMode - The system can’t complete the requested operation because the player enabled Lockdown Mode on the device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/lockdownMode
	GKErrorLockdownMode GKErrorCode = 0
	// GKErrorMatchNotConnected - The system can’t complete the requested operation because the match isn’t connected to other players.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/matchNotConnected
	GKErrorMatchNotConnected GKErrorCode = 0
	// GKErrorMatchRequestInvalid - The system can’t complete the requested operation because the match request is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/matchRequestInvalid
	GKErrorMatchRequestInvalid GKErrorCode = 0
	// GKErrorNotAuthenticated - The system can’t complete the requested operation because the system hasn’t authorized the player.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/notAuthenticated
	GKErrorNotAuthenticated GKErrorCode = 0
	// GKErrorNotAuthorized - The system can’t complete the requested operation because the system hasn’t authorized the player.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/notAuthorized
	GKErrorNotAuthorized GKErrorCode = 0
	// GKErrorNotSupported - The app doesn’t have Game Center enabled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/notSupported
	GKErrorNotSupported GKErrorCode = 0
	// GKErrorParentalControlsBlocked - The system can’t complete the requested operation because the user disabled this feature in Restrictions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/parentalControlsBlocked
	GKErrorParentalControlsBlocked GKErrorCode = 0
	// GKErrorPlayerPhotoFailure - The system can’t complete the requested operation to retrieve a player’s photo.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/playerPhotoFailure
	GKErrorPlayerPhotoFailure GKErrorCode = 0
	// GKErrorPlayerStatusExceedsMaximumLength - The player’s status exceeds the maximum length.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/playerStatusExceedsMaximumLength
	GKErrorPlayerStatusExceedsMaximumLength GKErrorCode = 0
	// GKErrorPlayerStatusInvalid - The player’s status is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/playerStatusInvalid
	GKErrorPlayerStatusInvalid GKErrorCode = 0
	// GKErrorRestrictedToAutomatch - The system can’t complete the requested operation because the player is using automatch.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/restrictedToAutomatch
	GKErrorRestrictedToAutomatch GKErrorCode = 0
	// GKErrorScoreNotSet - The system can’t complete the requested operation because the system hasn’t set the score.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/scoreNotSet
	GKErrorScoreNotSet GKErrorCode = 0
	// GKErrorTurnBasedInvalidParticipant - The system can’t complete the requested operation because the specified participant is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/turnBasedInvalidParticipant
	GKErrorTurnBasedInvalidParticipant GKErrorCode = 0
	// GKErrorTurnBasedInvalidState - The system can’t complete the requested operation because the session is in an invalid state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/turnBasedInvalidState
	GKErrorTurnBasedInvalidState GKErrorCode = 0
	// GKErrorTurnBasedInvalidTurn - The system can’t complete the requested operation because the participant doesn’t have the required turn state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/turnBasedInvalidTurn
	GKErrorTurnBasedInvalidTurn GKErrorCode = 0
	// GKErrorTurnBasedMatchDataTooLarge - The system can’t complete the requested operation because the match data is too large.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/turnBasedMatchDataTooLarge
	GKErrorTurnBasedMatchDataTooLarge GKErrorCode = 0
	// GKErrorTurnBasedTooManySessions - The system can’t complete the requested operation because it exceeds the maximum number of sessions.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/turnBasedTooManySessions
	GKErrorTurnBasedTooManySessions GKErrorCode = 0
	// GKErrorUbiquityContainerUnavailable - The system can’t complete the requested operation because the user hasn’t signed in to iCloud or hasn’t enabled iCloud Drive.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/ubiquityContainerUnavailable
	GKErrorUbiquityContainerUnavailable GKErrorCode = 0
	// GKErrorUnderage - The system can’t complete the requested operation because this feature isn’t available to underage players.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/underage
	GKErrorUnderage GKErrorCode = 0
	// GKErrorUnexpectedConnection - An unexpected player has connected to a match.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/unexpectedConnection
	GKErrorUnexpectedConnection GKErrorCode = 0
	// GKErrorUnknown - The system can’t complete the requested operation due to an unknown error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/unknown
	GKErrorUnknown GKErrorCode = 0
	// GKErrorUserDenied - The system can’t complete the requested operation because the user denied it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKError/Code/userDenied
	GKErrorUserDenied GKErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum GKGameActivityState (4 cases) */
// GKGameActivityState - The state of a game activity.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/State-swift.enum
type GKGameActivityState uint

const (
	// GKGameActivityStateActive - The game activity is active.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/State-swift.enum/active
	GKGameActivityStateActive GKGameActivityState = 0
	// GKGameActivityStateEnded - The game activity has ended. This is a terminal state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/State-swift.enum/ended
	GKGameActivityStateEnded GKGameActivityState = 0
	// GKGameActivityStateInitialized - The game activity is initialized but has not started.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/State-swift.enum/initialized
	GKGameActivityStateInitialized GKGameActivityState = 0
	// GKGameActivityStatePaused - The game activity is paused.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivity/State-swift.enum/paused
	GKGameActivityStatePaused GKGameActivityState = 0
)

/* debug [enums.gen.go]: Processing enum GKLeaderboardType (2 cases) */
// GKLeaderboardType - Specifies whether a leaderboard is recurring.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/LeaderboardType
type GKLeaderboardType uint

const (
	// GKLeaderboardTypeClassic - A leaderboard that never expires, showing all-time rankings of all players.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/LeaderboardType/classic
	GKLeaderboardTypeClassic GKLeaderboardType = 0
	// GKLeaderboardTypeRecurring - A leaderboard that recurs, allowing players a fresh start to compete and earn higher ranks in each ocurrence.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/LeaderboardType/recurring
	GKLeaderboardTypeRecurring GKLeaderboardType = 0
)

/* debug [enums.gen.go]: Processing enum GKLeaderboardPlayerScope (2 cases) */
// GKLeaderboardPlayerScope - Specifies the type of players for filtering data.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/PlayerScope-swift.enum
type GKLeaderboardPlayerScope uint

const (
	// GKLeaderboardPlayerScopeFriendsOnly - Loads only data for friends of the local player.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/PlayerScope-swift.enum/friendsOnly
	GKLeaderboardPlayerScopeFriendsOnly GKLeaderboardPlayerScope = 0
	// GKLeaderboardPlayerScopeGlobal - Loads data for all players of the game.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/PlayerScope-swift.enum/global
	GKLeaderboardPlayerScopeGlobal GKLeaderboardPlayerScope = 0
)

/* debug [enums.gen.go]: Processing enum GKLeaderboardTimeScope (3 cases) */
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
	// GKLeaderboardTimeScopeWeek - Loads data for the past week.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboard/TimeScope-swift.enum/week
	GKLeaderboardTimeScopeWeek GKLeaderboardTimeScope = 0
)

/* debug [enums.gen.go]: Processing enum GKMatchSendDataMode (2 cases) */
// GKMatchSendDataMode - The mechanism used to transmit data to other players.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatch/SendDataMode
type GKMatchSendDataMode uint

const (
	// GKMatchSendDataReliable - Sends data continuously until the recipients successfully receive it or the connection times out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatch/SendDataMode/reliable
	GKMatchSendDataReliable GKMatchSendDataMode = 0
	// GKMatchSendDataUnreliable - Sends data once even if an error occurs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatch/SendDataMode/unreliable
	GKMatchSendDataUnreliable GKMatchSendDataMode = 0
)

/* debug [enums.gen.go]: Processing enum GKPhotoSize (2 cases) */
// GKPhotoSize - The size of a photo that Game Center loads.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPlayer/PhotoSize
type GKPhotoSize uint

const (
	// GKPhotoSizeNormal - Loads a normal-sized photo.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPlayer/PhotoSize/normal
	GKPhotoSizeNormal GKPhotoSize = 0
	// GKPhotoSizeSmall - Loads a small photo.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPlayer/PhotoSize/small
	GKPhotoSizeSmall GKPhotoSize = 0
)

/* debug [enums.gen.go]: Processing enum GKTurnBasedMatchOutcome (11 cases) */
// GKTurnBasedMatchOutcome - The state of a participant when they forfeit a match or when a match ends.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/Outcome
type GKTurnBasedMatchOutcome uint

const (
	// GKTurnBasedMatchOutcomeCustomRange - The participant reaches a game-specific outcome.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/Outcome/customRange
	GKTurnBasedMatchOutcomeCustomRange GKTurnBasedMatchOutcome = 0
	// GKTurnBasedMatchOutcomeFirst - The participant finishes in first place.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/Outcome/first
	GKTurnBasedMatchOutcomeFirst GKTurnBasedMatchOutcome = 0
	// GKTurnBasedMatchOutcomeFourth - The participant finishes in fourth place.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/Outcome/fourth
	GKTurnBasedMatchOutcomeFourth GKTurnBasedMatchOutcome = 0
	// GKTurnBasedMatchOutcomeLost - The participant loses the match.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/Outcome/lost
	GKTurnBasedMatchOutcomeLost GKTurnBasedMatchOutcome = 0
	// GKTurnBasedMatchOutcomeNone - The participant doesn’t reach an outcome.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/Outcome/none
	GKTurnBasedMatchOutcomeNone GKTurnBasedMatchOutcome = 0
	// GKTurnBasedMatchOutcomeQuit - The participant forfeits the match.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/Outcome/quit
	GKTurnBasedMatchOutcomeQuit GKTurnBasedMatchOutcome = 0
	// GKTurnBasedMatchOutcomeSecond - The participant finishes in second place.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/Outcome/second
	GKTurnBasedMatchOutcomeSecond GKTurnBasedMatchOutcome = 0
	// GKTurnBasedMatchOutcomeThird - The participant finishes in third place.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/Outcome/third
	GKTurnBasedMatchOutcomeThird GKTurnBasedMatchOutcome = 0
	// GKTurnBasedMatchOutcomeTied - The participant ties the match.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/Outcome/tied
	GKTurnBasedMatchOutcomeTied GKTurnBasedMatchOutcome = 0
	// GKTurnBasedMatchOutcomeTimeExpired - The match ends because the time limit expires.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/Outcome/timeExpired
	GKTurnBasedMatchOutcomeTimeExpired GKTurnBasedMatchOutcome = 0
	// GKTurnBasedMatchOutcomeWon - The participant wins the match.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/Outcome/won
	GKTurnBasedMatchOutcomeWon GKTurnBasedMatchOutcome = 0
)

/* debug [enums.gen.go]: Processing enum GKTurnBasedMatchStatus (4 cases) */
// GKTurnBasedMatchStatus - The states of a match from when it’s created to when it ends.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/Status-swift.enum
type GKTurnBasedMatchStatus uint

const (
	// GKTurnBasedMatchStatusEnded - A match that finishes.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/Status-swift.enum/ended
	GKTurnBasedMatchStatusEnded GKTurnBasedMatchStatus = 0
	// GKTurnBasedMatchStatusMatching - A match with empty slots that Game Center is actively filling.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/Status-swift.enum/matching
	GKTurnBasedMatchStatusMatching GKTurnBasedMatchStatus = 0
	// GKTurnBasedMatchStatusOpen - A match that participants are actively playing.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/Status-swift.enum/open
	GKTurnBasedMatchStatusOpen GKTurnBasedMatchStatus = 0
	// GKTurnBasedMatchStatusUnknown - A match that is in an unknown state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/Status-swift.enum/unknown
	GKTurnBasedMatchStatusUnknown GKTurnBasedMatchStatus = 0
)

/* debug [enums.gen.go]: Processing enum GKTurnBasedParticipantStatus (6 cases) */
// GKTurnBasedParticipantStatus - The state the participant is in during the match.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedParticipant/Status-swift.enum
type GKTurnBasedParticipantStatus uint

const (
	// GKTurnBasedParticipantStatusActive - The participant joins the match and is an active player.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedParticipant/Status-swift.enum/active
	GKTurnBasedParticipantStatusActive GKTurnBasedParticipantStatus = 0
	// GKTurnBasedParticipantStatusDeclined - The participant declines the invitation to join the match, automatically terminating the match.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedParticipant/Status-swift.enum/declined
	GKTurnBasedParticipantStatusDeclined GKTurnBasedParticipantStatus = 0
	// GKTurnBasedParticipantStatusDone - The participant leaves the match.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedParticipant/Status-swift.enum/done
	GKTurnBasedParticipantStatusDone GKTurnBasedParticipantStatus = 0
	// GKTurnBasedParticipantStatusInvited - The participant is invited to the match, but hasn’t responded to the invitation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedParticipant/Status-swift.enum/invited
	GKTurnBasedParticipantStatusInvited GKTurnBasedParticipantStatus = 0
	// GKTurnBasedParticipantStatusMatching - The participant represents an unfilled position in the match that Game Center promises to fill when needed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedParticipant/Status-swift.enum/matching
	GKTurnBasedParticipantStatusMatching GKTurnBasedParticipantStatus = 0
	// GKTurnBasedParticipantStatusUnknown - The participant is in an unexpected state.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedParticipant/Status-swift.enum/unknown
	GKTurnBasedParticipantStatusUnknown GKTurnBasedParticipantStatus = 0
)

/* debug [enums.gen.go]: Processing enum GKChallengeState (4 cases) */
// GKChallengeState - The state of a challenge.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeState
type GKChallengeState uint

const (
	// GKChallengeStateCompleted - The player successfully completed the challenge.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeState/completed
	GKChallengeStateCompleted GKChallengeState = 0
	// GKChallengeStateDeclined - The player declined the challenge.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeState/declined
	GKChallengeStateDeclined GKChallengeState = 0
	// GKChallengeStateInvalid - The challenge isn’t valid because an error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeState/invalid
	GKChallengeStateInvalid GKChallengeState = 0
	// GKChallengeStatePending - The player issued a challenge, but the other player hasn’t accepted or refused it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeState/pending
	GKChallengeStatePending GKChallengeState = 0
)

/* debug [enums.gen.go]: Processing enum GKConnectionState (2 cases) */
// GKConnectionState - Possible connection states for a player
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKConnectionState
type GKConnectionState uint

const (
	// GKConnectionStateConnected - The player is connected to the game session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKConnectionState/connected
	GKConnectionStateConnected GKConnectionState = 0
	// GKConnectionStateNotConnected - The player is not connected to the game session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKConnectionState/notConnected
	GKConnectionStateNotConnected GKConnectionState = 0
)

/* debug [enums.gen.go]: Processing enum GKFriendsAuthorizationStatus (4 cases) */
// GKFriendsAuthorizationStatus - Constants that indicate if the local player grants access to their friends list.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKFriendsAuthorizationStatus
type GKFriendsAuthorizationStatus uint

const (
	// GKFriendsAuthorizationStatusAuthorized - The player authorized your game to access their list of friends.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKFriendsAuthorizationStatus/authorized
	GKFriendsAuthorizationStatusAuthorized GKFriendsAuthorizationStatus = 0
	// GKFriendsAuthorizationStatusDenied - Access to the player’s friends’ data denied.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKFriendsAuthorizationStatus/denied
	GKFriendsAuthorizationStatusDenied GKFriendsAuthorizationStatus = 0
	// GKFriendsAuthorizationStatusNotDetermined - The player hasn’t choosen whether your game may access their friends list.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKFriendsAuthorizationStatus/notDetermined
	GKFriendsAuthorizationStatusNotDetermined GKFriendsAuthorizationStatus = 0
	// GKFriendsAuthorizationStatusRestricted - Access to the player’s list of friends restricted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKFriendsAuthorizationStatus/restricted
	GKFriendsAuthorizationStatusRestricted GKFriendsAuthorizationStatus = 0
)

/* debug [enums.gen.go]: Processing enum GKGameActivityPlayStyle (3 cases) */
// GKGameActivityPlayStyle - Play Style of the game activity. It can be either Asynchronous or Synchronous.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityPlayStyle
type GKGameActivityPlayStyle uint

const (
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityPlayStyle/asynchronous
	GKGameActivityPlayStyleAsynchronous GKGameActivityPlayStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityPlayStyle/synchronous
	GKGameActivityPlayStyleSynchronous GKGameActivityPlayStyle = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameActivityPlayStyle/unspecified
	GKGameActivityPlayStyleUnspecified GKGameActivityPlayStyle = 0
)

/* debug [enums.gen.go]: Processing enum GKGameCenterViewControllerState (7 cases) */
// GKGameCenterViewControllerState - The type of content for the view controller to present.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewControllerState
type GKGameCenterViewControllerState uint

const (
	// GKGameCenterViewControllerStateAchievements - The view controller should present a list of achievements.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewControllerState/achievements
	GKGameCenterViewControllerStateAchievements GKGameCenterViewControllerState = 0
	// GKGameCenterViewControllerStateChallenges - The view controller should present a list of challenges.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewControllerState/challenges
	GKGameCenterViewControllerStateChallenges GKGameCenterViewControllerState = 0
	// GKGameCenterViewControllerStateDashboard - The view controller should present the dashboard.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewControllerState/dashboard
	GKGameCenterViewControllerStateDashboard GKGameCenterViewControllerState = 0
	// GKGameCenterViewControllerStateDefault - The view controller should present the default screen.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewControllerState/default
	GKGameCenterViewControllerStateDefault GKGameCenterViewControllerState = 0
	// GKGameCenterViewControllerStateLeaderboards - The view controller should present leaderboard sets or leaderboards if there are no sets.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewControllerState/leaderboards
	GKGameCenterViewControllerStateLeaderboards GKGameCenterViewControllerState = 0
	// GKGameCenterViewControllerStateLocalPlayerFriendsList - The view controller should present the friends list.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewControllerState/localPlayerFriendsList
	GKGameCenterViewControllerStateLocalPlayerFriendsList GKGameCenterViewControllerState = 0
	// GKGameCenterViewControllerStateLocalPlayerProfile - The view controller should present the local player’s profile.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewControllerState/localPlayerProfile
	GKGameCenterViewControllerStateLocalPlayerProfile GKGameCenterViewControllerState = 0
)

/* debug [enums.gen.go]: Processing enum GKGameSessionErrorCode (16 cases) */
// GKGameSessionErrorCode - Error codes for the game session domain.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSessionError/Code
type GKGameSessionErrorCode uint

const (
	// GKGameSessionErrorBadContainer - The requested operation could not be completed because the iCloud container is invalid.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSessionError/Code/badContainer
	GKGameSessionErrorBadContainer GKGameSessionErrorCode = 0
	// GKGameSessionErrorCloudDriveDisabled - The requested operation could not be completed because iCloud Drive has been disabled for the application.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSessionError/Code/cloudDriveDisabled
	GKGameSessionErrorCloudDriveDisabled GKGameSessionErrorCode = 0
	// GKGameSessionErrorCloudQuotaExceeded - The requested operation could not be completed because the user’s iCloud quota would be exceeded.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSessionError/Code/cloudQuotaExceeded
	GKGameSessionErrorCloudQuotaExceeded GKGameSessionErrorCode = 0
	// GKGameSessionErrorConnectionCancelledByUser - The requested operation could not be completed because the connection to the session was cancelled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSessionError/Code/connectionCancelledByUser
	GKGameSessionErrorConnectionCancelledByUser GKGameSessionErrorCode = 0
	// GKGameSessionErrorConnectionFailed - The requested operation could not be completed because the session could not find other players to connect to.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSessionError/Code/connectionFailed
	GKGameSessionErrorConnectionFailed GKGameSessionErrorCode = 0
	// GKGameSessionErrorInvalidSession - The requested operation could not be completed because the Game Session does not exist or the player is not part of the game session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSessionError/Code/invalidSession
	GKGameSessionErrorInvalidSession GKGameSessionErrorCode = 0
	// GKGameSessionErrorNetworkFailure - The requested operation could not be completed due to an error communicating with the server.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSessionError/Code/networkFailure
	GKGameSessionErrorNetworkFailure GKGameSessionErrorCode = 0
	// GKGameSessionErrorNotAuthenticated - The requested operation could not be completed because you are not signed in to iCloud.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSessionError/Code/notAuthenticated
	GKGameSessionErrorNotAuthenticated GKGameSessionErrorCode = 0
	// GKGameSessionErrorSendDataNoRecipients - The requested operation could not be completed because there are no recipients connected to session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSessionError/Code/sendDataNoRecipients
	GKGameSessionErrorSendDataNoRecipients GKGameSessionErrorCode = 0
	// GKGameSessionErrorSendDataNotConnected - The requested operation could not be completed because you are not connected to the session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSessionError/Code/sendDataNotConnected
	GKGameSessionErrorSendDataNotConnected GKGameSessionErrorCode = 0
	// GKGameSessionErrorSendDataNotReachable - The requested operation could not be completed because one or more players is not reachable.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSessionError/Code/sendDataNotReachable
	GKGameSessionErrorSendDataNotReachable GKGameSessionErrorCode = 0
	// GKGameSessionErrorSendRateLimitReached - The requested operation could not be completed because you have reached the limits for save data request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSessionError/Code/sendRateLimitReached
	GKGameSessionErrorSendRateLimitReached GKGameSessionErrorCode = 0
	// GKGameSessionErrorSessionConflict - The requested operation could not be completed because the session has been updated on the server, causing a conflict.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSessionError/Code/sessionConflict
	GKGameSessionErrorSessionConflict GKGameSessionErrorCode = 0
	// GKGameSessionErrorSessionHasMaxConnectedPlayers - The requested operation could not be completed because the session has reached the maximum number of connected players.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSessionError/Code/sessionHasMaxConnectedPlayers
	GKGameSessionErrorSessionHasMaxConnectedPlayers GKGameSessionErrorCode = 0
	// GKGameSessionErrorSessionNotShared - The requested operation could not be completed because this session has not been shared with other players.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSessionError/Code/sessionNotShared
	GKGameSessionErrorSessionNotShared GKGameSessionErrorCode = 0
	// GKGameSessionErrorUnknown - The requested operation could not be completed due to an unknown error.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSessionError/Code/unknown
	GKGameSessionErrorUnknown GKGameSessionErrorCode = 0
)

/* debug [enums.gen.go]: Processing enum GKInviteRecipientResponse (12 cases) */
// GKInviteRecipientResponse - A player’s response to an invitation to join a match.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKInviteRecipientResponse
type GKInviteRecipientResponse uint

const (
	// GKInviteRecipientResponseAccepted - A response when the player accepts the invitation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKInviteRecipientResponse/accepted
	GKInviteRecipientResponseAccepted GKInviteRecipientResponse = 0
	// GKInviteRecipientResponseDeclined - A response when the player rejects the invitation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKInviteRecipientResponse/declined
	GKInviteRecipientResponseDeclined GKInviteRecipientResponse = 0
	// GKInviteRecipientResponseFailed - A response when the system fails to deliver the invitation to the player.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKInviteRecipientResponse/failed
	GKInviteRecipientResponseFailed GKInviteRecipientResponse = 0
	// GKInviteRecipientResponseIncompatible - A response when the player isn’t running a compatible version of the game.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKInviteRecipientResponse/incompatible
	GKInviteRecipientResponseIncompatible GKInviteRecipientResponse = 0
	// GKInviteeResponseAccepted - The player accepted the invitation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKInviteRecipientResponse/inviteeResponseAccepted
	GKInviteeResponseAccepted GKInviteRecipientResponse = 0
	// GKInviteeResponseDeclined - The player rejected the invitation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKInviteRecipientResponse/inviteeResponseDeclined
	GKInviteeResponseDeclined GKInviteRecipientResponse = 0
	// GKInviteeResponseFailed - The invitation was unable to be delivered.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKInviteRecipientResponse/inviteeResponseFailed
	GKInviteeResponseFailed GKInviteRecipientResponse = 0
	// GKInviteeResponseIncompatible - The invitee isn’t running a compatible version of your game.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKInviteRecipientResponse/inviteeResponseIncompatible
	GKInviteeResponseIncompatible GKInviteRecipientResponse = 0
	// GKInviteeResponseNoAnswer - The invitation timed out without an answer.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKInviteRecipientResponse/inviteeResponseNoAnswer
	GKInviteeResponseNoAnswer GKInviteRecipientResponse = 0
	// GKInviteeResponseUnableToConnect - The invitee couldn’t be contacted.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKInviteRecipientResponse/inviteeResponseUnableToConnect
	GKInviteeResponseUnableToConnect GKInviteRecipientResponse = 0
	// GKInviteRecipientResponseNoAnswer - A response when the invitation times out because the player doesn’t answer it.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKInviteRecipientResponse/noAnswer
	GKInviteRecipientResponseNoAnswer GKInviteRecipientResponse = 0
	// GKInviteRecipientResponseUnableToConnect - A response when the system can’t contact the player.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKInviteRecipientResponse/unableToConnect
	GKInviteRecipientResponseUnableToConnect GKInviteRecipientResponse = 0
)

/* debug [enums.gen.go]: Processing enum GKMatchmakingMode (4 cases) */
// GKMatchmakingMode - Possible modes that a multiplayer game uses to find matches.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakingMode
type GKMatchmakingMode uint

const (
	// GKMatchmakingModeAutomatchOnly - A mode that matches the local player only with players who are also actively looking for a match.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakingMode/automatchOnly
	GKMatchmakingModeAutomatchOnly GKMatchmakingMode = 0
	// GKMatchmakingModeDefault - The default matchmaking mode.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakingMode/default
	GKMatchmakingModeDefault GKMatchmakingMode = 0
	// GKMatchmakingModeInviteOnly - A mode that matches the local player only with players who they invite, and doesn’t use automatch to fill empty slots.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakingMode/inviteOnly
	GKMatchmakingModeInviteOnly GKMatchmakingMode = 0
	// GKMatchmakingModeNearbyOnly - A mode that matches the local player only with nearby players.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmakingMode/nearbyOnly
	GKMatchmakingModeNearbyOnly GKMatchmakingMode = 0
)

/* debug [enums.gen.go]: Processing enum GKMatchType (3 cases) */
// GKMatchType - The kind of match managed by Game Center.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchType
type GKMatchType uint

const (
	// GKMatchTypeHosted - A match hosted on your private server.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchType/hosted
	GKMatchTypeHosted GKMatchType = 0
	// GKMatchTypePeerToPeer - A peer-to-peer match hosted by Game Center.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchType/peerToPeer
	GKMatchTypePeerToPeer GKMatchType = 0
	// GKMatchTypeTurnBased - A turn-based match hosted by Game Center.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchType/turnBased
	GKMatchTypeTurnBased GKMatchType = 0
)

/* debug [enums.gen.go]: Processing enum GKPeerConnectionState (6 cases) */
// GKPeerConnectionState - The state of a peer known to the session.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPeerConnectionState
type GKPeerConnectionState uint

const (
	// GKPeerStateAvailable - A peer not connected to the session, but one that the session can connect to.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPeerConnectionState/stateAvailable
	GKPeerStateAvailable GKPeerConnectionState = 0
	// GKPeerStateConnected - A peer connected to the session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPeerConnectionState/stateConnected
	GKPeerStateConnected GKPeerConnectionState = 0
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPeerConnectionState/stateConnectedRelay
	GKPeerStateConnectedRelay GKPeerConnectionState = 0
	// GKPeerStateConnecting - A peer attempting to connect to the session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPeerConnectionState/stateConnecting
	GKPeerStateConnecting GKPeerConnectionState = 0
	// GKPeerStateDisconnected - A peer that disconnected from the session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPeerConnectionState/stateDisconnected
	GKPeerStateDisconnected GKPeerConnectionState = 0
	// GKPeerStateUnavailable - A peer that is no longer interested in receiving connections.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPeerConnectionState/stateUnavailable
	GKPeerStateUnavailable GKPeerConnectionState = 0
)

/* debug [enums.gen.go]: Processing enum GKPeerPickerConnectionType (2 cases) */
// GKPeerPickerConnectionType - Network connections available to the peer picker dialog.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPeerPickerConnectionType
type GKPeerPickerConnectionType uint

const (
	// GKPeerPickerConnectionTypeNearby - A Bluetooth connection to a device.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPeerPickerConnectionType/nearby
	GKPeerPickerConnectionTypeNearby GKPeerPickerConnectionType = 0
	// GKPeerPickerConnectionTypeOnline - An Internet-based connection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPeerPickerConnectionType/online
	GKPeerPickerConnectionTypeOnline GKPeerPickerConnectionType = 0
)

/* debug [enums.gen.go]: Processing enum GKPlayerConnectionState (3 cases) */
// GKPlayerConnectionState - The possible states of a connection to a match.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPlayerConnectionState
type GKPlayerConnectionState uint

const (
	// GKPlayerStateConnected - A state in which a player connects to the match and can receive data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPlayerConnectionState/connected
	GKPlayerStateConnected GKPlayerConnectionState = 0
	// GKPlayerStateDisconnected - A state in which a player disconnects from the match and can’t receive data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPlayerConnectionState/disconnected
	GKPlayerStateDisconnected GKPlayerConnectionState = 0
	// GKPlayerStateUnknown - An undetermined state in which the player can’t receive data.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPlayerConnectionState/unknown
	GKPlayerStateUnknown GKPlayerConnectionState = 0
)

/* debug [enums.gen.go]: Processing enum GKReleaseState (3 cases) */
// GKReleaseState - Describes the release state of an App Store Connect resource, such as an Achievement or Leaderboard.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKReleaseState
type GKReleaseState uint

const (
	// GKReleaseStateUnknown - The system can’t determine the release state of the resource.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKReleaseState/GKReleaseStateUnknown
	GKReleaseStateUnknown GKReleaseState = 0
	// GKReleaseStatePrereleased - The resource has been created in App Store Connect but isn’t yet associated with a released version of an App.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKReleaseState/prereleased
	GKReleaseStatePrereleased GKReleaseState = 0
	// GKReleaseStateReleased - The resource is associated with a release in App Store Connect. This has no relationship with the “archived” state of a resource (i.e., A resource can be release   archived).
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKReleaseState/released
	GKReleaseStateReleased GKReleaseState = 0
)

/* debug [enums.gen.go]: Processing enum GKSendDataMode (2 cases) */
// GKSendDataMode - The mechanism used to transmit data to other peers.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSendDataMode
type GKSendDataMode uint

const (
	// GKSendDataReliable - The data is sent continuously until it is successfully received by the intended recipients or the connection times out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSendDataMode/reliable
	GKSendDataReliable GKSendDataMode = 0
	// GKSendDataUnreliable - The data is sent once and is not sent again if a transmission error occurred.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSendDataMode/unreliable
	GKSendDataUnreliable GKSendDataMode = 0
)

/* debug [enums.gen.go]: Processing enum GKSessionError (16 cases) */
// GKSessionError - Error codes for the session error domain.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSessionError-swift.struct/Code
type GKSessionError uint

const (
	// GKSessionCancelledError - A peer that invited the session to connect to them canceled the connection request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSessionError-swift.struct/Code/cancelledError
	GKSessionCancelledError GKSessionError = 0
	// GKSessionCannotEnableError - Bluetooth is not currently available.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSessionError-swift.struct/Code/cannotEnableError
	GKSessionCannotEnableError GKSessionError = 0
	// GKSessionConnectionClosedError - The connection to another peer closed unexpectedly.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSessionError-swift.struct/Code/connectionClosedError
	GKSessionConnectionClosedError GKSessionError = 0
	// GKSessionConnectionFailedError - The attempt to establish a connection with another peer failed.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSessionError-swift.struct/Code/connectionFailedError
	GKSessionConnectionFailedError GKSessionError = 0
	// GKSessionConnectivityError - An error occurred in the   object’s connection code.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSessionError-swift.struct/Code/connectivityError
	GKSessionConnectivityError GKSessionError = 0
	// GKSessionDataTooBigError - The data your application attempted to send was too large for the session to transmit in a single call.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSessionError-swift.struct/Code/dataTooBigError
	GKSessionDataTooBigError GKSessionError = 0
	// GKSessionDeclinedError - The peer your application tried to connect to refused the connection.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSessionError-swift.struct/Code/declinedError
	GKSessionDeclinedError GKSessionError = 0
	// GKSessionInProgressError - The peer your application attempted to connect to has already requested a connection to your session.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSessionError-swift.struct/Code/inProgressError
	GKSessionInProgressError GKSessionError = 0
	// GKSessionInternalError - A serious error occurred inside  .
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSessionError-swift.struct/Code/internalError
	GKSessionInternalError GKSessionError = 0
	// GKSessionInvalidParameterError - A parameter had an unexpected value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSessionError-swift.struct/Code/invalidParameterError
	GKSessionInvalidParameterError GKSessionError = 0
	// GKSessionNotConnectedError - Reserved for future use.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSessionError-swift.struct/Code/notConnectedError
	GKSessionNotConnectedError GKSessionError = 0
	// GKSessionPeerNotFoundError - A peer with the specified   string could not be found.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSessionError-swift.struct/Code/peerNotFoundError
	GKSessionPeerNotFoundError GKSessionError = 0
	// GKSessionSystemError - An error occurred outside of the   object’s control, such as memory allocation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSessionError-swift.struct/Code/systemError
	GKSessionSystemError GKSessionError = 0
	// GKSessionTimedOutError - The operation could not be completed in the specified timeout period.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSessionError-swift.struct/Code/timedOutError
	GKSessionTimedOutError GKSessionError = 0
	// GKSessionTransportError - An error occurred in the   object’s transport code.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSessionError-swift.struct/Code/transportError
	GKSessionTransportError GKSessionError = 0
	// GKSessionUnknownError - Reserved for when the error does not fit in another category above.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSessionError-swift.struct/Code/unknownError
	GKSessionUnknownError GKSessionError = 0
)

/* debug [enums.gen.go]: Processing enum GKSessionMode (3 cases) */
// GKSessionMode - Modes that determine how a session interacts with other peers.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSessionMode
type GKSessionMode uint

const (
	// GKSessionModeClient - A client searches for servers advertising the same   property.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSessionMode/client
	GKSessionModeClient GKSessionMode = 0
	// GKSessionModePeer - A peer advertises like a server and searches like a client.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSessionMode/peer
	GKSessionModePeer GKSessionMode = 0
	// GKSessionModeServer - A server advertises itself to local devices using its   property.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKSessionMode/server
	GKSessionModeServer GKSessionMode = 0
)

/* debug [enums.gen.go]: Processing enum GKTransportType (2 cases) */
// GKTransportType - The mechanism used to send messages to other players in a game session.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTransportType
type GKTransportType uint

const (
	// GKTransportTypeReliable - The data is sent continuously until it is successfully received by the intended recipients or the connection times out.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTransportType/reliable
	GKTransportTypeReliable GKTransportType = 0
	// GKTransportTypeUnreliable - The data is sent once and is not sent again if a transmission error occurs.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTransportType/unreliable
	GKTransportTypeUnreliable GKTransportType = 0
)

/* debug [enums.gen.go]: Processing enum GKTurnBasedExchangeStatus (5 cases) */
// GKTurnBasedExchangeStatus - The status of an exchange or reply.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchangeStatus
type GKTurnBasedExchangeStatus uint

const (
	// GKTurnBasedExchangeStatusActive - GameKit sent the exchange request to recipients but not all recipients replied.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchangeStatus/active
	GKTurnBasedExchangeStatusActive GKTurnBasedExchangeStatus = 0
	// GKTurnBasedExchangeStatusCanceled - The sender canceled the exchange request.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedExchangeStatus/canceled
	GKTurnBasedExchangeStatusCanceled GKTurnBasedExchangeStatus = 0
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

/* debug [enums.gen.go]: Processing enum GKVoiceChatPlayerState (5 cases) */
// GKVoiceChatPlayerState - The state of a player in a voice chat.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChat/PlayerState
type GKVoiceChatPlayerState uint

const (
	// GKVoiceChatPlayerConnected - The state when the player connects to the channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChat/PlayerState/connected
	GKVoiceChatPlayerConnected GKVoiceChatPlayerState = 0
	// GKVoiceChatPlayerConnecting - The state when the player is connecting to the channel, but isn’t connected yet.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChat/PlayerState/connecting
	GKVoiceChatPlayerConnecting GKVoiceChatPlayerState = 0
	// GKVoiceChatPlayerDisconnected - The state when the player left the channel.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChat/PlayerState/disconnected
	GKVoiceChatPlayerDisconnected GKVoiceChatPlayerState = 0
	// GKVoiceChatPlayerSilent - The state when the player isn’t speaking.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChat/PlayerState/silent
	GKVoiceChatPlayerSilent GKVoiceChatPlayerState = 0
	// GKVoiceChatPlayerSpeaking - The state when the player speaks.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChat/PlayerState/speaking
	GKVoiceChatPlayerSpeaking GKVoiceChatPlayerState = 0
)

/* debug [enums.gen.go]: Processing enum GKVoiceChatServiceError (17 cases) */
// GKVoiceChatServiceError - Error codes for the voice chat service error domain.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatServiceError-swift.struct/Code
type GKVoiceChatServiceError uint

const (
	// GKVoiceChatServiceAudioUnavailableError - The audio hardware is unavailable to the voice chat service.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatServiceError-swift.struct/Code/audioUnavailableError
	GKVoiceChatServiceAudioUnavailableError GKVoiceChatServiceError = 0
	// GKVoiceChatServiceClientMissingRequiredMethodsError - The voice chat service did not find an expected method defined by the client.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatServiceError-swift.struct/Code/clientMissingRequiredMethodsError
	GKVoiceChatServiceClientMissingRequiredMethodsError GKVoiceChatServiceError = 0
	// GKVoiceChatServiceInternalError - A serious error occurred inside the voice chat service.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatServiceError-swift.struct/Code/internalError
	GKVoiceChatServiceInternalError GKVoiceChatServiceError = 0
	// GKVoiceChatServiceInvalidCallIDError - The voice chat service didn’t recognize the call identifier.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatServiceError-swift.struct/Code/invalidCallIDError
	GKVoiceChatServiceInvalidCallIDError GKVoiceChatServiceError = 0
	// GKVoiceChatServiceInvalidParameterError - A parameter had an unrecognized value.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatServiceError-swift.struct/Code/invalidParameterError
	GKVoiceChatServiceInvalidParameterError GKVoiceChatServiceError = 0
	// GKVoiceChatServiceMethodCurrentlyInvalidError - A method on the voice chat service was called when it was not allowed to be called (for example, attempting to connect when the voice chat service was already connected).
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatServiceError-swift.struct/Code/methodCurrentlyInvalidError
	GKVoiceChatServiceMethodCurrentlyInvalidError GKVoiceChatServiceError = 0
	// GKVoiceChatServiceNetworkConfigurationError - The voice chat service had problems accessing the network.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatServiceError-swift.struct/Code/networkConfigurationError
	GKVoiceChatServiceNetworkConfigurationError GKVoiceChatServiceError = 0
	// GKVoiceChatServiceNoRemotePacketsError - The voice chat service stopped receiving packets from the remote participant.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatServiceError-swift.struct/Code/noRemotePacketsError
	GKVoiceChatServiceNoRemotePacketsError GKVoiceChatServiceError = 0
	// GKVoiceChatServiceOutOfMemoryError - The voice chat service was unable to allocate memory required to operate.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatServiceError-swift.struct/Code/outOfMemoryError
	GKVoiceChatServiceOutOfMemoryError GKVoiceChatServiceError = 0
	// GKVoiceChatServiceRemoteParticipantBusyError - The remote participant is already connected to a voice chat.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatServiceError-swift.struct/Code/remoteParticipantBusyError
	GKVoiceChatServiceRemoteParticipantBusyError GKVoiceChatServiceError = 0
	// GKVoiceChatServiceRemoteParticipantCancelledError - A remote participant attempted to start a voice chat, then canceled.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatServiceError-swift.struct/Code/remoteParticipantCancelledError
	GKVoiceChatServiceRemoteParticipantCancelledError GKVoiceChatServiceError = 0
	// GKVoiceChatServiceRemoteParticipantDeclinedInviteError - A remote participant declined an invitation.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatServiceError-swift.struct/Code/remoteParticipantDeclinedInviteError
	GKVoiceChatServiceRemoteParticipantDeclinedInviteError GKVoiceChatServiceError = 0
	// GKVoiceChatServiceRemoteParticipantHangupError - The remote participant in a voice chat stopped the chat.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatServiceError-swift.struct/Code/remoteParticipantHangupError
	GKVoiceChatServiceRemoteParticipantHangupError GKVoiceChatServiceError = 0
	// GKVoiceChatServiceRemoteParticipantResponseInvalidError - Invalid data was received from a remote participant.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatServiceError-swift.struct/Code/remoteParticipantResponseInvalidError
	GKVoiceChatServiceRemoteParticipantResponseInvalidError GKVoiceChatServiceError = 0
	// GKVoiceChatServiceUnableToConnectError - The voice chat service was unable to establish a connection with another user.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatServiceError-swift.struct/Code/unableToConnectError
	GKVoiceChatServiceUnableToConnectError GKVoiceChatServiceError = 0
	// GKVoiceChatServiceUninitializedClientError - The application did not set a client before calling voice chat service methods.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatServiceError-swift.struct/Code/uninitializedClientError
	GKVoiceChatServiceUninitializedClientError GKVoiceChatServiceError = 0
	// GKVoiceChatServiceUnsupportedRemoteVersionError - The other participant is running a different version of the voice chat service.
	//
	// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKVoiceChatServiceError-swift.struct/Code/unsupportedRemoteVersionError
	GKVoiceChatServiceUnsupportedRemoteVersionError GKVoiceChatServiceError = 0
)


