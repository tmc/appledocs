// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// PCXCallDirectoryExtensionContextDelegate is the CXCallDirectoryExtensionContextDelegate protocol interface.
//
// A collection of methods a Call Directory extension context object calls when a request fails.
//
// Availability:
//   - Mac Catalyst 10.0+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.callkit/documentation/CallKit/CXCallDirectoryExtensionContextDelegate
type PCXCallDirectoryExtensionContextDelegate interface {
	// Required methods
	RequestFailedForExtensionContextWithError(extensionContext ICXCallDirectoryExtensionContext, error_ objc.IObject /* cross-framework: Error */)/* debug [protocol_interface/required_method]: RequestFailedForExtensionContextWithError */
}

// CXCallDirectoryExtensionContextDelegate is a delegate implementation builder for the PCXCallDirectoryExtensionContextDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type CXCallDirectoryExtensionContextDelegate struct {
	_RequestFailedForExtensionContextWithError func(extensionContext ICXCallDirectoryExtensionContext, error_ objc.IObject /* cross-framework: Error */)
}

// SetRequestFailedForExtensionContextWithError sets the handler for the RequestFailedForExtensionContextWithError delegate method.
//
// Called when a Call Directory app extension request fails.
func (d *CXCallDirectoryExtensionContextDelegate) SetRequestFailedForExtensionContextWithError(f func(extensionContext ICXCallDirectoryExtensionContext, error_ objc.IObject /* cross-framework: Error */)) {
	d._RequestFailedForExtensionContextWithError = f
}

// RequestFailedForExtensionContextWithError implements the PCXCallDirectoryExtensionContextDelegate interface.
func (d *CXCallDirectoryExtensionContextDelegate) RequestFailedForExtensionContextWithError(extensionContext ICXCallDirectoryExtensionContext, error_ objc.IObject /* cross-framework: Error */) {
	if d._RequestFailedForExtensionContextWithError != nil {
		d._RequestFailedForExtensionContextWithError(extensionContext, error_)
	}
}

// HasRequestFailedForExtensionContextWithError returns true if a handler for RequestFailedForExtensionContextWithError has been set.
func (d *CXCallDirectoryExtensionContextDelegate) HasRequestFailedForExtensionContextWithError() bool {
	return d._RequestFailedForExtensionContextWithError != nil
}
