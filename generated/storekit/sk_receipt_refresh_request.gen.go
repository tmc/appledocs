// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (
	"sync"
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"
	"github.com/tmc/appledocs/generated/foundation"
)

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

// An interface definition for the [ReceiptRefreshRequest] class.
type IReceiptRefreshRequest interface {
	IRequest
	// properties:
	AppStoreReceiptURL() objc.IObject /* cross-framework: URL */
	SetAppStoreReceiptURL(value objc.IObject /* cross-framework: URL */)
	SKReceiptPropertyIsExpired() objc.IObject /* cross-framework: NSString */
	SKReceiptPropertyIsRevoked() objc.IObject /* cross-framework: NSString */
	SKReceiptPropertyIsVolumePurchase() objc.IObject /* cross-framework: NSString */
	ReceiptProperties() objc.IObject /* cross-framework: NSString */
	SetReceiptProperties(value objc.IObject /* cross-framework: NSString */)
	Delegate() RequestDelegate /* not a class type */
	SetDelegate(value RequestDelegate /* not a class type */)
	// methods:
}

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

// Alloc allocates a new instance without initialization.
func (rc _ReceiptRefreshRequestClass) Alloc() ReceiptRefreshRequest {
	rv := objc.Send[ReceiptRefreshRequest](objc.ID(rc.class), objc.Sel("alloc"))
	return rv
}

// New creates and returns a new autoreleased instance (equivalent to [[Class alloc] init]).
// Note: Despite the name, this returns an autoreleased object for consistency with Go patterns.
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



// The file URL for the bundle’s App Store receipt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/appStoreReceiptURL
func (r_ ReceiptRefreshRequest) AppStoreReceiptURL() objc.IObject /* cross-framework: URL */ {
	rv := objc.Send[foundation.URL](r_.ID, objc.Sel("appStoreReceiptURL"))
	return rv
}


// The file URL for the bundle’s App Store receipt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/appStoreReceiptURL
func (r_ ReceiptRefreshRequest) SetAppStoreReceiptURL(value objc.IObject /* cross-framework: URL */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setAppStoreReceiptURL:"), value)
}


// A key with a value that indicates whether the receipt is in an expired state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skreceiptpropertyisexpired
func (r_ ReceiptRefreshRequest) SKReceiptPropertyIsExpired() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("SKReceiptPropertyIsExpired"))
	return rv
}


// A key with a value that indicates whether the receipt is in a revoked state.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skreceiptpropertyisrevoked
func (r_ ReceiptRefreshRequest) SKReceiptPropertyIsRevoked() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("SKReceiptPropertyIsRevoked"))
	return rv
}


// A key with a value that indicates whether the receipt is a Volume Purchase Plan receipt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skreceiptpropertyisvolumepurchase
func (r_ ReceiptRefreshRequest) SKReceiptPropertyIsVolumePurchase() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("SKReceiptPropertyIsVolumePurchase"))
	return rv
}


// The properties of the receipt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skreceiptrefreshrequest/receiptproperties
func (r_ ReceiptRefreshRequest) ReceiptProperties() objc.IObject /* cross-framework: NSString */ {
	rv := objc.Send[foundation.NSString](r_.ID, objc.Sel("receiptProperties"))
	return rv
}


// The properties of the receipt.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skreceiptrefreshrequest/receiptproperties
func (r_ ReceiptRefreshRequest) SetReceiptProperties(value objc.IObject /* cross-framework: NSString */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setReceiptProperties:"), value)
}


// The delegate of the request object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skrequest/delegate
func (r_ ReceiptRefreshRequest) Delegate() RequestDelegate /* not a class type */ {
	rv := objc.Send[RequestDelegate](r_.ID, objc.Sel("delegate"))
	return rv
}


// The delegate of the request object.
//
// [Full Topic]
// [Full Topic]: https://developer.apple.com/documentation/storekit/skrequest/delegate
func (r_ ReceiptRefreshRequest) SetDelegate(value RequestDelegate /* not a class type */) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDelegate:"), value)
}



