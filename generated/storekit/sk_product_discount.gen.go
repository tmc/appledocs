// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SKProductDiscount */


/* debug [class_header]: Header for SKProductDiscount */
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
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ProductDiscount */
// An interface definition for the [ProductDiscount] class.
type IProductDiscount interface {
	objectivec.IObject
	
/* debug [class_interface_properties]: Properties for ProductDiscount */
	// properties:
	Identifier() objc.IObject /* cross-framework: NSString */
	NumberOfPeriods() uint
	PaymentMode() ProductDiscountPaymentMode
	Price() foundation.DecimalNumber
	PriceLocale() foundation.Locale
	SubscriptionPeriod() ISKProductSubscriptionPeriod
	Type() ProductDiscountType
	Discounts() ISKProductDiscount
	SetDiscounts(value ISKProductDiscount)
	IntroductoryPrice() ISKProductDiscount
	SetIntroductoryPrice(value ISKProductDiscount)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ProductDiscount */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ProductDiscount */
// Alloc allocates a new instance without initialization.
func (pc _ProductDiscountClass) Alloc() ProductDiscount {
	rv := objc.Send[ProductDiscount](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ProductDiscount */
// The details of an introductory offer or a promotional offer for an auto-renewable subscription.
//
// You set up introductory and promotional offers in App Store Connect. contains the offer information as retrieved from the App Store. For more information about setting up offers, see and .


// The details of an introductory offer or a promotional offer for an auto-renewable subscription.
//
// [Full Topic]
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
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ProductDiscount *//* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ProductDiscount */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ProductDiscount */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ProductDiscount */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ProductDiscount */

// A string used to uniquely identify a discount offer for a product.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount/identifier
func (p_ ProductDiscount) Identifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("identifier"))
	return rv
}/* debug [instance_properties/getter]: identifier */


// An integer that indicates the number of periods the product discount is available.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount/numberOfPeriods
func (p_ ProductDiscount) NumberOfPeriods() uint {
	rv := objc.Send[uint](p_.ID, objc.Sel("numberOfPeriods"))
	return rv
}/* debug [instance_properties/getter]: numberOfPeriods */


// The payment mode for this product discount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount/paymentMode-swift.property
func (p_ ProductDiscount) PaymentMode() ProductDiscountPaymentMode {
	rv := objc.Send[ProductDiscountPaymentMode](p_.ID, objc.Sel("paymentMode"))
	return rv
}/* debug [instance_properties/getter]: paymentMode */


// The discount price of the product in the local currency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount/price
func (p_ ProductDiscount) Price() foundation.DecimalNumber {
	rv := objc.Send[foundation.DecimalNumber](p_.ID, objc.Sel("price"))
	return rv
}/* debug [instance_properties/getter]: price */


// The locale used to format the discount price of the product.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount/priceLocale
func (p_ ProductDiscount) PriceLocale() foundation.Locale {
	rv := objc.Send[foundation.Locale](p_.ID, objc.Sel("priceLocale"))
	return rv
}/* debug [instance_properties/getter]: priceLocale */


// An object that defines the period for the product discount.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount/subscriptionPeriod
func (p_ ProductDiscount) SubscriptionPeriod() ISKProductSubscriptionPeriod {
	rv := objc.Send[ProductSubscriptionPeriod](p_.ID, objc.Sel("subscriptionPeriod"))
	return rv
}/* debug [instance_properties/getter]: subscriptionPeriod */


// The type of discount offer.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductDiscount/type-swift.property
func (p_ ProductDiscount) Type() ProductDiscountType {
	rv := objc.Send[ProductDiscountType](p_.ID, objc.Sel("type"))
	return rv
}/* debug [instance_properties/getter]: type */


// An array of subscription offers available for the auto-renewable subscription.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/discounts
func (p_ ProductDiscount) Discounts() ISKProductDiscount {
	rv := objc.Send[ProductDiscount](p_.ID, objc.Sel("discounts"))
	return rv
}/* debug [instance_properties/getter]: discounts */


// An array of subscription offers available for the auto-renewable subscription.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/discounts
func (p_ ProductDiscount) SetDiscounts(value ISKProductDiscount) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDiscounts:"), value)
}/* debug [instance_properties/setter]: discounts */


// The object containing introductory price information for the product.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/introductoryprice
func (p_ ProductDiscount) IntroductoryPrice() ISKProductDiscount {
	rv := objc.Send[ProductDiscount](p_.ID, objc.Sel("introductoryPrice"))
	return rv
}/* debug [instance_properties/getter]: introductoryPrice */


// The object containing introductory price information for the product.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/introductoryprice
func (p_ ProductDiscount) SetIntroductoryPrice(value ISKProductDiscount) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIntroductoryPrice:"), value)
}/* debug [instance_properties/setter]: introductoryPrice */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SKProductDiscount */



