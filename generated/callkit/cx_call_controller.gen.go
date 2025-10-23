// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CXCallController] class.
var (
	CXCallControllerClass     _CXCallControllerClass
	CXCallControllerClassOnce sync.Once
)

func getCXCallControllerClass() _CXCallControllerClass {
	CXCallControllerClassOnce.Do(func() {
		CXCallControllerClass = _CXCallControllerClass{objc.GetClass("CXCallController")}
	})
	return CXCallControllerClass
}

type _CXCallControllerClass struct {
	class objc.Class
}

// An interface definition for the [CXCallController] class.
type ICXCallController interface {
	objectivec.IObject
	RequestTransactionCompletion(transaction ICXTransaction, completion unsafe.Pointer)
	RequestTransactionWithActionsCompletion(actions []CXAction, completion unsafe.Pointer)
	RequestTransactionWithActionCompletion(action ICXAction, completion unsafe.Pointer)
	CallObserver() CXCallObserver
	CXErrorDomainRequestTransaction() string
}

// A programmatic interface for interacting with and observing calls.
//
// A object interacts with calls by performing actions, which are represented by instances of subclasses. You can request that one or more actions be performed in a single object using the method. A transaction may be rejected by the system for one of the reasons listed in the enumeration. Each object manages a object, which can be accessed using the property. You can provide an object conforming to the protocol to the call observer in order to be notified of any changes to active calls.


// A programmatic interface for interacting with and observing calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallController
type CXCallController struct {
	objectivec.Object
}

// CXCallControllerFrom constructs a [CXCallController] from an unsafe.Pointer.
//
// A programmatic interface for interacting with and observing calls.
func CXCallControllerFrom(ptr unsafe.Pointer) CXCallController {
	return CXCallController{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CXCallControllerClass) Alloc() CXCallController {
	rv := objc.Send[CXCallController](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CXCallControllerClass) New() CXCallController {
	rv := objc.Send[CXCallController](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXCallController) Init() CXCallController {
	rv := objc.Send[CXCallController](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXCallController) Autorelease() CXCallController {
	rv := objc.Send[CXCallController](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXCallController creates a new CXCallController instance.
func NewCXCallController() CXCallController {
	return getCXCallControllerClass().New()
}



// Initializes a new call controller with a specified queue, which is used for calling completion blocks.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallController/init(queue:)
func NewCXCallControllerWithQueue(queue unsafe.Pointer) CXCallController {
	instance := getCXCallControllerClass().Alloc()
	rv := objc.Send[CXCallController](instance.ID, objc.Sel("initWithQueue:"), queue)
	rv.Autorelease()
	return rv
}



// Requests that the actions in the specified transaction be asynchronously performed by the telephony provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallController/request(_:completion:)
func (c_ CXCallController) RequestTransactionCompletion(transaction ICXTransaction, completion unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("requestTransaction:completion:"), transaction, completion)
}


// Requests that the transaction that contains the specified actions be asynchronously performed by the telephony provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallController/requestTransaction(with:completion:)-4o1m4
func (c_ CXCallController) RequestTransactionWithActionsCompletion(actions []CXAction, completion unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("requestTransactionWithActions:completion:"), actions, completion)
}


// Requests that the transaction that contains the specified action be asynchronously performed by the telephony provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallController/requestTransaction(with:completion:)-ffme
func (c_ CXCallController) RequestTransactionWithActionCompletion(action ICXAction, completion unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("requestTransactionWithAction:completion:"), action, completion)
}


// Returns an observer for active calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallController/callObserver
func (c_ CXCallController) CallObserver() CXCallObserver {
	rv := objc.Send[CXCallObserver](c_.ID, objc.Sel("callObserver"))
	return rv
}


// Domain for errors when requesting a transaction from a call controller.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxerrordomainrequesttransaction
func (c_ CXCallController) CXErrorDomainRequestTransaction() string {
	rv := objc.Send[string](c_.ID, objc.Sel("CXErrorDomainRequestTransaction"))
	return rv
}


