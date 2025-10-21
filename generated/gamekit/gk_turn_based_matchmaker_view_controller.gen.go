// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

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

// An interface definition for the [TurnBasedMatchmakerViewController] class.
type ITurnBasedMatchmakerViewController interface {
	appkit.IViewController
}

// An interface that allows a player to invite other players to a turn-based match and automatch to fill any empty slots.
//
// Before you create a object, create a object and configure it according to the parameters of your game. Then, pass the match request to the initializer to create the view controller. Configure the view controller and set its delegate before you present it to the local player. The view controller allows the local player to choose other players and optionally fill empty slots using automatch. The interface also allows players to select an existing match, forfeit a match, or view a completed match. Implement the protocol to handle when a player selects players, cancels matchmaking, or encounters an error. Implement the delegate method to dismiss the view controller when the local player invites players. Register as a listener of the protocol and implement methods that handle other turn-based events. For example, implement the to update match data and present the gameplay interface for the local player to take their turn. In iOS, you present and dismiss the view controller from another view controller in your game, using the methods provided by the class. If you use SwiftUI, you can get the root view controller from the object. In macOS, you use the class to present and dismiss the view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatchmakerViewController
type TurnBasedMatchmakerViewController struct {
	appkit.ViewController
}

// TurnBasedMatchmakerViewControllerFrom constructs a [TurnBasedMatchmakerViewController] from an unsafe.Pointer.
//
// An interface that allows a player to invite other players to a turn-based match and automatch to fill any empty slots.
func TurnBasedMatchmakerViewControllerFrom(ptr unsafe.Pointer) TurnBasedMatchmakerViewController {
	return TurnBasedMatchmakerViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (tc _TurnBasedMatchmakerViewControllerClass) Alloc() TurnBasedMatchmakerViewController {
	rv := objc.Send[TurnBasedMatchmakerViewController](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Creates a matchmaker view controller for the local player to start inviting other players to a turn-based game.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatchmakerViewController/init(matchRequest:)
func NewTurnBasedMatchmakerViewControllerWithMatchRequest(request unsafe.Pointer) TurnBasedMatchmakerViewController {
	instance := getTurnBasedMatchmakerViewControllerClass().Alloc()
	rv := objc.Send[TurnBasedMatchmakerViewController](instance.ID, objc.Sel("initWithMatchRequest:"), request)
	rv.Autorelease()
	return rv
}


// A Boolean value that determines whether the view controller shows existing matches.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatchmakerViewController/showExistingMatches
func (t_ TurnBasedMatchmakerViewController) ShowExistingMatches() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("showExistingMatches"))
	return rv
}


// SetShowExistingMatches sets the value of the showExistingMatches property.
// A Boolean value that determines whether the view controller shows existing matches.

//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatchmakerViewController/showExistingMatches
func (t_ TurnBasedMatchmakerViewController) SetShowExistingMatches(value bool) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setShowExistingMatches:"), value)
}

// The object that handles turn-based matchmaker view controller changes.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatchmakerViewController/turnBasedMatchmakerDelegate
func (t_ TurnBasedMatchmakerViewController) TurnBasedMatchmakerDelegate() objc.ID {
	rv := objc.Send[objc.ID](t_.ID, objc.Sel("turnBasedMatchmakerDelegate"))
	return rv
}


// SetTurnBasedMatchmakerDelegate sets the value of the turnBasedMatchmakerDelegate property.
// The object that handles turn-based matchmaker view controller changes.

//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedMatchmakerViewController/turnBasedMatchmakerDelegate
func (t_ TurnBasedMatchmakerViewController) SetTurnBasedMatchmakerDelegate(value objc.ID) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setTurnBasedMatchmakerDelegate:"), value)
}


