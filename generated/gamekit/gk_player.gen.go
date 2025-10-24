// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [Player] class.
type IPlayer interface {
	IBasePlayer
	// properties:
	Alias() objc.IObject /* cross-framework: NSString */
	DisplayName() objc.IObject /* cross-framework: NSString */
	SetDisplayName(value objc.IObject /* cross-framework: NSString */)
	GamePlayerID() objc.IObject /* cross-framework: NSString */
	SetGamePlayerID(value objc.IObject /* cross-framework: NSString */)
	GuestIdentifier() objc.IObject /* cross-framework: NSString */
	SetGuestIdentifier(value objc.IObject /* cross-framework: NSString */)
	IsFriend() bool
	SetIsFriend(value bool)
	IsInvitable() bool
	SetIsInvitable(value bool)
	PlayerID() objc.IObject /* cross-framework: NSString */
	SetPlayerID(value objc.IObject /* cross-framework: NSString */)
	TeamPlayerID() objc.IObject /* cross-framework: NSString */
	SetTeamPlayerID(value objc.IObject /* cross-framework: NSString */)
	GKPlayerIDNoLongerAvailable() objc.IObject /* cross-framework: NSString */
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (pc _PlayerClass) Alloc() Player {
	rv := objc.Send[Player](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// A string the player chooses to identify themself to other players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPlayer/alias
func (p_ Player) Alias() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("alias"))
	return rv
}


// A string to display for the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/displayname
func (p_ Player) DisplayName() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("displayName"))
	return rv
}


// A string to display for the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/displayname
func (p_ Player) SetDisplayName(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDisplayName:"), value)
}


// A unique identifier for a player of the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/gameplayerid
func (p_ Player) GamePlayerID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("gamePlayerID"))
	return rv
}


// A unique identifier for a player of the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/gameplayerid
func (p_ Player) SetGamePlayerID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGamePlayerID:"), value)
}


// A developer-created string that identifies a guest player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/guestidentifier
func (p_ Player) GuestIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("guestIdentifier"))
	return rv
}


// A developer-created string that identifies a guest player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/guestidentifier
func (p_ Player) SetGuestIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGuestIdentifier:"), value)
}


// A Boolean value that indicates whether the player is a friend of the local player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/isfriend
func (p_ Player) IsFriend() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isFriend"))
	return rv
}


// A Boolean value that indicates whether the player is a friend of the local player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/isfriend
func (p_ Player) SetIsFriend(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsFriend:"), value)
}


// A Boolean value that indicates whether the local player can send an invitation to the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/isinvitable
func (p_ Player) IsInvitable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isInvitable"))
	return rv
}


// A Boolean value that indicates whether the local player can send an invitation to the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/isinvitable
func (p_ Player) SetIsInvitable(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsInvitable:"), value)
}


// A unique identifier for a player of the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/playerid
func (p_ Player) PlayerID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("playerID"))
	return rv
}


// A unique identifier for a player of the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/playerid
func (p_ Player) SetPlayerID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlayerID:"), value)
}


// A unique identifier for a player of all the games that you distribute using your developer account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/teamplayerid
func (p_ Player) TeamPlayerID() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("teamPlayerID"))
	return rv
}


// A unique identifier for a player of all the games that you distribute using your developer account.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/teamplayerid
func (p_ Player) SetTeamPlayerID(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTeamPlayerID:"), value)
}


// A constant for a player ID that’s no longer available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayeridnolongeravailable
func (p_ Player) GKPlayerIDNoLongerAvailable() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("GKPlayerIDNoLongerAvailable"))
	return rv
}



