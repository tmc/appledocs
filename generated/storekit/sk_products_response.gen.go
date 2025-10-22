// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

// The class instance for the [ProductsResponse] class.
var (
	ProductsResponseClass     _ProductsResponseClass
	ProductsResponseClassOnce sync.Once
)

func getProductsResponseClass() _ProductsResponseClass {
	ProductsResponseClassOnce.Do(func() {
		ProductsResponseClass = _ProductsResponseClass{objc.GetClass("SKProductsResponse")}
	})
	return ProductsResponseClass
}

type _ProductsResponseClass struct {
	class objc.Class
}

// An interface definition for the [ProductsResponse] class.
type IProductsResponse interface {
	objectivec.IObject
	InvalidProductIdentifiers() []string
	Products() []Product
}

// An App Store response to a request for information about a list of products.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductsResponse
type ProductsResponse struct {
	objectivec.Object
}

// ProductsResponseFrom constructs a [ProductsResponse] from an unsafe.Pointer.
//
// An App Store response to a request for information about a list of products.
func ProductsResponseFrom(ptr unsafe.Pointer) ProductsResponse {
	return ProductsResponse{objectivec.Object{objc.ID(ptr)}}
}

// Alloc allocates a new instance without initialization.
func (pc _ProductsResponseClass) Alloc() ProductsResponse {
	rv := objc.Send[ProductsResponse](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
func (pc _ProductsResponseClass) New() ProductsResponse {
	rv := objc.Send[ProductsResponse](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ ProductsResponse) Init() ProductsResponse {
	rv := objc.Send[ProductsResponse](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ ProductsResponse) Autorelease() ProductsResponse {
	rv := objc.Send[ProductsResponse](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewProductsResponse creates a new ProductsResponse instance.
func NewProductsResponse() ProductsResponse {
	return getProductsResponseClass().New()
}


// An array of product identifier strings that the App Store doesn’t recognize.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductsResponse/invalidProductIdentifiers
func (p_ ProductsResponse) InvalidProductIdentifiers() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("invalidProductIdentifiers"))
	return rv
}

// A list of products, one product for each valid product identifier provided in the original request.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductsResponse/products
func (p_ ProductsResponse) Products() []Product {
	rv := objc.Send[[]Product](p_.ID, objc.Sel("products"))
	return rv
}



