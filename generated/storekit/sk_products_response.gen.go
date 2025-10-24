// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SKProductsResponse */

/* debug [class_header]: Header for SKProductsResponse */
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

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for ProductsResponse */
// An interface definition for the [ProductsResponse] class.
type IProductsResponse interface {
	objectivec.IObject

	/* debug [class_interface_properties]: Properties for ProductsResponse */
	// properties:
	InvalidProductIdentifiers() []string
	Products() []Product
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for ProductsResponse */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for ProductsResponse */
// Alloc allocates a new instance without initialization.
func (pc _ProductsResponseClass) Alloc() ProductsResponse {
	rv := objc.Send[ProductsResponse](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
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

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for ProductsResponse */
// An App Store response to a request for information about a list of products.

// An App Store response to a request for information about a list of products.
//
// [Full Topic]
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

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for ProductsResponse */ /* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for ProductsResponse */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for ProductsResponse */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for ProductsResponse */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for ProductsResponse */

// An array of product identifier strings that the App Store doesn’t recognize.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductsResponse/invalidProductIdentifiers
func (p_ ProductsResponse) InvalidProductIdentifiers() []string {
	rv := objc.Send[[]string](p_.ID, objc.Sel("invalidProductIdentifiers"))
	return rv
} /* debug [instance_properties/getter]: invalidProductIdentifiers */

// A list of products, one product for each valid product identifier provided in the original request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductsResponse/products
func (p_ ProductsResponse) Products() []Product {
	rv := objc.Send[[]Product](p_.ID, objc.Sel("products"))
	return rv
} /* debug [instance_properties/getter]: products */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class SKProductsResponse */
