// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [Payment] class.
type IPayment interface {
	objectivec.IObject
}

// A request to the App Store to process payment for additional functionality that your app offers.
//
// A payment object identifies a product and the quantity of those items the user would like to purchase.
//
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

// Alloc allocates a new instance without initialization.
func (pc _PaymentClass) Alloc() Payment {
	rv := objc.Send[Payment](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Returns a new payment for the specified product.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPayment/init(product:)
func NewPaymentWithProduct(product unsafe.Pointer) Payment {
	rv := objc.Send[Payment](objc.ID(getPaymentClass().class), objc.Sel("paymentWithProduct:"), product)
	return rv
}


// Returns a new payment for the specified product.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPayment/init(product:)
func (pc _PaymentClass) PaymentWithProduct(product unsafe.Pointer) unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](objc.ID(pc.class), objc.Sel("paymentWithProduct:"), product)
	return rv
}

// Returns a new payment with the specified product identifier.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPayment/paymentWithProductIdentifier:
func (pc _PaymentClass) PaymentWithProductIdentifier(identifier string) objc.ID {
	rv := objc.Send[objc.ID](objc.ID(pc.class), objc.Sel("paymentWithProductIdentifier:"), objc.String(identifier))
	return rv
}

// A string that associates the transaction with a user account on your service.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPayment/applicationUsername
func (p_ Payment) ApplicationUsername() string {
	rv := objc.Send[string](p_.ID, objc.Sel("applicationUsername"))
	return rv
}

// The details of the discount offer to apply to the payment.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPayment/paymentDiscount
func (p_ Payment) PaymentDiscount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("paymentDiscount"))
	return rv
}

// A string used to identify a product that can be purchased from within your app.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPayment/productIdentifier
func (p_ Payment) ProductIdentifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("productIdentifier"))
	return rv
}

// The number of items the user wants to purchase.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPayment/quantity
func (p_ Payment) Quantity() int {
	rv := objc.Send[int](p_.ID, objc.Sel("quantity"))
	return rv
}

// Reserved for future use.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPayment/requestData
func (p_ Payment) RequestData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("requestData"))
	return rv
}

// A Boolean value that produces an “ask to buy” flow for this payment in the sandbox.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPayment/simulatesAskToBuyInSandbox
func (p_ Payment) SimulatesAskToBuyInSandbox() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("simulatesAskToBuyInSandbox"))
	return rv
}


