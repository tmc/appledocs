// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SKPaymentQueue */


/* debug [class_header]: Header for SKPaymentQueue */
// The class instance for the [PaymentQueue] class.
var (
	PaymentQueueClass     _PaymentQueueClass
	PaymentQueueClassOnce sync.Once
)

func getPaymentQueueClass() _PaymentQueueClass {
	PaymentQueueClassOnce.Do(func() {
		PaymentQueueClass = _PaymentQueueClass{objc.GetClass("SKPaymentQueue")}
	})
	return PaymentQueueClass
}

type _PaymentQueueClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for PaymentQueue */
// An interface definition for the [PaymentQueue] class.
type IPaymentQueue interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for PaymentQueue */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
	Storefront() objc.IObject /* cross-framework: Storefront */
	TransactionObservers() []objc.ID
	Transactions() []PaymentTransaction
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for PaymentQueue */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for PaymentQueue */
// Alloc allocates a new instance without initialization.
func (pc _PaymentQueueClass) Alloc() PaymentQueue {
	rv := objc.Send[PaymentQueue](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PaymentQueueClass) New() PaymentQueue {
	rv := objc.Send[PaymentQueue](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PaymentQueue) Init() PaymentQueue {
	rv := objc.Send[PaymentQueue](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PaymentQueue) Autorelease() PaymentQueue {
	rv := objc.Send[PaymentQueue](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPaymentQueue creates a new PaymentQueue instance.
func NewPaymentQueue() PaymentQueue {
	return getPaymentQueueClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for PaymentQueue */
// A queue of payment transactions for the App Store to process.
//
// The payment queue communicates with the App Store and presents a user interface so that the user can authorize payment. The contents of the queue are persistent between launches of your app. To process a payment, first add at least one observer object ( ) to the queue (see ). Then, add a payment object ( ) for the item the user wants to purchase. Each time you add a payment object, the queue creates a transaction object ( ) to process that payment and enqueues it to be processed. After payment is fulfilled, the queue updates the transaction object and then calls any observer objects to provide them the updated transaction. Your observer should process the transaction and then remove it from the queue. The exact mechanism you use to process a processed transaction depends on the design of your app and the product being purchased. Here are a few common examples: If the product is a feature already built into your app, your app enables the feature to process the transaction. If the product includes downloadable content provided by the App Store, your app retrieves the objects from the transaction and ask the payment queue to download them. You provide the actual content files to be served by the App Store to App Store Connect when you create the product information. If the product represents downloadable content provided by your own server, your app might open a network connection to your server and download the content from there. For more information on designing the payment processing portion of your app, see .


// A queue of payment transactions for the App Store to process.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentQueue
type PaymentQueue struct {
	objectivec.Object
}

// PaymentQueueFrom constructs a [PaymentQueue] from an unsafe.Pointer.
//
// A queue of payment transactions for the App Store to process.
func PaymentQueueFrom(ptr unsafe.Pointer) PaymentQueue {
	return PaymentQueue{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for PaymentQueue *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for PaymentQueue */

// A method that indicates whether the person can make purchases.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentQueue/canMakePayments()
func (pc _PaymentQueueClass) CanMakePayments() bool {
	rv := objc.Send[bool](objc.ID(pc.class), objc.Sel("canMakePayments"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=CanMakePayments) */


// Returns the default payment queue instance.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentQueue/default()
func (pc _PaymentQueueClass) DefaultQueue() objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("defaultQueue"))
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=DefaultQueue) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for PaymentQueue */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for PaymentQueue */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for PaymentQueue */

// A delegate that provides information needed to complete transactions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentQueue/delegate
func (p_ PaymentQueue) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// A delegate that provides information needed to complete transactions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentQueue/delegate
func (p_ PaymentQueue) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */


// The App Store storefront of the device.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentQueue/storefront
func (p_ PaymentQueue) Storefront() objc.IObject /* cross-framework: Storefront */ {
	rv := objc.Send[objc.ID](p_.ID, objc.Sel("storefront"))
	return rv
}/* debug [instance_properties/getter]: storefront */


// An array of all active payment queue observers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentQueue/transactionObservers
func (p_ PaymentQueue) TransactionObservers() []objc.ID {
	rv := objc.Send[[]objc.ID](p_.ID, objc.Sel("transactionObservers"))
	return rv
}/* debug [instance_properties/getter]: transactionObservers */


// Returns an array of pending transactions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentQueue/transactions
func (p_ PaymentQueue) Transactions() []PaymentTransaction {
	rv := objc.Send[[]PaymentTransaction](p_.ID, objc.Sel("transactions"))
	return rv
}/* debug [instance_properties/getter]: transactions */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SKPaymentQueue */


