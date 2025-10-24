// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SKProduct */

/* debug [class_header]: Header for SKProduct */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for Product */
// An interface definition for the [Product] class.
type IProduct interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for Product */
	// properties:
	ContentLengths() []foundation.Number
	ContentVersion() objc.IObject /* cross-framework: NSString */
	Discounts() []ProductDiscount
	Downloadable() bool
	DownloadContentLengths() []foundation.Number
	DownloadContentVersion() objc.IObject /* cross-framework: NSString */
	IntroductoryPrice() ISKProductDiscount
	IsDownloadable() bool
	IsFamilyShareable() bool
	LocalizedDescription() objc.IObject /* cross-framework: NSString */
	LocalizedTitle() objc.IObject       /* cross-framework: NSString */
	Price() foundation.DecimalNumber
	PriceLocale() foundation.Locale
	ProductIdentifier() objc.IObject           /* cross-framework: NSString */
	SubscriptionGroupIdentifier() objc.IObject /* cross-framework: NSString */
	SubscriptionPeriod() ISKProductSubscriptionPeriod
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for Product */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for Product */
// Alloc allocates a new instance without initialization.
func (pc _ProductClass) Alloc() Product {
	rv := objc.Send[Product](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for Product */
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for Product */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for Product */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for Product */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for Product */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for Product */

// The total size of the content, in bytes.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/contentLengths
func (p_ Product) ContentLengths() []foundation.Number {
	rv := objc.Send[[]foundation.Number](p_.ID, objc.Sel("contentLengths"))
	return rv
} /* debug [instance_properties/getter]: contentLengths */

// A string that identifies the version of the content.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/contentVersion
func (p_ Product) ContentVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("contentVersion"))
	return rv
} /* debug [instance_properties/getter]: contentVersion */

// An array of subscription offers available for the auto-renewable subscription.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/discounts
func (p_ Product) Discounts() []ProductDiscount {
	rv := objc.Send[[]ProductDiscount](p_.ID, objc.Sel("discounts"))
	return rv
} /* debug [instance_properties/getter]: discounts */

// A Boolean value that indicates whether the App Store has downloadable content for this product.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/downloadable
func (p_ Product) Downloadable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("downloadable"))
	return rv
} /* debug [instance_properties/getter]: downloadable */

// The lengths of the downloadable files available for this product.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/downloadContentLengths
func (p_ Product) DownloadContentLengths() []foundation.Number {
	rv := objc.Send[[]foundation.Number](p_.ID, objc.Sel("downloadContentLengths"))
	return rv
} /* debug [instance_properties/getter]: downloadContentLengths */

// A string that identifies which version of the content is available for download.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/downloadContentVersion
func (p_ Product) DownloadContentVersion() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("downloadContentVersion"))
	return rv
} /* debug [instance_properties/getter]: downloadContentVersion */

// The object containing introductory price information for the product.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/introductoryPrice
func (p_ Product) IntroductoryPrice() ISKProductDiscount {
	rv := objc.Send[ProductDiscount](p_.ID, objc.Sel("introductoryPrice"))
	return rv
} /* debug [instance_properties/getter]: introductoryPrice */

// A Boolean value that indicates whether the App Store has downloadable content for this product.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/isDownloadable
func (p_ Product) IsDownloadable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isDownloadable"))
	return rv
} /* debug [instance_properties/getter]: isDownloadable */

// A Boolean value that indicates whether the product is available for Family Sharing in App Store Connect.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/isFamilyShareable
func (p_ Product) IsFamilyShareable() bool {
	rv := objc.Send[bool](p_.ID, objc.Sel("isFamilyShareable"))
	return rv
} /* debug [instance_properties/getter]: isFamilyShareable */

// A description of the product.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/localizedDescription
func (p_ Product) LocalizedDescription() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("localizedDescription"))
	return rv
} /* debug [instance_properties/getter]: localizedDescription */

// The name of the product.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/localizedTitle
func (p_ Product) LocalizedTitle() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("localizedTitle"))
	return rv
} /* debug [instance_properties/getter]: localizedTitle */

// The cost of the product in the local currency.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/price
func (p_ Product) Price() foundation.DecimalNumber {
	rv := objc.Send[foundation.DecimalNumber](p_.ID, objc.Sel("price"))
	return rv
} /* debug [instance_properties/getter]: price */

// The locale used to format the price of the product.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/priceLocale
func (p_ Product) PriceLocale() foundation.Locale {
	rv := objc.Send[foundation.Locale](p_.ID, objc.Sel("priceLocale"))
	return rv
} /* debug [instance_properties/getter]: priceLocale */

// The string that identifies the product to the Apple App Store.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/productIdentifier
func (p_ Product) ProductIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("productIdentifier"))
	return rv
} /* debug [instance_properties/getter]: productIdentifier */

// The identifier of the subscription group to which the subscription belongs.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/subscriptionGroupIdentifier
func (p_ Product) SubscriptionGroupIdentifier() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](p_.ID, objc.Sel("subscriptionGroupIdentifier"))
	return rv
} /* debug [instance_properties/getter]: subscriptionGroupIdentifier */

// The period details for products that are subscriptions.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProduct/subscriptionPeriod
func (p_ Product) SubscriptionPeriod() ISKProductSubscriptionPeriod {
	rv := objc.Send[ProductSubscriptionPeriod](p_.ID, objc.Sel("subscriptionPeriod"))
	return rv
} /* debug [instance_properties/getter]: subscriptionPeriod */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class SKProduct */
