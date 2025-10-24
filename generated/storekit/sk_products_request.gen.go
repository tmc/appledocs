// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/objectivec"
)

/* debug [class.gen.go]: Generating class SKProductsRequest */


/* debug [class_header]: Header for SKProductsRequest */
// The class instance for the [ProductsRequest] class.
var (
	ProductsRequestClass     _ProductsRequestClass
	ProductsRequestClassOnce sync.Once
)

func getProductsRequestClass() _ProductsRequestClass {
	ProductsRequestClassOnce.Do(func() {
		ProductsRequestClass = _ProductsRequestClass{objc.GetClass("SKProductsRequest")}
	})
	return ProductsRequestClass
}

type _ProductsRequestClass struct {
	class objc.Class
}
/* debug [class_header]: End header */



/* debug [class_interface]: Interface for ProductsRequest */
// An interface definition for the [ProductsRequest] class.
type IProductsRequest interface {
	IRequest
	
/* debug [class_interface_properties]: Properties for ProductsRequest */
	// properties:
	Delegate() unsafe.Pointer
	SetDelegate(value unsafe.Pointer)
/* debug [class_interface_properties]: End properties */

	
/* debug [class_interface_methods]: Methods for ProductsRequest */
	// methods:
/* debug [class_interface_methods]: End methods */

}
/* debug [class_interface]: End interface */



/* debug [class_constructors]: Constructors for ProductsRequest */
// Alloc allocates a new instance without initialization.
func (pc _ProductsRequestClass) Alloc() ProductsRequest {
	rv := objc.Send[ProductsRequest](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (pc _ProductsRequestClass) New() ProductsRequest {
	rv := objc.Send[ProductsRequest](objc.ID(pc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (p_ ProductsRequest) Init() ProductsRequest {
	rv := objc.Send[ProductsRequest](p_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (p_ ProductsRequest) Autorelease() ProductsRequest {
	rv := objc.Send[ProductsRequest](p_.ID, objc.Sel("autorelease"))
	return rv
}

// NewProductsRequest creates a new ProductsRequest instance.
func NewProductsRequest() ProductsRequest {
	return getProductsRequestClass().New()
}
/* debug [class_constructors]: End constructors */



/* debug [class_struct]: Struct for ProductsRequest */
// An object that can retrieve localized information from the App Store about a specified list of products.
//
// Your app uses an object to present localized prices and other information to the user without having to maintain that list of product information itself. To use an object, you initialize it with a list of product identifier strings, attach a delegate, and then call the request’s method. When the request completes, your delegate receives an object.


// An object that can retrieve localized information from the App Store about a specified list of products.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductsRequest
type ProductsRequest struct {
	Request
}

// ProductsRequestFrom constructs a [ProductsRequest] from an unsafe.Pointer.
//
// An object that can retrieve localized information from the App Store about a specified list of products.
func ProductsRequestFrom(ptr unsafe.Pointer) ProductsRequest {
	return ProductsRequest{
		Request: RequestFrom(ptr),
	}
}
/* debug [class_struct]: End struct */



/* debug [class_init_methods]: Init methods for ProductsRequest */

// Initializes the request with the set of product identifiers.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductsRequest/init(productIdentifiers:)
func NewProductsRequestWithProductIdentifiers(productIdentifiers unsafe.Pointer) ProductsRequest {
	instance := getProductsRequestClass().Alloc()
	rv := objc.Send[ProductsRequest](instance.ID, objc.Sel("initWithProductIdentifiers:"), productIdentifiers)
	rv.Autorelease()
	return rv
}/* debug [class_init_methods/constructor]: NewProductsRequestWithProductIdentifiers */

/* debug [class_init_methods]: End init methods */



/* debug [class_methods]: Class methods for ProductsRequest */
/* debug [class_methods]: End class methods */



/* debug [class_properties_class]: Class properties for ProductsRequest */
/* debug [class_properties_class]: End class properties */



/* debug [instance_methods]: Instance methods for ProductsRequest */
/* debug [instance_methods]: End instance methods */



/* debug [instance_properties]: Instance properties for ProductsRequest */

// The delegate that receives the response of the app’s products request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductsRequest/delegate
func (p_ ProductsRequest) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](p_.ID, objc.Sel("delegate"))
	return rv
}/* debug [instance_properties/getter]: delegate */


// The delegate that receives the response of the app’s products request.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKProductsRequest/delegate
func (p_ ProductsRequest) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](p_.ID, objc.Sel("setDelegate:"), value)
}/* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */


/* debug [class.gen.go]: End class SKProductsRequest */


