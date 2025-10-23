// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [TurnBasedMatch] class.
type ITurnBasedMatch interface {
	objectivec.IObject
	// properties:
	Participants() []TurnBasedParticipant /* primitive/slice/pointer. */
	ActiveExchanges() IGKTurnBasedExchange
	SetActiveExchanges(value IGKTurnBasedExchange)
	CompletedExchanges() IGKTurnBasedExchange
	SetCompletedExchanges(value IGKTurnBasedExchange)
	CreationDate() foundation.objc.IObject /* cross-framework: Date */
	SetCreationDate(value foundation.objc.IObject /* cross-framework: Date */)
	CurrentParticipant() IGKTurnBasedParticipant
	SetCurrentParticipant(value IGKTurnBasedParticipant)
	ExchangeDataMaximumSize() int /* primitive/slice/pointer. */
	SetExchangeDataMaximumSize(value int /* primitive/slice/pointer. */)
	ExchangeMaxInitiatedExchangesPerPlayer() int /* primitive/slice/pointer. */
	SetExchangeMaxInitiatedExchangesPerPlayer(value int /* primitive/slice/pointer. */)
	Exchanges() IGKTurnBasedExchange
	SetExchanges(value IGKTurnBasedExchange)
	MatchData() foundation.objc.IObject /* cross-framework: Data */
	SetMatchData(value foundation.objc.IObject /* cross-framework: Data */)
	MatchDataMaximumSize() int /* primitive/slice/pointer. */
	SetMatchDataMaximumSize(value int /* primitive/slice/pointer. */)
	MatchID() string /* primitive/slice/pointer. */
	SetMatchID(value string /* primitive/slice/pointer. */)
	Message() string /* primitive/slice/pointer. */
	SetMessage(value string /* primitive/slice/pointer. */)
	Status() unsafe.Pointer
	SetStatus(value unsafe.Pointer)
	// methods:
	SendExchangeToParticipantsDataLocalizableMessageKeyArgumentsTimeoutCompletionHandler(participants []TurnBasedParticipant /* primitive/slice/pointer. */, data foundation.objc.IObject /* cross-framework NSData */, key string /* primitive/slice/pointer. */, arguments []string /* primitive/slice/pointer. */, timeout foundation.TimeInterval /* not a class type */, completionHandler unsafe.Pointer)
	SetLocalizableMessageWithKeyArguments(key string /* primitive/slice/pointer. */, arguments []string /* primitive/slice/pointer. */)
}

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

// Alloc allocates a new instance without initialization.
func (tc _TurnBasedMatchClass) Alloc() TurnBasedMatch {
	rv := objc.Send[TurnBasedMatch](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Loads a specific match with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/load(withID:withCompletionHandler:)
func (tc _TurnBasedMatchClass) LoadMatchWithIDWithCompletionHandler(matchID string /* primitive/slice/pointer. */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("loadMatchWithID:withCompletionHandler:"), objc.String(matchID), completionHandler)
}


// Sends an exchange request that contains your game data to one or more participants.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/sendExchange(to:data:localizableMessageKey:arguments:timeout:completionHandler:)
func (t_ TurnBasedMatch) SendExchangeToParticipantsDataLocalizableMessageKeyArgumentsTimeoutCompletionHandler(participants []TurnBasedParticipant /* primitive/slice/pointer. */, data foundation.objc.IObject /* cross-framework NSData */, key string /* primitive/slice/pointer. */, arguments []string /* primitive/slice/pointer. */, timeout foundation.TimeInterval /* not a class type */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("sendExchangeToParticipants:data:localizableMessageKey:arguments:timeout:completionHandler:"), participants, data, objc.String(key), arguments, timeout, completionHandler)
}


// Sends a localized message from the current participant to all other participants when you end a turn, forfeit a match, or end a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/setLocalizableMessageWithKey(_:arguments:)
func (t_ TurnBasedMatch) SetLocalizableMessageWithKeyArguments(key string /* primitive/slice/pointer. */, arguments []string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLocalizableMessageWithKey:arguments:"), objc.String(key), arguments)
}


// The players that participate in a turn-based match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/participants
func (t_ TurnBasedMatch) Participants() []TurnBasedParticipant /* primitive/slice/pointer. */ {
	rv := objc.Send[[]TurnBasedParticipant](t_.ID, objc.Sel("participants"))
	return rv
}


// The exchanges that the local player needs to accept or reject.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/activeexchanges
func (t_ TurnBasedMatch) ActiveExchanges() IGKTurnBasedExchange {
	rv := objc.Send[TurnBasedExchange](t_.ID, objc.Sel("activeExchanges"))
	return rv
}


// The exchanges that the local player needs to accept or reject.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/activeexchanges
func (t_ TurnBasedMatch) SetActiveExchanges(value IGKTurnBasedExchange) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setActiveExchanges:"), value)
}


// The exchange requests that all recipients replied to and the current participant needs to save.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/completedexchanges
func (t_ TurnBasedMatch) CompletedExchanges() IGKTurnBasedExchange {
	rv := objc.Send[TurnBasedExchange](t_.ID, objc.Sel("completedExchanges"))
	return rv
}


// The exchange requests that all recipients replied to and the current participant needs to save.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/completedexchanges
func (t_ TurnBasedMatch) SetCompletedExchanges(value IGKTurnBasedExchange) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCompletedExchanges:"), value)
}


// The date that Game Center created the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/creationdate
func (t_ TurnBasedMatch) CreationDate() foundation.objc.IObject /* cross-framework: Date */ {
	rv := objc.Send[foundation.Date](t_.ID, objc.Sel("creationDate"))
	return rv
}


// The date that Game Center created the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/creationdate
func (t_ TurnBasedMatch) SetCreationDate(value foundation.objc.IObject /* cross-framework: Date */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCreationDate:"), value)
}


// The participant whose turn it is.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/currentparticipant
func (t_ TurnBasedMatch) CurrentParticipant() IGKTurnBasedParticipant {
	rv := objc.Send[TurnBasedParticipant](t_.ID, objc.Sel("currentParticipant"))
	return rv
}


// The participant whose turn it is.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/currentparticipant
func (t_ TurnBasedMatch) SetCurrentParticipant(value IGKTurnBasedParticipant) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setCurrentParticipant:"), value)
}


// The maximum size of the exchange data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/exchangedatamaximumsize
func (t_ TurnBasedMatch) ExchangeDataMaximumSize() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](t_.ID, objc.Sel("exchangeDataMaximumSize"))
	return rv
}


// The maximum size of the exchange data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/exchangedatamaximumsize
func (t_ TurnBasedMatch) SetExchangeDataMaximumSize(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setExchangeDataMaximumSize:"), value)
}


// The maximum number of exchanges the local player can initiate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/exchangemaxinitiatedexchangesperplayer
func (t_ TurnBasedMatch) ExchangeMaxInitiatedExchangesPerPlayer() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](t_.ID, objc.Sel("exchangeMaxInitiatedExchangesPerPlayer"))
	return rv
}


// The maximum number of exchanges the local player can initiate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/exchangemaxinitiatedexchangesperplayer
func (t_ TurnBasedMatch) SetExchangeMaxInitiatedExchangesPerPlayer(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setExchangeMaxInitiatedExchangesPerPlayer:"), value)
}


// The exchange requests that are active or complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/exchanges
func (t_ TurnBasedMatch) Exchanges() IGKTurnBasedExchange {
	rv := objc.Send[TurnBasedExchange](t_.ID, objc.Sel("exchanges"))
	return rv
}


// The exchange requests that are active or complete.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/exchanges
func (t_ TurnBasedMatch) SetExchanges(value IGKTurnBasedExchange) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setExchanges:"), value)
}


// The game-specific data that you store in Game Center and pass between participants through a match object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/matchdata
func (t_ TurnBasedMatch) MatchData() foundation.objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](t_.ID, objc.Sel("matchData"))
	return rv
}


// The game-specific data that you store in Game Center and pass between participants through a match object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/matchdata
func (t_ TurnBasedMatch) SetMatchData(value foundation.objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMatchData:"), value)
}


// The maximum size of the match data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/matchdatamaximumsize
func (t_ TurnBasedMatch) MatchDataMaximumSize() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](t_.ID, objc.Sel("matchDataMaximumSize"))
	return rv
}


// The maximum size of the match data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/matchdatamaximumsize
func (t_ TurnBasedMatch) SetMatchDataMaximumSize(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMatchDataMaximumSize:"), value)
}


// A unique identifier for the turn-based match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/matchid
func (t_ TurnBasedMatch) MatchID() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](t_.ID, objc.Sel("matchID"))
	return rv
}


// A unique identifier for the turn-based match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/matchid
func (t_ TurnBasedMatch) SetMatchID(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMatchID:"), objc.String(value))
}


// A message from the current participant to all other participants when you end a turn, forfeit a match, or end a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/message
func (t_ TurnBasedMatch) Message() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](t_.ID, objc.Sel("message"))
	return rv
}


// A message from the current participant to all other participants when you end a turn, forfeit a match, or end a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/message
func (t_ TurnBasedMatch) SetMessage(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMessage:"), objc.String(value))
}


// The state of the match, such as whether the match is open or has ended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/status-swift.property
func (t_ TurnBasedMatch) Status() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("status"))
	return rv
}


// The state of the match, such as whether the match is open or has ended.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedmatch/status-swift.property
func (t_ TurnBasedMatch) SetStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setStatus:"), value)
}



