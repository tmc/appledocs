// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class GKChallengesViewController */


/* debug [class_header]: Header for GKChallengesViewController */
// The class instance for the [ChallengesViewController] class.
var (
	ChallengesViewControllerClass     _ChallengesViewControllerClass
	ChallengesViewControllerClassOnce sync.Once
)

func getChallengesViewControllerClass() _ChallengesViewControllerClass {
	ChallengesViewControllerClassOnce.Do(func() {
		ChallengesViewControllerClass = _ChallengesViewControllerClass{objc.GetClass("GKChallengesViewController")}
	})
	return ChallengesViewControllerClass
}

type _ChallengesViewControllerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ChallengesViewController */
// An interface definition for the [ChallengesViewController] class.
type IChallengesViewController interface {
	IViewController
	
/* debug [class_interface_properties]: Properties for ChallengesViewController */
	// properties:
	ChallengeDelegate() unsafe.Pointer
	SetChallengeDelegate(value unsafe.Pointer)
	Delegate() ObjectProtocol /* not a class type */
	SetDelegate(value ObjectProtocol /* not a class type */)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ChallengesViewController */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ChallengesViewController */
// Alloc allocates a new instance without initialization.
func (cc _ChallengesViewControllerClass) Alloc() ChallengesViewController {
	rv := objc.Send[ChallengesViewController](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ChallengesViewControllerClass) New() ChallengesViewController {
	rv := objc.Send[ChallengesViewController](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ChallengesViewController) Init() ChallengesViewController {
	rv := objc.Send[ChallengesViewController](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ChallengesViewController) Autorelease() ChallengesViewController {
	rv := objc.Send[ChallengesViewController](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewChallengesViewController creates a new ChallengesViewController instance.
func NewChallengesViewController() ChallengesViewController {
	return getChallengesViewControllerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ChallengesViewController */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengesViewController
type ChallengesViewController struct {
	ViewController
}

// ChallengesViewControllerFrom constructs a [ChallengesViewController] from an unsafe.Pointer.
func ChallengesViewControllerFrom(ptr unsafe.Pointer) ChallengesViewController {
	return ChallengesViewController{
		ViewController: ViewControllerFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ChallengesViewController *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ChallengesViewController */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ChallengesViewController */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ChallengesViewController */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ChallengesViewController */

// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengesViewController/challengeDelegate
func (c_ ChallengesViewController) ChallengeDelegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("challengeDelegate"))
	return rv
}/* debug [instance_properties/getter]: challengeDelegate */


// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengesViewController/challengeDelegate
func (c_ ChallengesViewController) SetChallengeDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setChallengeDelegate:"), value)
}/* debug [instance_properties/setter]: challengeDelegate */


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (c_ ChallengesViewController) Delegate() ObjectProtocol /* not a class type */ {
	rv := objc.Send[ObjectProtocol](c_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/gamekit/gkturnbasedeventhandler/delegate
func (c_ ChallengesViewController) SetDelegate(value ObjectProtocol /* not a class type */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKChallengesViewController */



