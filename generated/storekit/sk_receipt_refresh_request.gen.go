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
}

// A request to the App Store to get the app receipt, which represents the customer’s transactions with your app.
//
// Use this API to request a new app receipt from the App Store if the receipt is invalid or missing from its expected location, . To request the receipt using the object, you initialize it, attach a , and then call the request’s method. When the request completes successfully, your delegate receives an object in its method. Locate the app receipt using the property. For information about validating the receipt, see . If the request fails and calls your delegate’s method, your app needs to release the request and not attempt to call it a second time. Requests can fail when a user doesn’t authenticate or chooses to cancel the request. Without a validated receipt, assume the user doesn’t have access to premium content. In the sandbox environment, you can initialize a receipt with any combination of properties for testing when you call .
//
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




// Creates a receipt refresh request with optional properties.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKReceiptRefreshRequest/init(receiptProperties:)
func NewReceiptRefreshRequestWithReceiptProperties(properties unsafe.Pointer) ReceiptRefreshRequest {
	instance := getReceiptRefreshRequestClass().Alloc()
	rv := objc.Send[ReceiptRefreshRequest](instance.ID, objc.Sel("initWithReceiptProperties:"), properties)
	rv.Autorelease()
	return rv
}


// The properties of the receipt.
//
// [Full Topic]: https://developer.apple.com/documentation/StoreKit/SKReceiptRefreshRequest/receiptProperties
func (r_ ReceiptRefreshRequest) ReceiptProperties() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("receiptProperties"))
	return rv
}

// The file URL for the bundle’s App Store receipt.
//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/appStoreReceiptURL
func (r_ ReceiptRefreshRequest) AppStoreReceiptURL() foundation.URL {
	rv := objc.Send[foundation.URL](r_.ID, objc.Sel("appStoreReceiptURL"))
	return rv
}


// SetAppStoreReceiptURL sets the value of the appStoreReceiptURL property.
// The file URL for the bundle’s App Store receipt.

//
// [Full Topic]: https://developer.apple.com/documentation/Foundation/Bundle/appStoreReceiptURL
func (r_ ReceiptRefreshRequest) SetAppStoreReceiptURL(value foundation.IURL) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setAppStoreReceiptURL:"), value)
}

// A key with a value that indicates whether the receipt is in an expired state.
//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skreceiptpropertyisexpired
func (r_ ReceiptRefreshRequest) SKReceiptPropertyIsExpired() appkit.string {
	rv := objc.Send[appkit.string](r_.ID, objc.Sel("SKReceiptPropertyIsExpired"))
	return rv
}

// A key with a value that indicates whether the receipt is in a revoked state.
//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skreceiptpropertyisrevoked
func (r_ ReceiptRefreshRequest) SKReceiptPropertyIsRevoked() appkit.string {
	rv := objc.Send[appkit.string](r_.ID, objc.Sel("SKReceiptPropertyIsRevoked"))
	return rv
}

// A key with a value that indicates whether the receipt is a Volume Purchase Plan receipt.
//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skreceiptpropertyisvolumepurchase
func (r_ ReceiptRefreshRequest) SKReceiptPropertyIsVolumePurchase() appkit.string {
	rv := objc.Send[appkit.string](r_.ID, objc.Sel("SKReceiptPropertyIsVolumePurchase"))
	return rv
}

// The delegate of the request object.
//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skrequest/delegate
func (r_ ReceiptRefreshRequest) Delegate() unsafe.Pointer {
	rv := objc.Send[unsafe.Pointer](r_.ID, objc.Sel("delegate"))
	return rv
}


// SetDelegate sets the value of the delegate property.
// The delegate of the request object.

//
// [Full Topic]: https://developer.apple.com/documentation/storekit/skrequest/delegate
func (r_ ReceiptRefreshRequest) SetDelegate(value unsafe.Pointer) {
	objc.Send[objc.ID](r_.ID, objc.Sel("setDelegate:"), value)
}


