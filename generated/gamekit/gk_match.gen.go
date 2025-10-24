// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKMatch */


/* debug [class_header]: Header for GKMatch */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Match */
// An interface definition for the [Match] class.
type IMatch interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Match */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	ExpectedPlayerCount() uint
	PlayerIDs() []string
	PlayerProperties() foundation.IDictionary
	Players() []Player
	Properties() MatchProperties /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Match */
	// methods:
	ChooseBestHostingPlayerWithCompletionHandler(completionHandler unsafe.Pointer)
	Disconnect()
	RematchWithCompletionHandler(completionHandler unsafe.Pointer)
	SendDataToPlayersDataModeError(data objc.IObject /* cross-framework: NSData */, players []Player, mode MatchSendDataMode, error_ unsafe.Pointer) bool
	SendDataToAllPlayersWithDataModeError(data objc.IObject /* cross-framework: NSData */, mode MatchSendDataMode, error_ unsafe.Pointer) bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Match */
// Alloc allocates a new instance without initialization.
func (mc _MatchClass) Alloc() Match {
	rv := objc.Send[Match](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Match */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Match *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Match */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Match */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Match */

// Determines the best player in the game to act as the server for a client-server topology.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatch/chooseBestHostingPlayer(completionHandler:)
func (m_ Match) ChooseBestHostingPlayerWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("chooseBestHostingPlayerWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: ChooseBestHostingPlayerWithCompletionHandler */


// Disconnects the local player from the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatch/disconnect()
func (m_ Match) Disconnect() {
	objc.Send[objc.ID](m_.ID, objc.Sel("disconnect"))
}/* debug [instance_methods/method]: Disconnect */


// Creates a new match with the players from an existing match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatch/rematch(completionHandler:)
func (m_ Match) RematchWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("rematchWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: RematchWithCompletionHandler */


// Transmits data to one or more players connected to the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatch/send(_:to:dataMode:)
func (m_ Match) SendDataToPlayersDataModeError(data objc.IObject /* cross-framework: NSData */, players []Player, mode MatchSendDataMode, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("sendData:toPlayers:dataMode:error:"), data, players, mode, error_)
	return rv
}/* debug [instance_methods/method]: SendDataToPlayersDataModeError */


// Transmits data to all players connected to the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatch/sendData(toAllPlayers:with:)
func (m_ Match) SendDataToAllPlayersWithDataModeError(data objc.IObject /* cross-framework: NSData */, mode MatchSendDataMode, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("sendDataToAllPlayers:withDataMode:error:"), data, mode, error_)
	return rv
}/* debug [instance_methods/method]: SendDataToAllPlayersWithDataModeError */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Match */

// The delegate that handles communication between players in a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatch/delegate
func (m_ Match) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate that handles communication between players in a match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatch/delegate
func (m_ Match) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The remaining number of players invited but not yet connected to the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatch/expectedPlayerCount
func (m_ Match) ExpectedPlayerCount() uint {
	rv := objc.Send[uint](m_.ID, objc.Sel("expectedPlayerCount"))
	return rv
}/* debug [instance_properties/getter]: expectedPlayerCount */


// The player identifiers for remote players in the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatch/playerIDs
func (m_ Match) PlayerIDs() []string {
	rv := objc.Send[[]string](m_.ID, objc.Sel("playerIDs"))
	return rv
}/* debug [instance_properties/getter]: playerIDs */


// The properties for other players that matchmaking rules uses to find players, with some additions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatch/playerProperties
func (m_ Match) PlayerProperties() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("playerProperties"))
	return rv
}/* debug [instance_properties/getter]: playerProperties */


// The players that join the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatch/players
func (m_ Match) Players() []Player {
	rv := objc.Send[[]Player](m_.ID, objc.Sel("players"))
	return rv
}/* debug [instance_properties/getter]: players */


// The local player’s properties that matchmaking rules used to find the players with some additions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatch/properties
func (m_ Match) Properties() MatchProperties /* not a class type */ {
	rv := objc.Send[MatchProperties](m_.ID, objc.Sel("properties"))
	return rv
}/* debug [instance_properties/getter]: properties */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKMatch */



