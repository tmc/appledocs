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


// A string that identifies the version of the content.
//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/contentversion
func (p_ Product) ContentVersion() string {
	rv := objc.Send[string](p_.ID, objc.Sel("contentVersion"))
	return rv
}


// SetContentVersion sets the value of the contentVersion property.
// A string that identifies the version of the content.

//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/contentversion
func (p_ Product) SetContentVersion(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentVersion:"), objc.String(value))
}

// The name of the product.
//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/localizedtitle
func (p_ Product) LocalizedTitle() string {
	rv := objc.Send[string](p_.ID, objc.Sel("localizedTitle"))
	return rv
}


// SetLocalizedTitle sets the value of the localizedTitle property.
// The name of the product.

//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/localizedtitle
func (p_ Product) SetLocalizedTitle(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLocalizedTitle:"), objc.String(value))
}

// A string that identifies which version of the content is available for download.
//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/downloadcontentversion
func (p_ Product) DownloadContentVersion() string {
	rv := objc.Send[string](p_.ID, objc.Sel("downloadContentVersion"))
	return rv
}


// SetDownloadContentVersion sets the value of the downloadContentVersion property.
// A string that identifies which version of the content is available for download.

//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/downloadcontentversion
func (p_ Product) SetDownloadContentVersion(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDownloadContentVersion:"), objc.String(value))
}

// A Boolean value that indicates whether the App Store has downloadable content for this product.
//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/isdownloadable
func (p_ Product) IsDownloadable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isDownloadable"))
	return rv
}


// SetIsDownloadable sets the value of the isDownloadable property.
// A Boolean value that indicates whether the App Store has downloadable content for this product.

//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/isdownloadable
func (p_ Product) SetIsDownloadable(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setIsDownloadable:"), value)
}

// A Boolean value that indicates whether the App Store has downloadable content for this product.
//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/downloadable
func (p_ Product) Downloadable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("downloadable"))
	return rv
}


// SetDownloadable sets the value of the downloadable property.
// A Boolean value that indicates whether the App Store has downloadable content for this product.

//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/downloadable
func (p_ Product) SetDownloadable(value bool) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDownloadable:"), value)
}

// The total size of the content, in bytes.
//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/contentlengths
func (p_ Product) ContentLengths() foundation.Number {
	rv := objc.Send[foundation.Number](p_.ID, objc.Sel("contentLengths"))
	return rv
}


// SetContentLengths sets the value of the contentLengths property.
// The total size of the content, in bytes.

//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/contentlengths
func (p_ Product) SetContentLengths(value foundation.Number) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setContentLengths:"), value)
}

// The lengths of the downloadable files available for this product.
//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/downloadcontentlengths
func (p_ Product) DownloadContentLengths() foundation.Number {
	rv := objc.Send[foundation.Number](p_.ID, objc.Sel("downloadContentLengths"))
	return rv
}


// SetDownloadContentLengths sets the value of the downloadContentLengths property.
// The lengths of the downloadable files available for this product.

//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/downloadcontentlengths
func (p_ Product) SetDownloadContentLengths(value foundation.Number) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDownloadContentLengths:"), value)
}

// A description of the product.
//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/localizeddescription
func (p_ Product) LocalizedDescription() string {
	rv := objc.Send[string](p_.ID, objc.Sel("localizedDescription"))
	return rv
}


// SetLocalizedDescription sets the value of the localizedDescription property.
// A description of the product.

//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/localizeddescription
func (p_ Product) SetLocalizedDescription(value string) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setLocalizedDescription:"), objc.String(value))
}

// The period details for products that are subscriptions.
//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/subscriptionperiod
func (p_ Product) SubscriptionPeriod() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("subscriptionPeriod"))
	return rv
}


// SetSubscriptionPeriod sets the value of the subscriptionPeriod property.
// The period details for products that are subscriptions.

//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skproduct/subscriptionperiod
func (p_ Product) SetSubscriptionPeriod(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setSubscriptionPeriod:"), value)
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



