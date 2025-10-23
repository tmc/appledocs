// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
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
	// properties:
	AuthenticateHandler() func(error objc.ID) /* already interface */
	SetAuthenticateHandler(value func(error objc.ID) /* already interface */)
	Authenticated() bool /* primitive/slice/pointer. */
	MultiplayerGamingRestricted() bool /* primitive/slice/pointer. */
	PersonalizedCommunicationRestricted() bool /* primitive/slice/pointer. */
	IsPresentingFriendRequestViewController() bool /* primitive/slice/pointer. */
	IsAuthenticated() bool /* primitive/slice/pointer. */
	SetIsAuthenticated(value bool /* primitive/slice/pointer. */)
	IsMultiplayerGamingRestricted() bool /* primitive/slice/pointer. */
	SetIsMultiplayerGamingRestricted(value bool /* primitive/slice/pointer. */)
	IsPersonalizedCommunicationRestricted() bool /* primitive/slice/pointer. */
	SetIsPersonalizedCommunicationRestricted(value bool /* primitive/slice/pointer. */)
	IsUnderage() bool /* primitive/slice/pointer. */
	SetIsUnderage(value bool /* primitive/slice/pointer. */)
	// methods:
	FetchItemsForIdentityVerificationSignature(completionHandler unsafe.Pointer)
	LoadFriends(completionHandler unsafe.Pointer)
	SaveGameDataWithNameCompletionHandler(data foundation.objc.IObject /* cross-framework NSData */, name string /* primitive/slice/pointer. */, handler unsafe.Pointer)
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



// Generates a signature that you can use to authenticate the local player on your own server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/fetchItems(forIdentityVerificationSignature:)
func (l_ LocalPlayer) FetchItemsForIdentityVerificationSignature(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("fetchItemsForIdentityVerificationSignature:"), completionHandler)
}


// Loads the local player’s friends list if the local player and their friends grant access.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/loadFriends(_:)
func (l_ LocalPlayer) LoadFriends(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("loadFriends:"), completionHandler)
}


// Saves game data with the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/saveGameData(_:withName:completionHandler:)
func (l_ LocalPlayer) SaveGameDataWithNameCompletionHandler(data foundation.objc.IObject /* cross-framework NSData */, name string /* primitive/slice/pointer. */, handler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("saveGameData:withName:completionHandler:"), data, objc.String(name), handler)
}


// A handler that GameKit calls while initializing the local player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/authenticateHandler
func (l_ LocalPlayer) AuthenticateHandler() func(error objc.ID) /* already interface */ {
	rv := objc.Send[func(error objc.ID)](l_.ID, objc.Sel("authenticateHandler"))
	return rv
}


// A handler that GameKit calls while initializing the local player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/authenticateHandler
func (l_ LocalPlayer) SetAuthenticateHandler(value func(error objc.ID) /* already interface */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setAuthenticateHandler:"), value)
}


// A Boolean value that indicates whether a local player has signed in to Game Center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/isAuthenticated
func (l_ LocalPlayer) Authenticated() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](l_.ID, objc.Sel("authenticated"))
	return rv
}


// A Boolean value that indicates whether the player can join multiplayer games.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/isMultiplayerGamingRestricted
func (l_ LocalPlayer) MultiplayerGamingRestricted() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](l_.ID, objc.Sel("multiplayerGamingRestricted"))
	return rv
}


// A Boolean value that indicates whether the player can use personalized communication on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/isPersonalizedCommunicationRestricted
func (l_ LocalPlayer) PersonalizedCommunicationRestricted() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](l_.ID, objc.Sel("personalizedCommunicationRestricted"))
	return rv
}


// A Boolean value that indicates whether your game presents the friends request view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/isPresentingFriendRequestViewController
func (l_ LocalPlayer) IsPresentingFriendRequestViewController() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](l_.ID, objc.Sel("isPresentingFriendRequestViewController"))
	return rv
}


// A Boolean value that indicates whether a local player has signed in to Game Center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/isauthenticated
func (l_ LocalPlayer) IsAuthenticated() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](l_.ID, objc.Sel("isAuthenticated"))
	return rv
}


// A Boolean value that indicates whether a local player has signed in to Game Center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/isauthenticated
func (l_ LocalPlayer) SetIsAuthenticated(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsAuthenticated:"), value)
}


// A Boolean value that indicates whether the player can join multiplayer games.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/ismultiplayergamingrestricted
func (l_ LocalPlayer) IsMultiplayerGamingRestricted() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](l_.ID, objc.Sel("isMultiplayerGamingRestricted"))
	return rv
}


// A Boolean value that indicates whether the player can join multiplayer games.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/ismultiplayergamingrestricted
func (l_ LocalPlayer) SetIsMultiplayerGamingRestricted(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsMultiplayerGamingRestricted:"), value)
}


// A Boolean value that indicates whether the player can use personalized communication on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/ispersonalizedcommunicationrestricted
func (l_ LocalPlayer) IsPersonalizedCommunicationRestricted() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](l_.ID, objc.Sel("isPersonalizedCommunicationRestricted"))
	return rv
}


// A Boolean value that indicates whether the player can use personalized communication on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/ispersonalizedcommunicationrestricted
func (l_ LocalPlayer) SetIsPersonalizedCommunicationRestricted(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsPersonalizedCommunicationRestricted:"), value)
}


// A Boolean value that indicates whether the local player is underage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/isunderage
func (l_ LocalPlayer) IsUnderage() bool /* primitive/slice/pointer. */ {
	rv := objc.Send[bool](l_.ID, objc.Sel("isUnderage"))
	return rv
}


// A Boolean value that indicates whether the local player is underage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/isunderage
func (l_ LocalPlayer) SetIsUnderage(value bool /* primitive/slice/pointer. */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsUnderage:"), value)
}



