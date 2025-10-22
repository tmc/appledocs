// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [PaymentTransaction] class.
var (
	PaymentTransactionClass     _PaymentTransactionClass
	PaymentTransactionClassOnce sync.Once
)

func getPaymentTransactionClass() _PaymentTransactionClass {
	PaymentTransactionClassOnce.Do(func() {
		PaymentTransactionClass = _PaymentTransactionClass{objc.GetClass("SKPaymentTransaction")}
	})
	return PaymentTransactionClass
}

type _PaymentTransactionClass struct {
	class objc.Class
}

// An interface definition for the [PaymentTransaction] class.
type IPaymentTransaction interface {
	objectivec.IObject
	Downloads() []Download
	Error() foundation.Error
	OriginalTransaction() SKPaymentTransaction
	Payment() SKPayment
	TransactionDate() foundation.NSDate
	TransactionIdentifier() string
	TransactionReceipt() foundation.NSData
	TransactionState() PaymentTransactionState
	Original() SKPaymentTransaction
	SetOriginal(value ISKPaymentTransaction)
}

// An object in the payment queue.
//
// A payment transaction is created whenever a payment is added to the payment queue. The system delivers transactions to your app when the App Store finishes processing the payment. Completed transactions provide a receipt and transaction identifier that your app can use to save a permanent record of the processed payment.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransaction
type PaymentTransaction struct {
	objectivec.Object
}

// PaymentTransactionFrom constructs a [PaymentTransaction] from an unsafe.Pointer.
//
// An object in the payment queue.
func PaymentTransactionFrom(ptr unsafe.Pointer) PaymentTransaction {
	return PaymentTransaction{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _PaymentTransactionClass) Alloc() PaymentTransaction {
	rv := objc.Send[PaymentTransaction](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _PaymentTransactionClass) New() PaymentTransaction {
	rv := objc.Send[PaymentTransaction](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PaymentTransaction) Init() PaymentTransaction {
	rv := objc.Send[PaymentTransaction](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PaymentTransaction) Autorelease() PaymentTransaction {
	rv := objc.Send[PaymentTransaction](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPaymentTransaction creates a new PaymentTransaction instance.
func NewPaymentTransaction() PaymentTransaction {
	return getPaymentTransactionClass().New()
}


// An array of download objects representing the downloadable content associated with the transaction.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransaction/downloads
func (p_ PaymentTransaction) Downloads() []Download {
	rv := objc.Send[[]Download](p_.ID, objc.Sel("downloads"))
	return rv
}

// An object describing the error that occurred while processing the transaction.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransaction/error
func (p_ PaymentTransaction) Error() foundation.Error {
	rv := objc.Send[foundation.Error](p_.ID, objc.Sel("error"))
	return rv
}

// The transaction that was restored by the App Store.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransaction/original
func (p_ PaymentTransaction) OriginalTransaction() SKPaymentTransaction {
	rv := objc.Send[SKPaymentTransaction](p_.ID, objc.Sel("originalTransaction"))
	return rv
}

// The payment for the transaction.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransaction/payment
func (p_ PaymentTransaction) Payment() SKPayment {
	rv := objc.Send[SKPayment](p_.ID, objc.Sel("payment"))
	return rv
}

// The date the transaction was added to the App Store’s payment queue.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransaction/transactionDate
func (p_ PaymentTransaction) TransactionDate() foundation.NSDate {
	rv := objc.Send[foundation.NSDate](p_.ID, objc.Sel("transactionDate"))
	return rv
}

// A string that uniquely identifies a successful payment transaction.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransaction/transactionIdentifier
func (p_ PaymentTransaction) TransactionIdentifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("transactionIdentifier"))
	return rv
}

// A signed receipt that records all information about a successful payment transaction.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransaction/transactionReceipt
func (p_ PaymentTransaction) TransactionReceipt() foundation.NSData {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("transactionReceipt"))
	return rv
}

// The current state of the transaction.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransaction/transactionState
func (p_ PaymentTransaction) TransactionState() PaymentTransactionState {
	rv := objc.Send[PaymentTransactionState](p_.ID, objc.Sel("transactionState"))
	return rv
}

// The transaction that was restored by the App Store.
//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skpaymenttransaction/original
func (p_ PaymentTransaction) Original() SKPaymentTransaction {
	rv := objc.Send[SKPaymentTransaction](p_.ID, objc.Sel("original"))
	return rv
}


// SetOriginal sets the value of the original property.
// The transaction that was restored by the App Store.

//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skpaymenttransaction/original
func (p_ PaymentTransaction) SetOriginal(value ISKPaymentTransaction) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOriginal:"), value)
}



