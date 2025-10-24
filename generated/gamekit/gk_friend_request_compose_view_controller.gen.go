// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class GKFriendRequestComposeViewController */


/* debug [class_header]: Header for GKFriendRequestComposeViewController */
// The class instance for the [FriendRequestComposeViewController] class.
var (
	FriendRequestComposeViewControllerClass     _FriendRequestComposeViewControllerClass
	FriendRequestComposeViewControllerClassOnce sync.Once
)

func getFriendRequestComposeViewControllerClass() _FriendRequestComposeViewControllerClass {
	FriendRequestComposeViewControllerClassOnce.Do(func() {
		FriendRequestComposeViewControllerClass = _FriendRequestComposeViewControllerClass{objc.GetClass("GKFriendRequestComposeViewController")}
	})
	return FriendRequestComposeViewControllerClass
}

type _FriendRequestComposeViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for FriendRequestComposeViewController */
// An interface definition for the [FriendRequestComposeViewController] class.
type IFriendRequestComposeViewController interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for FriendRequestComposeViewController */
	// properties:
	ComposeViewDelegate() unsafe.Pointer
	SetComposeViewDelegate(value unsafe.Pointer)
	Delegate() ObjectProtocol /* not a class type */
	SetDelegate(value ObjectProtocol /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for FriendRequestComposeViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for FriendRequestComposeViewController */
// Alloc allocates a new instance without initialization.
func (fc _FriendRequestComposeViewControllerClass) Alloc() FriendRequestComposeViewController {
	rv := objc.Send[FriendRequestComposeViewController](objc.ID(fc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (fc _FriendRequestComposeViewControllerClass) New() FriendRequestComposeViewController {
	rv := objc.Send[FriendRequestComposeViewController](objc.ID(fc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (f_ FriendRequestComposeViewController) Init() FriendRequestComposeViewController {
	rv := objc.Send[FriendRequestComposeViewController](f_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (f_ FriendRequestComposeViewController) Autorelease() FriendRequestComposeViewController {
	rv := objc.Send[FriendRequestComposeViewController](f_.ID, objc.Sel("autorelease"))
	return rv
}

// NewFriendRequestComposeViewController creates a new FriendRequestComposeViewController instance.
func NewFriendRequestComposeViewController() FriendRequestComposeViewController {
	return getFriendRequestComposeViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for FriendRequestComposeViewController */
// Your game uses the class to present a screen that allows the local player to send friend requests to other players.
//
// To show a friend request, initialize a new object and set the delegate. Optionally, you can customize the request by adding a text message or a list of recipients. Then, present the new view controller and wait for the delegate to be called. Once the delegate is called, dismiss the view controller. On iOS, you present and dismiss the view controller from another view controller in your game, using the methods provided by the class. In macOS, you use the class to present and dismiss the view controller. The listing below shows one way your view controller can allow a player to send a request to other players. For this method, an array of objects is passed in as a parameter. The method instantiates a object, sets its delegate, and adds the list of players intended to receive the invitation. The view controller then presents the friend request and returns.


// Your game uses the class to present a screen that allows the local player to send friend requests to other players.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKFriendRequestComposeViewController
type FriendRequestComposeViewController struct {
	ViewController
}

// FriendRequestComposeViewControllerFrom constructs a [FriendRequestComposeViewController] from an unsafe.Pointer.
//
// Your game uses the class to present a screen that allows the local player to send friend requests to other players.
func FriendRequestComposeViewControllerFrom(ptr unsafe.Pointer) FriendRequestComposeViewController {
	return FriendRequestComposeViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for FriendRequestComposeViewController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for FriendRequestComposeViewController */

// Returns the maximum number of recipients permitted in a single request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKFriendRequestComposeViewController/maxNumberOfRecipients()
func (fc _FriendRequestComposeViewControllerClass) MaxNumberOfRecipients() uint {
	rv := objc.Send[uint](objc.ID(fc.class), objc.Sel("maxNumberOfRecipients"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=MaxNumberOfRecipients) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for FriendRequestComposeViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for FriendRequestComposeViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for FriendRequestComposeViewController */

// The view controller’s delegate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKFriendRequestComposeViewController/composeViewDelegate
func (f_ FriendRequestComposeViewController) ComposeViewDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](f_.ID, objc.Sel("composeViewDelegate"))
	return rv
}/* debug [instance_properties/getter]: composeViewDelegate */


// The view controller’s delegate
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKFriendRequestComposeViewController/composeViewDelegate
func (f_ FriendRequestComposeViewController) SetComposeViewDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setComposeViewDelegate:"), value)
}/* debug [instance_properties/setter]: composeViewDelegate */


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (f_ FriendRequestComposeViewController) Delegate() ObjectProtocol /* not a class type */ {
	rv := objc.Send[ObjectProtocol](f_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (f_ FriendRequestComposeViewController) SetDelegate(value ObjectProtocol /* not a class type */) {
	objc.Send[objc.ID](f_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKFriendRequestComposeViewController */



