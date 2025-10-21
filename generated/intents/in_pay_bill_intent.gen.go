// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INPayBillIntent] class.
var (
	INPayBillIntentClass     _INPayBillIntentClass
	INPayBillIntentClassOnce sync.Once
)

func getINPayBillIntentClass() _INPayBillIntentClass {
	INPayBillIntentClassOnce.Do(func() {
		INPayBillIntentClass = _INPayBillIntentClass{objc.GetClass("INPayBillIntent")}
	})
	return INPayBillIntentClass
}

type _INPayBillIntentClass struct {
	class objc.Class
}

// An interface definition for the [INPayBillIntent] class.
type IINPayBillIntent interface {
	IINIntent
}

// A request to transfer money to facilitate payment of a bill.
//
// Siri creates an object when the user asks to pay a bill for a designated payee. A pay bill intent object includes the payment amount, the payment date, and the recipient of the payment. Use that information to validate the transaction and to schedule the payment. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with the results of scheduling the bill payment. This intent object represents a financial transaction between the user and an entity (such as a utility company or credit card bill) defined in your app. You’re responsible for configuring and managing the entities that accept the payment of bills.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INPayBillIntent
type INPayBillIntent struct {
	INIntent
}

// INPayBillIntentFrom constructs a [INPayBillIntent] from an unsafe.Pointer.
//
// A request to transfer money to facilitate payment of a bill.
func INPayBillIntentFrom(ptr unsafe.Pointer) INPayBillIntent {
	return INPayBillIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INPayBillIntentClass) Alloc() INPayBillIntent {
	rv := objc.Send[INPayBillIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INPayBillIntentClass) New() INPayBillIntent {
	rv := objc.Send[INPayBillIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INPayBillIntent) Init() INPayBillIntent {
	rv := objc.Send[INPayBillIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INPayBillIntent) Autorelease() INPayBillIntent {
	rv := objc.Send[INPayBillIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINPayBillIntent creates a new INPayBillIntent instance.
func NewINPayBillIntent() INPayBillIntent {
	return getINPayBillIntentClass().New()
}




