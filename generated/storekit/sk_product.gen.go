// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objectivec"
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
	// properties:
	ContentLengths() []objc.IObject /* cross-framework: Number */
	ContentVersion() objc.IObject /* cross-framework: NSString */
	Discounts() []IProductDiscount
	DownloadContentLengths() []objc.IObject /* cross-framework: Number */
	DownloadContentVersion() objc.IObject /* cross-framework: NSString */
	Downloadable() bool
	IntroductoryPrice() ISKProductDiscount
	IsDownloadable() bool
	IsFamilyShareable() bool
	LocalizedDescription() objc.IObject /* cross-framework: NSString */
	LocalizedTitle() objc.IObject /* cross-framework: NSString */
	Price() objc.IObject /* cross-framework: DecimalNumber */
	ProductIdentifier() objc.IObject /* cross-framework: NSString */
	SubscriptionGroupIdentifier() objc.IObject /* cross-framework: NSString */
	SubscriptionPeriod() ISKProductSubscriptionPeriod
	PriceLocale() objc.IObject /* cross-framework: Locale */
	SetPriceLocale(value objc.IObject /* cross-framework: Locale */)
	// methods:
}

// Information about a registered product in App Store Connect.
//
// objects are returned as part of an object.


// Information about a registered product in App Store Connect.
//
// [Full Topic]
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



// The total size of the content, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/contentLengths
func (p_ Product) ContentLengths() []objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[[]foundation.Number](p_.ID, objc.Sel("contentLengths"))
	return rv
}


// A string that identifies the version of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/contentVersion
func (p_ Product) ContentVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("contentVersion"))
	return rv
}


// An array of subscription offers available for the auto-renewable subscription.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/discounts
func (p_ Product) Discounts() []IProductDiscount {
	rv := objc.Send[[]ProductDiscount](p_.ID, objc.Sel("discounts"))
	return rv
}


// The lengths of the downloadable files available for this product.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/downloadContentLengths
func (p_ Product) DownloadContentLengths() []objc.IObject /* cross-framework: Number */ {
	rv := objc.Send[[]foundation.Number](p_.ID, objc.Sel("downloadContentLengths"))
	return rv
}


// A string that identifies which version of the content is available for download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/downloadContentVersion
func (p_ Product) DownloadContentVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("downloadContentVersion"))
	return rv
}


// A Boolean value that indicates whether the App Store has downloadable content for this product.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/downloadable
func (p_ Product) Downloadable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("downloadable"))
	return rv
}


// The object containing introductory price information for the product.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/introductoryPrice
func (p_ Product) IntroductoryPrice() ISKProductDiscount {
	rv := objc.Send[ProductDiscount](p_.ID, objc.Sel("introductoryPrice"))
	return rv
}


// A Boolean value that indicates whether the App Store has downloadable content for this product.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/isDownloadable
func (p_ Product) IsDownloadable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isDownloadable"))
	return rv
}


// A Boolean value that indicates whether the product is available for Family Sharing in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/isFamilyShareable
func (p_ Product) IsFamilyShareable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isFamilyShareable"))
	return rv
}


// A description of the product.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/localizedDescription
func (p_ Product) LocalizedDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("localizedDescription"))
	return rv
}


// The name of the product.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/localizedTitle
func (p_ Product) LocalizedTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("localizedTitle"))
	return rv
}


// The cost of the product in the local currency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/price
func (p_ Product) Price() objc.IObject /* cross-framework: DecimalNumber */ {
	rv := objc.Send[foundation.DecimalNumber](p_.ID, objc.Sel("price"))
	return rv
}


// The string that identifies the product to the Apple App Store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/productIdentifier
func (p_ Product) ProductIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("productIdentifier"))
	return rv
}


// The identifier of the subscription group to which the subscription belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/subscriptionGroupIdentifier
func (p_ Product) SubscriptionGroupIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("subscriptionGroupIdentifier"))
	return rv
}


// The period details for products that are subscriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/subscriptionPeriod
func (p_ Product) SubscriptionPeriod() ISKProductSubscriptionPeriod {
	rv := objc.Send[ProductSubscriptionPeriod](p_.ID, objc.Sel("subscriptionPeriod"))
	return rv
}


// The locale used to format the price of the product.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/pricelocale
func (p_ Product) PriceLocale() objc.IObject /* cross-framework: Locale */ {
	rv := objc.Send[foundation.Locale](p_.ID, objc.Sel("priceLocale"))
	return rv
}


// The locale used to format the price of the product.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/pricelocale
func (p_ Product) SetPriceLocale(value objc.IObject /* cross-framework: Locale */) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setPriceLocale:"), value)
}



