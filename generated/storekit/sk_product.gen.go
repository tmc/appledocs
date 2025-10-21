// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
	"github.com/tmc/appledocs/generated/foundation"
)

// The class instance for the [Product] class.
var (
	ProductClass     _ProductClass
	ProductClassOnce sync.Once
)

func getProductClass() _ProductClass {
	ProductClassOnce.Do(func() {
		ProductClass = _ProductClass{objc.GetClass("SKProduct")}
	})
	return ProductClass
}

type _ProductClass struct {
	class objc.Class
}

// An interface definition for the [Product] class.
type IProduct interface {
	objectivec.IObject
}

// Information about a registered product in App Store Connect.
//
// objects are returned as part of an object.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct
type Product struct {
	objectivec.Object
}

// ProductFrom constructs a [Product] from an unsafe.Pointer.
//
// Information about a registered product in App Store Connect.
func ProductFrom(ptr unsafe.Pointer) Product {
	return Product{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _ProductClass) Alloc() Product {
	rv := objc.Send[Product](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _ProductClass) New() Product {
	rv := objc.Send[Product](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ Product) Init() Product {
	rv := objc.Send[Product](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ Product) Autorelease() Product {
	rv := objc.Send[Product](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewProduct creates a new Product instance.
func NewProduct() Product {
	return getProductClass().New()
}


// An array of subscription offers available for the auto-renewable subscription.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/discounts
func (p_ Product) Discounts() []ProductDiscount {
	rv := objc.Send[[]ProductDiscount](p_.ID, objc.Sel("discounts"))
	return rv
}

// The object containing introductory price information for the product.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/introductoryPrice
func (p_ Product) IntroductoryPrice() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("introductoryPrice"))
	return rv
}

// A Boolean value that indicates whether the product is available for Family Sharing in App Store Connect.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/isFamilyShareable
func (p_ Product) IsFamilyShareable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isFamilyShareable"))
	return rv
}

// The cost of the product in the local currency.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/price
func (p_ Product) Price() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("price"))
	return rv
}

// The locale used to format the price of the product.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/priceLocale
func (p_ Product) PriceLocale() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("priceLocale"))
	return rv
}

// The string that identifies the product to the Apple App Store.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/productIdentifier
func (p_ Product) ProductIdentifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("productIdentifier"))
	return rv
}

// The identifier of the subscription group to which the subscription belongs.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/subscriptionGroupIdentifier
func (p_ Product) SubscriptionGroupIdentifier() string {
	rv := objc.Send[string](p_.ID, objc.Sel("subscriptionGroupIdentifier"))
	return rv
}



