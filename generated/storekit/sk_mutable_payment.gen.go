// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class SKMutablePayment */

/* debug [class_header]: Header for SKMutablePayment */
// The class instance for the [MutablePayment] class.
var (
	MutablePaymentClass     _MutablePaymentClass
	MutablePaymentClassOnce sync.Once
)

func getMutablePaymentClass() _MutablePaymentClass {
	MutablePaymentClassOnce.Do(func() {
		MutablePaymentClass = _MutablePaymentClass{objc.GetClass("SKMutablePayment")}
	})
	return MutablePaymentClass
}

type _MutablePaymentClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for MutablePayment */
// An interface definition for the [MutablePayment] class.
type IMutablePayment interface {
	IPayment

	/* debug [class_interface_properties]: Properties for MutablePayment */
	// properties:
	ApplicationUsername() objc.IObject /* cross-framework: NSString */
	SetApplicationUsername(value objc.IObject /* cross-framework: NSString */)
	PaymentDiscount() ISKPaymentDiscount
	SetPaymentDiscount(value ISKPaymentDiscount)
	ProductIdentifier() objc.IObject /* cross-framework: NSString */
	SetProductIdentifier(value objc.IObject /* cross-framework: NSString */)
	Quantity() int
	SetQuantity(value int)
	RequestData() objc.IObject /* cross-framework: NSData */
	SetRequestData(value objc.IObject /* cross-framework: NSData */)
	SimulatesAskToBuyInSandbox() bool
	SetSimulatesAskToBuyInSandbox(value bool)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for MutablePayment */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for MutablePayment */
// Alloc allocates a new instance without initialization.
func (mc _MutablePaymentClass) Alloc() MutablePayment {
	rv := objc.Send[MutablePayment](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (mc _MutablePaymentClass) New() MutablePayment {
	rv := objc.Send[MutablePayment](objc.ID(mc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (m_ MutablePayment) Init() MutablePayment {
	rv := objc.Send[MutablePayment](m_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (m_ MutablePayment) Autorelease() MutablePayment {
	rv := objc.Send[MutablePayment](m_.ID, objc.Sel("autorelease"))
	return rv
}

// NewMutablePayment creates a new MutablePayment instance.
func NewMutablePayment() MutablePayment {
	return getMutablePaymentClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for MutablePayment */
// A mutable request to the App Store to process payment for additional functionality that your app offers.
//
// A mutable payment object identifies a product and the quantity of that item the user would like to purchase. When a mutable payment is added to the payment queue, the payment queue copies the contents into an immutable request before queueing the request. Your app can safely change the contents of the mutable payment object.

// A mutable request to the App Store to process payment for additional functionality that your app offers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKMutablePayment
type MutablePayment struct {
	Payment
}

// MutablePaymentFrom constructs a [MutablePayment] from an unsafe.Pointer.
//
// A mutable request to the App Store to process payment for additional functionality that your app offers.
func MutablePaymentFrom(ptr unsafe.Pointer) MutablePayment {
	return MutablePayment{
		Payment: PaymentFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for MutablePayment */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for MutablePayment */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for MutablePayment */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for MutablePayment */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for MutablePayment */

// A string that associates the transaction with a user account on your service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKMutablePayment/applicationUsername
func (m_ MutablePayment) ApplicationUsername() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("applicationUsername"))
	return rv
} /* debug [instance_properties/getter]: applicationUsername */

// A string that associates the transaction with a user account on your service.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKMutablePayment/applicationUsername
func (m_ MutablePayment) SetApplicationUsername(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setApplicationUsername:"), value)
} /* debug [instance_properties/setter]: applicationUsername */

// The details of the discount offer to apply to the payment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKMutablePayment/paymentDiscount
func (m_ MutablePayment) PaymentDiscount() ISKPaymentDiscount {
	rv := objc.Send[PaymentDiscount](m_.ID, objc.Sel("paymentDiscount"))
	return rv
} /* debug [instance_properties/getter]: paymentDiscount */

// The details of the discount offer to apply to the payment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKMutablePayment/paymentDiscount
func (m_ MutablePayment) SetPaymentDiscount(value ISKPaymentDiscount) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPaymentDiscount:"), value)
} /* debug [instance_properties/setter]: paymentDiscount */

// A string that identifies a product that can be purchased from within your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKMutablePayment/productIdentifier
func (m_ MutablePayment) ProductIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](m_.ID, objc.Sel("productIdentifier"))
	return rv
} /* debug [instance_properties/getter]: productIdentifier */

// A string that identifies a product that can be purchased from within your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKMutablePayment/productIdentifier
func (m_ MutablePayment) SetProductIdentifier(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductIdentifier:"), value)
} /* debug [instance_properties/setter]: productIdentifier */

// The number of items the user wants to purchase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKMutablePayment/quantity
func (m_ MutablePayment) Quantity() int {
	rv := objc.Send[int](m_.ID, objc.Sel("quantity"))
	return rv
} /* debug [instance_properties/getter]: quantity */

// The number of items the user wants to purchase.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKMutablePayment/quantity
func (m_ MutablePayment) SetQuantity(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setQuantity:"), value)
} /* debug [instance_properties/setter]: quantity */

// Reserved for future use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKMutablePayment/requestData
func (m_ MutablePayment) RequestData() objc.IObject /* cross-framework: NSData */ {
	rv := objc.Send[foundation.NSData](m_.ID, objc.Sel("requestData"))
	return rv
} /* debug [instance_properties/getter]: requestData */

// Reserved for future use.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKMutablePayment/requestData
func (m_ MutablePayment) SetRequestData(value objc.IObject /* cross-framework: NSData */) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequestData:"), value)
} /* debug [instance_properties/setter]: requestData */

// A Boolean value that produces an “ask to buy” flow for this payment in the sandbox.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKMutablePayment/simulatesAskToBuyInSandbox
func (m_ MutablePayment) SimulatesAskToBuyInSandbox() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("simulatesAskToBuyInSandbox"))
	return rv
} /* debug [instance_properties/getter]: simulatesAskToBuyInSandbox */

// A Boolean value that produces an “ask to buy” flow for this payment in the sandbox.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKMutablePayment/simulatesAskToBuyInSandbox
func (m_ MutablePayment) SetSimulatesAskToBuyInSandbox(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSimulatesAskToBuyInSandbox:"), value)
} /* debug [instance_properties/setter]: simulatesAskToBuyInSandbox */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class SKMutablePayment */
