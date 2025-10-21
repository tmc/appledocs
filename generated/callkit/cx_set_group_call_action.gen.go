// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CXSetGroupCallAction] class.
var (
	CXSetGroupCallActionClass     _CXSetGroupCallActionClass
	CXSetGroupCallActionClassOnce sync.Once
)

func getCXSetGroupCallActionClass() _CXSetGroupCallActionClass {
	CXSetGroupCallActionClassOnce.Do(func() {
		CXSetGroupCallActionClass = _CXSetGroupCallActionClass{objc.GetClass("CXSetGroupCallAction")}
	})
	return CXSetGroupCallActionClass
}

type _CXSetGroupCallActionClass struct {
	class objc.Class
}

// An interface definition for the [CXSetGroupCallAction] class.
type ICXSetGroupCallAction interface {
	ICXCallAction
}

// An encapsulation of the act of grouping or ungrouping calls.
//
// is a concrete subclass of . When the user or the system groups a call with another call, the provider sends to its delegate. The provider’s delegate calls the method to indicate that the action was successfully performed. A group call allows more than two recipients to simultaneously communicate with one another.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetGroupCallAction
type CXSetGroupCallAction struct {
	CXCallAction
}

// CXSetGroupCallActionFrom constructs a [CXSetGroupCallAction] from an unsafe.Pointer.
//
// An encapsulation of the act of grouping or ungrouping calls.
func CXSetGroupCallActionFrom(ptr unsafe.Pointer) CXSetGroupCallAction {
	return CXSetGroupCallAction{
		CXCallAction: CXCallActionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CXSetGroupCallActionClass) Alloc() CXSetGroupCallAction {
	rv := objc.Send[CXSetGroupCallAction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CXSetGroupCallActionClass) New() CXSetGroupCallAction {
	rv := objc.Send[CXSetGroupCallAction](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXSetGroupCallAction) Init() CXSetGroupCallAction {
	rv := objc.Send[CXSetGroupCallAction](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXSetGroupCallAction) Autorelease() CXSetGroupCallAction {
	rv := objc.Send[CXSetGroupCallAction](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXSetGroupCallAction creates a new CXSetGroupCallAction instance.
func NewCXSetGroupCallAction() CXSetGroupCallAction {
	return getCXSetGroupCallActionClass().New()
}




// Initializes a new action for a call identified by a given UUID, as well as a call to group with identified by another UUID.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetGroupCallAction/init(call:callUUIDToGroupWith:)
func NewCXSetGroupCallActionWithCallUUIDCallUUIDToGroupWith(callUUID foundation.IUUID, callUUIDToGroupWith foundation.IUUID) CXSetGroupCallAction {
	instance := getCXSetGroupCallActionClass().Alloc()
	rv := objc.Send[CXSetGroupCallAction](instance.ID, objc.Sel("initWithCallUUID:callUUIDToGroupWith:"), callUUID, callUUIDToGroupWith)
	rv.Autorelease()
	return rv
}



// Creates a new action to group calls with data in an unarchiver.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetGroupCallAction/init(coder:)
func NewCXSetGroupCallActionWithCoder(aDecoder foundation.ICoder) CXSetGroupCallAction {
	instance := getCXSetGroupCallActionClass().Alloc()
	rv := objc.Send[CXSetGroupCallAction](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}


// The unique identifier of the call to be grouped with the call associated with the receiver.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetGroupCallAction/callUUIDToGroupWith
func (c_ CXSetGroupCallAction) CallUUIDToGroupWith() foundation.UUID {
	rv := objc.Send[foundation.UUID](c_.ID, objc.Sel("callUUIDToGroupWith"))
	return rv
}


// SetCallUUIDToGroupWith sets the value of the callUUIDToGroupWith property.
// The unique identifier of the call to be grouped with the call associated with the receiver.

//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXSetGroupCallAction/callUUIDToGroupWith
func (c_ CXSetGroupCallAction) SetCallUUIDToGroupWith(value foundation.IUUID) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCallUUIDToGroupWith:"), value)
}


