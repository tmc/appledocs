// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class CXCall */


/* debug [class_header]: Header for CXCall */
// The class instance for the [CXCall] class.
var (
	CXCallClass     _CXCallClass
	CXCallClassOnce sync.Once
)

func getCXCallClass() _CXCallClass {
	CXCallClassOnce.Do(func() {
		CXCallClass = _CXCallClass{objc.GetClass("CXCall")}
	})
	return CXCallClass
}

type _CXCallClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for CXCall */
// An interface definition for the [CXCall] class.
type ICXCall interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for CXCall */
	// properties:
	IsOnHold() bool
	SetIsOnHold(value bool)
	IsOutgoing() bool
	SetIsOutgoing(value bool)
	Calls() ICXCall
	SetCalls(value ICXCall)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for CXCall */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for CXCall */
// Alloc allocates a new instance without initialization.
func (cc _CXCallClass) Alloc() CXCall {
	rv := objc.Send[CXCall](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (cc _CXCallClass) New() CXCall {
	rv := objc.Send[CXCall](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXCall) Init() CXCall {
	rv := objc.Send[CXCall](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXCall) Autorelease() CXCall {
	rv := objc.Send[CXCall](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXCall creates a new CXCall instance.
func NewCXCall() CXCall {
	return getCXCallClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for CXCall */
// A telephony call.
//
// You don’t instantiate objects directly. Instead, objects are created by the telephony provider when an incoming call is received or an outgoing call is initiated. Each object is uniquely identified by a . You primarily interact with calls by passing their unique identifiers to CallKit APIs. For example, to place a call on hold, you create an instance of with passing the of the call and , create a object containing the action, and then pass the transaction to an instance of using the method. You can use the managed by a to access instances for active calls using the property, or provide an object conforming to the protocol to be notified anytime a call is updated.


// A telephony call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCall
type CXCall struct {
	objectivec.Object
}

// CXCallFrom constructs a [CXCall] from an unsafe.Pointer.
//
// A telephony call.
func CXCallFrom(ptr unsafe.Pointer) CXCall {
	return CXCall{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for CXCall *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for CXCall */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for CXCall */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for CXCall */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for CXCall */

// A Boolean value that indicates whether the call is on hold.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxcall/isonhold
func (c_ CXCall) IsOnHold() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isOnHold"))
	return rv
}/* debug [instance_properties/getter]: isOnHold */


// A Boolean value that indicates whether the call is on hold.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxcall/isonhold
func (c_ CXCall) SetIsOnHold(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsOnHold:"), value)
}/* debug [instance_properties/setter]: isOnHold */


// A Boolean value that indicates whether the call is outgoing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxcall/isoutgoing
func (c_ CXCall) IsOutgoing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isOutgoing"))
	return rv
}/* debug [instance_properties/getter]: isOutgoing */


// A Boolean value that indicates whether the call is outgoing.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxcall/isoutgoing
func (c_ CXCall) SetIsOutgoing(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsOutgoing:"), value)
}/* debug [instance_properties/setter]: isOutgoing */


// Returns the active calls of the telephony provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxcallobserver/calls
func (c_ CXCall) Calls() ICXCall {
	rv := objc.Send[CXCall](c_.ID, objc.Sel("calls"))
	return rv
}/* debug [instance_properties/getter]: calls */


// Returns the active calls of the telephony provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxcallobserver/calls
func (c_ CXCall) SetCalls(value ICXCall) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCalls:"), value)
}/* debug [instance_properties/setter]: calls */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class CXCall */


