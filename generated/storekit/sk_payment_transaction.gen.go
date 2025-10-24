// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/coretelephony"
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
	// properties:
	TransactionDate() objc.IObject /* cross-framework: NSDate */
	TransactionIdentifier() objc.IObject /* cross-framework: NSString */
	Downloads() objc.IObject /* cross-framework: Download */
	SetDownloads(value objc.IObject /* cross-framework: Download */)
	Error() objc.IObject /* cross-framework: Error */
	SetError(value objc.IObject /* cross-framework: Error */)
	Original() ISKPaymentTransaction
	SetOriginal(value ISKPaymentTransaction)
	Payment() ISKPayment
	SetPayment(value ISKPayment)
	TransactionReceipt() objc.IObject /* cross-framework: Data */
	SetTransactionReceipt(value objc.IObject /* cross-framework: Data */)
	TransactionState() PaymentTransactionState /* not a class type */
	SetTransactionState(value PaymentTransactionState /* not a class type */)
	// methods:
}

// An object in the payment queue.
//
// A payment transaction is created whenever a payment is added to the payment queue. The system delivers transactions to your app when the App Store finishes processing the payment. Completed transactions provide a receipt and transaction identifier that your app can use to save a permanent record of the processed payment.


// An object in the payment queue.
//
// [Full Topic]
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



// The date the transaction was added to the App Store’s payment queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransaction/transactionDate
func (p_ PaymentTransaction) TransactionDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](p_.ID, objc.Sel("transactionDate"))
	return rv
}


// A string that uniquely identifies a successful payment transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransaction/transactionIdentifier
func (p_ PaymentTransaction) TransactionIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("transactionIdentifier"))
	return rv
}


// An array of download objects representing the downloadable content associated with the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skpaymenttransaction/downloads
func (p_ PaymentTransaction) Downloads() objc.IObject /* cross-framework: Download */ {
	rv := objc.Send[Download](p_.ID, objc.Sel("downloads"))
	return rv
}


// An array of download objects representing the downloadable content associated with the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skpaymenttransaction/downloads
func (p_ PaymentTransaction) SetDownloads(value objc.IObject /* cross-framework: Download */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDownloads:"), value)
}


// An object describing the error that occurred while processing the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skpaymenttransaction/error
func (p_ PaymentTransaction) Error() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](p_.ID, objc.Sel("error"))
	return rv
}


// An object describing the error that occurred while processing the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skpaymenttransaction/error
func (p_ PaymentTransaction) SetError(value objc.IObject /* cross-framework: Error */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setError:"), value)
}


// The transaction that was restored by the App Store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skpaymenttransaction/original
func (p_ PaymentTransaction) Original() ISKPaymentTransaction {
	rv := objc.Send[PaymentTransaction](p_.ID, objc.Sel("original"))
	return rv
}


// The transaction that was restored by the App Store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skpaymenttransaction/original
func (p_ PaymentTransaction) SetOriginal(value ISKPaymentTransaction) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOriginal:"), value)
}


// The payment for the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skpaymenttransaction/payment
func (p_ PaymentTransaction) Payment() ISKPayment {
	rv := objc.Send[Payment](p_.ID, objc.Sel("payment"))
	return rv
}


// The payment for the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skpaymenttransaction/payment
func (p_ PaymentTransaction) SetPayment(value ISKPayment) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPayment:"), value)
}


// A signed receipt that records all information about a successful payment transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skpaymenttransaction/transactionreceipt
func (p_ PaymentTransaction) TransactionReceipt() objc.IObject /* cross-framework: Data */ {
	rv := objc.Send[foundation.Data](p_.ID, objc.Sel("transactionReceipt"))
	return rv
}


// A signed receipt that records all information about a successful payment transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skpaymenttransaction/transactionreceipt
func (p_ PaymentTransaction) SetTransactionReceipt(value objc.IObject /* cross-framework: Data */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTransactionReceipt:"), value)
}


// The current state of the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skpaymenttransaction/transactionstate
func (p_ PaymentTransaction) TransactionState() PaymentTransactionState /* not a class type */ {
	rv := objc.Send[PaymentTransactionState](p_.ID, objc.Sel("transactionState"))
	return rv
}


// The current state of the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skpaymenttransaction/transactionstate
func (p_ PaymentTransaction) SetTransactionState(value PaymentTransactionState /* not a class type */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setTransactionState:"), value)
}



