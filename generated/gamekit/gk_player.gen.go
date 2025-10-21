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
func (p_ Player) Alias() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("alias"))
	return rv
}



