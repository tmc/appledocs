// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CXAnswerCallAction] class.
var (
	CXAnswerCallActionClass     _CXAnswerCallActionClass
	CXAnswerCallActionClassOnce sync.Once
)

func getCXAnswerCallActionClass() _CXAnswerCallActionClass {
	CXAnswerCallActionClassOnce.Do(func() {
		CXAnswerCallActionClass = _CXAnswerCallActionClass{objc.GetClass("CXAnswerCallAction")}
	})
	return CXAnswerCallActionClass
}

type _CXAnswerCallActionClass struct {
	class objc.Class
}

// An interface definition for the [CXAnswerCallAction] class.
type ICXAnswerCallAction interface {
	ICXCallAction
	// properties:
	// methods:
	FulfillWithDateConnected(dateConnected foundation.objc.IObject /* cross-framework NSDate */)
}

// An encapsulation of the act of answering an incoming call.
//
// is a concrete subclass of . When an incoming call is allowed by the system and approved by the user, the provider sends to its delegate. The provider’s delegate calls the method to indicate that the action was successfully performed. To indicate that the call connected at a time other than the current time, you can instead call the .


// An encapsulation of the act of answering an incoming call.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXAnswerCallAction
type CXAnswerCallAction struct {
	CXCallAction
}

// CXAnswerCallActionFrom constructs a [CXAnswerCallAction] from an unsafe.Pointer.
//
// An encapsulation of the act of answering an incoming call.
func CXAnswerCallActionFrom(ptr unsafe.Pointer) CXAnswerCallAction {
	return CXAnswerCallAction{
		CXCallAction: CXCallActionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CXAnswerCallActionClass) Alloc() CXAnswerCallAction {
	rv := objc.Send[CXAnswerCallAction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CXAnswerCallActionClass) New() CXAnswerCallAction {
	rv := objc.Send[CXAnswerCallAction](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXAnswerCallAction) Init() CXAnswerCallAction {
	rv := objc.Send[CXAnswerCallAction](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXAnswerCallAction) Autorelease() CXAnswerCallAction {
	rv := objc.Send[CXAnswerCallAction](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXAnswerCallAction creates a new CXAnswerCallAction instance.
func NewCXAnswerCallAction() CXAnswerCallAction {
	return getCXAnswerCallActionClass().New()
}



// Reports the successful execution of the action at the specified time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXAnswerCallAction/fulfill(withDateConnected:)
func (c_ CXAnswerCallAction) FulfillWithDateConnected(dateConnected foundation.objc.IObject /* cross-framework NSDate */) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fulfillWithDateConnected:"), dateConnected)
}



