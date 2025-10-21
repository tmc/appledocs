// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSendPaymentIntentResponse] class.
var (
	INSendPaymentIntentResponseClass     _INSendPaymentIntentResponseClass
	INSendPaymentIntentResponseClassOnce sync.Once
)

func getINSendPaymentIntentResponseClass() _INSendPaymentIntentResponseClass {
	INSendPaymentIntentResponseClassOnce.Do(func() {
		INSendPaymentIntentResponseClass = _INSendPaymentIntentResponseClass{objc.GetClass("INSendPaymentIntentResponse")}
	})
	return INSendPaymentIntentResponseClass
}

type _INSendPaymentIntentResponseClass struct {
	class objc.Class
}

// An interface definition for the [INSendPaymentIntentResponse] class.
type IINSendPaymentIntentResponse interface {
	IINIntentResponse
}

// Your app’s response to a send payment intent.
//
// Use an object to specify the details of the financial transaction that you perform. After creating the object, assign the details of the payment transaction to the property, the details of which Siri communicates to the user at appropriate times. You create an object in the and methods of your handler object. For more information about implementing your handler object, see .
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendPaymentIntentResponse
type INSendPaymentIntentResponse struct {
	INIntentResponse
}

// INSendPaymentIntentResponseFrom constructs a [INSendPaymentIntentResponse] from an unsafe.Pointer.
//
// Your app’s response to a send payment intent.
func INSendPaymentIntentResponseFrom(ptr unsafe.Pointer) INSendPaymentIntentResponse {
	return INSendPaymentIntentResponse{
		INIntentResponse: INIntentResponseFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSendPaymentIntentResponseClass) Alloc() INSendPaymentIntentResponse {
	rv := objc.Send[INSendPaymentIntentResponse](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSendPaymentIntentResponseClass) New() INSendPaymentIntentResponse {
	rv := objc.Send[INSendPaymentIntentResponse](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSendPaymentIntentResponse) Init() INSendPaymentIntentResponse {
	rv := objc.Send[INSendPaymentIntentResponse](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSendPaymentIntentResponse) Autorelease() INSendPaymentIntentResponse {
	rv := objc.Send[INSendPaymentIntentResponse](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSendPaymentIntentResponse creates a new INSendPaymentIntentResponse instance.
func NewINSendPaymentIntentResponse() INSendPaymentIntentResponse {
	return getINSendPaymentIntentResponseClass().New()
}


// The code indicating whether you successfully handled the intent.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendPaymentIntentResponse/code
func (i_ INSendPaymentIntentResponse) Code() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("code"))
	return rv
}

// The details of the payment transaction.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insendpaymentintentresponse/paymentrecord
func (i_ INSendPaymentIntentResponse) PaymentRecord() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("paymentRecord"))
	return rv
}


// SetPaymentRecord sets the value of the paymentRecord property.
// The details of the payment transaction.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insendpaymentintentresponse/paymentrecord
func (i_ INSendPaymentIntentResponse) SetPaymentRecord(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPaymentRecord:"), value)
}



