// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
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
	Alias() string
	DisplayName() string
	SetDisplayName(value string)
	GamePlayerID() string
	SetGamePlayerID(value string)
	GuestIdentifier() string
	SetGuestIdentifier(value string)
	IsFriend() bool
	SetIsFriend(value bool)
	IsInvitable() bool
	SetIsInvitable(value bool)
	PlayerID() string
	SetPlayerID(value string)
	TeamPlayerID() string
	SetTeamPlayerID(value string)
	GKPlayerIDNoLongerAvailable() string
}

// A remote player who the local player running your game can invite and communicate with through Game Center.
//
// Before using Game Center for the first time, players create a single account that identifies them across all Game Center games. The player only needs to sign in to Game Center once per device to start using GameKit features in your game. A player sets a nickname and avatar in their account that provide a consistent and familiar look in your game. Game Center then uses the account to record leaderboard scores and achievements, and to start games with other players. In your code, represents remote or other players who the local player running your app can invite and communicate with. is also the superclass for the local player class that provides common data and methods for all players. For example, use the property to get the nickname for a player. To load the player avatars, use the method. To create a guest player who doesn’t have a Game Center account, use the method. GameKit treats guest players similar to Game Center players except they can’t earn achievements, post to leaderboards, or participate in challenges. Use the property as a unique identifier for just your game, and the property as a unique identifier for all games that you offer through your developer account. For more information, see .
//
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
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKPlayer/alias
func (p_ Player) Alias() string {
	rv := objc.Send[string](p_.ID, objc.Sel("alias"))
	return rv
}

// A string to display for the player.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/displayname
func (p_ Player) DisplayName() string {
	rv := objc.Send[string](p_.ID, objc.Sel("displayName"))
	return rv
}


// SetDisplayName sets the value of the displayName property.
// A string to display for the player.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/displayname
func (p_ Player) SetDisplayName(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDisplayName:"), objc.String(value))
}

// A unique identifier for a player of the game.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/gameplayerid
func (p_ Player) GamePlayerID() string {
	rv := objc.Send[string](p_.ID, objc.Sel("gamePlayerID"))
	return rv
}


// SetGamePlayerID sets the value of the gamePlayerID property.
// A unique identifier for a player of the game.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/gameplayerid
func (p_ Player) SetGamePlayerID(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGamePlayerID:"), objc.String(value))
}

// A developer-created string that identifies a guest player.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/guestidentifier
func (p_ Player) GuestIdentifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("guestIdentifier"))
	return rv
}


// SetGuestIdentifier sets the value of the guestIdentifier property.
// A developer-created string that identifies a guest player.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/guestidentifier
func (p_ Player) SetGuestIdentifier(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setGuestIdentifier:"), objc.String(value))
}

// A Boolean value that indicates whether the player is a friend of the local player.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/isfriend
func (p_ Player) IsFriend() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isFriend"))
	return rv
}


// SetIsFriend sets the value of the isFriend property.
// A Boolean value that indicates whether the player is a friend of the local player.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/isfriend
func (p_ Player) SetIsFriend(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsFriend:"), value)
}

// A Boolean value that indicates whether the local player can send an invitation to the player.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/isinvitable
func (p_ Player) IsInvitable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isInvitable"))
	return rv
}


// SetIsInvitable sets the value of the isInvitable property.
// A Boolean value that indicates whether the local player can send an invitation to the player.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/isinvitable
func (p_ Player) SetIsInvitable(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsInvitable:"), value)
}

// A unique identifier for a player of the game.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/playerid
func (p_ Player) PlayerID() string {
	rv := objc.Send[string](p_.ID, objc.Sel("playerID"))
	return rv
}


// SetPlayerID sets the value of the playerID property.
// A unique identifier for a player of the game.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/playerid
func (p_ Player) SetPlayerID(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPlayerID:"), objc.String(value))
}

// A unique identifier for a player of all the games that you distribute using your developer account.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/teamplayerid
func (p_ Player) TeamPlayerID() string {
	rv := objc.Send[string](p_.ID, objc.Sel("teamPlayerID"))
	return rv
}


// SetTeamPlayerID sets the value of the teamPlayerID property.
// A unique identifier for a player of all the games that you distribute using your developer account.

//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayer/teamplayerid
func (p_ Player) SetTeamPlayerID(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTeamPlayerID:"), objc.String(value))
}

// A constant for a player ID that’s no longer available.
//
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkplayeridnolongeravailable
func (p_ Player) GKPlayerIDNoLongerAvailable() string {
	rv := objc.Send[string](p_.ID, objc.Sel("GKPlayerIDNoLongerAvailable"))
	return rv
}



