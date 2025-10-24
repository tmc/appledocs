// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKMatchedPlayers */


/* debug [class_header]: Header for GKMatchedPlayers */
// The class instance for the [MatchedPlayers] class.
var (
	MatchedPlayersClass     _MatchedPlayersClass
	MatchedPlayersClassOnce sync.Once
)

func getMatchedPlayersClass() _MatchedPlayersClass {
	MatchedPlayersClassOnce.Do(func() {
		MatchedPlayersClass = _MatchedPlayersClass{objc.GetClass("GKMatchedPlayers")}
	})
	return MatchedPlayersClass
}

type _MatchedPlayersClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for MatchedPlayers */
// An interface definition for the [MatchedPlayers] class.
type IMatchedPlayers interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for MatchedPlayers */
	// properties:
	PlayerProperties() foundation.IDictionary
	Players() []Player
	Properties() MatchProperties /* not a class type */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for MatchedPlayers */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for MatchedPlayers */
// Alloc allocates a new instance without initialization.
func (mc _MatchedPlayersClass) Alloc() MatchedPlayers {
	rv := objc.Send[MatchedPlayers](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MatchedPlayersClass) New() MatchedPlayers {
	rv := objc.Send[MatchedPlayers](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MatchedPlayers) Init() MatchedPlayers {
	rv := objc.Send[MatchedPlayers](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MatchedPlayers) Autorelease() MatchedPlayers {
	rv := objc.Send[MatchedPlayers](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMatchedPlayers creates a new MatchedPlayers instance.
func NewMatchedPlayers() MatchedPlayers {
	return getMatchedPlayersClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for MatchedPlayers */
// An object that represents matchmaking results, including the players that join the match and their properties that matchmaking rules uses.
//
// If the  and   properties are , Game Center didn’t use matchmaking rules to find the players. For more information, see .


// An object that represents matchmaking results, including the players that join the match and their properties that matchmaking rules uses.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchedPlayers
type MatchedPlayers struct {
	objectivec.Object
}

// MatchedPlayersFrom constructs a [MatchedPlayers] from an unsafe.Pointer.
//
// An object that represents matchmaking results, including the players that join the match and their properties that matchmaking rules uses.
func MatchedPlayersFrom(ptr unsafe.Pointer) MatchedPlayers {
	return MatchedPlayers{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for MatchedPlayers *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for MatchedPlayers */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for MatchedPlayers */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for MatchedPlayers */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for MatchedPlayers */

// The properties for other players that matchmaking rules uses to find players, with some additions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchedPlayers/playerProperties
func (m_ MatchedPlayers) PlayerProperties() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](m_.ID, objc.Sel("playerProperties"))
	return rv
}/* debug [instance_properties/getter]: playerProperties */


// The players that join the match.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchedPlayers/players
func (m_ MatchedPlayers) Players() []Player {
	rv := objc.Send[[]Player](m_.ID, objc.Sel("players"))
	return rv
}/* debug [instance_properties/getter]: players */


// The local player’s properties that matchmaking rules uses to find the players, with some additions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKMatchedPlayers/properties
func (m_ MatchedPlayers) Properties() MatchProperties /* not a class type */ {
	rv := objc.Send[MatchProperties](m_.ID, objc.Sel("properties"))
	return rv
}/* debug [instance_properties/getter]: properties */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKMatchedPlayers */



