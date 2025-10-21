// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	SendExchangeToParticipantsDataLocalizableMessageKeyArgumentsTimeoutCompletionHandler(participants unsafe.Pointer, data unsafe.Pointer, key string, arguments unsafe.Pointer, timeout TimeInterval, completionHandler unsafe.Pointer)
	SetLocalizableMessageWithKeyArguments(key string, arguments unsafe.Pointer)
}

// An object that encapsulates the match data for games where players take turns.
//
// A object represents a match in a turn-based game that Game Center stores and forwards to participants in the match. In a turn-based game, participants take turns to advance gameplay until they reach an outcome. You end the match when all participants reach an outcome or they can no longer continue. A turn-based match object contains the status of the match, list of participants, the participant whose turn it is, a message about the last turn, and your game-specific data. You can get more details about the participants through the objects in the property. You don’t create turn-based match objects directly. When a match event occurs, GameKit passes the match object to listeners that conform to the protocol. Retain the match object or its match ID in the protocol methods, so you can get the latest match data later during gameplay. Using the match object passed to protocol methods, you can perform these actions on behalf of the local player: Save game data End a turn Forfeit a match End a match Send a reminder to the participant whose turn it is Exchange data between participants Remove a completed match from Game Center When you end a turn, forfeit a match, or end a match, you update the match data and if gameplay can continue, choose the next participant. If you end a match, you set the individual participant outcomes as well. If you present a object for players to manage their turn-based matches, the player can start a match, accept an invitation, open an existing match, and forfeit a match.
//
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
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/load(withID:withCompletionHandler:)
func (tc _TurnBasedMatchClass) LoadMatchWithIDWithCompletionHandler(matchID string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(tc.class), objc.Sel("loadMatchWithID:withCompletionHandler:"), objc.String(matchID), completionHandler)
}

// Sends an exchange request that contains your game data to one or more participants.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/sendExchange(to:data:localizableMessageKey:arguments:timeout:completionHandler:)
func (t_ TurnBasedMatch) SendExchangeToParticipantsDataLocalizableMessageKeyArgumentsTimeoutCompletionHandler(participants unsafe.Pointer, data unsafe.Pointer, key string, arguments unsafe.Pointer, timeout TimeInterval, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("sendExchangeToParticipants:data:localizableMessageKey:arguments:timeout:completionHandler:"), participants, data, objc.String(key), arguments, timeout, completionHandler)
}

// Sends a localized message from the current participant to all other participants when you end a turn, forfeit a match, or end a match.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatch/setLocalizableMessageWithKey(_:arguments:)
func (t_ TurnBasedMatch) SetLocalizableMessageWithKeyArguments(key string, arguments unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setLocalizableMessageWithKey:arguments:"), objc.String(key), arguments)
}



