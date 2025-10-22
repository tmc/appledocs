// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/appkit"
)

// The class instance for the [GameSessionSharingViewController] class.
var (
	GameSessionSharingViewControllerClass     _GameSessionSharingViewControllerClass
	GameSessionSharingViewControllerClassOnce sync.Once
)

func getGameSessionSharingViewControllerClass() _GameSessionSharingViewControllerClass {
	GameSessionSharingViewControllerClassOnce.Do(func() {
		GameSessionSharingViewControllerClass = _GameSessionSharingViewControllerClass{objc.GetClass("GKGameSessionSharingViewController")}
	})
	return GameSessionSharingViewControllerClass
}

type _GameSessionSharingViewControllerClass struct {
	class objc.Class
}

// An interface definition for the [GameSessionSharingViewController] class.
type IGameSessionSharingViewController interface {
	appkit.IViewController
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	Session() GKGameSession
}

// A user interface you can use to invite other users into a tvOS game session.
//
// The sharing view controller on tvOS presents a user’s Game Center friend list, along with other people with whom the user has recently played the game. Users can select a person from the list and send them an invite using the Send button.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSessionSharingViewController
type GameSessionSharingViewController struct {
	appkit.ViewController
}

// GameSessionSharingViewControllerFrom constructs a [GameSessionSharingViewController] from an unsafe.Pointer.
//
// A user interface you can use to invite other users into a tvOS game session.
func GameSessionSharingViewControllerFrom(ptr unsafe.Pointer) GameSessionSharingViewController {
	return GameSessionSharingViewController{
		ViewController: appkit.ViewControllerFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (gc _GameSessionSharingViewControllerClass) Alloc() GameSessionSharingViewController {
	rv := objc.Send[GameSessionSharingViewController](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (gc _GameSessionSharingViewControllerClass) New() GameSessionSharingViewController {
	rv := objc.Send[GameSessionSharingViewController](objc.ID(gc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (g_ GameSessionSharingViewController) Init() GameSessionSharingViewController {
	rv := objc.Send[GameSessionSharingViewController](g_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (g_ GameSessionSharingViewController) Autorelease() GameSessionSharingViewController {
	rv := objc.Send[GameSessionSharingViewController](g_.ID, objc.Sel("autorelease"))
	return rv
}

// NewGameSessionSharingViewController creates a new GameSessionSharingViewController instance.
func NewGameSessionSharingViewController() GameSessionSharingViewController {
	return getGameSessionSharingViewControllerClass().New()
}




// Creates a new sharing view controller for a specified session.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSessionSharingViewController/init(session:)
func NewGameSessionSharingViewControllerWithSession(session IGKGameSession) GameSessionSharingViewController {
	instance := getGameSessionSharingViewControllerClass().Alloc()
	rv := objc.Send[GameSessionSharingViewController](instance.ID, objc.Sel("initWithSession:"), session)
	rv.Autorelease()
	return rv
}


// The delegate for the sharing view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSessionSharingViewController/delegate
func (g_ GameSessionSharingViewController) Delegate() objc.ID {
	rv := objc.Send[objc.ID](g_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate for the sharing view controller.

//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSessionSharingViewController/delegate
func (g_ GameSessionSharingViewController) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](g_.ID, objc.Sel("setDelegate:"), value)
}

// The game session associated with the view controller.
//
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSessionSharingViewController/session
func (g_ GameSessionSharingViewController) Session() GKGameSession {
	rv := objc.Send[GKGameSession](g_.ID, objc.Sel("session"))
	return rv
}


