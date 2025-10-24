// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKTurnBasedMatch */


/* debug [class_header]: Header for GKTurnBasedMatch */
// The class instance for the [TurnBasedMatch] class.
var (
	TurnBasedMatchClass     _TurnBasedMatchClass
	TurnBasedMatchClassOnce sync.Once
)

func getTurnBasedMatchClass() _TurnBasedMatchClass {
	TurnBasedMatchClassOnce.Do(func() {
		TurnBasedMatchClass = _TurnBasedMatchClass{objc.GetClass("GKTurnBasedMatch")}
	})
	return TurnBasedMatchClass
}

type _TurnBasedMatchClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TurnBasedMatch */
// An interface definition for the [TurnBasedMatch] class.
type ITurnBasedMatch interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TurnBasedMatch */
	// properties:
	ActiveExchanges() []TurnBasedExchange
	CompletedExchanges() []TurnBasedExchange
	CreationDate() objc.IObject /* cross-framework: NSDate */
	CurrentParticipant() IGKTurnBasedParticipant
	ExchangeDataMaximumSize() uint
	ExchangeMaxInitiatedExchangesPerPlayer() uint
	Exchanges() []TurnBasedExchange
	MatchData() objc.IObject /* cross-framework: NSData */
	MatchDataMaximumSize() uint
	MatchID() objc.IObject /* cross-framework: NSString */
	Message() objc.IObject /* cross-framework: NSString */
	SetMessage(value objc.IObject /* cross-framework: NSString */)
	Participants() []TurnBasedParticipant
	Status() TurnBasedMatchStatus
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TurnBasedMatch */
	// methods:
	AcceptInviteWithCompletionHandler(completionHandler unsafe.Pointer)
	DeclineInviteWithCompletionHandler(completionHandler unsafe.Pointer)
	EndMatchInTurnWithMatchDataCompletionHandler(matchData objc.IObject /* cross-framework: NSData */, completionHandler unsafe.Pointer)
	EndMatchInTurnWithMatchDataLeaderboardScoresAchievementsCompletionHandler(matchData objc.IObject /* cross-framework: NSData */, scores []LeaderboardScore, achievements objc.IObject /* cross-framework: NSArray */, completionHandler unsafe.Pointer)
	EndTurnWithNextParticipantsTurnTimeoutMatchDataCompletionHandler(nextParticipants []TurnBasedParticipant, timeout float64, matchData objc.IObject /* cross-framework: NSData */, completionHandler unsafe.Pointer)
	LoadMatchDataWithCompletionHandler(completionHandler unsafe.Pointer)
	ParticipantQuitInTurnWithOutcomeNextParticipantsTurnTimeoutMatchDataCompletionHandler(matchOutcome TurnBasedMatchOutcome, nextParticipants []TurnBasedParticipant, timeout float64, matchData objc.IObject /* cross-framework: NSData */, completionHandler unsafe.Pointer)
	ParticipantQuitOutOfTurnWithOutcomeWithCompletionHandler(matchOutcome TurnBasedMatchOutcome, completionHandler unsafe.Pointer)
	RematchWithCompletionHandler(completionHandler unsafe.Pointer)
	RemoveWithCompletionHandler(completionHandler unsafe.Pointer)
	SaveCurrentTurnWithMatchDataCompletionHandler(matchData objc.IObject /* cross-framework: NSData */, completionHandler unsafe.Pointer)
	SaveMergedMatchDataWithResolvedExchangesCompletionHandler(matchData objc.IObject /* cross-framework: NSData */, exchanges []TurnBasedExchange, completionHandler unsafe.Pointer)
	SendExchangeToParticipantsDataLocalizableMessageKeyArgumentsTimeoutCompletionHandler(participants []TurnBasedParticipant, data objc.IObject /* cross-framework: NSData */, key objc.IObject /* cross-framework: NSString */, arguments []string, timeout float64, completionHandler unsafe.Pointer)
	SendReminderToParticipantsLocalizableMessageKeyArgumentsCompletionHandler(participants []TurnBasedParticipant, key objc.IObject /* cross-framework: NSString */, arguments []string, completionHandler unsafe.Pointer)
	SetLocalizableMessageWithKeyArguments(key objc.IObject /* cross-framework: NSString */, arguments []string)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TurnBasedMatch */
