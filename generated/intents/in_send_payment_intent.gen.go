// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSendPaymentIntent] class.
var (
	INSendPaymentIntentClass     _INSendPaymentIntentClass
	INSendPaymentIntentClassOnce sync.Once
)

func getINSendPaymentIntentClass() _INSendPaymentIntentClass {
	INSendPaymentIntentClassOnce.Do(func() {
		INSendPaymentIntentClass = _INSendPaymentIntentClass{objc.GetClass("INSendPaymentIntent")}
	})
	return INSendPaymentIntentClass
}

type _INSendPaymentIntentClass struct {
	class objc.Class
}

// An interface definition for the [INSendPaymentIntent] class.
type IINSendPaymentIntent interface {
	IINIntent
}

// A request to transfer money from the current user’s account to a different user’s account.
//
// Siri creates an object when the current user asks to transfer money to another user. A send payment intent object includes the payment amount and the recipient of the payment. Use that information to validate the transaction and transfer the funds. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with the result of sending the money to the specified user. This intent object represents a financial transaction between two users.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSendPaymentIntent
type INSendPaymentIntent struct {
	INIntent
}

// INSendPaymentIntentFrom constructs a [INSendPaymentIntent] from an unsafe.Pointer.
//
// A request to transfer money from the current user’s account to a different user’s account.
func INSendPaymentIntentFrom(ptr unsafe.Pointer) INSendPaymentIntent {
	return INSendPaymentIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSendPaymentIntentClass) Alloc() INSendPaymentIntent {
	rv := objc.Send[INSendPaymentIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSendPaymentIntentClass) New() INSendPaymentIntent {
	rv := objc.Send[INSendPaymentIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSendPaymentIntent) Init() INSendPaymentIntent {
	rv := objc.Send[INSendPaymentIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSendPaymentIntent) Autorelease() INSendPaymentIntent {
	rv := objc.Send[INSendPaymentIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSendPaymentIntent creates a new INSendPaymentIntent instance.
func NewINSendPaymentIntent() INSendPaymentIntent {
	return getINSendPaymentIntentClass().New()
}


// The amount of the payment.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insendpaymentintent/currencyamount
func (i_ INSendPaymentIntent) CurrencyAmount() INCurrencyAmount {
	rv := objc.Send[INCurrencyAmount](i_.ID, objc.Sel("currencyAmount"))
	return rv
}


// SetCurrencyAmount sets the value of the currencyAmount property.
// The amount of the payment.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insendpaymentintent/currencyamount
func (i_ INSendPaymentIntent) SetCurrencyAmount(value INCurrencyAmount) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setCurrencyAmount:"), value)
}

// A note associated with the payment.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insendpaymentintent/note
func (i_ INSendPaymentIntent) Note() appkit.string {
	rv := objc.Send[appkit.string](i_.ID, objc.Sel("note"))
	return rv
}


// SetNote sets the value of the note property.
// A note associated with the payment.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insendpaymentintent/note
func (i_ INSendPaymentIntent) SetNote(value appkit.string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setNote:"), value)
}

// The recipient of the payment.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insendpaymentintent/payee
func (i_ INSendPaymentIntent) Payee() INPerson {
	rv := objc.Send[INPerson](i_.ID, objc.Sel("payee"))
	return rv
}


// SetPayee sets the value of the payee property.
// The recipient of the payment.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insendpaymentintent/payee
func (i_ INSendPaymentIntent) SetPayee(value INPerson) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPayee:"), value)
}



