// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class GKLocalPlayer */


/* debug [class_header]: Header for GKLocalPlayer */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for LocalPlayer */
// An interface definition for the [LocalPlayer] class.
type ILocalPlayer interface {
	IPlayer
	
/* debug [class_interface_properties]: Properties for LocalPlayer */
	// properties:
	AuthenticateHandler() func(unsafe.Pointer)
	SetAuthenticateHandler(value func(unsafe.Pointer))
	Friends() []string
	Authenticated() bool
	MultiplayerGamingRestricted() bool
	PersonalizedCommunicationRestricted() bool
	IsPresentingFriendRequestViewController() bool
	Underage() bool
	IsAuthenticated() bool
	SetIsAuthenticated(value bool)
	IsMultiplayerGamingRestricted() bool
	SetIsMultiplayerGamingRestricted(value bool)
	IsPersonalizedCommunicationRestricted() bool
	SetIsPersonalizedCommunicationRestricted(value bool)
	IsUnderage() bool
	SetIsUnderage(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for LocalPlayer */
	// methods:
	DeleteSavedGamesWithNameCompletionHandler(name objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer)
	FetchItemsForIdentityVerificationSignature(completionHandler unsafe.Pointer)
	FetchSavedGamesWithCompletionHandler(handler unsafe.Pointer)
	LoadChallengableFriendsWithCompletionHandler(completionHandler unsafe.Pointer)
	LoadDefaultLeaderboardIdentifierWithCompletionHandler(completionHandler unsafe.Pointer)
	LoadFriends(completionHandler unsafe.Pointer)
	LoadFriendsWithIdentifiersCompletionHandler(identifiers []string, completionHandler unsafe.Pointer)
	LoadFriendsAuthorizationStatus(completionHandler unsafe.Pointer)
	LoadRecentPlayersWithCompletionHandler(completionHandler unsafe.Pointer)
	PresentFriendRequestCreatorFromWindowError(window appkit.Window, error_ unsafe.Pointer) bool
	RegisterListener(listener unsafe.Pointer)
	ResolveConflictingSavedGamesWithDataCompletionHandler(conflictingSavedGames []SavedGame, data objc.IObject /* cross-framework: NSData */, handler unsafe.Pointer)
	SaveGameDataWithNameCompletionHandler(data objc.IObject /* cross-framework: NSData */, name objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer)
	SetDefaultLeaderboardIdentifierCompletionHandler(leaderboardIdentifier objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer)
	UnregisterAllListeners()
	UnregisterListener(listener unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for LocalPlayer */
// Alloc allocates a new instance without initialization.
func (lc _LocalPlayerClass) Alloc() LocalPlayer {
	rv := objc.Send[LocalPlayer](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for LocalPlayer */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for LocalPlayer *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for LocalPlayer */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for LocalPlayer */

// The shared instance of the local player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/local-1mzi0
func (lc _LocalPlayerClass) Local() LocalPlayer {
	rv := objc.Send[LocalPlayer](objc.ID(lc.class), objc.Sel("local"))
	return rv
}/* debug [class_properties_class/property]: local */

// The shared instance of the local player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/local-oaa8
func (lc _LocalPlayerClass) LocalPlayer() LocalPlayer {
	rv := objc.Send[LocalPlayer](objc.ID(lc.class), objc.Sel("localPlayer"))
	return rv
}/* debug [class_properties_class/property]: localPlayer */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for LocalPlayer */

// Deletes saved games with the specified filename.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/deleteSavedGames(withName:completionHandler:)
func (l_ LocalPlayer) DeleteSavedGamesWithNameCompletionHandler(name objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("deleteSavedGamesWithName:completionHandler:"), name, handler)
}/* debug [instance_methods/method]: DeleteSavedGamesWithNameCompletionHandler */


// Generates a signature that you can use to authenticate the local player on your own server.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/fetchItems(forIdentityVerificationSignature:)
func (l_ LocalPlayer) FetchItemsForIdentityVerificationSignature(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("fetchItemsForIdentityVerificationSignature:"), completionHandler)
}/* debug [instance_methods/method]: FetchItemsForIdentityVerificationSignature */


// Retrieves all available saved games.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/fetchSavedGames(completionHandler:)
func (l_ LocalPlayer) FetchSavedGamesWithCompletionHandler(handler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("fetchSavedGamesWithCompletionHandler:"), handler)
}/* debug [instance_methods/method]: FetchSavedGamesWithCompletionHandler */


// Loads players to whom the local player can issue a challenge.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/loadChallengableFriends(completionHandler:)
func (l_ LocalPlayer) LoadChallengableFriendsWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("loadChallengableFriendsWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: LoadChallengableFriendsWithCompletionHandler */


// Loads the identifier for the local player’s default leaderboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/loadDefaultLeaderboardIdentifier(completionHandler:)
func (l_ LocalPlayer) LoadDefaultLeaderboardIdentifierWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("loadDefaultLeaderboardIdentifierWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: LoadDefaultLeaderboardIdentifierWithCompletionHandler */


// Loads the local player’s friends list if the local player and their friends grant access.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/loadFriends(_:)
func (l_ LocalPlayer) LoadFriends(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("loadFriends:"), completionHandler)
}/* debug [instance_methods/method]: LoadFriends */


// Loads the player’s friends list, scoped by the identifiers, if the player and their friends grant access.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/loadFriends(identifiedBy:completionHandler:)
func (l_ LocalPlayer) LoadFriendsWithIdentifiersCompletionHandler(identifiers []string, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("loadFriendsWithIdentifiers:completionHandler:"), identifiers, completionHandler)
}/* debug [instance_methods/method]: LoadFriendsWithIdentifiersCompletionHandler */


// Returns whether the player authorizes your game to access their friends list.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/loadFriendsAuthorizationStatus(_:)
func (l_ LocalPlayer) LoadFriendsAuthorizationStatus(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("loadFriendsAuthorizationStatus:"), completionHandler)
}/* debug [instance_methods/method]: LoadFriendsAuthorizationStatus */


// Loads players from the friends list or players that recently participated in a game with the local player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/loadRecentPlayers(completionHandler:)
func (l_ LocalPlayer) LoadRecentPlayersWithCompletionHandler(completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("loadRecentPlayersWithCompletionHandler:"), completionHandler)
}/* debug [instance_methods/method]: LoadRecentPlayersWithCompletionHandler */


// Opens the Messages app with a sheet for the player to request friends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/presentFriendRequestCreator(from:)-7clh6
func (l_ LocalPlayer) PresentFriendRequestCreatorFromWindowError(window appkit.Window, error_ unsafe.Pointer) bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("presentFriendRequestCreatorFromWindow:error:"), window, error_)
	return rv
}/* debug [instance_methods/method]: PresentFriendRequestCreatorFromWindowError */


// Registers a listener for a particular event.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/register(_:)
func (l_ LocalPlayer) RegisterListener(listener unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("registerListener:"), listener)
}/* debug [instance_methods/method]: RegisterListener */


// Replaces duplicate saved games that use the same filename with one file containing the specified game data.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/resolveConflictingSavedGames(_:with:completionHandler:)
func (l_ LocalPlayer) ResolveConflictingSavedGamesWithDataCompletionHandler(conflictingSavedGames []SavedGame, data objc.IObject /* cross-framework: NSData */, handler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("resolveConflictingSavedGames:withData:completionHandler:"), conflictingSavedGames, data, handler)
}/* debug [instance_methods/method]: ResolveConflictingSavedGamesWithDataCompletionHandler */


// Saves game data with the specified name.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/saveGameData(_:withName:completionHandler:)
func (l_ LocalPlayer) SaveGameDataWithNameCompletionHandler(data objc.IObject /* cross-framework: NSData */, name objc.IObject /* cross-framework: NSString */, handler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("saveGameData:withName:completionHandler:"), data, name, handler)
}/* debug [instance_methods/method]: SaveGameDataWithNameCompletionHandler */


// Sets the local player’s default leaderboard.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/setDefaultLeaderboardIdentifier(_:completionHandler:)
func (l_ LocalPlayer) SetDefaultLeaderboardIdentifierCompletionHandler(leaderboardIdentifier objc.IObject /* cross-framework: NSString */, completionHandler unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDefaultLeaderboardIdentifier:completionHandler:"), leaderboardIdentifier, completionHandler)
}/* debug [instance_methods/method]: SetDefaultLeaderboardIdentifierCompletionHandler */


// Unregisters all listeners in your game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/unregisterAllListeners()
func (l_ LocalPlayer) UnregisterAllListeners() {
	objc.Send[objc.ID](l_.ID, objc.Sel("unregisterAllListeners"))
}/* debug [instance_methods/method]: UnregisterAllListeners */


// Unregisters a listener object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/unregisterListener(_:)
func (l_ LocalPlayer) UnregisterListener(listener unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("unregisterListener:"), listener)
}/* debug [instance_methods/method]: UnregisterListener */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for LocalPlayer */

// A handler that GameKit calls while initializing the local player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/authenticateHandler
func (l_ LocalPlayer) AuthenticateHandler() func(unsafe.Pointer) {
	rv := objc.Send[func(unsafe.Pointer)](l_.ID, objc.Sel("authenticateHandler"))
	return rv
}/* debug [instance_properties/getter]: authenticateHandler */


// A handler that GameKit calls while initializing the local player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/authenticateHandler
func (l_ LocalPlayer) SetAuthenticateHandler(value func(unsafe.Pointer)) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setAuthenticateHandler:"), value)
}/* debug [instance_properties/setter]: authenticateHandler */


// The player identifiers for the local player’s friends.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/friends
func (l_ LocalPlayer) Friends() []string {
	rv := objc.Send[[]string](l_.ID, objc.Sel("friends"))
	return rv
}/* debug [instance_properties/getter]: friends */


// A Boolean value that indicates whether a local player has signed in to Game Center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/isAuthenticated
func (l_ LocalPlayer) Authenticated() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("authenticated"))
	return rv
}/* debug [instance_properties/getter]: authenticated */


// A Boolean value that indicates whether the player can join multiplayer games.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/isMultiplayerGamingRestricted
func (l_ LocalPlayer) MultiplayerGamingRestricted() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("multiplayerGamingRestricted"))
	return rv
}/* debug [instance_properties/getter]: multiplayerGamingRestricted */


// A Boolean value that indicates whether the player can use personalized communication on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/isPersonalizedCommunicationRestricted
func (l_ LocalPlayer) PersonalizedCommunicationRestricted() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("personalizedCommunicationRestricted"))
	return rv
}/* debug [instance_properties/getter]: personalizedCommunicationRestricted */


// A Boolean value that indicates whether your game presents the friends request view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/isPresentingFriendRequestViewController
func (l_ LocalPlayer) IsPresentingFriendRequestViewController() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isPresentingFriendRequestViewController"))
	return rv
}/* debug [instance_properties/getter]: isPresentingFriendRequestViewController */


// A Boolean value that indicates whether the local player is underage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/isUnderage
func (l_ LocalPlayer) Underage() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("underage"))
	return rv
}/* debug [instance_properties/getter]: underage */


// The shared instance of the local player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/local-1mzi0
func (l_ LocalPlayer) Local() IGKLocalPlayer {
	rv := objc.Send[LocalPlayer](l_.ID, objc.Sel("local"))
	return rv
}/* debug [instance_properties/getter]: local */


// The shared instance of the local player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLocalPlayer/local-oaa8
func (l_ LocalPlayer) LocalPlayer() IGKLocalPlayer {
	rv := objc.Send[LocalPlayer](l_.ID, objc.Sel("localPlayer"))
	return rv
}/* debug [instance_properties/getter]: localPlayer */


// A Boolean value that indicates whether a local player has signed in to Game Center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/isauthenticated
func (l_ LocalPlayer) IsAuthenticated() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isAuthenticated"))
	return rv
}/* debug [instance_properties/getter]: isAuthenticated */


// A Boolean value that indicates whether a local player has signed in to Game Center.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/isauthenticated
func (l_ LocalPlayer) SetIsAuthenticated(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsAuthenticated:"), value)
}/* debug [instance_properties/setter]: isAuthenticated */


// A Boolean value that indicates whether the player can join multiplayer games.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/ismultiplayergamingrestricted
func (l_ LocalPlayer) IsMultiplayerGamingRestricted() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isMultiplayerGamingRestricted"))
	return rv
}/* debug [instance_properties/getter]: isMultiplayerGamingRestricted */


// A Boolean value that indicates whether the player can join multiplayer games.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/ismultiplayergamingrestricted
func (l_ LocalPlayer) SetIsMultiplayerGamingRestricted(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsMultiplayerGamingRestricted:"), value)
}/* debug [instance_properties/setter]: isMultiplayerGamingRestricted */


// A Boolean value that indicates whether the player can use personalized communication on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/ispersonalizedcommunicationrestricted
func (l_ LocalPlayer) IsPersonalizedCommunicationRestricted() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isPersonalizedCommunicationRestricted"))
	return rv
}/* debug [instance_properties/getter]: isPersonalizedCommunicationRestricted */


// A Boolean value that indicates whether the player can use personalized communication on the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/ispersonalizedcommunicationrestricted
func (l_ LocalPlayer) SetIsPersonalizedCommunicationRestricted(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsPersonalizedCommunicationRestricted:"), value)
}/* debug [instance_properties/setter]: isPersonalizedCommunicationRestricted */


// A Boolean value that indicates whether the local player is underage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/isunderage
func (l_ LocalPlayer) IsUnderage() bool {
	rv := objc.Send[bool](l_.ID, objc.Sel("isUnderage"))
	return rv
}/* debug [instance_properties/getter]: isUnderage */


// A Boolean value that indicates whether the local player is underage.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gklocalplayer/isunderage
func (l_ LocalPlayer) SetIsUnderage(value bool) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setIsUnderage:"), value)
}/* debug [instance_properties/setter]: isUnderage */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKLocalPlayer */


