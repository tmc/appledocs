// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [MutablePayment] class.
type IMutablePayment interface {
	IPayment
}

// A mutable request to the App Store to process payment for additional functionality that your app offers.
//
// A mutable payment object identifies a product and the quantity of that item the user would like to purchase. When a mutable payment is added to the payment queue, the payment queue copies the contents into an immutable request before queueing the request. Your app can safely change the contents of the mutable payment object.
//
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

// Alloc allocates a new instance without initialization.
func (mc _MutablePaymentClass) Alloc() MutablePayment {
	rv := objc.Send[MutablePayment](objc.ID(mc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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


// A string that associates the transaction with a user account on your service.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKMutablePayment/applicationUsername
func (m_ MutablePayment) ApplicationUsername() string {
	rv := objc.Send[string](m_.ID, objc.Sel("applicationUsername"))
	return rv
}


// SetApplicationUsername sets the value of the applicationUsername property.
// A string that associates the transaction with a user account on your service.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKMutablePayment/applicationUsername
func (m_ MutablePayment) SetApplicationUsername(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setApplicationUsername:"), objc.String(value))
}
// The details of the discount offer to apply to the payment.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKMutablePayment/paymentDiscount
func (m_ MutablePayment) PaymentDiscount() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("paymentDiscount"))
	return rv
}


// SetPaymentDiscount sets the value of the paymentDiscount property.
// The details of the discount offer to apply to the payment.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKMutablePayment/paymentDiscount
func (m_ MutablePayment) SetPaymentDiscount(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setPaymentDiscount:"), value)
}
// A string that identifies a product that can be purchased from within your app.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKMutablePayment/productIdentifier
func (m_ MutablePayment) ProductIdentifier() string {
	rv := objc.Send[string](m_.ID, objc.Sel("productIdentifier"))
	return rv
}


// SetProductIdentifier sets the value of the productIdentifier property.
// A string that identifies a product that can be purchased from within your app.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKMutablePayment/productIdentifier
func (m_ MutablePayment) SetProductIdentifier(value string) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setProductIdentifier:"), objc.String(value))
}
// The number of items the user wants to purchase.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKMutablePayment/quantity
func (m_ MutablePayment) Quantity() int {
	rv := objc.Send[int](m_.ID, objc.Sel("quantity"))
	return rv
}


// SetQuantity sets the value of the quantity property.
// The number of items the user wants to purchase.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKMutablePayment/quantity
func (m_ MutablePayment) SetQuantity(value int) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setQuantity:"), value)
}
// Reserved for future use.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKMutablePayment/requestData
func (m_ MutablePayment) RequestData() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](m_.ID, objc.Sel("requestData"))
	return rv
}


// SetRequestData sets the value of the requestData property.
// Reserved for future use.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKMutablePayment/requestData
func (m_ MutablePayment) SetRequestData(value unsafe.Pointer) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setRequestData:"), value)
}
// A Boolean value that produces an “ask to buy” flow for this payment in the sandbox.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKMutablePayment/simulatesAskToBuyInSandbox
func (m_ MutablePayment) SimulatesAskToBuyInSandbox() bool {
	rv := objc.Send[bool](m_.ID, objc.Sel("simulatesAskToBuyInSandbox"))
	return rv
}


// SetSimulatesAskToBuyInSandbox sets the value of the simulatesAskToBuyInSandbox property.
// A Boolean value that produces an “ask to buy” flow for this payment in the sandbox.

//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKMutablePayment/simulatesAskToBuyInSandbox
func (m_ MutablePayment) SetSimulatesAskToBuyInSandbox(value bool) {
	objc.Send[objc.ID](m_.ID, objc.Sel("setSimulatesAskToBuyInSandbox:"), value)
}


