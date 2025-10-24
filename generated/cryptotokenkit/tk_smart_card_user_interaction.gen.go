// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class TKSmartCardUserInteraction */


/* debug [class_header]: Header for TKSmartCardUserInteraction */
// The class instance for the [TKSmartCardUserInteraction] class.
var (
	TKSmartCardUserInteractionClass     _TKSmartCardUserInteractionClass
	TKSmartCardUserInteractionClassOnce sync.Once
)

func getTKSmartCardUserInteractionClass() _TKSmartCardUserInteractionClass {
	TKSmartCardUserInteractionClassOnce.Do(func() {
		TKSmartCardUserInteractionClass = _TKSmartCardUserInteractionClass{objc.GetClass("TKSmartCardUserInteraction")}
	})
	return TKSmartCardUserInteractionClass
}

type _TKSmartCardUserInteractionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for TKSmartCardUserInteraction */
// An interface definition for the [TKSmartCardUserInteraction] class.
type ITKSmartCardUserInteraction interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for TKSmartCardUserInteraction */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	InitialTimeout() float64
	SetInitialTimeout(value float64)
	InteractionTimeout() float64
	SetInteractionTimeout(value float64)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for TKSmartCardUserInteraction */
	// methods:
	Cancel() bool
	RunWithReply(reply unsafe.Pointer)
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for TKSmartCardUserInteraction */
// Alloc allocates a new instance without initialization.
func (tc _TKSmartCardUserInteractionClass) Alloc() TKSmartCardUserInteraction {
	rv := objc.Send[TKSmartCardUserInteraction](objc.ID(tc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (tc _TKSmartCardUserInteractionClass) New() TKSmartCardUserInteraction {
	rv := objc.Send[TKSmartCardUserInteraction](objc.ID(tc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (t_ TKSmartCardUserInteraction) Init() TKSmartCardUserInteraction {
	rv := objc.Send[TKSmartCardUserInteraction](t_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (t_ TKSmartCardUserInteraction) Autorelease() TKSmartCardUserInteraction {
	rv := objc.Send[TKSmartCardUserInteraction](t_.ID, objc.Sel("autorelease"))
	return rv
}

// NewTKSmartCardUserInteraction creates a new TKSmartCardUserInteraction instance.
func NewTKSmartCardUserInteraction() TKSmartCardUserInteraction {
	return getTKSmartCardUserInteractionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for TKSmartCardUserInteraction */
// The base class for encapsulating user interaction with a Smart Card reader.
//
// There are two types of user interactions: those for secure PIN change and those for secure PIN validation. These interactions are instances of the , or subclasses of , respectively. is a subclass of . You interact with instances of one of the subclasses of when calling the and methods on an object.


// The base class for encapsulating user interaction with a Smart Card reader.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction
type TKSmartCardUserInteraction struct {
	objectivec.Object
}

// TKSmartCardUserInteractionFrom constructs a [TKSmartCardUserInteraction] from an unsafe.Pointer.
//
// The base class for encapsulating user interaction with a Smart Card reader.
func TKSmartCardUserInteractionFrom(ptr unsafe.Pointer) TKSmartCardUserInteraction {
	return TKSmartCardUserInteraction{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for TKSmartCardUserInteraction *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for TKSmartCardUserInteraction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for TKSmartCardUserInteraction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for TKSmartCardUserInteraction */

// Attempts to cancel an interaction started by calling . For certain interactions, cancellation may not be available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction/cancel()
func (t_ TKSmartCardUserInteraction) Cancel() bool {
	rv := objc.Send[bool](t_.ID, objc.Sel("cancel"))
	return rv
}/* debug [instance_methods/method]: Cancel */


// Runs the user interaction and asynchronously receives a reply.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction/run(reply:)
func (t_ TKSmartCardUserInteraction) RunWithReply(reply unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("runWithReply:"), reply)
}/* debug [instance_methods/method]: RunWithReply */

/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for TKSmartCardUserInteraction */

// The delegate for observing events that occur during the user interaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction/delegate
func (t_ TKSmartCardUserInteraction) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](t_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate for observing events that occur during the user interaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction/delegate
func (t_ TKSmartCardUserInteraction) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The timeout, in seconds, for initial interaction. If set to , the reader-defined default timeout is used. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction/initialTimeout
func (t_ TKSmartCardUserInteraction) InitialTimeout() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("initialTimeout"))
	return rv
}/* debug [instance_properties/getter]: initialTimeout */


// The timeout, in seconds, for initial interaction. If set to , the reader-defined default timeout is used. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction/initialTimeout
func (t_ TKSmartCardUserInteraction) SetInitialTimeout(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setInitialTimeout:"), value)
}/* debug [instance_properties/setter]: initialTimeout */


// The timeout, in seconds, after the first key stroke. If set to , the reader-defined default timeout is used. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction/interactionTimeout
func (t_ TKSmartCardUserInteraction) InteractionTimeout() float64 {
	rv := objc.Send[float64](t_.ID, objc.Sel("interactionTimeout"))
	return rv
}/* debug [instance_properties/getter]: interactionTimeout */


// The timeout, in seconds, after the first key stroke. If set to , the reader-defined default timeout is used. by default.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CryptoTokenKit/TKSmartCardUserInteraction/interactionTimeout
func (t_ TKSmartCardUserInteraction) SetInteractionTimeout(value float64) {
	objc.Send[objc.ID](t_.ID, objc.Sel("setInteractionTimeout:"), value)
}/* debug [instance_properties/setter]: interactionTimeout */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class TKSmartCardUserInteraction */



