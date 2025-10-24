// Code generated from Apple documentation for GameKit. DO NOT EDIT.

package gamekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class GKTurnBasedEventHandler */


/* debug [class_header]: Header for GKTurnBasedEventHandler */
// The class instance for the [TurnBasedEventHandler] class.
var (
	TurnBasedEventHandlerClass     _TurnBasedEventHandlerClass
	TurnBasedEventHandlerClassOnce sync.Once
)

func getTurnBasedEventHandlerClass() _TurnBasedEventHandlerClass {
	TurnBasedEventHandlerClassOnce.Do(func() {
		TurnBasedEventHandlerClass = _TurnBasedEventHandlerClass{objc.GetClass("GKTurnBasedEventHandler")}
	})
	return TurnBasedEventHandlerClass
}

type _TurnBasedEventHandlerClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TurnBasedEventHandler */
// An interface definition for the [TurnBasedEventHandler] class.
type ITurnBasedEventHandler interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TurnBasedEventHandler */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TurnBasedEventHandler */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TurnBasedEventHandler */
// Alloc allocates a new instance without initialization.
func (tc _TurnBasedEventHandlerClass) Alloc() TurnBasedEventHandler {
	rv := objc.Send[TurnBasedEventHandler](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TurnBasedEventHandlerClass) New() TurnBasedEventHandler {
	rv := objc.Send[TurnBasedEventHandler](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TurnBasedEventHandler) Init() TurnBasedEventHandler {
	rv := objc.Send[TurnBasedEventHandler](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TurnBasedEventHandler) Autorelease() TurnBasedEventHandler {
	rv := objc.Send[TurnBasedEventHandler](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTurnBasedEventHandler creates a new TurnBasedEventHandler instance.
func NewTurnBasedEventHandler() TurnBasedEventHandler {
	return getTurnBasedEventHandlerClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TurnBasedEventHandler */
// The class is used to respond to important messages related to turn-based matches. To use it, call the class method to get the singleton instance and assign an object that implements the protocol to its property. All methods are called on the main thread.
//
// This framework has been deprecated in iOS 7. Use .


// The class is used to respond to important messages related to turn-based matches. To use it, call the class method to get the singleton instance and assign an object that implements the protocol to its property. All methods are called on the main thread.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedEventHandler
type TurnBasedEventHandler struct {
	objectivec.Object
}

// TurnBasedEventHandlerFrom constructs a [TurnBasedEventHandler] from an unsafe.Pointer.
//
// The class is used to respond to important messages related to turn-based matches. To use it, call the class method to get the singleton instance and assign an object that implements the protocol to its property. All methods are called on the main thread.
func TurnBasedEventHandlerFrom(ptr unsafe.Pointer) TurnBasedEventHandler {
	return TurnBasedEventHandler{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TurnBasedEventHandler *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TurnBasedEventHandler */

// Returns the shared instance of the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedEventHandler/shared()
func (tc _TurnBasedEventHandlerClass) SharedTurnBasedEventHandler() ITurnBasedEventHandler {
	rv := objc.Send[TurnBasedEventHandler](objc.ID(tc.class), objc.Sel("sharedTurnBasedEventHandler"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=SharedTurnBasedEventHandler) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TurnBasedEventHandler */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TurnBasedEventHandler */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TurnBasedEventHandler */

// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedEventHandler/delegate
func (t_ TurnBasedEventHandler) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate for the event handler.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/GameKit/GKTurnBasedEventHandler/delegate
func (t_ TurnBasedEventHandler) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class GKTurnBasedEventHandler */



