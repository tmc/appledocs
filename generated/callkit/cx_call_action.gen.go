// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CXCallAction] class.
var (
	CXCallActionClass     _CXCallActionClass
	CXCallActionClassOnce sync.Once
)

func getCXCallActionClass() _CXCallActionClass {
	CXCallActionClassOnce.Do(func() {
		CXCallActionClass = _CXCallActionClass{objc.GetClass("CXCallAction")}
	})
	return CXCallActionClass
}

type _CXCallActionClass struct {
	class objc.Class
}

// An interface definition for the [CXCallAction] class.
type ICXCallAction interface {
	ICXAction
	CallUUID() foundation.UUID
}

// A programmatic interface for objects that represent a telephony action associated with a call object.
//
// The CallKit framework provides the following concrete subclasses. To perform one or more actions, you add them to a new object and pass the transaction to an instance of using the method. After each action is performed by the telephony provider, the provider’s delegate calls either the method, indicating that the action was successfully performed, or the method, to indicate that an error occurred; both of these methods set the property of the action to .


// A programmatic interface for objects that represent a telephony action associated with a call object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallAction
type CXCallAction struct {
	CXAction
}

// CXCallActionFrom constructs a [CXCallAction] from an unsafe.Pointer.
//
// A programmatic interface for objects that represent a telephony action associated with a call object.
func CXCallActionFrom(ptr unsafe.Pointer) CXCallAction {
	return CXCallAction{
		CXAction: CXActionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CXCallActionClass) Alloc() CXCallAction {
	rv := objc.Send[CXCallAction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CXCallActionClass) New() CXCallAction {
	rv := objc.Send[CXCallAction](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXCallAction) Init() CXCallAction {
	rv := objc.Send[CXCallAction](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXCallAction) Autorelease() CXCallAction {
	rv := objc.Send[CXCallAction](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXCallAction creates a new CXCallAction instance.
func NewCXCallAction() CXCallAction {
	return getCXCallActionClass().New()
}



// Initializes a new action for a call identified by a given UUID.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallAction/init(call:)
func NewCXCallActionWithCallUUID(callUUID foundation.IUUID) CXCallAction {
	instance := getCXCallActionClass().Alloc()
	rv := objc.Send[CXCallAction](instance.ID, objc.Sel("initWithCallUUID:"), callUUID)
	rv.Autorelease()
	return rv
}


// Creates a new action for a call with data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallAction/init(coder:)
func NewCXCallActionWithCoder(aDecoder foundation.ICoder) CXCallAction {
	instance := getCXCallActionClass().Alloc()
	rv := objc.Send[CXCallAction](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}



// The unique identifier for the call associated with the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallAction/callUUID
func (c_ CXCallAction) CallUUID() foundation.UUID {
	rv := objc.Send[foundation.UUID](c_.ID, objc.Sel("callUUID"))
	return rv
}


