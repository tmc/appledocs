// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [CXCall] class.
type ICXCall interface {
	objectivec.IObject
	IsEqualToCall(call unsafe.Pointer) bool
}

// A telephony call.
//
// You don’t instantiate objects directly. Instead, objects are created by the telephony provider when an incoming call is received or an outgoing call is initiated. Each object is uniquely identified by a . You primarily interact with calls by passing their unique identifiers to CallKit APIs. For example, to place a call on hold, you create an instance of with passing the of the call and , create a object containing the action, and then pass the transaction to an instance of using the method. You can use the managed by a to access instances for active calls using the property, or provide an object conforming to the protocol to be notified anytime a call is updated.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CXCallClass) Alloc() CXCall {
	rv := objc.Send[CXCall](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// Returns a Boolean value that indicates whether a given call is equal to the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCall/isEqualToCall:
func (c_ CXCall) IsEqualToCall(call unsafe.Pointer) bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isEqualToCall:"), call)
	return rv
}

// A Boolean value that indicates whether the call has connected.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCall/hasConnected
func (c_ CXCall) HasConnected() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasConnected"))
	return rv
}

// A Boolean value that indicates whether the call has ended.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCall/hasEnded
func (c_ CXCall) HasEnded() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("hasEnded"))
	return rv
}

// A Boolean value that indicates whether the call is on hold.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCall/isOnHold
func (c_ CXCall) OnHold() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("onHold"))
	return rv
}

// A Boolean value that indicates whether the call is outgoing.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCall/isOutgoing
func (c_ CXCall) Outgoing() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("outgoing"))
	return rv
}

// The unique identifier for the call.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCall/uuid
func (c_ CXCall) UUID() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](c_.ID, objc.Sel("UUID"))
	return rv
}



