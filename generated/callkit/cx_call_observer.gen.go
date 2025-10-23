// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [CXCallObserver] class.
var (
	CXCallObserverClass     _CXCallObserverClass
	CXCallObserverClassOnce sync.Once
)

func getCXCallObserverClass() _CXCallObserverClass {
	CXCallObserverClassOnce.Do(func() {
		CXCallObserverClass = _CXCallObserverClass{objc.GetClass("CXCallObserver")}
	})
	return CXCallObserverClass
}

type _CXCallObserverClass struct {
	class objc.Class
}

// An interface definition for the [CXCallObserver] class.
type ICXCallObserver interface {
	objectivec.IObject
	SetDelegateQueue(delegate objectivec.IObject, queue unsafe.Pointer)
	Calls() []CXCall
	CallObserver() CXCallObserver
	SetCallObserver(value ICXCallObserver)
}

// A programmatic interface for an object that manages a list of active calls and observes call changes.
//
// You can retrieve a list of active calls on an object using the property. You can also provide an object conforming to the protocol as the call observer delegate using the method to respond to any active call changes. VoIP apps typically interact with the object returned by the property of a instance. However, any app can create a new object to be notified of any calls activity on the system.


// A programmatic interface for an object that manages a list of active calls and observes call changes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallObserver
type CXCallObserver struct {
	objectivec.Object
}

// CXCallObserverFrom constructs a [CXCallObserver] from an unsafe.Pointer.
//
// A programmatic interface for an object that manages a list of active calls and observes call changes.
func CXCallObserverFrom(ptr unsafe.Pointer) CXCallObserver {
	return CXCallObserver{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (cc _CXCallObserverClass) Alloc() CXCallObserver {
	rv := objc.Send[CXCallObserver](objc.ID(cc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (cc _CXCallObserverClass) New() CXCallObserver {
	rv := objc.Send[CXCallObserver](objc.ID(cc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (c_ CXCallObserver) Init() CXCallObserver {
	rv := objc.Send[CXCallObserver](c_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (c_ CXCallObserver) Autorelease() CXCallObserver {
	rv := objc.Send[CXCallObserver](c_.ID, objc.Sel("autorelease"))
	return rv
}

// NewCXCallObserver creates a new CXCallObserver instance.
func NewCXCallObserver() CXCallObserver {
	return getCXCallObserverClass().New()
}



// Sets a call observer delegate, specifying an optional queue on which to execute delegate methods.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallObserver/setDelegate(_:queue:)
func (c_ CXCallObserver) SetDelegateQueue(delegate objectivec.IObject, queue unsafe.Pointer) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setDelegate:queue:"), delegate, queue)
}


// Returns the active calls of the telephony provider.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CallKit/CXCallObserver/calls
func (c_ CXCallObserver) Calls() []CXCall {
	rv := objc.Send[[]CXCall](c_.ID, objc.Sel("calls"))
	return rv
}


// Returns an observer for active calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxcallcontroller/callobserver
func (c_ CXCallObserver) CallObserver() CXCallObserver {
	rv := objc.Send[CXCallObserver](c_.ID, objc.Sel("callObserver"))
	return rv
}


// Returns an observer for active calls.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/callkit/cxcallcontroller/callobserver
func (c_ CXCallObserver) SetCallObserver(value ICXCallObserver) {
	objc.Send[objc.ID](c_.ID, objc.Sel("setCallObserver:"), value)
}



