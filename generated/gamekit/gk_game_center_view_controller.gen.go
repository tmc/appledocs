// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [GameCenterViewController] class.
type IGameCenterViewController interface {
	appkit.IViewController
	// properties:
	GameCenterDelegate() GameCenterControllerDelegate /* not a class type */
	SetGameCenterDelegate(value GameCenterControllerDelegate /* not a class type */)
	Delegate() ObjectProtocol /* not a class type */
	SetDelegate(value ObjectProtocol /* not a class type */)
	// methods:
}

// The dashboard that allows players to access their Game Center data in your game.
//
// This view controller presents the dashboard from which players can browse and manage their Game Center data. You can present the dashboard in a specific state from which players can navigate to other areas, including their profile. Your game should pause other activities before presenting the dashboard. To present the dashboard, initialize a new object and set its delegate. Optionally, initialize a view controller in a specific state, to show a leaderboard with scores from a set of players or during a time period, or to show a specific achievement. Then present the view controller to the player, and GameKit calls your delegate when the player dismisses it. For visionOS games, the dashboard appears anchored to the window, scene, or view relative to where you present the view controller. For immersive games, set the parent window to a separate window group than the immersive space window group. For the visionOS location of the dashboard when using the access point, see .


// The dashboard that allows players to access their Game Center data in your game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewController
type GameCenterViewController struct {
	appkit.ViewController
}

// GameCenterViewControllerFrom constructs a [GameCenterViewController] from an unsafe.Pointer.
//
// The dashboard that allows players to access their Game Center data in your game.
func GameCenterViewControllerFrom(ptr unsafe.Pointer) GameCenterViewController {
	return GameCenterViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GameCenterViewControllerClass) Alloc() GameCenterViewController {
	rv := objc.Send[GameCenterViewController](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// Creates a view controller that presents an achievement.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewController/init(achievementID:)
func NewGameCenterViewControllerWithAchievementID(achievementID objc.IObject /* cross-framework: NSString */) GameCenterViewController {
	instance := getGameCenterViewControllerClass().Alloc()
	rv := objc.Send[GameCenterViewController](instance.ID, objc.Sel("initWithAchievementID:"), achievementID)
	rv.Autorelease()
	return rv
}


// Creates a view controller that presents a leaderboard set.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewController/init(leaderboardSetID:)
func NewGameCenterViewControllerWithLeaderboardSetID(leaderboardSetID objc.IObject /* cross-framework: NSString */) GameCenterViewController {
	instance := getGameCenterViewControllerClass().Alloc()
	rv := objc.Send[GameCenterViewController](instance.ID, objc.Sel("initWithLeaderboardSetID:"), leaderboardSetID)
	rv.Autorelease()
	return rv
}


// Creates a view controller that presents a player’s Game Center profile.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewController/init(player:)
func NewGameCenterViewControllerWithPlayer(player IGKPlayer) GameCenterViewController {
	instance := getGameCenterViewControllerClass().Alloc()
	rv := objc.Send[GameCenterViewController](instance.ID, objc.Sel("initWithPlayer:"), player)
	rv.Autorelease()
	return rv
}


// Creates a view controller that presents the specified Game Center content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameCenterViewController/init(state:)
func NewGameCenterViewControllerWithState(state GameCenterViewControllerState /* not a class type */) GameCenterViewController {
	instance := getGameCenterViewControllerClass().Alloc()
	rv := objc.Send[GameCenterViewController](instance.ID, objc.Sel("initWithState:"), state)
	rv.Autorelease()
	return rv
}



// The view controller’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgamecenterviewcontroller/gamecenterdelegate
func (g_ GameCenterViewController) GameCenterDelegate() GameCenterControllerDelegate /* not a class type */ {
	rv := objc.Send[GameCenterControllerDelegate](g_.ID, objc.Sel("gameCenterDelegate"))
	return rv
}


// The view controller’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkgamecenterviewcontroller/gamecenterdelegate
func (g_ GameCenterViewController) SetGameCenterDelegate(value GameCenterControllerDelegate /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setGameCenterDelegate:"), value)
}


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (g_ GameCenterViewController) Delegate() ObjectProtocol /* not a class type */ {
	rv := objc.Send[ObjectProtocol](g_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (g_ GameCenterViewController) SetDelegate(value ObjectProtocol /* not a class type */) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDelegate:"), value)
}


