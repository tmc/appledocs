// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SKPaymentDiscount */

/* debug [class_header]: Header for SKPaymentDiscount */
// The class instance for the [PaymentDiscount] class.
var (
	PaymentDiscountClass     _PaymentDiscountClass
	PaymentDiscountClassOnce sync.Once
)

func getPaymentDiscountClass() _PaymentDiscountClass {
	PaymentDiscountClassOnce.Do(func() {
		PaymentDiscountClass = _PaymentDiscountClass{objc.GetClass("SKPaymentDiscount")}
	})
	return PaymentDiscountClass
}

type _PaymentDiscountClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for PaymentDiscount */
// An interface definition for the [PaymentDiscount] class.
type IPaymentDiscount interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for PaymentDiscount */
	// properties:
	Identifier() objc.IObject    /* cross-framework: NSString */
	KeyIdentifier() objc.IObject /* cross-framework: NSString */
	Nonce() foundation.UUID
	Signature() objc.IObject /* cross-framework: NSString */
	Timestamp() objc.IObject /* cross-framework: NSNumber */
	PaymentDiscount() ISKPaymentDiscount
	SetPaymentDiscount(value ISKPaymentDiscount)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for PaymentDiscount */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for PaymentDiscount */
// Alloc allocates a new instance without initialization.
func (pc _PaymentDiscountClass) Alloc() PaymentDiscount {
	rv := objc.Send[PaymentDiscount](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _PaymentDiscountClass) New() PaymentDiscount {
	rv := objc.Send[PaymentDiscount](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ PaymentDiscount) Init() PaymentDiscount {
	rv := objc.Send[PaymentDiscount](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ PaymentDiscount) Autorelease() PaymentDiscount {
	rv := objc.Send[PaymentDiscount](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewPaymentDiscount creates a new PaymentDiscount instance.
func NewPaymentDiscount() PaymentDiscount {
	return getPaymentDiscountClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for PaymentDiscount */
// The signed discount to apply to a payment.
//
// The contains the details of a promotional offer discount that you want to apply to a . Include the signature that you generated in this object. For guidance, see . The App Store uses this signature and the parameters to validate the promotional offer. Keep in mind that the signature must correspond to the parameters in the payment for a transaction to be successful.

// The signed discount to apply to a payment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentDiscount
type PaymentDiscount struct {
	objectivec.Object
}

// PaymentDiscountFrom constructs a [PaymentDiscount] from an unsafe.Pointer.
//
// The signed discount to apply to a payment.
func PaymentDiscountFrom(ptr unsafe.Pointer) PaymentDiscount {
	return PaymentDiscount{objectivec.Object{objc.ID(ptr)}}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for PaymentDiscount */

// Initializes the payment discount with a signature and the parameters used by the signature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentDiscount/init(identifier:keyIdentifier:nonce:signature:timestamp:)
func NewPaymentDiscountWithIdentifierKeyIdentifierNonceSignatureTimestamp(identifier objc.IObject /* cross-framework: NSString */, keyIdentifier objc.IObject /* cross-framework: NSString */, nonce foundation.UUID, signature objc.IObject /* cross-framework: NSString */, timestamp objc.IObject /* cross-framework: NSNumber */) PaymentDiscount {
	instance := getPaymentDiscountClass().Alloc()
	rv := objc.Send[PaymentDiscount](instance.ID, objc.Sel("initWithIdentifier:keyIdentifier:nonce:signature:timestamp:"), identifier, keyIdentifier, nonce, signature, timestamp)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewPaymentDiscountWithIdentifierKeyIdentifierNonceSignatureTimestamp */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for PaymentDiscount */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for PaymentDiscount */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for PaymentDiscount */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for PaymentDiscount */

// A string used to uniquely identify a discount offer for a product.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentDiscount/identifier
func (p_ PaymentDiscount) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("identifier"))
	return rv
} /* debug [instance_properties/getter]: identifier */

// A string that identifies the key used to generate the signature.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentDiscount/keyIdentifier
func (p_ PaymentDiscount) KeyIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("keyIdentifier"))
	return rv
} /* debug [instance_properties/getter]: keyIdentifier */

// A universally unique ID (UUID) value that you define.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentDiscount/nonce
func (p_ PaymentDiscount) Nonce() foundation.UUID {
	rv := objc.Send[foundation.UUID](p_.ID, objc.Sel("nonce"))
	return rv
} /* debug [instance_properties/getter]: nonce */

// A string representing the properties of a specific promotional offer, cryptographically signed.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentDiscount/signature
func (p_ PaymentDiscount) Signature() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("signature"))
	return rv
} /* debug [instance_properties/getter]: signature */

// The date and time of the signature’s creation in milliseconds, formatted in Unix epoch time.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentDiscount/timestamp
func (p_ PaymentDiscount) Timestamp() objc.IObject /* cross-framework: NSNumber */ {
	rv := objc.Send[foundation.NSNumber](p_.ID, objc.Sel("timestamp"))
	return rv
} /* debug [instance_properties/getter]: timestamp */

// The details of the discount offer to apply to the payment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skpayment/paymentdiscount
func (p_ PaymentDiscount) PaymentDiscount() ISKPaymentDiscount {
	rv := objc.Send[PaymentDiscount](p_.ID, objc.Sel("paymentDiscount"))
	return rv
} /* debug [instance_properties/getter]: paymentDiscount */

// The details of the discount offer to apply to the payment.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skpayment/paymentdiscount
func (p_ PaymentDiscount) SetPaymentDiscount(value ISKPaymentDiscount) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPaymentDiscount:"), value)
} /* debug [instance_properties/setter]: paymentDiscount */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class SKPaymentDiscount */
