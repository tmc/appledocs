// Code generated from Apple documentation for CoreTelephony. DO NOT EDIT.

package coretelephony

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Subscriber] class.
var (
	SubscriberClass     _SubscriberClass
	SubscriberClassOnce sync.Once
)

func getSubscriberClass() _SubscriberClass {
	SubscriberClassOnce.Do(func() {
		SubscriberClass = _SubscriberClass{objc.GetClass("CTSubscriber")}
	})
	return SubscriberClass
}

type _SubscriberClass struct {
	class objc.Class
}

// An interface definition for the [Subscriber] class.
type ISubscriber interface {
	objectivec.IObject
	Delegate() objc.ID
	SetDelegate(value objc.ID)
	CarrierToken() foundation.Data
	SetCarrierToken(value foundation.IData)
	Identifier() string
	SetIdentifier(value string)
	IsSIMInserted() bool
	SetIsSIMInserted(value bool)
	CTSubscriberTokenRefreshed() string
}

// A cellular network subscriber.


// A cellular network subscriber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTSubscriber
type Subscriber struct {
	objectivec.Object
}

// SubscriberFrom constructs a [Subscriber] from an unsafe.Pointer.
//
// A cellular network subscriber.
func SubscriberFrom(ptr unsafe.Pointer) Subscriber {
	return Subscriber{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _SubscriberClass) Alloc() Subscriber {
	rv := objc.Send[Subscriber](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _SubscriberClass) New() Subscriber {
	rv := objc.Send[Subscriber](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Subscriber) Init() Subscriber {
	rv := objc.Send[Subscriber](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Subscriber) Autorelease() Subscriber {
	rv := objc.Send[Subscriber](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewSubscriber creates a new Subscriber instance.
func NewSubscriber() Subscriber {
	return getSubscriberClass().New()
}



// A delegate that receives updates on the subscriber information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTSubscriber/delegate
func (s_ Subscriber) Delegate() objc.ID {
	rv := objc.Send[objc.ID](s_.ID, objc.Sel("delegate"))
	return rv
}


// A delegate that receives updates on the subscriber information.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/CoreTelephony/CTSubscriber/delegate
func (s_ Subscriber) SetDelegate(value objc.ID) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setDelegate:"), value)
}


// A data object containing authorization information about the subscriber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctsubscriber/carriertoken
func (s_ Subscriber) CarrierToken() foundation.Data {
	rv := objc.Send[foundation.Data](s_.ID, objc.Sel("carrierToken"))
	return rv
}


// A data object containing authorization information about the subscriber.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctsubscriber/carriertoken
func (s_ Subscriber) SetCarrierToken(value foundation.IData) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setCarrierToken:"), value)
}


// An implementation-defined identifier used to correlate this subscriber with information vended by other APIs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctsubscriber/identifier
func (s_ Subscriber) Identifier() string {
	rv := objc.Send[string](s_.ID, objc.Sel("identifier"))
	return rv
}


// An implementation-defined identifier used to correlate this subscriber with information vended by other APIs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctsubscriber/identifier
func (s_ Subscriber) SetIdentifier(value string) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIdentifier:"), objc.String(value))
}


// A Boolean property that indicates whether a SIM is present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctsubscriber/issiminserted
func (s_ Subscriber) IsSIMInserted() bool {
	rv := objc.Send[bool](s_.ID, objc.Sel("isSIMInserted"))
	return rv
}


// A Boolean property that indicates whether a SIM is present.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctsubscriber/issiminserted
func (s_ Subscriber) SetIsSIMInserted(value bool) {
	objc.Send[objc.ID](s_.ID, objc.Sel("setIsSIMInserted:"), value)
}


// The name of the notification indicating that the carrier token is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/coretelephony/ctsubscribertokenrefreshed
func (s_ Subscriber) CTSubscriberTokenRefreshed() string {
	rv := objc.Send[string](s_.ID, objc.Sel("CTSubscriberTokenRefreshed"))
	return rv
}



