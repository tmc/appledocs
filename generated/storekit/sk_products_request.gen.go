// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
)

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

// An interface definition for the [ProductsRequest] class.
type IProductsRequest interface {
	IRequest
}

// An object that can retrieve localized information from the App Store about a specified list of products.
//
// Your app uses an object to present localized prices and other information to the user without having to maintain that list of product information itself. To use an object, you initialize it with a list of product identifier strings, attach a delegate, and then call the request’s method. When the request completes, your delegate receives an object.
//
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

// Alloc allocates a new instance without initialization.
func (pc _ProductsRequestClass) Alloc() ProductsRequest {
	rv := objc.Send[ProductsRequest](objc.ID(pc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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




