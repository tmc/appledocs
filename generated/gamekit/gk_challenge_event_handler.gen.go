// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKChallengeEventHandler */


/* debug [class_header]: Header for GKChallengeEventHandler */
// The class instance for the [ChallengeEventHandler] class.
var (
	ChallengeEventHandlerClass     _ChallengeEventHandlerClass
	ChallengeEventHandlerClassOnce sync.Once
)

func getChallengeEventHandlerClass() _ChallengeEventHandlerClass {
	ChallengeEventHandlerClassOnce.Do(func() {
		ChallengeEventHandlerClass = _ChallengeEventHandlerClass{objc.GetClass("GKChallengeEventHandler")}
	})
	return ChallengeEventHandlerClass
}

type _ChallengeEventHandlerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ChallengeEventHandler */
// An interface definition for the [ChallengeEventHandler] class.
type IChallengeEventHandler interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ChallengeEventHandler */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ChallengeEventHandler */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ChallengeEventHandler */
// Alloc allocates a new instance without initialization.
func (cc _ChallengeEventHandlerClass) Alloc() ChallengeEventHandler {
	rv := objc.Send[ChallengeEventHandler](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _ChallengeEventHandlerClass) New() ChallengeEventHandler {
	rv := objc.Send[ChallengeEventHandler](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ ChallengeEventHandler) Init() ChallengeEventHandler {
	rv := objc.Send[ChallengeEventHandler](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ ChallengeEventHandler) Autorelease() ChallengeEventHandler {
	rv := objc.Send[ChallengeEventHandler](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewChallengeEventHandler creates a new ChallengeEventHandler instance.
func NewChallengeEventHandler() ChallengeEventHandler {
	return getChallengeEventHandlerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ChallengeEventHandler */
// The class is used to respond to events related to challenges sent or received by the local player.
//
// To use it, call the class method to get the instance and assign an object that implements the protocol to its property. You should assign a challenge event handler immediately after initializing the local player, because your game may have launched in response to a challenge notification being received by the player.


// The class is used to respond to events related to challenges sent or received by the local player.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeEventHandler
type ChallengeEventHandler struct {
	objectivec.Object
}

// ChallengeEventHandlerFrom constructs a [ChallengeEventHandler] from an unsafe.Pointer.
//
// The class is used to respond to events related to challenges sent or received by the local player.
func ChallengeEventHandlerFrom(ptr unsafe.Pointer) ChallengeEventHandler {
	return ChallengeEventHandler{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ChallengeEventHandler *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ChallengeEventHandler */

// Returns the shared instance of the event handler
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeEventHandler/challengeEventHandler
func (cc _ChallengeEventHandlerClass) ChallengeEventHandler() IChallengeEventHandler {
	rv := objc.Send[ChallengeEventHandler](objc.ID(cc.class), objc.Sel("challengeEventHandler"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=ChallengeEventHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ChallengeEventHandler */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ChallengeEventHandler */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ChallengeEventHandler */

// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeEventHandler/delegate
func (c_ ChallengeEventHandler) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKChallengeEventHandler/delegate
func (c_ ChallengeEventHandler) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKChallengeEventHandler */



