// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GKTurnBasedMatchmakerViewController */


/* debug [class_header]: Header for GKTurnBasedMatchmakerViewController */
// The class instance for the [TurnBasedMatchmakerViewController] class.
var (
	TurnBasedMatchmakerViewControllerClass     _TurnBasedMatchmakerViewControllerClass
	TurnBasedMatchmakerViewControllerClassOnce sync.Once
)

func getTurnBasedMatchmakerViewControllerClass() _TurnBasedMatchmakerViewControllerClass {
	TurnBasedMatchmakerViewControllerClassOnce.Do(func() {
		TurnBasedMatchmakerViewControllerClass = _TurnBasedMatchmakerViewControllerClass{objc.GetClass("GKTurnBasedMatchmakerViewController")}
	})
	return TurnBasedMatchmakerViewControllerClass
}

type _TurnBasedMatchmakerViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TurnBasedMatchmakerViewController */
// An interface definition for the [TurnBasedMatchmakerViewController] class.
type ITurnBasedMatchmakerViewController interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for TurnBasedMatchmakerViewController */
	// properties:
	MatchmakingMode() MatchmakingMode
	SetMatchmakingMode(value MatchmakingMode)
	ShowExistingMatches() bool
	SetShowExistingMatches(value bool)
	TurnBasedMatchmakerDelegate() unsafe.Pointer
	SetTurnBasedMatchmakerDelegate(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TurnBasedMatchmakerViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TurnBasedMatchmakerViewController */
// Alloc allocates a new instance without initialization.
func (tc _TurnBasedMatchmakerViewControllerClass) Alloc() TurnBasedMatchmakerViewController {
	rv := objc.Send[TurnBasedMatchmakerViewController](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TurnBasedMatchmakerViewControllerClass) New() TurnBasedMatchmakerViewController {
	rv := objc.Send[TurnBasedMatchmakerViewController](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TurnBasedMatchmakerViewController) Init() TurnBasedMatchmakerViewController {
	rv := objc.Send[TurnBasedMatchmakerViewController](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TurnBasedMatchmakerViewController) Autorelease() TurnBasedMatchmakerViewController {
	rv := objc.Send[TurnBasedMatchmakerViewController](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTurnBasedMatchmakerViewController creates a new TurnBasedMatchmakerViewController instance.
func NewTurnBasedMatchmakerViewController() TurnBasedMatchmakerViewController {
	return getTurnBasedMatchmakerViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TurnBasedMatchmakerViewController */
// An interface that allows a player to invite other players to a turn-based match and automatch to fill any empty slots.
//
// Before you create a object, create a object and configure it according to the parameters of your game. Then, pass the match request to the initializer to create the view controller. Configure the view controller and set its delegate before you present it to the local player. The view controller allows the local player to choose other players and optionally fill empty slots using automatch. The interface also allows players to select an existing match, forfeit a match, or view a completed match. Implement the protocol to handle when a player selects players, cancels matchmaking, or encounters an error. Implement the delegate method to dismiss the view controller when the local player invites players. Register as a listener of the protocol and implement methods that handle other turn-based events. For example, implement the to update match data and present the gameplay interface for the local player to take their turn. In iOS, you present and dismiss the view controller from another view controller in your game, using the methods provided by the class. If you use SwiftUI, you can get the root view controller from the object. In macOS, you use the class to present and dismiss the view controller.


// An interface that allows a player to invite other players to a turn-based match and automatch to fill any empty slots.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatchmakerViewController
type TurnBasedMatchmakerViewController struct {
	ViewController
}

// TurnBasedMatchmakerViewControllerFrom constructs a [TurnBasedMatchmakerViewController] from an unsafe.Pointer.
//
// An interface that allows a player to invite other players to a turn-based match and automatch to fill any empty slots.
func TurnBasedMatchmakerViewControllerFrom(ptr unsafe.Pointer) TurnBasedMatchmakerViewController {
	return TurnBasedMatchmakerViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TurnBasedMatchmakerViewController */

// Creates a matchmaker view controller for the local player to start inviting other players to a turn-based game.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatchmakerViewController/init(matchRequest:)
func NewTurnBasedMatchmakerViewControllerWithMatchRequest(request IGKMatchRequest) TurnBasedMatchmakerViewController {
	instance := getTurnBasedMatchmakerViewControllerClass().Alloc()
	rv := objc.Send[TurnBasedMatchmakerViewController](instance.ID, objc.Sel("initWithMatchRequest:"), request)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewTurnBasedMatchmakerViewControllerWithMatchRequest */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TurnBasedMatchmakerViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TurnBasedMatchmakerViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TurnBasedMatchmakerViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TurnBasedMatchmakerViewController */

// The mode that a multiplayer game uses to find players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatchmakerViewController/matchmakingMode
func (t_ TurnBasedMatchmakerViewController) MatchmakingMode() MatchmakingMode {
	rv := objc.Send[MatchmakingMode](t_.ID, objc.Sel("matchmakingMode"))
	return rv
}/* debug [instance_properties/getter]: matchmakingMode */


// The mode that a multiplayer game uses to find players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatchmakerViewController/matchmakingMode
func (t_ TurnBasedMatchmakerViewController) SetMatchmakingMode(value MatchmakingMode) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setMatchmakingMode:"), value)
}/* debug [instance_properties/setter]: matchmakingMode */


// A Boolean value that determines whether the view controller shows existing matches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatchmakerViewController/showExistingMatches
func (t_ TurnBasedMatchmakerViewController) ShowExistingMatches() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("showExistingMatches"))
	return rv
}/* debug [instance_properties/getter]: showExistingMatches */


// A Boolean value that determines whether the view controller shows existing matches.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatchmakerViewController/showExistingMatches
func (t_ TurnBasedMatchmakerViewController) SetShowExistingMatches(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setShowExistingMatches:"), value)
}/* debug [instance_properties/setter]: showExistingMatches */


// The object that handles turn-based matchmaker view controller changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatchmakerViewController/turnBasedMatchmakerDelegate
func (t_ TurnBasedMatchmakerViewController) TurnBasedMatchmakerDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("turnBasedMatchmakerDelegate"))
	return rv
}/* debug [instance_properties/getter]: turnBasedMatchmakerDelegate */


// The object that handles turn-based matchmaker view controller changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatchmakerViewController/turnBasedMatchmakerDelegate
func (t_ TurnBasedMatchmakerViewController) SetTurnBasedMatchmakerDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTurnBasedMatchmakerDelegate:"), value)
}/* debug [instance_properties/setter]: turnBasedMatchmakerDelegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKTurnBasedMatchmakerViewController */


