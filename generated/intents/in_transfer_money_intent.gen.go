// Code generated from Apple documentation for Intents. DO NOT EDIT.

package intents

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

// The class instance for the [INTransferMoneyIntent] class.
var (
	INTransferMoneyIntentClass     _INTransferMoneyIntentClass
	INTransferMoneyIntentClassOnce sync.Once
)

func getINTransferMoneyIntentClass() _INTransferMoneyIntentClass {
	INTransferMoneyIntentClassOnce.Do(func() {
		INTransferMoneyIntentClass = _INTransferMoneyIntentClass{objc.GetClass("INTransferMoneyIntent")}
	})
	return INTransferMoneyIntentClass
}

type _INTransferMoneyIntentClass struct {
	class objc.Class
}

// An interface definition for the [INTransferMoneyIntent] class.
type IINTransferMoneyIntent interface {
	IINIntent
	FromAccount() unsafe.Pointer
	ToAccount() unsafe.Pointer
	SetToAccount(value unsafe.Pointer)
	TransactionAmount() unsafe.Pointer
	SetTransactionAmount(value unsafe.Pointer)
	TransactionNote() string
	SetTransactionNote(value string)
	TransactionScheduledDate() INDateComponentsRange
	SetTransactionScheduledDate(value INDateComponentsRange)
}

// A request to transfer money between two accounts.
//
// Siri creates an object when the user asks to transfer money between two accounts. Transfers can occur only between accounts associated with the user. Use the information provided by the intent object to identify the involved accounts and the amount to transfer. To handle this intent, the handler object in your Intents extension must adopt the protocol. Your handler should confirm the request and create an object with the transaction details. For successful transfers, Siri offers a way for the user to view the results.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INTransferMoneyIntent
type INTransferMoneyIntent struct {
	INIntent
}

// INTransferMoneyIntentFrom constructs a [INTransferMoneyIntent] from an unsafe.Pointer.
//
// A request to transfer money between two accounts.
func INTransferMoneyIntentFrom(ptr unsafe.Pointer) INTransferMoneyIntent {
	return INTransferMoneyIntent{
		INIntent: INIntentFrom(ptr),
	}
}

// Alloc allocates a new instance without initialization.
func (ic _INTransferMoneyIntentClass) Alloc() INTransferMoneyIntent {
	rv := objc.Send[INTransferMoneyIntent](objc.ID(ic.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (ic _INTransferMoneyIntentClass) New() INTransferMoneyIntent {
	rv := objc.Send[INTransferMoneyIntent](objc.ID(ic.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (i_ INTransferMoneyIntent) Init() INTransferMoneyIntent {
	rv := objc.Send[INTransferMoneyIntent](i_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (i_ INTransferMoneyIntent) Autorelease() INTransferMoneyIntent {
	rv := objc.Send[INTransferMoneyIntent](i_.ID, objc.Sel("autorelease"))
	return rv
}

// NewINTransferMoneyIntent creates a new INTransferMoneyIntent instance.
func NewINTransferMoneyIntent() INTransferMoneyIntent {
	return getINTransferMoneyIntentClass().New()
}


// The account containing the funds to transfer.
//
// [Full Topic]: https://developer.apple.com/documentation/Intents/INTransferMoneyIntent/fromAccount
func (i_ INTransferMoneyIntent) FromAccount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("fromAccount"))
	return rv
}

// The account receiving the funds.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/intransfermoneyintent/toaccount
func (i_ INTransferMoneyIntent) ToAccount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("toAccount"))
	return rv
}


// SetToAccount sets the value of the toAccount property.
// The account receiving the funds.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/intransfermoneyintent/toaccount
func (i_ INTransferMoneyIntent) SetToAccount(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setToAccount:"), value)
}

// The amount to transfer.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/intransfermoneyintent/transactionamount
func (i_ INTransferMoneyIntent) TransactionAmount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](i_.ID, objc.Sel("transactionAmount"))
	return rv
}


// SetTransactionAmount sets the value of the transactionAmount property.
// The amount to transfer.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/intransfermoneyintent/transactionamount
func (i_ INTransferMoneyIntent) SetTransactionAmount(value unsafe.Pointer) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransactionAmount:"), value)
}

// An optional note associated with the transaction.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/intransfermoneyintent/transactionnote
func (i_ INTransferMoneyIntent) TransactionNote() string {
	rv := objc.Send[string](i_.ID, objc.Sel("transactionNote"))
	return rv
}


// SetTransactionNote sets the value of the transactionNote property.
// An optional note associated with the transaction.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/intransfermoneyintent/transactionnote
func (i_ INTransferMoneyIntent) SetTransactionNote(value string) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransactionNote:"), objc.String(value))
}

// The date on which to transfer the funds.
//
// [Full Topic]: https://developer.apple.com/documentation/intents/intransfermoneyintent/transactionscheduleddate
func (i_ INTransferMoneyIntent) TransactionScheduledDate() INDateComponentsRange {
	rv := objc.Send[INDateComponentsRange](i_.ID, objc.Sel("transactionScheduledDate"))
	return rv
}


// SetTransactionScheduledDate sets the value of the transactionScheduledDate property.
// The date on which to transfer the funds.

//
// [Full Topic]: https://developer.apple.com/documentation/intents/intransfermoneyintent/transactionscheduleddate
func (i_ INTransferMoneyIntent) SetTransactionScheduledDate(value INDateComponentsRange) {
	objc.Send[objc.ID](i_.ID, objc.Sel("setTransactionScheduledDate:"), value)
}



