// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class GKLeaderboardViewController */


/* debug [class_header]: Header for GKLeaderboardViewController */
// The class instance for the [LeaderboardViewController] class.
var (
	LeaderboardViewControllerClass     _LeaderboardViewControllerClass
	LeaderboardViewControllerClassOnce sync.Once
)

func getLeaderboardViewControllerClass() _LeaderboardViewControllerClass {
	LeaderboardViewControllerClassOnce.Do(func() {
		LeaderboardViewControllerClass = _LeaderboardViewControllerClass{objc.GetClass("GKLeaderboardViewController")}
	})
	return LeaderboardViewControllerClass
}

type _LeaderboardViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for LeaderboardViewController */
// An interface definition for the [LeaderboardViewController] class.
type ILeaderboardViewController interface {
	IGameCenterViewController
	
/* debug [class_interface_properties]: Properties for LeaderboardViewController */
	// properties:
	Category() objc.IObject /* cross-framework: NSString */
	SetCategory(value objc.IObject /* cross-framework: NSString */)
	LeaderboardDelegate() unsafe.Pointer
	SetLeaderboardDelegate(value unsafe.Pointer)
	TimeScope() LeaderboardTimeScope
	SetTimeScope(value LeaderboardTimeScope)
	Delegate() ObjectProtocol /* not a class type */
	SetDelegate(value ObjectProtocol /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for LeaderboardViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for LeaderboardViewController */
// Alloc allocates a new instance without initialization.
func (lc _LeaderboardViewControllerClass) Alloc() LeaderboardViewController {
	rv := objc.Send[LeaderboardViewController](objc.ID(lc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (lc _LeaderboardViewControllerClass) New() LeaderboardViewController {
	rv := objc.Send[LeaderboardViewController](objc.ID(lc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (l_ LeaderboardViewController) Init() LeaderboardViewController {
	rv := objc.Send[LeaderboardViewController](l_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (l_ LeaderboardViewController) Autorelease() LeaderboardViewController {
	rv := objc.Send[LeaderboardViewController](l_.ID, objc.Sel("autorelease"))
	return rv
}

// NewLeaderboardViewController creates a new LeaderboardViewController instance.
func NewLeaderboardViewController() LeaderboardViewController {
	return getLeaderboardViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for LeaderboardViewController */
// The class provides a standard user interface that displays leaderboard scores to the player. If the class is available, you should use it instead.
//
// To show a leaderboard screen, initialize a new object and set the delegate. Optionally, you can configure the view controller to display specific data to the player. Then, present the new view controller and wait for the delegate to be called. Once the delegate is called, dismiss the view controller. On iOS, you present and dismiss the view controller from another view controller in your game, using the methods provided by the class. In macOS, you use the class to present and dismiss the view controller. Your game should pause other activities before presenting the leaderboard.


// The class provides a standard user interface that displays leaderboard scores to the player. If the class is available, you should use it instead.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboardViewController
type LeaderboardViewController struct {
	GameCenterViewController
}

// LeaderboardViewControllerFrom constructs a [LeaderboardViewController] from an unsafe.Pointer.
//
// The class provides a standard user interface that displays leaderboard scores to the player. If the class is available, you should use it instead.
func LeaderboardViewControllerFrom(ptr unsafe.Pointer) LeaderboardViewController {
	return LeaderboardViewController{
		GameCenterViewController: GameCenterViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for LeaderboardViewController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for LeaderboardViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for LeaderboardViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for LeaderboardViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for LeaderboardViewController */

// The named leaderboard that is displayed by the view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboardViewController/category
func (l_ LeaderboardViewController) Category() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](l_.ID, objc.Sel("category"))
	return rv
}/* debug [instance_properties/getter]: category */


// The named leaderboard that is displayed by the view controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboardViewController/category
func (l_ LeaderboardViewController) SetCategory(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setCategory:"), value)
}/* debug [instance_properties/setter]: category */


// The view controller’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboardViewController/leaderboardDelegate
func (l_ LeaderboardViewController) LeaderboardDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](l_.ID, objc.Sel("leaderboardDelegate"))
	return rv
}/* debug [instance_properties/getter]: leaderboardDelegate */


// The view controller’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboardViewController/leaderboardDelegate
func (l_ LeaderboardViewController) SetLeaderboardDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setLeaderboardDelegate:"), value)
}/* debug [instance_properties/setter]: leaderboardDelegate */


// A time filter used to restrict which scores are displayed to the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboardViewController/timeScope
func (l_ LeaderboardViewController) TimeScope() LeaderboardTimeScope {
	rv := objc.Send[LeaderboardTimeScope](l_.ID, objc.Sel("timeScope"))
	return rv
}/* debug [instance_properties/getter]: timeScope */


// A time filter used to restrict which scores are displayed to the player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKLeaderboardViewController/timeScope
func (l_ LeaderboardViewController) SetTimeScope(value LeaderboardTimeScope) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setTimeScope:"), value)
}/* debug [instance_properties/setter]: timeScope */


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (l_ LeaderboardViewController) Delegate() ObjectProtocol /* not a class type */ {
	rv := objc.Send[ObjectProtocol](l_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (l_ LeaderboardViewController) SetDelegate(value ObjectProtocol /* not a class type */) {
	objc.Send[objc.ID](l_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKLeaderboardViewController */



