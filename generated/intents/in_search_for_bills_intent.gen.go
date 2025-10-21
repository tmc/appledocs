// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INSearchForBillsIntent] class.
var (
	INSearchForBillsIntentClass     _INSearchForBillsIntentClass
	INSearchForBillsIntentClassOnce sync.Once
)

func getINSearchForBillsIntentClass() _INSearchForBillsIntentClass {
	INSearchForBillsIntentClassOnce.Do(func() {
		INSearchForBillsIntentClass = _INSearchForBillsIntentClass{objc.GetClass("INSearchForBillsIntent")}
	})
	return INSearchForBillsIntentClass
}

type _INSearchForBillsIntentClass struct {
	class objc.Class
}

// An interface definition for the [INSearchForBillsIntent] class.
type IINSearchForBillsIntent interface {
	IINIntent
}

// A request for the list of bills matching the specified criteria.
//
// Siri creates an object when the user asks to see pending or already paid bills. The intent object contains the values to match when searching for bills. Users can search for bills based on the due date, the payee, the type, and whether they’re pending or already paid. When performing the search, use only the provided parameters to filter the search results and ignore any parameters that have a missing or unknown value. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with the results of the search. For successful searches, Siri offers a way for the user to view the results.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForBillsIntent
type INSearchForBillsIntent struct {
	INIntent
}

// INSearchForBillsIntentFrom constructs a [INSearchForBillsIntent] from an unsafe.Pointer.
//
// A request for the list of bills matching the specified criteria.
func INSearchForBillsIntentFrom(ptr unsafe.Pointer) INSearchForBillsIntent {
	return INSearchForBillsIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INSearchForBillsIntentClass) Alloc() INSearchForBillsIntent {
	rv := objc.Send[INSearchForBillsIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INSearchForBillsIntentClass) New() INSearchForBillsIntent {
	rv := objc.Send[INSearchForBillsIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INSearchForBillsIntent) Init() INSearchForBillsIntent {
	rv := objc.Send[INSearchForBillsIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INSearchForBillsIntent) Autorelease() INSearchForBillsIntent {
	rv := objc.Send[INSearchForBillsIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINSearchForBillsIntent creates a new INSearchForBillsIntent instance.
func NewINSearchForBillsIntent() INSearchForBillsIntent {
	return getINSearchForBillsIntentClass().New()
}




// Initializes an intent object that describes a search for bill details with the specified search parameters.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForBillsIntent/init(billPayee:paymentDateRange:billType:status:dueDateRange:)
func NewINSearchForBillsIntentWithBillPayeePaymentDateRangeBillTypeStatusDueDateRange(billPayee unsafe.Pointer, paymentDateRange unsafe.Pointer, billType unsafe.Pointer, status unsafe.Pointer, dueDateRange unsafe.Pointer) INSearchForBillsIntent {
	instance := getINSearchForBillsIntentClass().Alloc()
	rv := objc.Send[INSearchForBillsIntent](instance.ID, objc.Sel("initWithBillPayee:paymentDateRange:billType:status:dueDateRange:"), billPayee, paymentDateRange, billType, status, dueDateRange)
	rv.Autorelease()
	return rv
}


// The recipient of the payment.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforbillsintent/billpayee
func (i_ INSearchForBillsIntent) BillPayee() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("billPayee"))
	return rv
}


// SetBillPayee sets the value of the billPayee property.
// The recipient of the payment.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforbillsintent/billpayee
func (i_ INSearchForBillsIntent) SetBillPayee(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBillPayee:"), value)
}

// The range of payment dates in which to search for bills.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforbillsintent/paymentdaterange
func (i_ INSearchForBillsIntent) PaymentDateRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("paymentDateRange"))
	return rv
}


// SetPaymentDateRange sets the value of the paymentDateRange property.
// The range of payment dates in which to search for bills.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforbillsintent/paymentdaterange
func (i_ INSearchForBillsIntent) SetPaymentDateRange(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setPaymentDateRange:"), value)
}

// The status of the bill.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforbillsintent/status
func (i_ INSearchForBillsIntent) Status() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("status"))
	return rv
}


// SetStatus sets the value of the status property.
// The status of the bill.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/insearchforbillsintent/status
func (i_ INSearchForBillsIntent) SetStatus(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setStatus:"), value)
}

// The type of the bill.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForBillsIntent/billType
func (i_ INSearchForBillsIntent) BillType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("billType"))
	return rv
}

// The range of due dates in which to search for bills.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INSearchForBillsIntent/dueDateRange
func (i_ INSearchForBillsIntent) DueDateRange() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("dueDateRange"))
	return rv
}


