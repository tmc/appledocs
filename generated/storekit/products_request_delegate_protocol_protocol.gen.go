// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (

	"github.com/tmc/appledocs/generated/objc"
)

// PProductsRequestDelegate is the SKProductsRequestDelegate protocol interface.
//
// A set of methods the delegate implements so it receives the product information your app requests.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 18.0)
//   - iOS 3.0+ (Deprecated in 18.0)
//   - iPadOS 3.0+ (Deprecated in 18.0)
//   - macOS 10.7+ (Deprecated in 15.0)
//   - tvOS + (Deprecated in 18.0)
//   - visionOS 1.0+ (Deprecated in 2.0)
//   - watchOS 6.2+ (Deprecated in 11.0)
//
// See: doc://com.apple.storekit/documentation/StoreKit/SKProductsRequestDelegate
type PProductsRequestDelegate interface {
	// Required methods
	ProductsRequestDidReceiveResponse(request ISKProductsRequest, response ISKProductsResponse)/* debug [protocol_interface/required_method]: ProductsRequestDidReceiveResponse */
}

// ProductsRequestDelegate is a delegate implementation builder for the PProductsRequestDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type ProductsRequestDelegate struct {
	_ProductsRequestDidReceiveResponse func(request ISKProductsRequest, response ISKProductsResponse)
}

// SetProductsRequestDidReceiveResponse sets the handler for the ProductsRequestDidReceiveResponse delegate method.
//
// Accepts the App Store response that contains the app-requested product information.
func (d *ProductsRequestDelegate) SetProductsRequestDidReceiveResponse(f func(request ISKProductsRequest, response ISKProductsResponse)) {
	d._ProductsRequestDidReceiveResponse = f
}

// ProductsRequestDidReceiveResponse implements the PProductsRequestDelegate interface.
func (d *ProductsRequestDelegate) ProductsRequestDidReceiveResponse(request ISKProductsRequest, response ISKProductsResponse) {
	if d._ProductsRequestDidReceiveResponse != nil {
		d._ProductsRequestDidReceiveResponse(request, response)
	}
}

// HasProductsRequestDidReceiveResponse returns true if a handler for ProductsRequestDidReceiveResponse has been set.
func (d *ProductsRequestDelegate) HasProductsRequestDidReceiveResponse() bool {
	return d._ProductsRequestDidReceiveResponse != nil
}
