// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Matchmaker] class.
var (
	MatchmakerClass     _MatchmakerClass
	MatchmakerClassOnce sync.Once
)

func getMatchmakerClass() _MatchmakerClass {
	MatchmakerClassOnce.Do(func() {
		MatchmakerClass = _MatchmakerClass{objc.GetClass("GKMatchmaker")}
	})
	return MatchmakerClass
}

type _MatchmakerClass struct {
	class objc.Class
}

// An interface definition for the [Matchmaker] class.
type IMatchmaker interface {
	objectivec.IObject
	// properties:
	ExpectedPlayerCount() int /* primitive/slice/pointer. */
	SetExpectedPlayerCount(value int /* primitive/slice/pointer. */)
	// methods:
	AddPlayersToMatchMatchRequestCompletionHandler(match IGKMatch, matchRequest IGKMatchRequest, completionHandler unsafe.Pointer)
}

// An object that creates matches with other players without presenting an interface to the players.
//
// Use the class to auto-match players for a quicker game start, programmatically invite specific players, or implement your own interface for players to invite other players. If you want to present a familiar matchmaking GameKit interface to players, instead use either the or class. If you host a game on your own server, you can also use this class to find Game Center players. That is, you implement the networking and communication between the players through your own servers not Game Center. To find players using this class, create a object and configure it according to the parameters of your game. Then, pass the match request and a handler using the method, or the method for hosted games, to the shared object. GameKit calls the handler when players accept their invitations. Implement the handler to set the delegate of the object that GameKit sends and start the game when there are enough players. If the match doesn’t have enough players (for example, some players decline their invitations), you can create another match request and call the method repeatedly until the match’s property is zero. When you have enough players to start the match, call the method to end the matchmaking process. If you provide a SharePlay interface for inviting players, use the and methods to create a group activity on behalf of the player.


// An object that creates matches with other players without presenting an interface to the players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmaker
type Matchmaker struct {
	objectivec.Object
}

// MatchmakerFrom constructs a [Matchmaker] from an unsafe.Pointer.
//
// An object that creates matches with other players without presenting an interface to the players.
func MatchmakerFrom(ptr unsafe.Pointer) Matchmaker {
	return Matchmaker{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MatchmakerClass) Alloc() Matchmaker {
	rv := objc.Send[Matchmaker](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MatchmakerClass) New() Matchmaker {
	rv := objc.Send[Matchmaker](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ Matchmaker) Init() Matchmaker {
	rv := objc.Send[Matchmaker](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ Matchmaker) Autorelease() Matchmaker {
	rv := objc.Send[Matchmaker](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatchmaker creates a new Matchmaker instance.
func NewMatchmaker() Matchmaker {
	return getMatchmakerClass().New()
}



// Invites additional players to an existing match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchmaker/addPlayers(to:matchRequest:completionHandler:)
func (m_ Matchmaker) AddPlayersToMatchMatchRequestCompletionHandler(match IGKMatch, matchRequest IGKMatchRequest, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("addPlayersToMatch:matchRequest:completionHandler:"), match, matchRequest, completionHandler)
}


// The remaining number of players invited but not yet connected to the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatch/expectedplayercount
func (m_ Matchmaker) ExpectedPlayerCount() int /* primitive/slice/pointer. */ {
	rv := objc.Send[int](m_.ID, objc.Sel("expectedPlayerCount"))
	return rv
}


// The remaining number of players invited but not yet connected to the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatch/expectedplayercount
func (m_ Matchmaker) SetExpectedPlayerCount(value int /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setExpectedPlayerCount:"), value)
}



