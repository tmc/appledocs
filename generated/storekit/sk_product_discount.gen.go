// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [ProductDiscount] class.
var (
	ProductDiscountClass     _ProductDiscountClass
	ProductDiscountClassOnce sync.Once
)

func getProductDiscountClass() _ProductDiscountClass {
	ProductDiscountClassOnce.Do(func() {
		ProductDiscountClass = _ProductDiscountClass{objc.GetClass("SKProductDiscount")}
	})
	return ProductDiscountClass
}

type _ProductDiscountClass struct {
	class objc.Class
}

// An interface definition for the [ProductDiscount] class.
type IProductDiscount interface {
	objectivec.IObject
}

// The details of an introductory offer or a promotional offer for an auto-renewable subscription.
//
// You set up introductory and promotional offers in App Store Connect. contains the offer information as retrieved from the App Store. For more information about setting up offers, see and .
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount
type ProductDiscount struct {
	objectivec.Object
}

// ProductDiscountFrom constructs a [ProductDiscount] from an unsafe.Pointer.
//
// The details of an introductory offer or a promotional offer for an auto-renewable subscription.
func ProductDiscountFrom(ptr unsafe.Pointer) ProductDiscount {
	return ProductDiscount{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _ProductDiscountClass) Alloc() ProductDiscount {
	rv := objc.Send[ProductDiscount](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _ProductDiscountClass) New() ProductDiscount {
	rv := objc.Send[ProductDiscount](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ ProductDiscount) Init() ProductDiscount {
	rv := objc.Send[ProductDiscount](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ ProductDiscount) Autorelease() ProductDiscount {
	rv := objc.Send[ProductDiscount](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewProductDiscount creates a new ProductDiscount instance.
func NewProductDiscount() ProductDiscount {
	return getProductDiscountClass().New()
}


// A string used to uniquely identify a discount offer for a product.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount/identifier
func (p_ ProductDiscount) Identifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("identifier"))
	return rv
}

// An integer that indicates the number of periods the product discount is available.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount/numberOfPeriods
func (p_ ProductDiscount) NumberOfPeriods() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("numberOfPeriods"))
	return rv
}

// The payment mode for this product discount.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount/paymentMode-swift.property
func (p_ ProductDiscount) PaymentMode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("paymentMode"))
	return rv
}

// The discount price of the product in the local currency.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount/price
func (p_ ProductDiscount) Price() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("price"))
	return rv
}

// The locale used to format the discount price of the product.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount/priceLocale
func (p_ ProductDiscount) PriceLocale() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("priceLocale"))
	return rv
}

// An object that defines the period for the product discount.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount/subscriptionPeriod
func (p_ ProductDiscount) SubscriptionPeriod() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("subscriptionPeriod"))
	return rv
}

// The type of discount offer.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount/type-swift.property
func (p_ ProductDiscount) Type() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("type"))
	return rv
}



