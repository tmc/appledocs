// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [CXSetHeldCallAction] class.
type ICXSetHeldCallAction interface {
	ICXCallAction
}

// An encapsulation of the act of placing a call on hold or removing a call from hold.
//
// is a concrete subclass of . When a caller places a call on hold, callers are unable to communicate with one another until the holding caller removes the call from hold. Placing a call on hold doesn’t end the call. When the user or the system places a call on hold, the provider sends to its delegate. The provider’s delegate calls the method to indicate that the action was successfully performed.
//
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

// Alloc allocates a new instance without initialization.
func (cc _CXSetHeldCallActionClass) Alloc() CXSetHeldCallAction {
	rv := objc.Send[CXSetHeldCallAction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Initializes a new action for a call identified by a given UUID, as well as whether the call is on hold.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetHeldCallAction/init(call:onHold:)
func NewCXSetHeldCallActionWithCallUUIDOnHold(callUUID unsafe.Pointer, onHold bool) CXSetHeldCallAction {
	instance := getCXSetHeldCallActionClass().Alloc()
	rv := objc.Send[CXSetHeldCallAction](instance.ID, objc.Sel("initWithCallUUID:onHold:"), callUUID, onHold)
	rv.Autorelease()
	return rv
}



// Creates a new action to place a call on hold with data in an unarchiver.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetHeldCallAction/init(coder:)
func NewCXSetHeldCallActionWithCoder(aDecoder unsafe.Pointer) CXSetHeldCallAction {
	instance := getCXSetHeldCallActionClass().Alloc()
	rv := objc.Send[CXSetHeldCallAction](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}


// A Boolean value that indicates whether the call is placed on hold.
//
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxsetheldcallaction/isonhold
func (c_ CXSetHeldCallAction) IsOnHold() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isOnHold"))
	return rv
}


// SetIsOnHold sets the value of the isOnHold property.
// A Boolean value that indicates whether the call is placed on hold.

//
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxsetheldcallaction/isonhold
func (c_ CXSetHeldCallAction) SetIsOnHold(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsOnHold:"), value)
}

// A Boolean value that indicates whether the call is placed on hold.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetHeldCallAction/isOnHold
func (c_ CXSetHeldCallAction) OnHold() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("onHold"))
	return rv
}


// SetOnHold sets the value of the onHold property.
// A Boolean value that indicates whether the call is placed on hold.

//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetHeldCallAction/isOnHold
func (c_ CXSetHeldCallAction) SetOnHold(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setOnHold:"), value)
}


