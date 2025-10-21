// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INRequestPaymentIntent] class.
var (
	INRequestPaymentIntentClass     _INRequestPaymentIntentClass
	INRequestPaymentIntentClassOnce sync.Once
)

func getINRequestPaymentIntentClass() _INRequestPaymentIntentClass {
	INRequestPaymentIntentClassOnce.Do(func() {
		INRequestPaymentIntentClass = _INRequestPaymentIntentClass{objc.GetClass("INRequestPaymentIntent")}
	})
	return INRequestPaymentIntentClass
}

type _INRequestPaymentIntentClass struct {
	class objc.Class
}

// An interface definition for the [INRequestPaymentIntent] class.
type IINRequestPaymentIntent interface {
	IINIntent
}

// An intent for requesting money from another user’s account.
//
// Siri creates an object when the current user requests a payment from another user. A request payment intent object includes the payment amount and the person receiving the request. This intent represents only a request for payment and shouldn’t initiate any payments. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with the results of making the request.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INRequestPaymentIntent
type INRequestPaymentIntent struct {
	INIntent
}

// INRequestPaymentIntentFrom constructs a [INRequestPaymentIntent] from an unsafe.Pointer.
//
// An intent for requesting money from another user’s account.
func INRequestPaymentIntentFrom(ptr unsafe.Pointer) INRequestPaymentIntent {
	return INRequestPaymentIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INRequestPaymentIntentClass) Alloc() INRequestPaymentIntent {
	rv := objc.Send[INRequestPaymentIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INRequestPaymentIntentClass) New() INRequestPaymentIntent {
	rv := objc.Send[INRequestPaymentIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INRequestPaymentIntent) Init() INRequestPaymentIntent {
	rv := objc.Send[INRequestPaymentIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INRequestPaymentIntent) Autorelease() INRequestPaymentIntent {
	rv := objc.Send[INRequestPaymentIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINRequestPaymentIntent creates a new INRequestPaymentIntent instance.
func NewINRequestPaymentIntent() INRequestPaymentIntent {
	return getINRequestPaymentIntentClass().New()
}




