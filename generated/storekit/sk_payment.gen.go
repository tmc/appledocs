// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SKPayment */


/* debug [class_header]: Header for SKPayment */
// The class instance for the [Payment] class.
var (
	PaymentClass     _PaymentClass
	PaymentClassOnce sync.Once
)

func getPaymentClass() _PaymentClass {
	PaymentClassOnce.Do(func() {
		PaymentClass = _PaymentClass{objc.GetClass("SKPayment")}
	})
	return PaymentClass
}

type _PaymentClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for Payment */
// An interface definition for the [Payment] class.
type IPayment interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for Payment */
	// properties:
	ApplicationUsername() objc.IObject /* cross-framework: NSString */
	PaymentDiscount() ISKPaymentDiscount
	ProductIdentifier() objc.IObject /* cross-framework: NSString */
	Quantity() int
	RequestData() objc.IObject /* cross-framework: NSData */
	SimulatesAskToBuyInSandbox() bool
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for Payment */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for Payment */
// Alloc allocates a new instance without initialization.
func (pc _PaymentClass) Alloc() Payment {
	rv := objc.Send[Payment](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PaymentClass) New() Payment {
	rv := objc.Send[Payment](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Payment) Init() Payment {
	rv := objc.Send[Payment](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Payment) Autorelease() Payment {
	rv := objc.Send[Payment](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPayment creates a new Payment instance.
func NewPayment() Payment {
	return getPaymentClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for Payment */
// A request to the App Store to process payment for additional functionality that your app offers.
//
// A payment object identifies a product and the quantity of those items the user would like to purchase.


// A request to the App Store to process payment for additional functionality that your app offers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPayment
type Payment struct {
	objectivec.Object
}

// PaymentFrom constructs a [Payment] from an unsafe.Pointer.
//
// A request to the App Store to process payment for additional functionality that your app offers.
func PaymentFrom(ptr unsafe.Pointer) Payment {
	return Payment{objectivec.Object{objc.ID(ptr)}}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for Payment */

// Returns a new payment for the specified product.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPayment/init(product:)
func NewPaymentWithProduct(product ISKProduct) Payment {
	rv := objc.Send[Payment](objc.ID(getPaymentClass().class), objc.Sel("paymentWithProduct:"), product)
	return rv
}/* debug [class_init_methods/constructor]: NewPaymentWithProduct */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for Payment */

// Returns a new payment for the specified product.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPayment/init(product:)
func (pc _PaymentClass) PaymentWithProduct(product ISKProduct) objectivec.IObject {
	rv := objc.Send[objectivec.IObject](objc.ID(pc.class), objc.Sel("paymentWithProduct:"), product)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PaymentWithProduct) */


// Returns a new payment with the specified product identifier.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPayment/paymentWithProductIdentifier:
func (pc _PaymentClass) PaymentWithProductIdentifier(identifier objc.IObject /* cross-framework: NSString */) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("paymentWithProductIdentifier:"), identifier)
	return rv
}/* debug [class_methods/method]: Class method for%!(EXTRA string=PaymentWithProductIdentifier) */

/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for Payment */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for Payment */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for Payment */

// A string that associates the transaction with a user account on your service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPayment/applicationUsername
func (p_ Payment) ApplicationUsername() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("applicationUsername"))
	return rv
}/* debug [instance_properties/getter]: applicationUsername */


// The details of the discount offer to apply to the payment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPayment/paymentDiscount
func (p_ Payment) PaymentDiscount() ISKPaymentDiscount {
	rv := objc.Send[PaymentDiscount](p_.ID, objc.Sel("paymentDiscount"))
	return rv
}/* debug [instance_properties/getter]: paymentDiscount */


// A string used to identify a product that can be purchased from within your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPayment/productIdentifier
func (p_ Payment) ProductIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("productIdentifier"))
	return rv
}/* debug [instance_properties/getter]: productIdentifier */


// The number of items the user wants to purchase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPayment/quantity
func (p_ Payment) Quantity() int {
	rv := objc.Send[int](p_.ID, objc.Sel("quantity"))
	return rv
}/* debug [instance_properties/getter]: quantity */


// Reserved for future use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPayment/requestData
func (p_ Payment) RequestData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](p_.ID, objc.Sel("requestData"))
	return rv
}/* debug [instance_properties/getter]: requestData */


// A Boolean value that produces an “ask to buy” flow for this payment in the sandbox.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPayment/simulatesAskToBuyInSandbox
func (p_ Payment) SimulatesAskToBuyInSandbox() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("simulatesAskToBuyInSandbox"))
	return rv
}/* debug [instance_properties/getter]: simulatesAskToBuyInSandbox */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SKPayment */


