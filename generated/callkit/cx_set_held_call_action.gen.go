// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

/* debug [class.gen.go]: Generating class CXSetHeldCallAction */


/* debug [class_header]: Header for CXSetHeldCallAction */
// The class instance for the [CXSetHeldCallAction] class.
var (
	CXSetHeldCallActionClass     _CXSetHeldCallActionClass
	CXSetHeldCallActionClassOnce sync.Once
)

func getCXSetHeldCallActionClass() _CXSetHeldCallActionClass {
	CXSetHeldCallActionClassOnce.Do(func() {
		CXSetHeldCallActionClass = _CXSetHeldCallActionClass{objc.GetClass("CXSetHeldCallAction")}
	})
	return CXSetHeldCallActionClass
}

type _CXSetHeldCallActionClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CXSetHeldCallAction */
// An interface definition for the [CXSetHeldCallAction] class.
type ICXSetHeldCallAction interface {
	ICXCallAction
	
/* debug [class_interface_properties]: Properties for CXSetHeldCallAction */
	// properties:
	IsOnHold() bool
	SetIsOnHold(value bool)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CXSetHeldCallAction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CXSetHeldCallAction */
// Alloc allocates a new instance without initialization.
func (cc _CXSetHeldCallActionClass) Alloc() CXSetHeldCallAction {
	rv := objc.Send[CXSetHeldCallAction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CXSetHeldCallActionClass) New() CXSetHeldCallAction {
	rv := objc.Send[CXSetHeldCallAction](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXSetHeldCallAction) Init() CXSetHeldCallAction {
	rv := objc.Send[CXSetHeldCallAction](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXSetHeldCallAction) Autorelease() CXSetHeldCallAction {
	rv := objc.Send[CXSetHeldCallAction](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXSetHeldCallAction creates a new CXSetHeldCallAction instance.
func NewCXSetHeldCallAction() CXSetHeldCallAction {
	return getCXSetHeldCallActionClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CXSetHeldCallAction */
// An encapsulation of the act of placing a call on hold or removing a call from hold.
//
// is a concrete subclass of . When a caller places a call on hold, callers are unable to communicate with one another until the holding caller removes the call from hold. Placing a call on hold doesn’t end the call. When the user or the system places a call on hold, the provider sends to its delegate. The provider’s delegate calls the method to indicate that the action was successfully performed.


// An encapsulation of the act of placing a call on hold or removing a call from hold.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetHeldCallAction
type CXSetHeldCallAction struct {
	CXCallAction
}

// CXSetHeldCallActionFrom constructs a [CXSetHeldCallAction] from an unsafe.Pointer.
//
// An encapsulation of the act of placing a call on hold or removing a call from hold.
func CXSetHeldCallActionFrom(ptr unsafe.Pointer) CXSetHeldCallAction {
	return CXSetHeldCallAction{
		CXCallAction: CXCallActionFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CXSetHeldCallAction */

// Initializes a new action for a call identified by a given UUID, as well as whether the call is on hold.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetHeldCallAction/init(call:onHold:)
func NewCXSetHeldCallActionWithCallUUIDOnHold(callUUID foundation.UUID, onHold bool) CXSetHeldCallAction {
	instance := getCXSetHeldCallActionClass().Alloc()
	rv := objc.Send[CXSetHeldCallAction](instance.ID, objc.Sel("initWithCallUUID:onHold:"), callUUID, onHold)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCXSetHeldCallActionWithCallUUIDOnHold */


// Creates a new action to place a call on hold with data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetHeldCallAction/init(coder:)
func NewCXSetHeldCallActionWithCoder(aDecoder foundation.Coder) CXSetHeldCallAction {
	instance := getCXSetHeldCallActionClass().Alloc()
	rv := objc.Send[CXSetHeldCallAction](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewCXSetHeldCallActionWithCoder */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CXSetHeldCallAction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CXSetHeldCallAction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CXSetHeldCallAction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CXSetHeldCallAction */

// A Boolean value that indicates whether the call is placed on hold.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxsetheldcallaction/isonhold
func (c_ CXSetHeldCallAction) IsOnHold() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isOnHold"))
	return rv
}/* debug [instance_properties/getter]: isOnHold */


// A Boolean value that indicates whether the call is placed on hold.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxsetheldcallaction/isonhold
func (c_ CXSetHeldCallAction) SetIsOnHold(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsOnHold:"), value)
}/* debug [instance_properties/setter]: isOnHold */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CXSetHeldCallAction */