// Alloc allocates a new instance without initialization.
func (tc _TurnBasedMatchClass) Alloc() TurnBasedMatch {
	rv := objc.Send[TurnBasedMatch](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TurnBasedMatchClass) New() TurnBasedMatch {
	rv := objc.Send[TurnBasedMatch](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TurnBasedMatch) Init() TurnBasedMatch {
	rv := objc.Send[TurnBasedMatch](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TurnBasedMatch) Autorelease() TurnBasedMatch {
	rv := objc.Send[TurnBasedMatch](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTurnBasedMatch creates a new TurnBasedMatch instance.
func NewTurnBasedMatch() TurnBasedMatch {
	return getTurnBasedMatchClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TurnBasedMatch */
// An object that encapsulates the match data for games where players take turns.
//
// A object represents a match in a turn-based game that Game Center stores and forwards to participants in the match. In a turn-based game, participants take turns to advance gameplay until they reach an outcome. You end the match when all participants reach an outcome or they can no longer continue. A turn-based match object contains the status of the match, list of participants, the participant whose turn it is, a message about the last turn, and your game-specific data. You can get more details about the participants through the objects in the property. You don’t create turn-based match objects directly. When a match event occurs, GameKit passes the match object to listeners that conform to the protocol. Retain the match object or its match ID in the protocol methods, so you can get the latest match data later during gameplay. Using the match object passed to protocol methods, you can perform these actions on behalf of the local player: Save game data End a turn Forfeit a match End a match Send a reminder to the participant whose turn it is Exchange data between participants Remove a completed match from Game Center When you end a turn, forfeit a match, or end a match, you update the match data and if gameplay can continue, choose the next participant. If you end a match, you set the individual participant outcomes as well. If you present a object for players to manage their turn-based matches, the player can start a match, accept an invitation, open an existing match, and forfeit a match.


// An object that encapsulates the match data for games where players take turns.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch
type TurnBasedMatch struct {
	objectivec.Object
}

// TurnBasedMatchFrom constructs a [TurnBasedMatch] from an unsafe.Pointer.
//
// An object that encapsulates the match data for games where players take turns.
func TurnBasedMatchFrom(ptr unsafe.Pointer) TurnBasedMatch {
	return TurnBasedMatch{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TurnBasedMatch *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TurnBasedMatch */

// Creates a new match or finds an existing match that needs a player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/find(for:withCompletionHandler:)
func (tc _TurnBasedMatchClass) FindMatchForRequestWithCompletionHandler(request IGKMatchRequest, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("findMatchForRequest:withCompletionHandler:"), request, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=FindMatchForRequestWithCompletionHandler) */


// Loads a specific match with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/load(withID:withCompletionHandler:)
func (tc _TurnBasedMatchClass) LoadMatchWithIDWithCompletionHandler(matchID objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("loadMatchWithID:withCompletionHandler:"), matchID, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadMatchWithIDWithCompletionHandler) */


// Fetches the turn-based matches from Game Center that the local player participates in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/loadMatches(completionHandler:)
func (tc _TurnBasedMatchClass) LoadMatchesWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("loadMatchesWithCompletionHandler:"), completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadMatchesWithCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TurnBasedMatch */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TurnBasedMatch */

// Accepts an invitation for the local player to join a turn-based match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/acceptInvite(completionHandler:)
func (t_ TurnBasedMatch) AcceptInviteWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("acceptInviteWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: AcceptInviteWithCompletionHandler */


// Declines an invitation for the local player to join a turn-based match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/declineInvite(completionHandler:)
func (t_ TurnBasedMatch) DeclineInviteWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("declineInviteWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: DeclineInviteWithCompletionHandler */


// Ends the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/endMatchInTurn(withMatch:completionHandler:)
func (t_ TurnBasedMatch) EndMatchInTurnWithMatchDataCompletionHandler(matchData objc.IObject /* cross-framework: NSData */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("endMatchInTurnWithMatchData:completionHandler:"), matchData, completionHandler)
}/* debug [instance_methods/method]: EndMatchInTurnWithMatchDataCompletionHandler */


// Ends the match while submitting scores and achievements for all of the participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/endMatchInTurn(withMatch:leaderboardScores:achievements:completionHandler:)
func (t_ TurnBasedMatch) EndMatchInTurnWithMatchDataLeaderboardScoresAchievementsCompletionHandler(matchData objc.IObject /* cross-framework: NSData */, scores []LeaderboardScore, achievements objc.IObject /* cross-framework: NSArray */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("endMatchInTurnWithMatchData:leaderboardScores:achievements:completionHandler:"), matchData, scores, achievements, completionHandler)
}/* debug [instance_methods/method]: EndMatchInTurnWithMatchDataLeaderboardScoresAchievementsCompletionHandler */


// Passes the turn from the current participant to the next participant.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/endTurn(withNextParticipants:turnTimeout:match:completionHandler:)
func (t_ TurnBasedMatch) EndTurnWithNextParticipantsTurnTimeoutMatchDataCompletionHandler(nextParticipants []TurnBasedParticipant, timeout float64, matchData objc.IObject /* cross-framework: NSData */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("endTurnWithNextParticipants:turnTimeout:matchData:completionHandler:"), nextParticipants, timeout, matchData, completionHandler)
}/* debug [instance_methods/method]: EndTurnWithNextParticipantsTurnTimeoutMatchDataCompletionHandler */


// Fetches your game-specific data that you store in Game Center when ending a turn, saving a turn, or leaving a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/loadMatchData(completionHandler:)
func (t_ TurnBasedMatch) LoadMatchDataWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("loadMatchDataWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: LoadMatchDataWithCompletionHandler */


// Forfeits the match on behalf of the local player when it’s their turn.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/participantQuitInTurn(with:nextParticipants:turnTimeout:match:completionHandler:)
func (t_ TurnBasedMatch) ParticipantQuitInTurnWithOutcomeNextParticipantsTurnTimeoutMatchDataCompletionHandler(matchOutcome TurnBasedMatchOutcome, nextParticipants []TurnBasedParticipant, timeout float64, matchData objc.IObject /* cross-framework: NSData */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("participantQuitInTurnWithOutcome:nextParticipants:turnTimeout:matchData:completionHandler:"), matchOutcome, nextParticipants, timeout, matchData, completionHandler)
}/* debug [instance_methods/method]: ParticipantQuitInTurnWithOutcomeNextParticipantsTurnTimeoutMatchDataCompletionHandler */


// Forfeits the match on behalf of the local player when it’s not their turn.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/participantQuitOutOfTurn(with:withCompletionHandler:)
func (t_ TurnBasedMatch) ParticipantQuitOutOfTurnWithOutcomeWithCompletionHandler(matchOutcome TurnBasedMatchOutcome, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("participantQuitOutOfTurnWithOutcome:withCompletionHandler:"), matchOutcome, completionHandler)
}/* debug [instance_methods/method]: ParticipantQuitOutOfTurnWithOutcomeWithCompletionHandler */


// Creates a new turn-based match with the same participants from an existing match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/rematch(completionHandler:)
func (t_ TurnBasedMatch) RematchWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("rematchWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: RematchWithCompletionHandler */


// Removes a match from Game Center that the local player participants in.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/remove(completionHandler:)
func (t_ TurnBasedMatch) RemoveWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("removeWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: RemoveWithCompletionHandler */


// Saves your match data in Game Center without ending the turn.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/saveCurrentTurn(withMatch:completionHandler:)
func (t_ TurnBasedMatch) SaveCurrentTurnWithMatchDataCompletionHandler(matchData objc.IObject /* cross-framework: NSData */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("saveCurrentTurnWithMatchData:completionHandler:"), matchData, completionHandler)
}/* debug [instance_methods/method]: SaveCurrentTurnWithMatchDataCompletionHandler */


// Saves match data for completed exchanges without ending the turn.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/saveMergedMatch(_:withResolvedExchanges:completionHandler:)
func (t_ TurnBasedMatch) SaveMergedMatchDataWithResolvedExchangesCompletionHandler(matchData objc.IObject /* cross-framework: NSData */, exchanges []TurnBasedExchange, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("saveMergedMatchData:withResolvedExchanges:completionHandler:"), matchData, exchanges, completionHandler)
}/* debug [instance_methods/method]: SaveMergedMatchDataWithResolvedExchangesCompletionHandler */


// Sends an exchange request that contains your game data to one or more participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/sendExchange(to:data:localizableMessageKey:arguments:timeout:completionHandler:)
func (t_ TurnBasedMatch) SendExchangeToParticipantsDataLocalizableMessageKeyArgumentsTimeoutCompletionHandler(participants []TurnBasedParticipant, data objc.IObject /* cross-framework: NSData */, key objc.IObject /* cross-framework: NSString */, arguments []string, timeout float64, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("sendExchangeToParticipants:data:localizableMessageKey:arguments:timeout:completionHandler:"), participants, data, key, arguments, timeout, completionHandler)
}/* debug [instance_methods/method]: SendExchangeToParticipantsDataLocalizableMessageKeyArgumentsTimeoutCompletionHandler */


// Sends a reminder from one participant to a specific set of other participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/sendReminder(to:localizableMessageKey:arguments:completionHandler:)
func (t_ TurnBasedMatch) SendReminderToParticipantsLocalizableMessageKeyArgumentsCompletionHandler(participants []TurnBasedParticipant, key objc.IObject /* cross-framework: NSString */, arguments []string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("sendReminderToParticipants:localizableMessageKey:arguments:completionHandler:"), participants, key, arguments, completionHandler)
}/* debug [instance_methods/method]: SendReminderToParticipantsLocalizableMessageKeyArgumentsCompletionHandler */


// Sends a localized message from the current participant to all other participants when you end a turn, forfeit a match, or end a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/setLocalizableMessageWithKey(_:arguments:)
func (t_ TurnBasedMatch) SetLocalizableMessageWithKeyArguments(key objc.IObject /* cross-framework: NSString */, arguments []string) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLocalizableMessageWithKey:arguments:"), key, arguments)
}/* debug [instance_methods/method]: SetLocalizableMessageWithKeyArguments */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TurnBasedMatch */

// The exchanges that the local player needs to accept or reject.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/activeExchanges
func (t_ TurnBasedMatch) ActiveExchanges() []TurnBasedExchange {
	rv := objc.Send[[]TurnBasedExchange](t_.ID, objc.Sel("activeExchanges"))
	return rv
}/* debug [instance_properties/getter]: activeExchanges */


// The exchange requests that all recipients replied to and the current participant needs to save.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/completedExchanges
func (t_ TurnBasedMatch) CompletedExchanges() []TurnBasedExchange {
	rv := objc.Send[[]TurnBasedExchange](t_.ID, objc.Sel("completedExchanges"))
	return rv
}/* debug [instance_properties/getter]: completedExchanges */


// The date that Game Center created the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/creationDate
func (t_ TurnBasedMatch) CreationDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](t_.ID, objc.Sel("creationDate"))
	return rv
}/* debug [instance_properties/getter]: creationDate */


// The participant whose turn it is.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/currentParticipant
func (t_ TurnBasedMatch) CurrentParticipant() IGKTurnBasedParticipant {
	rv := objc.Send[TurnBasedParticipant](t_.ID, objc.Sel("currentParticipant"))
	return rv
}/* debug [instance_properties/getter]: currentParticipant */


// The maximum size of the exchange data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/exchangeDataMaximumSize
func (t_ TurnBasedMatch) ExchangeDataMaximumSize() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("exchangeDataMaximumSize"))
	return rv
}/* debug [instance_properties/getter]: exchangeDataMaximumSize */


// The maximum number of exchanges the local player can initiate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/exchangeMaxInitiatedExchangesPerPlayer
func (t_ TurnBasedMatch) ExchangeMaxInitiatedExchangesPerPlayer() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("exchangeMaxInitiatedExchangesPerPlayer"))
	return rv
}/* debug [instance_properties/getter]: exchangeMaxInitiatedExchangesPerPlayer */


// The exchange requests that are active or complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/exchanges
func (t_ TurnBasedMatch) Exchanges() []TurnBasedExchange {
	rv := objc.Send[[]TurnBasedExchange](t_.ID, objc.Sel("exchanges"))
	return rv
}/* debug [instance_properties/getter]: exchanges */


// The game-specific data that you store in Game Center and pass between participants through a match object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/matchData
func (t_ TurnBasedMatch) MatchData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](t_.ID, objc.Sel("matchData"))
	return rv
}/* debug [instance_properties/getter]: matchData */


// The maximum size of the match data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/matchDataMaximumSize
func (t_ TurnBasedMatch) MatchDataMaximumSize() uint {
	rv := objc.Send[uint](t_.ID, objc.Sel("matchDataMaximumSize"))
	return rv
}/* debug [instance_properties/getter]: matchDataMaximumSize */


// A unique identifier for the turn-based match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/matchID
func (t_ TurnBasedMatch) MatchID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("matchID"))
	return rv
}/* debug [instance_properties/getter]: matchID */


// A message from the current participant to all other participants when you end a turn, forfeit a match, or end a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/message
func (t_ TurnBasedMatch) Message() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](t_.ID, objc.Sel("message"))
	return rv
}/* debug [instance_properties/getter]: message */


// A message from the current participant to all other participants when you end a turn, forfeit a match, or end a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/message
func (t_ TurnBasedMatch) SetMessage(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMessage:"), value)
}/* debug [instance_properties/setter]: message */


// The players that participate in a turn-based match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/participants
func (t_ TurnBasedMatch) Participants() []TurnBasedParticipant {
	rv := objc.Send[[]TurnBasedParticipant](t_.ID, objc.Sel("participants"))
	return rv
}/* debug [instance_properties/getter]: participants */


// The state of the match, such as whether the match is open or has ended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/status-swift.property
func (t_ TurnBasedMatch) Status() TurnBasedMatchStatus {
	rv := objc.Send[TurnBasedMatchStatus](t_.ID, objc.Sel("status"))
	return rv
}/* debug [instance_properties/getter]: status */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKTurnBasedMatch */


