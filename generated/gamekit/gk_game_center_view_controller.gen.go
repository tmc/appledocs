// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class GKGameCenterViewController */


/* debug [class_header]: Header for GKGameCenterViewController */
// The class instance for the [GameCenterViewController] class.
var (
	GameCenterViewControllerClass     _GameCenterViewControllerClass
	GameCenterViewControllerClassOnce sync.Once
)

func getGameCenterViewControllerClass() _GameCenterViewControllerClass {
	GameCenterViewControllerClassOnce.Do(func() {
		GameCenterViewControllerClass = _GameCenterViewControllerClass{objc.GetClass("GKGameCenterViewController")}
	})
	return GameCenterViewControllerClass
}

type _GameCenterViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GameCenterViewController */
// An interface definition for the [GameCenterViewController] class.
type IGameCenterViewController interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for GameCenterViewController */
	// properties:
	GameCenterDelegate() unsafe.Pointer
	SetGameCenterDelegate(value unsafe.Pointer)
	LeaderboardCategory() objc.IObject /* cross-framework: NSString */
	SetLeaderboardCategory(value objc.IObject /* cross-framework: NSString */)
	LeaderboardIdentifier() objc.IObject /* cross-framework: NSString */
	SetLeaderboardIdentifier(value objc.IObject /* cross-framework: NSString */)
	LeaderboardTimeScope() LeaderboardTimeScope
	SetLeaderboardTimeScope(value LeaderboardTimeScope)
	ViewState() GameCenterViewControllerState
	SetViewState(value GameCenterViewControllerState)
	Delegate() ObjectProtocol /* not a class type */
	SetDelegate(value ObjectProtocol /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GameCenterViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GameCenterViewController */
// Alloc allocates a new instance without initialization.
func (gc _GameCenterViewControllerClass) Alloc() GameCenterViewController {
	rv := objc.Send[GameCenterViewController](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (gc _GameCenterViewControllerClass) New() GameCenterViewController {
	rv := objc.Send[GameCenterViewController](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GameCenterViewController) Init() GameCenterViewController {
	rv := objc.Send[GameCenterViewController](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GameCenterViewController) Autorelease() GameCenterViewController {
	rv := objc.Send[GameCenterViewController](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGameCenterViewController creates a new GameCenterViewController instance.
func NewGameCenterViewController() GameCenterViewController {
	return getGameCenterViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GameCenterViewController */
// The dashboard that allows players to access their Game Center data in your game.
//
// This view controller presents the dashboard from which players can browse and manage their Game Center data. You can present the dashboard in a specific state from which players can navigate to other areas, including their profile. Your game should pause other activities before presenting the dashboard. To present the dashboard, initialize a new object and set its delegate. Optionally, initialize a view controller in a specific state, to show a leaderboard with scores from a set of players or during a time period, or to show a specific achievement. Then present the view controller to the player, and GameKit calls your delegate when the player dismisses it. For visionOS games, the dashboard appears anchored to the window, scene, or view relative to where you present the view controller. For immersive games, set the parent window to a separate window group than the immersive space window group. For the visionOS location of the dashboard when using the access point, see .


// The dashboard that allows players to access their Game Center data in your game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewController
type GameCenterViewController struct {
	ViewController
}

// GameCenterViewControllerFrom constructs a [GameCenterViewController] from an unsafe.Pointer.
//
// The dashboard that allows players to access their Game Center data in your game.
func GameCenterViewControllerFrom(ptr unsafe.Pointer) GameCenterViewController {
	return GameCenterViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GameCenterViewController */

// Creates a view controller that presents an achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewController/init(achievementID:)
func NewGameCenterViewControllerWithAchievementID(achievementID objc.IObject /* cross-framework: NSString */) GameCenterViewController {
	instance := getGameCenterViewControllerClass().Alloc()
	rv := objc.Send[GameCenterViewController](instance.ID, objc.Sel("initWithAchievementID:"), achievementID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGameCenterViewControllerWithAchievementID */


// Creates a view controller that presents a leaderboard with data from the specified players and time period.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewController/init(leaderboardID:playerScope:timeScope:)
func NewGameCenterViewControllerWithLeaderboardIDPlayerScopeTimeScope(leaderboardID objc.IObject /* cross-framework: NSString */, playerScope LeaderboardPlayerScope, timeScope LeaderboardTimeScope) GameCenterViewController {
	instance := getGameCenterViewControllerClass().Alloc()
	rv := objc.Send[GameCenterViewController](instance.ID, objc.Sel("initWithLeaderboardID:playerScope:timeScope:"), leaderboardID, playerScope, timeScope)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGameCenterViewControllerWithLeaderboardIDPlayerScopeTimeScope */


// Creates a view controller that presents a leaderboard with data for the specified players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewController/init(leaderboard:playerScope:)
func NewGameCenterViewControllerWithLeaderboardPlayerScope(leaderboard IGKLeaderboard, playerScope LeaderboardPlayerScope) GameCenterViewController {
	instance := getGameCenterViewControllerClass().Alloc()
	rv := objc.Send[GameCenterViewController](instance.ID, objc.Sel("initWithLeaderboard:playerScope:"), leaderboard, playerScope)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGameCenterViewControllerWithLeaderboardPlayerScope */


// Creates a view controller that presents a leaderboard set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewController/init(leaderboardSetID:)
func NewGameCenterViewControllerWithLeaderboardSetID(leaderboardSetID objc.IObject /* cross-framework: NSString */) GameCenterViewController {
	instance := getGameCenterViewControllerClass().Alloc()
	rv := objc.Send[GameCenterViewController](instance.ID, objc.Sel("initWithLeaderboardSetID:"), leaderboardSetID)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGameCenterViewControllerWithLeaderboardSetID */


// Creates a view controller that presents a player’s Game Center profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewController/init(player:)
func NewGameCenterViewControllerWithPlayer(player IGKPlayer) GameCenterViewController {
	instance := getGameCenterViewControllerClass().Alloc()
	rv := objc.Send[GameCenterViewController](instance.ID, objc.Sel("initWithPlayer:"), player)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGameCenterViewControllerWithPlayer */


// Creates a view controller that presents the specified Game Center content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewController/init(state:)
func NewGameCenterViewControllerWithState(state GameCenterViewControllerState) GameCenterViewController {
	instance := getGameCenterViewControllerClass().Alloc()
	rv := objc.Send[GameCenterViewController](instance.ID, objc.Sel("initWithState:"), state)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGameCenterViewControllerWithState */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GameCenterViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GameCenterViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GameCenterViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GameCenterViewController */

// The view controller’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewController/gameCenterDelegate
func (g_ GameCenterViewController) GameCenterDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](g_.ID, objc.Sel("gameCenterDelegate"))
	return rv
}/* debug [instance_properties/getter]: gameCenterDelegate */


// The view controller’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewController/gameCenterDelegate
func (g_ GameCenterViewController) SetGameCenterDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGameCenterDelegate:"), value)
}/* debug [instance_properties/setter]: gameCenterDelegate */


// The named leaderboard that the view controller displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewController/leaderboardCategory
func (g_ GameCenterViewController) LeaderboardCategory() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("leaderboardCategory"))
	return rv
}/* debug [instance_properties/getter]: leaderboardCategory */


// The named leaderboard that the view controller displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewController/leaderboardCategory
func (g_ GameCenterViewController) SetLeaderboardCategory(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLeaderboardCategory:"), value)
}/* debug [instance_properties/setter]: leaderboardCategory */


// The named leaderboard that the view controller displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewController/leaderboardIdentifier
func (g_ GameCenterViewController) LeaderboardIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](g_.ID, objc.Sel("leaderboardIdentifier"))
	return rv
}/* debug [instance_properties/getter]: leaderboardIdentifier */


// The named leaderboard that the view controller displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewController/leaderboardIdentifier
func (g_ GameCenterViewController) SetLeaderboardIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLeaderboardIdentifier:"), value)
}/* debug [instance_properties/setter]: leaderboardIdentifier */


// A time filter that restricts the scores to display to the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewController/leaderboardTimeScope
func (g_ GameCenterViewController) LeaderboardTimeScope() LeaderboardTimeScope {
	rv := objc.Send[LeaderboardTimeScope](g_.ID, objc.Sel("leaderboardTimeScope"))
	return rv
}/* debug [instance_properties/getter]: leaderboardTimeScope */


// A time filter that restricts the scores to display to the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewController/leaderboardTimeScope
func (g_ GameCenterViewController) SetLeaderboardTimeScope(value LeaderboardTimeScope) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setLeaderboardTimeScope:"), value)
}/* debug [instance_properties/setter]: leaderboardTimeScope */


// The content that the Game Center controller displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewController/viewState
func (g_ GameCenterViewController) ViewState() GameCenterViewControllerState {
	rv := objc.Send[GameCenterViewControllerState](g_.ID, objc.Sel("viewState"))
	return rv
}/* debug [instance_properties/getter]: viewState */


// The content that the Game Center controller displays.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewController/viewState
func (g_ GameCenterViewController) SetViewState(value GameCenterViewControllerState) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setViewState:"), value)
}/* debug [instance_properties/setter]: viewState */


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (g_ GameCenterViewController) Delegate() ObjectProtocol /* not a class type */ {
	rv := objc.Send[ObjectProtocol](g_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (g_ GameCenterViewController) SetDelegate(value ObjectProtocol /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKGameCenterViewController */


