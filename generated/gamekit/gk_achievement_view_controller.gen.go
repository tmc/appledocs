// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GKAchievementViewController */


/* debug [class_header]: Header for GKAchievementViewController */
// The class instance for the [AchievementViewController] class.
var (
	AchievementViewControllerClass     _AchievementViewControllerClass
	AchievementViewControllerClassOnce sync.Once
)

func getAchievementViewControllerClass() _AchievementViewControllerClass {
	AchievementViewControllerClassOnce.Do(func() {
		AchievementViewControllerClass = _AchievementViewControllerClass{objc.GetClass("GKAchievementViewController")}
	})
	return AchievementViewControllerClass
}

type _AchievementViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for AchievementViewController */
// An interface definition for the [AchievementViewController] class.
type IAchievementViewController interface {
	IGameCenterViewController
	
/* debug [class_interface_properties]: Properties for AchievementViewController */
	// properties:
	AchievementDelegate() unsafe.Pointer
	SetAchievementDelegate(value unsafe.Pointer)
	Delegate() ObjectProtocol /* not a class type */
	SetDelegate(value ObjectProtocol /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for AchievementViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for AchievementViewController */
// Alloc allocates a new instance without initialization.
func (ac _AchievementViewControllerClass) Alloc() AchievementViewController {
	rv := objc.Send[AchievementViewController](objc.ID(ac.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (ac _AchievementViewControllerClass) New() AchievementViewController {
	rv := objc.Send[AchievementViewController](objc.ID(ac.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (a_ AchievementViewController) Init() AchievementViewController {
	rv := objc.Send[AchievementViewController](a_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (a_ AchievementViewController) Autorelease() AchievementViewController {
	rv := objc.Send[AchievementViewController](a_.ID, objc.Sel("autorelease"))
	return rv
}

// NewAchievementViewController creates a new AchievementViewController instance.
func NewAchievementViewController() AchievementViewController {
	return getAchievementViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for AchievementViewController */
// An object provides a standard user interface to display achievement progress for the local player. If the class is available, you should use it instead.
//
// To show achievements for the local player, initialize a new object and set the delegate. Then present the new view controller and wait for the delegate to be called. Once the delegate is called, dismiss the view controller. On iOS, you present and dismiss the view controller from another view controller in your game, using the methods provided by the class. In macOS, you use the class to present and dismiss the view controller in a window.


// An object provides a standard user interface to display achievement progress for the local player. If the class is available, you should use it instead.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementViewController
type AchievementViewController struct {
	GameCenterViewController
}

// AchievementViewControllerFrom constructs a [AchievementViewController] from an unsafe.Pointer.
//
// An object provides a standard user interface to display achievement progress for the local player. If the class is available, you should use it instead.
func AchievementViewControllerFrom(ptr unsafe.Pointer) AchievementViewController {
	return AchievementViewController{
		GameCenterViewController: GameCenterViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for AchievementViewController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for AchievementViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for AchievementViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for AchievementViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for AchievementViewController */

// The achievement view controller’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementViewController/achievementDelegate
func (a_ AchievementViewController) AchievementDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](a_.ID, objc.Sel("achievementDelegate"))
	return rv
}/* debug [instance_properties/getter]: achievementDelegate */


// The achievement view controller’s delegate.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKAchievementViewController/achievementDelegate
func (a_ AchievementViewController) SetAchievementDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setAchievementDelegate:"), value)
}/* debug [instance_properties/setter]: achievementDelegate */


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (a_ AchievementViewController) Delegate() ObjectProtocol /* not a class type */ {
	rv := objc.Send[ObjectProtocol](a_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (a_ AchievementViewController) SetDelegate(value ObjectProtocol /* not a class type */) {
	objc.Send[objc.ID](a_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKAchievementViewController */



