// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

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

// An interface definition for the [PaymentDiscount] class.
type IPaymentDiscount interface {
	objectivec.IObject
}

// The signed discount to apply to a payment.
//
// The contains the details of a promotional offer discount that you want to apply to a . Include the signature that you generated in this object. For guidance, see . The App Store uses this signature and the parameters to validate the promotional offer. Keep in mind that the signature must correspond to the parameters in the payment for a transaction to be successful.
//
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

// Alloc allocates a new instance without initialization.
func (pc _PaymentDiscountClass) Alloc() PaymentDiscount {
	rv := objc.Send[PaymentDiscount](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




// Initializes the payment discount with a signature and the parameters used by the signature.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentDiscount/init(identifier:keyIdentifier:nonce:signature:timestamp:)
func NewPaymentDiscountWithIdentifierKeyIdentifierNonceSignatureTimestamp(identifier string, keyIdentifier string, nonce unsafe.Pointer, signature string, timestamp unsafe.Pointer) PaymentDiscount {
	instance := getPaymentDiscountClass().Alloc()
	rv := objc.Send[PaymentDiscount](instance.ID, objc.Sel("initWithIdentifier:keyIdentifier:nonce:signature:timestamp:"), objc.String(identifier), objc.String(keyIdentifier), nonce, objc.String(signature), timestamp)
	rv.Autorelease()
	return rv
}


// A string used to uniquely identify a discount offer for a product.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentDiscount/identifier
func (p_ PaymentDiscount) Identifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("identifier"))
	return rv
}

// A string that identifies the key used to generate the signature.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentDiscount/keyIdentifier
func (p_ PaymentDiscount) KeyIdentifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("keyIdentifier"))
	return rv
}

// A universally unique ID (UUID) value that you define.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentDiscount/nonce
func (p_ PaymentDiscount) Nonce() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("nonce"))
	return rv
}

// A string representing the properties of a specific promotional offer, cryptographically signed.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentDiscount/signature
func (p_ PaymentDiscount) Signature() string {
	rv := objc.Send[string](p_.ID, objc.Sel("signature"))
	return rv
}

// The date and time of the signature’s creation in milliseconds, formatted in Unix epoch time.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKPaymentDiscount/timestamp
func (p_ PaymentDiscount) Timestamp() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("timestamp"))
	return rv
}


