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
	BillPayee() unsafe.Pointer
	SetBillPayee(value unsafe.Pointer)
	BillType() unsafe.Pointer
	SetBillType(value unsafe.Pointer)
	DueDate() INDateComponentsRange
	SetDueDate(value INDateComponentsRange)
	FromAccount() unsafe.Pointer
	SetFromAccount(value unsafe.Pointer)
	TransactionAmount() unsafe.Pointer
	SetTransactionAmount(value unsafe.Pointer)
	TransactionNote() string
	SetTransactionNote(value string)
	TransactionScheduledDate() INDateComponentsRange
	SetTransactionScheduledDate(value INDateComponentsRange)
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


// The recipient of the payment.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inpaybillintent/billpayee
func (i_ INPayBillIntent) BillPayee() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("billPayee"))
	return rv
}


// SetBillPayee sets the value of the billPayee property.
// The recipient of the payment.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inpaybillintent/billpayee
func (i_ INPayBillIntent) SetBillPayee(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBillPayee:"), value)
}

// The type of the bill.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inpaybillintent/billtype
func (i_ INPayBillIntent) BillType() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("billType"))
	return rv
}


// SetBillType sets the value of the billType property.
// The type of the bill.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inpaybillintent/billtype
func (i_ INPayBillIntent) SetBillType(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setBillType:"), value)
}

// The due date of the payment.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inpaybillintent/duedate
func (i_ INPayBillIntent) DueDate() INDateComponentsRange {
	rv := objc.Send[INDateComponentsRange](i_.ID, objc.Sel("dueDate"))
	return rv
}


// SetDueDate sets the value of the dueDate property.
// The due date of the payment.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inpaybillintent/duedate
func (i_ INPayBillIntent) SetDueDate(value INDateComponentsRange) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setDueDate:"), value)
}

// The user account containing the funds for the payment.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inpaybillintent/fromaccount
func (i_ INPayBillIntent) FromAccount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("fromAccount"))
	return rv
}


// SetFromAccount sets the value of the fromAccount property.
// The user account containing the funds for the payment.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inpaybillintent/fromaccount
func (i_ INPayBillIntent) SetFromAccount(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setFromAccount:"), value)
}

// The amount to transfer from the user to the payee.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inpaybillintent/transactionamount
func (i_ INPayBillIntent) TransactionAmount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("transactionAmount"))
	return rv
}


// SetTransactionAmount sets the value of the transactionAmount property.
// The amount to transfer from the user to the payee.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inpaybillintent/transactionamount
func (i_ INPayBillIntent) SetTransactionAmount(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransactionAmount:"), value)
}

// A note to associate with the payment transaction.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inpaybillintent/transactionnote
func (i_ INPayBillIntent) TransactionNote() string {
	rv := objc.Send[string](i_.ID, objc.Sel("transactionNote"))
	return rv
}


// SetTransactionNote sets the value of the transactionNote property.
// A note to associate with the payment transaction.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inpaybillintent/transactionnote
func (i_ INPayBillIntent) SetTransactionNote(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransactionNote:"), objc.String(value))
}

// The scheduled date for the payment, as requested by the user.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/inpaybillintent/transactionscheduleddate
func (i_ INPayBillIntent) TransactionScheduledDate() INDateComponentsRange {
	rv := objc.Send[INDateComponentsRange](i_.ID, objc.Sel("transactionScheduledDate"))
	return rv
}


// SetTransactionScheduledDate sets the value of the transactionScheduledDate property.
// The scheduled date for the payment, as requested by the user.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/inpaybillintent/transactionscheduleddate
func (i_ INPayBillIntent) SetTransactionScheduledDate(value INDateComponentsRange) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransactionScheduledDate:"), value)
}



