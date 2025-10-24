// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class GKPlayer */


/* debug [class_header]: Header for GKPlayer */
// The class instance for the [Player] class.
var (
	PlayerClass     _PlayerClass
	PlayerClassOnce sync.Once
)

func getPlayerClass() _PlayerClass {
	PlayerClassOnce.Do(func() {
		PlayerClass = _PlayerClass{objc.GetClass("GKPlayer")}
	})
	return PlayerClass
}

type _PlayerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Player */
// An interface definition for the [Player] class.
type IPlayer interface {
	IBasePlayer
	
/* debug [class_interface_properties]: Properties for Player */
	// properties:
	Alias() objc.IObject /* cross-framework: NSString */
	DisplayName() objc.IObject /* cross-framework: NSString */
	GamePlayerID() objc.IObject /* cross-framework: NSString */
	GuestIdentifier() objc.IObject /* cross-framework: NSString */
	IsFriend() bool
	IsInvitable() bool
	PlayerID() objc.IObject /* cross-framework: NSString */
	TeamPlayerID() objc.IObject /* cross-framework: NSString */
	GKPlayerIDNoLongerAvailable() objc.IObject /* cross-framework: NSString */
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Player */
	// methods:
	LoadPhotoForSizeWithCompletionHandler(size PhotoSize, completionHandler unsafe.Pointer)
	ScopedIDsArePersistent() bool
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Player */
// Alloc allocates a new instance without initialization.
func (pc _PlayerClass) Alloc() Player {
	rv := objc.Send[Player](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PlayerClass) New() Player {
	rv := objc.Send[Player](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Player) Init() Player {
	rv := objc.Send[Player](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Player) Autorelease() Player {
	rv := objc.Send[Player](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPlayer creates a new Player instance.
func NewPlayer() Player {
	return getPlayerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Player */
// A remote player who the local player running your game can invite and communicate with through Game Center.
//
// Before using Game Center for the first time, players create a single account that identifies them across all Game Center games. The player only needs to sign in to Game Center once per device to start using GameKit features in your game. A player sets a nickname and avatar in their account that provide a consistent and familiar look in your game. Game Center then uses the account to record leaderboard scores and achievements, and to start games with other players. In your code, represents remote or other players who the local player running your app can invite and communicate with. is also the superclass for the local player class that provides common data and methods for all players. For example, use the property to get the nickname for a player. To load the player avatars, use the method. To create a guest player who doesn’t have a Game Center account, use the method. GameKit treats guest players similar to Game Center players except they can’t earn achievements, post to leaderboards, or participate in challenges. Use the property as a unique identifier for just your game, and the property as a unique identifier for all games that you offer through your developer account. For more information, see .


// A remote player who the local player running your game can invite and communicate with through Game Center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPlayer
type Player struct {
	BasePlayer
}

// PlayerFrom constructs a [Player] from an unsafe.Pointer.
//
// A remote player who the local player running your game can invite and communicate with through Game Center.
func PlayerFrom(ptr unsafe.Pointer) Player {
	return Player{
		BasePlayer: BasePlayerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Player *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Player */

// Creates a guest player with the specified identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPlayer/anonymousGuestPlayer(withIdentifier:)
func (pc _PlayerClass) AnonymousGuestPlayerWithIdentifier(guestIdentifier objc.IObject /* cross-framework: NSString */) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("anonymousGuestPlayerWithIdentifier:"), guestIdentifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=AnonymousGuestPlayerWithIdentifier) */


// Loads information about a list of players from Game Center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPlayer/loadPlayers(forIdentifiers:withCompletionHandler:)
func (pc _PlayerClass) LoadPlayersForIdentifiersWithCompletionHandler(identifiers []string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("loadPlayersForIdentifiers:withCompletionHandler:"), identifiers, completionHandler)
}/* debug [class_methods/method]: Class method for%!(EXTRA string=LoadPlayersForIdentifiersWithCompletionHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Player */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Player */

// Loads a photo of the player from Game Center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPlayer/loadPhoto(for:withCompletionHandler:)
func (p_ Player) LoadPhotoForSizeWithCompletionHandler(size PhotoSize, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("loadPhotoForSize:withCompletionHandler:"), size, completionHandler)
}/* debug [instance_methods/method]: LoadPhotoForSizeWithCompletionHandler */


// Returns a Boolean value depending on whether the player identifiers are persistent across game instances or unique to the game instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPlayer/scopedIDsArePersistent()
func (p_ Player) ScopedIDsArePersistent() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("scopedIDsArePersistent"))
	return rv
}/* debug [instance_methods/method]: ScopedIDsArePersistent */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Player */

// A string the player chooses to identify themself to other players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPlayer/alias
func (p_ Player) Alias() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("alias"))
	return rv
}/* debug [instance_properties/getter]: alias */


// A string to display for the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPlayer/displayName
func (p_ Player) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("displayName"))
	return rv
}/* debug [instance_properties/getter]: displayName */


// A unique identifier for a player of the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPlayer/gamePlayerID
func (p_ Player) GamePlayerID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("gamePlayerID"))
	return rv
}/* debug [instance_properties/getter]: gamePlayerID */


// A developer-created string that identifies a guest player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPlayer/guestIdentifier
func (p_ Player) GuestIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("guestIdentifier"))
	return rv
}/* debug [instance_properties/getter]: guestIdentifier */


// A Boolean value that indicates whether the player is a friend of the local player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPlayer/isFriend
func (p_ Player) IsFriend() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isFriend"))
	return rv
}/* debug [instance_properties/getter]: isFriend */


// A Boolean value that indicates whether the local player can send an invitation to the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPlayer/isInvitable
func (p_ Player) IsInvitable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isInvitable"))
	return rv
}/* debug [instance_properties/getter]: isInvitable */


// A unique identifier for a player of the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPlayer/playerID
func (p_ Player) PlayerID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("playerID"))
	return rv
}/* debug [instance_properties/getter]: playerID */


// A unique identifier for a player of all the games that you distribute using your developer account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPlayer/teamPlayerID
func (p_ Player) TeamPlayerID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("teamPlayerID"))
	return rv
}/* debug [instance_properties/getter]: teamPlayerID */


// A constant for a player ID that’s no longer available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayeridnolongeravailable
func (p_ Player) GKPlayerIDNoLongerAvailable() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("GKPlayerIDNoLongerAvailable"))
	return rv
}/* debug [instance_properties/getter]: GKPlayerIDNoLongerAvailable */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKPlayer */



