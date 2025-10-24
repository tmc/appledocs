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

/* debug [class.gen.go]: Generating class SKPaymentTransaction */


/* debug [class_header]: Header for SKPaymentTransaction */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PaymentTransaction */
// An interface definition for the [PaymentTransaction] class.
type IPaymentTransaction interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PaymentTransaction */
	// properties:
	Downloads() []Download
	Error() objc.IObject /* cross-framework: Error */
	OriginalTransaction() ISKPaymentTransaction
	Payment() ISKPayment
	TransactionDate() objc.IObject /* cross-framework: NSDate */
	TransactionIdentifier() objc.IObject /* cross-framework: NSString */
	TransactionState() PaymentTransactionState
	Original() ISKPaymentTransaction
	SetOriginal(value ISKPaymentTransaction)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PaymentTransaction */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PaymentTransaction */
// Alloc allocates a new instance without initialization.
func (pc _PaymentTransactionClass) Alloc() PaymentTransaction {
	rv := objc.Send[PaymentTransaction](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PaymentTransaction */
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PaymentTransaction *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PaymentTransaction */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PaymentTransaction */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PaymentTransaction */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PaymentTransaction */

// An array of download objects representing the downloadable content associated with the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransaction/downloads
func (p_ PaymentTransaction) Downloads() []Download {
	rv := objc.Send[[]Download](p_.ID, objc.Sel("downloads"))
	return rv
}/* debug [instance_properties/getter]: downloads */


// An object describing the error that occurred while processing the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransaction/error
func (p_ PaymentTransaction) Error() objc.IObject /* cross-framework: Error */ {
	rv := objc.Send[coretelephony.Error](p_.ID, objc.Sel("error"))
	return rv
}/* debug [instance_properties/getter]: error */


// The transaction that was restored by the App Store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransaction/original
func (p_ PaymentTransaction) OriginalTransaction() ISKPaymentTransaction {
	rv := objc.Send[PaymentTransaction](p_.ID, objc.Sel("originalTransaction"))
	return rv
}/* debug [instance_properties/getter]: originalTransaction */


// The payment for the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransaction/payment
func (p_ PaymentTransaction) Payment() ISKPayment {
	rv := objc.Send[Payment](p_.ID, objc.Sel("payment"))
	return rv
}/* debug [instance_properties/getter]: payment */


// The date the transaction was added to the App Store’s payment queue.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransaction/transactionDate
func (p_ PaymentTransaction) TransactionDate() objc.IObject /* cross-framework: NSDate */ {
	rv := objc.Send[foundation.NSDate](p_.ID, objc.Sel("transactionDate"))
	return rv
}/* debug [instance_properties/getter]: transactionDate */


// A string that uniquely identifies a successful payment transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransaction/transactionIdentifier
func (p_ PaymentTransaction) TransactionIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("transactionIdentifier"))
	return rv
}/* debug [instance_properties/getter]: transactionIdentifier */


// The current state of the transaction.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentTransaction/transactionState
func (p_ PaymentTransaction) TransactionState() PaymentTransactionState {
	rv := objc.Send[PaymentTransactionState](p_.ID, objc.Sel("transactionState"))
	return rv
}/* debug [instance_properties/getter]: transactionState */


// The transaction that was restored by the App Store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skpaymenttransaction/original
func (p_ PaymentTransaction) Original() ISKPaymentTransaction {
	rv := objc.Send[PaymentTransaction](p_.ID, objc.Sel("original"))
	return rv
}/* debug [instance_properties/getter]: original */


// The transaction that was restored by the App Store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skpaymenttransaction/original
func (p_ PaymentTransaction) SetOriginal(value ISKPaymentTransaction) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setOriginal:"), value)
}/* debug [instance_properties/setter]: original */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SKPaymentTransaction */


