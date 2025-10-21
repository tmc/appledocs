// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [CXEndCallAction] class.
var (
	CXEndCallActionClass     _CXEndCallActionClass
	CXEndCallActionClassOnce sync.Once
)

func getCXEndCallActionClass() _CXEndCallActionClass {
	CXEndCallActionClassOnce.Do(func() {
		CXEndCallActionClass = _CXEndCallActionClass{objc.GetClass("CXEndCallAction")}
	})
	return CXEndCallActionClass
}

type _CXEndCallActionClass struct {
	class objc.Class
}

// An interface definition for the [CXEndCallAction] class.
type ICXEndCallAction interface {
	ICXCallAction
	FulfillWithDateEnded(dateEnded foundation.IDate)
}

// An encapsulation of the act of ending a call.
//
// is a concrete subclass of . When the user initiates an outgoing call, the provider sends to its delegate. The provider’s delegate calls the method to indicate that the action was successfully performed. To indicate that the call ended at a time other than the current time, you can instead call the
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXEndCallAction
type CXEndCallAction struct {
	CXCallAction
}

// CXEndCallActionFrom constructs a [CXEndCallAction] from an unsafe.Pointer.
//
// An encapsulation of the act of ending a call.
func CXEndCallActionFrom(ptr unsafe.Pointer) CXEndCallAction {
	return CXEndCallAction{
		CXCallAction: CXCallActionFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (cc _CXEndCallActionClass) Alloc() CXEndCallAction {
	rv := objc.Send[CXEndCallAction](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CXEndCallActionClass) New() CXEndCallAction {
	rv := objc.Send[CXEndCallAction](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXEndCallAction) Init() CXEndCallAction {
	rv := objc.Send[CXEndCallAction](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXEndCallAction) Autorelease() CXEndCallAction {
	rv := objc.Send[CXEndCallAction](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXEndCallAction creates a new CXEndCallAction instance.
func NewCXEndCallAction() CXEndCallAction {
	return getCXEndCallActionClass().New()
}


// Reports the successful execution of the action at the specified time.
//
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXEndCallAction/fulfill(withDateEnded:)
func (c_ CXEndCallAction) FulfillWithDateEnded(dateEnded foundation.IDate) {
	objc.Send[objc.ID](c_.ID, objc.Sel("fulfillWithDateEnded:"), dateEnded)
}



