// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CXPlayDTMFCallAction */


/* debug [class_header]: Header for CXPlayDTMFCallAction */
// The class instance for the [CXPlayDTMFCallAction] class.
var (
	CXPlayDTMFCallActionClass     _CXPlayDTMFCallActionClass
	CXPlayDTMFCallActionClassOnce sync.Once
)

func getCXPlayDTMFCallActionClass() _CXPlayDTMFCallActionClass {
	CXPlayDTMFCallActionClassOnce.Do(func() {
		CXPlayDTMFCallActionClass = _CXPlayDTMFCallActionClass{objc.GetClass("CXPlayDTMFCallAction")}
	})
	return CXPlayDTMFCallActionClass
}

type _CXPlayDTMFCallActionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CXPlayDTMFCallAction */
// An interface definition for the [CXPlayDTMFCallAction] class.
type ICXPlayDTMFCallAction interface {
	ICXCallAction
	
/* debug [class_interface_properties]: Properties for CXPlayDTMFCallAction */
	// properties:
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CXPlayDTMFCallAction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CXPlayDTMFCallAction */
// Alloc allocates a new instance without initialization.
func (cc _CXPlayDTMFCallActionClass) Alloc() CXPlayDTMFCallAction {
	rv := objc.Send[CXPlayDTMFCallAction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CXPlayDTMFCallActionClass) New() CXPlayDTMFCallAction {
	rv := objc.Send[CXPlayDTMFCallAction](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXPlayDTMFCallAction) Init() CXPlayDTMFCallAction {
	rv := objc.Send[CXPlayDTMFCallAction](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXPlayDTMFCallAction) Autorelease() CXPlayDTMFCallAction {
	rv := objc.Send[CXPlayDTMFCallAction](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXPlayDTMFCallAction creates a new CXPlayDTMFCallAction instance.
func NewCXPlayDTMFCallAction() CXPlayDTMFCallAction {
	return getCXPlayDTMFCallActionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CXPlayDTMFCallAction */
// An encapsulation of the act of playing a dual tone multifrequency (DTMF) sequence.
//
// is a concrete subclass of . Whenever digits are transmitted during a call, whether from a user interacting with a number pad or following a hard or soft pause, the provider sends to its delegate. The provider’s delegate calls the method to indicate that the action was successfully performed. The provider sends for successive actions only after the current action is fulfilled. When interacting with the number pad, each entered digit constitutes its own action. Digits following a hard or soft pause, however, are passed to as a single string of digits. For example, if a user taps the 4 button on the number pad, followed by the 2 button, the delegate is sent for the digit and waits for the action to be fulfilled; after the action is fulfilled, the delegate is sent for the digit . CallKit automatically plays the corresponding DTMF frequencies for any digits transmitted over a call. The app is responsible for managing the timing and handling of digits as part of fulfilling the action.


// An encapsulation of the act of playing a dual tone multifrequency (DTMF) sequence.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXPlayDTMFCallAction
type CXPlayDTMFCallAction struct {
	CXCallAction
}

// CXPlayDTMFCallActionFrom constructs a [CXPlayDTMFCallAction] from an unsafe.Pointer.
//
// An encapsulation of the act of playing a dual tone multifrequency (DTMF) sequence.
func CXPlayDTMFCallActionFrom(ptr unsafe.Pointer) CXPlayDTMFCallAction {
	return CXPlayDTMFCallAction{
		CXCallAction: CXCallActionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CXPlayDTMFCallAction */

// Initializes a new action for a call identified by a given UUID, as well as a specified type and sequence of digits.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXPlayDTMFCallAction/init(call:digits:type:)
func NewCXPlayDTMFCallActionWithCallUUIDDigitsType(callUUID foundation.UUID, digits objc.IObject /* cross-framework: NSString */, type_ CXPlayDTMFCallActionType) CXPlayDTMFCallAction {
	instance := getCXPlayDTMFCallActionClass().Alloc()
	rv := objc.Send[CXPlayDTMFCallAction](instance.ID, objc.Sel("initWithCallUUID:digits:type:"), callUUID, digits, type_)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCXPlayDTMFCallActionWithCallUUIDDigitsType */


// Creates a new action to play dual-tone multifrequency (DTMF) tones with data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXPlayDTMFCallAction/init(coder:)
func NewCXPlayDTMFCallActionWithCoder(aDecoder foundation.Coder) CXPlayDTMFCallAction {
	instance := getCXPlayDTMFCallActionClass().Alloc()
	rv := objc.Send[CXPlayDTMFCallAction](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCXPlayDTMFCallActionWithCoder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CXPlayDTMFCallAction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CXPlayDTMFCallAction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CXPlayDTMFCallAction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CXPlayDTMFCallAction */
/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CXPlayDTMFCallAction */


