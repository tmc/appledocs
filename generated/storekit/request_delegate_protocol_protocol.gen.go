// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PRequestDelegate is the SKRequestDelegate protocol interface.
//
// Common methods that are implemented by delegates for any subclass of the   abstract class.
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
// See: doc://com.apple.storekit/documentation/StoreKit/SKRequestDelegate
type PRequestDelegate interface {
	// Optional methods
	RequestDidFailWithError(request ISKRequest, error_ objc.IObject /* cross-framework: Error */)
	HasRequestDidFailWithError() bool
	RequestDidFinish(request ISKRequest)
	HasRequestDidFinish() bool
}

// RequestDelegate is a delegate implementation builder for the PRequestDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type RequestDelegate struct {
	_RequestDidFailWithError func(request ISKRequest, error_ objc.IObject /* cross-framework: Error */)
	_RequestDidFinish func(request ISKRequest)
}

// SetRequestDidFailWithError sets the handler for the RequestDidFailWithError delegate method.
//
// Tells the delegate that the request failed to execute.
func (d *RequestDelegate) SetRequestDidFailWithError(f func(request ISKRequest, error_ objc.IObject /* cross-framework: Error */)) {
	d._RequestDidFailWithError = f
}

// SetRequestDidFinish sets the handler for the RequestDidFinish delegate method.
//
// Tells the delegate that the request has completed.
func (d *RequestDelegate) SetRequestDidFinish(f func(request ISKRequest)) {
	d._RequestDidFinish = f
}

// RequestDidFailWithError implements the PRequestDelegate interface.
func (d *RequestDelegate) RequestDidFailWithError(request ISKRequest, error_ objc.IObject /* cross-framework: Error */) {
	if d._RequestDidFailWithError != nil {
		d._RequestDidFailWithError(request, error_)
	}
}

// HasRequestDidFailWithError returns true if a handler for RequestDidFailWithError has been set.
func (d *RequestDelegate) HasRequestDidFailWithError() bool {
	return d._RequestDidFailWithError != nil
}

// RequestDidFinish implements the PRequestDelegate interface.
func (d *RequestDelegate) RequestDidFinish(request ISKRequest) {
	if d._RequestDidFinish != nil {
		d._RequestDidFinish(request)
	}
}

// HasRequestDidFinish returns true if a handler for RequestDidFinish has been set.
func (d *RequestDelegate) HasRequestDidFinish() bool {
	return d._RequestDidFinish != nil
}
