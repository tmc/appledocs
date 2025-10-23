// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Match] class.
var (
	MatchClass     _MatchClass
	MatchClassOnce sync.Once
)

func getMatchClass() _MatchClass {
	MatchClassOnce.Do(func() {
		MatchClass = _MatchClass{objc.GetClass("GKMatch")}
	})
	return MatchClass
}

type _MatchClass struct {
	class objc.Class
}

// An interface definition for the [Match] class.
type IMatch interface {
	objectivec.IObject
	// properties:
	ExpectedPlayerCount() uint /* primitive/slice/pointer. */
	PlayerProperties() foundation.IDictionary /* already interface */
	Delegate() MatchDelegate /* not a class type */
	SetDelegate(value MatchDelegate /* not a class type */)
	PlayerIDs() string /* primitive/slice/pointer. */
	SetPlayerIDs(value string /* primitive/slice/pointer. */)
	Players() IGKPlayer
	SetPlayers(value IGKPlayer)
	Properties() string /* primitive/slice/pointer. */
	SetProperties(value string /* primitive/slice/pointer. */)
	// methods:
	Disconnect()
	SendDataToPlayersDataModeError(data foundation.objc.IObject /* cross-framework NSData */, players []Player /* primitive/slice/pointer. */, mode MatchSendDataMode /* not a class type */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */
}

// A peer-to-peer network between a group of players that sign into Game Center.
//
// Matches provide a mechanism for a player to send both game and voice data to other players. You never create a object directly. Instead, GameKit passes a match object to a method or a handler when you set up a multiplayer game. For details, see . If you use the class to find players, implement the delegate method to set the match delegate. If you use the class, set the match delegate in the handler you pass to the method. You can begin exchanging data when two or more players join the match. Implement the delegate method to track when players connect or disconnect from the match. Then use either the or the method to send data. To process the data on the recipient side, implement the delegate method. To implement voice chat, use the method to create one or more voice channels represented by the returned object. When you’re finished with a match, call the method and set the match’s delegate to . Otherwise, GameKit may send to the delegate until all players disconnect from the match.


// A peer-to-peer network between a group of players that sign into Game Center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatch
type Match struct {
	objectivec.Object
}

// MatchFrom constructs a [Match] from an unsafe.Pointer.
//
// A peer-to-peer network between a group of players that sign into Game Center.
func MatchFrom(ptr unsafe.Pointer) Match {
	return Match{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (mc _MatchClass) Alloc() Match {
	rv := objc.Send[Match](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (mc _MatchClass) New() Match {
	rv := objc.Send[Match](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ Match) Init() Match {
	rv := objc.Send[Match](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ Match) Autorelease() Match {
	rv := objc.Send[Match](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatch creates a new Match instance.
func NewMatch() Match {
	return getMatchClass().New()
}



// Disconnects the local player from the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatch/disconnect()
func (m_ Match) Disconnect() {
	objc.Send[objc.ID](m_.ID, objc.Sel("disconnect"))
}


// Transmits data to one or more players connected to the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatch/send(_:to:dataMode:)
func (m_ Match) SendDataToPlayersDataModeError(data foundation.objc.IObject /* cross-framework NSData */, players []Player /* primitive/slice/pointer. */, mode MatchSendDataMode /* not a class type */, error_ unsafe.Pointer) bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](m_.ID, objc.Sel("sendData:toPlayers:dataMode:error:"), data, players, mode, error_)
	return rv
}


// The remaining number of players invited but not yet connected to the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatch/expectedPlayerCount
func (m_ Match) ExpectedPlayerCount() uint /* primitive/slice/pointer. */ {
	rv := objc.Send[uint](m_.ID, objc.Sel("expectedPlayerCount"))
	return rv
}


// The properties for other players that matchmaking rules uses to find players, with some additions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatch/playerProperties
func (m_ Match) PlayerProperties() foundation.IDictionary /* already interface */ {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("playerProperties"))
	return rv
}


// The delegate that handles communication between players in a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatch/delegate
func (m_ Match) Delegate() MatchDelegate /* not a class type */ {
	rv := objc.Send[MatchDelegate](m_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate that handles communication between players in a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatch/delegate
func (m_ Match) SetDelegate(value MatchDelegate /* not a class type */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}


// The player identifiers for remote players in the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatch/playerids
func (m_ Match) PlayerIDs() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](m_.ID, objc.Sel("playerIDs"))
	return rv
}


// The player identifiers for remote players in the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatch/playerids
func (m_ Match) SetPlayerIDs(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlayerIDs:"), objc.String(value))
}


// The players that join the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatch/players
func (m_ Match) Players() IGKPlayer {
	rv := objc.Send[Player](m_.ID, objc.Sel("players"))
	return rv
}


// The players that join the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatch/players
func (m_ Match) SetPlayers(value IGKPlayer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPlayers:"), value)
}


// The local player’s properties that matchmaking rules used to find the players with some additions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatch/properties
func (m_ Match) Properties() string /* primitive/slice/pointer. */ {
	rv := objc.Send[string](m_.ID, objc.Sel("properties"))
	return rv
}


// The local player’s properties that matchmaking rules used to find the players with some additions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkmatch/properties
func (m_ Match) SetProperties(value string /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProperties:"), objc.String(value))
}



