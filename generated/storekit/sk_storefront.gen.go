// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [Storefront] class.
var (
	StorefrontClass     _StorefrontClass
	StorefrontClassOnce sync.Once
)

func getStorefrontClass() _StorefrontClass {
	StorefrontClassOnce.Do(func() {
		StorefrontClass = _StorefrontClass{objc.GetClass("SKStorefront")}
	})
	return StorefrontClass
}

type _StorefrontClass struct {
	class objc.Class
}

// An interface definition for the [Storefront] class.
type IStorefront interface {
	objectivec.IObject
}

// An object containing the location and unique identifier of an Apple App Store storefront.
//
// In-app products you create through App Store Connect are available for sale in every region with an App Store. You can use the storefront information to determine the customer’s region, and offer in-app products suitable for that region. StoreKit exposes storefront information as a read-only property in . You must maintain your own list of product identifiers and the storefronts in which you want to make them available.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKStorefront
type Storefront struct {
	objectivec.Object
}

// StorefrontFrom constructs a [Storefront] from an unsafe.Pointer.
//
// An object containing the location and unique identifier of an Apple App Store storefront.
func StorefrontFrom(ptr unsafe.Pointer) Storefront {
	return Storefront{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (sc _StorefrontClass) Alloc() Storefront {
	rv := objc.Send[Storefront](objc.ID(sc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (sc _StorefrontClass) New() Storefront {
	rv := objc.Send[Storefront](objc.ID(sc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (s_ Storefront) Init() Storefront {
	rv := objc.Send[Storefront](s_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (s_ Storefront) Autorelease() Storefront {
	rv := objc.Send[Storefront](s_.ID, objc.Sel("autorelease"))
	return rv
}

// NewStorefront creates a new Storefront instance.
func NewStorefront() Storefront {
	return getStorefrontClass().New()
}


// The three-letter code representing the country or region associated with the App Store storefront.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKStorefront/countryCode
func (s_ Storefront) CountryCode() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](s_.ID, objc.Sel("countryCode"))
	return rv
}



