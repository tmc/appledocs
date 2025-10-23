// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [LocalPlayer] class.
var (
	LocalPlayerClass     _LocalPlayerClass
	LocalPlayerClassOnce sync.Once
)

func getLocalPlayerClass() _LocalPlayerClass {
	LocalPlayerClassOnce.Do(func() {
		LocalPlayerClass = _LocalPlayerClass{objc.GetClass("GKLocalPlayer")}
	})
	return LocalPlayerClass
}

type _LocalPlayerClass struct {
	class objc.Class
}

// An interface definition for the [LocalPlayer] class.
type ILocalPlayer interface {
	IPlayer
	Authenticated() bool
	MultiplayerGamingRestricted() bool
	PersonalizedCommunicationRestricted() bool
	AuthenticateHandler() unsafe.Pointer
	SetAuthenticateHandler(value unsafe.Pointer)
	IsAuthenticated() bool
	SetIsAuthenticated(value bool)
	IsMultiplayerGamingRestricted() bool
	SetIsMultiplayerGamingRestricted(value bool)
	IsPersonalizedCommunicationRestricted() bool
	SetIsPersonalizedCommunicationRestricted(value bool)
	IsPresentingFriendRequestViewController() bool
	SetIsPresentingFriendRequestViewController(value bool)
	IsUnderage() bool
	SetIsUnderage(value bool)
}

// The local player who signs in to Game Center on the device running the game.
//
// Only one player can sign in to Game Center on a device at a time and that player is the . Before you can start a game that uses GameKit features, verify that the local player signs in to their Game Center account. You set the handler of the local player shared instance using the property. Then implement this method to handle the multiple times GameKit invokes it during the sign-in process. If the local player needs to create an account or sign in, GameKit provides a view controller that you present to the local player. If the local player successfully signs in, determine whether they have any account restrictions and adjust your game accordingly. For more information about the initialization of the local player, see . After the local player signs in, their account data and GameKit features are available. You can display the local player’s nickname and avatar, access their recent players and friends, and load their leaderboards and achievements. You can also register a listener object that GameKit calls when the local player sends or accepts invitations to play with others.


// The local player who signs in to Game Center on the device running the game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer
type LocalPlayer struct {
	Player
}

// LocalPlayerFrom constructs a [LocalPlayer] from an unsafe.Pointer.
//
// The local player who signs in to Game Center on the device running the game.
func LocalPlayerFrom(ptr unsafe.Pointer) LocalPlayer {
	return LocalPlayer{
		Player: PlayerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (lc _LocalPlayerClass) Alloc() LocalPlayer {
	rv := objc.Send[LocalPlayer](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (lc _LocalPlayerClass) New() LocalPlayer {
	rv := objc.Send[LocalPlayer](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LocalPlayer) Init() LocalPlayer {
	rv := objc.Send[LocalPlayer](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LocalPlayer) Autorelease() LocalPlayer {
	rv := objc.Send[LocalPlayer](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLocalPlayer creates a new LocalPlayer instance.
func NewLocalPlayer() LocalPlayer {
	return getLocalPlayerClass().New()
}



// A Boolean value that indicates whether a local player has signed in to Game Center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/isAuthenticated
func (l_ LocalPlayer) Authenticated() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("authenticated"))
	return rv
}


// A Boolean value that indicates whether the player can join multiplayer games.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/isMultiplayerGamingRestricted
func (l_ LocalPlayer) MultiplayerGamingRestricted() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("multiplayerGamingRestricted"))
	return rv
}


// A Boolean value that indicates whether the player can use personalized communication on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/isPersonalizedCommunicationRestricted
func (l_ LocalPlayer) PersonalizedCommunicationRestricted() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("personalizedCommunicationRestricted"))
	return rv
}


// A handler that GameKit calls while initializing the local player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/authenticatehandler
func (l_ LocalPlayer) AuthenticateHandler() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("authenticateHandler"))
	return rv
}


// A handler that GameKit calls while initializing the local player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/authenticatehandler
func (l_ LocalPlayer) SetAuthenticateHandler(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setAuthenticateHandler:"), value)
}


// A Boolean value that indicates whether a local player has signed in to Game Center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/isauthenticated
func (l_ LocalPlayer) IsAuthenticated() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isAuthenticated"))
	return rv
}


// A Boolean value that indicates whether a local player has signed in to Game Center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/isauthenticated
func (l_ LocalPlayer) SetIsAuthenticated(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsAuthenticated:"), value)
}


// A Boolean value that indicates whether the player can join multiplayer games.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/ismultiplayergamingrestricted
func (l_ LocalPlayer) IsMultiplayerGamingRestricted() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isMultiplayerGamingRestricted"))
	return rv
}


// A Boolean value that indicates whether the player can join multiplayer games.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/ismultiplayergamingrestricted
func (l_ LocalPlayer) SetIsMultiplayerGamingRestricted(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsMultiplayerGamingRestricted:"), value)
}


// A Boolean value that indicates whether the player can use personalized communication on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/ispersonalizedcommunicationrestricted
func (l_ LocalPlayer) IsPersonalizedCommunicationRestricted() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isPersonalizedCommunicationRestricted"))
	return rv
}


// A Boolean value that indicates whether the player can use personalized communication on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/ispersonalizedcommunicationrestricted
func (l_ LocalPlayer) SetIsPersonalizedCommunicationRestricted(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsPersonalizedCommunicationRestricted:"), value)
}


// A Boolean value that indicates whether your game presents the friends request view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/ispresentingfriendrequestviewcontroller
func (l_ LocalPlayer) IsPresentingFriendRequestViewController() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isPresentingFriendRequestViewController"))
	return rv
}


// A Boolean value that indicates whether your game presents the friends request view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/ispresentingfriendrequestviewcontroller
func (l_ LocalPlayer) SetIsPresentingFriendRequestViewController(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsPresentingFriendRequestViewController:"), value)
}


// A Boolean value that indicates whether the local player is underage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/isunderage
func (l_ LocalPlayer) IsUnderage() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isUnderage"))
	return rv
}


// A Boolean value that indicates whether the local player is underage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/isunderage
func (l_ LocalPlayer) SetIsUnderage(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsUnderage:"), value)
}



