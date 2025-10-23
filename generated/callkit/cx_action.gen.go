// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CXAction] class.
var (
	CXActionClass     _CXActionClass
	CXActionClassOnce sync.Once
)

func getCXActionClass() _CXActionClass {
	CXActionClassOnce.Do(func() {
		CXActionClass = _CXActionClass{objc.GetClass("CXAction")}
	})
	return CXActionClass
}

type _CXActionClass struct {
	class objc.Class
}

// An interface definition for the [CXAction] class.
type ICXAction interface {
	objectivec.IObject
	Fail()
	Fulfill()
	Complete() bool
	TimeoutDate() foundation.NSDate
	UUID() foundation.UUID
	IsComplete() bool
	SetIsComplete(value bool)
}

// An abstract class that declares a programmatic interface for objects that represent a telephony action.
//
// Each instance of is uniquely identified by a , which is generated on initialization. An action also tracks whether it has been completed or not. To perform one or more actions, you add them to a new object and pass the transaction to an instance of using the method. After each action is performed by the telephony provider, the provider’s delegate calls either the method, indicating that the action was successfully performed, or the method, to indicate that an error occurred; both of these methods set the property of the action to . The subclass is an abstract class that represents an action associated with a object. The CallKit framework provides several concrete subclasses to represent actions such as answering a call and putting a call on hold.


// An abstract class that declares a programmatic interface for objects that represent a telephony action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXAction
type CXAction struct {
	objectivec.Object
}

// CXActionFrom constructs a [CXAction] from an unsafe.Pointer.
//
// An abstract class that declares a programmatic interface for objects that represent a telephony action.
func CXActionFrom(ptr unsafe.Pointer) CXAction {
	return CXAction{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CXActionClass) Alloc() CXAction {
	rv := objc.Send[CXAction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CXActionClass) New() CXAction {
	rv := objc.Send[CXAction](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXAction) Init() CXAction {
	rv := objc.Send[CXAction](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXAction) Autorelease() CXAction {
	rv := objc.Send[CXAction](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXAction creates a new CXAction instance.
func NewCXAction() CXAction {
	return getCXActionClass().New()
}



// Creates a new telephony action with data in an unarchiver.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXAction/init(coder:)
func NewCXActionWithCoder(aDecoder foundation.ICoder) CXAction {
	instance := getCXActionClass().Alloc()
	rv := objc.Send[CXAction](instance.ID, objc.Sel("initWithCoder:"), aDecoder)
	rv.Autorelease()
	return rv
}



// Reports the failed execution of the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXAction/fail()
func (c_ CXAction) Fail() {
	objc.Send[objc.ID](c_.ID, objc.Sel("fail"))
}


// Reports the successful execution of the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXAction/fulfill()
func (c_ CXAction) Fulfill() {
	objc.Send[objc.ID](c_.ID, objc.Sel("fulfill"))
}


// A Boolean value that indicates whether the action has been performed by the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXAction/isComplete
func (c_ CXAction) Complete() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("complete"))
	return rv
}


// The time after which the action cannot be completed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXAction/timeoutDate
func (c_ CXAction) TimeoutDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](c_.ID, objc.Sel("timeoutDate"))
	return rv
}


// The unique identifier for the action.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXAction/uuid
func (c_ CXAction) UUID() foundation.UUID {
	rv := objc.Send[foundation.UUID](c_.ID, objc.Sel("UUID"))
	return rv
}


// A Boolean value that indicates whether the action has been performed by the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxaction/iscomplete
func (c_ CXAction) IsComplete() bool {
	rv := objc.Send[bool](c_.ID, objc.Sel("isComplete"))
	return rv
}


// A Boolean value that indicates whether the action has been performed by the provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxaction/iscomplete
func (c_ CXAction) SetIsComplete(value bool) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setIsComplete:"), value)
}


