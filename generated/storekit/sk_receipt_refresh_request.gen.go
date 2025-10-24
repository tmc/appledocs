// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
	"github.com/tmc/appledocs/generated/objc"
)

/* debug [class.gen.go]: Generating class SKReceiptRefreshRequest */

/* debug [class_header]: Header for SKReceiptRefreshRequest */
// The class instance for the [ReceiptRefreshRequest] class.
var (
	ReceiptRefreshRequestClass     _ReceiptRefreshRequestClass
	ReceiptRefreshRequestClassOnce sync.Once
)

func getReceiptRefreshRequestClass() _ReceiptRefreshRequestClass {
	ReceiptRefreshRequestClassOnce.Do(func() {
		ReceiptRefreshRequestClass = _ReceiptRefreshRequestClass{objc.GetClass("SKReceiptRefreshRequest")}
	})
	return ReceiptRefreshRequestClass
}

type _ReceiptRefreshRequestClass struct {
	class objc.Class
}

/* debug [class_header]: End header */

/* debug [class_interface]: Interface for ReceiptRefreshRequest */
// An interface definition for the [ReceiptRefreshRequest] class.
type IReceiptRefreshRequest interface {
	IRequest

	/* debug [class_interface_properties]: Properties for ReceiptRefreshRequest */
	// properties:
	ReceiptProperties() foundation.IDictionary
	AppStoreReceiptURL() foundation.URL
	SetAppStoreReceiptURL(value foundation.URL)
	SKReceiptPropertyIsExpired() objc.IObject        /* cross-framework: NSString */
	SKReceiptPropertyIsRevoked() objc.IObject        /* cross-framework: NSString */
	SKReceiptPropertyIsVolumePurchase() objc.IObject /* cross-framework: NSString */
	Delegate() objc.IObject                          /* cross-framework: RequestDelegate */
	SetDelegate(value objc.IObject /* cross-framework: RequestDelegate */)
	/* debug [class_interface_properties]: End properties */

	/* debug [class_interface_methods]: Methods for ReceiptRefreshRequest */
	// methods:
	/* debug [class_interface_methods]: End methods */

}

/* debug [class_interface]: End interface */

/* debug [class_constructors]: Constructors for ReceiptRefreshRequest */
// Alloc allocates a new instance without initialization.
func (rc _ReceiptRefreshRequestClass) Alloc() ReceiptRefreshRequest {
	rv := objc.Send[ReceiptRefreshRequest](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
func (rc _ReceiptRefreshRequestClass) New() ReceiptRefreshRequest {
	rv := objc.Send[ReceiptRefreshRequest](objc.ID(rc.class), objc.Sel("new"))
	rv.Autorelease()
	return rv
}

// Init initializes the instance.
func (r_ ReceiptRefreshRequest) Init() ReceiptRefreshRequest {
	rv := objc.Send[ReceiptRefreshRequest](r_.ID, objc.Sel("init"))
	return rv
}

// Autorelease adds the receiver to the current autorelease pool.
func (r_ ReceiptRefreshRequest) Autorelease() ReceiptRefreshRequest {
	rv := objc.Send[ReceiptRefreshRequest](r_.ID, objc.Sel("autorelease"))
	return rv
}

// NewReceiptRefreshRequest creates a new ReceiptRefreshRequest instance.
func NewReceiptRefreshRequest() ReceiptRefreshRequest {
	return getReceiptRefreshRequestClass().New()
}

/* debug [class_constructors]: End constructors */

/* debug [class_struct]: Struct for ReceiptRefreshRequest */
// A request to the App Store to get the app receipt, which represents the customer’s transactions with your app.
//
// Use this API to request a new app receipt from the App Store if the receipt is invalid or missing from its expected location, . To request the receipt using the object, you initialize it, attach a , and then call the request’s method. When the request completes successfully, your delegate receives an object in its method. Locate the app receipt using the property. For information about validating the receipt, see . If the request fails and calls your delegate’s method, your app needs to release the request and not attempt to call it a second time. Requests can fail when a user doesn’t authenticate or chooses to cancel the request. Without a validated receipt, assume the user doesn’t have access to premium content. In the sandbox environment, you can initialize a receipt with any combination of properties for testing when you call .

// A request to the App Store to get the app receipt, which represents the customer’s transactions with your app.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKReceiptRefreshRequest
type ReceiptRefreshRequest struct {
	Request
}

// ReceiptRefreshRequestFrom constructs a [ReceiptRefreshRequest] from an unsafe.Pointer.
//
// A request to the App Store to get the app receipt, which represents the customer’s transactions with your app.
func ReceiptRefreshRequestFrom(ptr unsafe.Pointer) ReceiptRefreshRequest {
	return ReceiptRefreshRequest{
		Request: RequestFrom(ptr),
	}
}

/* debug [class_struct]: End struct */

/* debug [class_init_methods]: Init methods for ReceiptRefreshRequest */

// Creates a receipt refresh request with optional properties.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKReceiptRefreshRequest/init(receiptProperties:)
func NewReceiptRefreshRequestWithReceiptProperties(properties foundation.IDictionary) ReceiptRefreshRequest {
	instance := getReceiptRefreshRequestClass().Alloc()
	rv := objc.Send[ReceiptRefreshRequest](instance.ID, objc.Sel("initWithReceiptProperties:"), properties)
	rv.Autorelease()
	return rv
} /* debug [class_init_methods/constructor]: NewReceiptRefreshRequestWithReceiptProperties */

/* debug [class_init_methods]: End init methods */

/* debug [class_methods]: Class methods for ReceiptRefreshRequest */
/* debug [class_methods]: End class methods */

/* debug [class_properties_class]: Class properties for ReceiptRefreshRequest */
/* debug [class_properties_class]: End class properties */

/* debug [instance_methods]: Instance methods for ReceiptRefreshRequest */
/* debug [instance_methods]: End instance methods */

/* debug [instance_properties]: Instance properties for ReceiptRefreshRequest */

// The properties of the receipt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKReceiptRefreshRequest/receiptProperties
func (r_ ReceiptRefreshRequest) ReceiptProperties() foundation.IDictionary {
	rv := objc.Send[foundation.IDictionary](r_.ID, objc.Sel("receiptProperties"))
	return rv
} /* debug [instance_properties/getter]: receiptProperties */

// The file URL for the bundle’s App Store receipt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/appStoreReceiptURL
func (r_ ReceiptRefreshRequest) AppStoreReceiptURL() foundation.URL {
	rv := objc.Send[foundation.URL](r_.ID, objc.Sel("appStoreReceiptURL"))
	return rv
} /* debug [instance_properties/getter]: appStoreReceiptURL */

// The file URL for the bundle’s App Store receipt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/appStoreReceiptURL
func (r_ ReceiptRefreshRequest) SetAppStoreReceiptURL(value foundation.URL) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setAppStoreReceiptURL:"), value)
} /* debug [instance_properties/setter]: appStoreReceiptURL */

// A key with a value that indicates whether the receipt is in an expired state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skreceiptpropertyisexpired
func (r_ ReceiptRefreshRequest) SKReceiptPropertyIsExpired() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("SKReceiptPropertyIsExpired"))
	return rv
} /* debug [instance_properties/getter]: SKReceiptPropertyIsExpired */

// A key with a value that indicates whether the receipt is in a revoked state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skreceiptpropertyisrevoked
func (r_ ReceiptRefreshRequest) SKReceiptPropertyIsRevoked() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("SKReceiptPropertyIsRevoked"))
	return rv
} /* debug [instance_properties/getter]: SKReceiptPropertyIsRevoked */

// A key with a value that indicates whether the receipt is a Volume Purchase Plan receipt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skreceiptpropertyisvolumepurchase
func (r_ ReceiptRefreshRequest) SKReceiptPropertyIsVolumePurchase() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("SKReceiptPropertyIsVolumePurchase"))
	return rv
} /* debug [instance_properties/getter]: SKReceiptPropertyIsVolumePurchase */

// The delegate of the request object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skrequest/delegate
func (r_ ReceiptRefreshRequest) Delegate() objc.IObject /* cross-framework: RequestDelegate */ {
	rv := objc.Send[objc.ID](r_.ID, objc.Sel("delegate"))
	return rv
} /* debug [instance_properties/getter]: delegate */

// The delegate of the request object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skrequest/delegate
func (r_ ReceiptRefreshRequest) SetDelegate(value objc.IObject /* cross-framework: RequestDelegate */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDelegate:"), value)
} /* debug [instance_properties/setter]: delegate */

/* debug [instance_properties]: End instance properties */

/* debug [class.gen.go]: End class SKReceiptRefreshRequest */
