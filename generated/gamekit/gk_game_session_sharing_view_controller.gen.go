// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GKGameSessionSharingViewController */


/* debug [class_header]: Header for GKGameSessionSharingViewController */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for GameSessionSharingViewController */
// An interface definition for the [GameSessionSharingViewController] class.
type IGameSessionSharingViewController interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for GameSessionSharingViewController */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for GameSessionSharingViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for GameSessionSharingViewController */
// Alloc allocates a new instance without initialization.
func (gc _GameSessionSharingViewControllerClass) Alloc() GameSessionSharingViewController {
	rv := objc.Send[GameSessionSharingViewController](objc.ID(gc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for GameSessionSharingViewController */
// A user interface you can use to invite other users into a tvOS game session.
//
// The sharing view controller on tvOS presents a user’s Game Center friend list, along with other people with whom the user has recently played the game. Users can select a person from the list and send them an invite using the Send button.


// A user interface you can use to invite other users into a tvOS game session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSessionSharingViewController
type GameSessionSharingViewController struct {
	ViewController
}

// GameSessionSharingViewControllerFrom constructs a [GameSessionSharingViewController] from an unsafe.Pointer.
//
// A user interface you can use to invite other users into a tvOS game session.
func GameSessionSharingViewControllerFrom(ptr unsafe.Pointer) GameSessionSharingViewController {
	return GameSessionSharingViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for GameSessionSharingViewController */

// Creates a new sharing view controller for a specified session.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKGameSessionSharingViewController/init(session:)
func NewGameSessionSharingViewControllerWithSession(session IGKGameSession) GameSessionSharingViewController {
	instance := getGameSessionSharingViewControllerClass().Alloc()
	rv := objc.Send[GameSessionSharingViewController](instance.ID, objc.Sel("initWithSession:"), session)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewGameSessionSharingViewControllerWithSession */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for GameSessionSharingViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for GameSessionSharingViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for GameSessionSharingViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for GameSessionSharingViewController */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKGameSessionSharingViewController */


