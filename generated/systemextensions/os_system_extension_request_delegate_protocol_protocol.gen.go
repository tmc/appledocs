// Code generated from Apple documentation for SystemExtensions. DO NOT EDIT.

package systemextensions

import (

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"
)

// POSSystemExtensionRequestDelegate is the OSSystemExtensionRequestDelegate protocol interface.
//
// A type that receives updates about the progress of a request.
//
// Availability:
//   - macOS 10.15+
//
// See: doc://com.apple.systemextensions/documentation/SystemExtensions/OSSystemExtensionRequestDelegate
type POSSystemExtensionRequestDelegate interface {
	// Required methods
	RequestActionForReplacingExtensionWithExtension(request IOSSystemExtensionRequest, existing IOSSystemExtensionProperties, ext IOSSystemExtensionProperties) OSSystemExtensionReplacementAction/* debug [protocol_interface/required_method]: RequestActionForReplacingExtensionWithExtension */
	RequestDidFailWithError(request IOSSystemExtensionRequest, error_ objc.IObject /* cross-framework: Error */)/* debug [protocol_interface/required_method]: RequestDidFailWithError */
	RequestDidFinishWithResult(request IOSSystemExtensionRequest, result OSSystemExtensionRequestResult)/* debug [protocol_interface/required_method]: RequestDidFinishWithResult */
	RequestNeedsUserApproval(request IOSSystemExtensionRequest)/* debug [protocol_interface/required_method]: RequestNeedsUserApproval */
	// Optional methods
	RequestFoundProperties(request IOSSystemExtensionRequest, properties []OSSystemExtensionProperties)
	HasRequestFoundProperties() bool
}

// OSSystemExtensionRequestDelegate is a delegate implementation builder for the POSSystemExtensionRequestDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type OSSystemExtensionRequestDelegate struct {
	_RequestFoundProperties func(request IOSSystemExtensionRequest, properties []OSSystemExtensionProperties)
	_RequestActionForReplacingExtensionWithExtension func(request IOSSystemExtensionRequest, existing IOSSystemExtensionProperties, ext IOSSystemExtensionProperties) OSSystemExtensionReplacementAction
	_RequestDidFailWithError func(request IOSSystemExtensionRequest, error_ objc.IObject /* cross-framework: Error */)
	_RequestDidFinishWithResult func(request IOSSystemExtensionRequest, result OSSystemExtensionRequestResult)
	_RequestNeedsUserApproval func(request IOSSystemExtensionRequest)
}

// SetRequestFoundProperties sets the handler for the RequestFoundProperties delegate method.
func (d *OSSystemExtensionRequestDelegate) SetRequestFoundProperties(f func(request IOSSystemExtensionRequest, properties []OSSystemExtensionProperties)) {
	d._RequestFoundProperties = f
}

// SetRequestActionForReplacingExtensionWithExtension sets the handler for the RequestActionForReplacingExtensionWithExtension delegate method.
//
// Tells the delegate that the user has a different version of the extension installed on their system.
func (d *OSSystemExtensionRequestDelegate) SetRequestActionForReplacingExtensionWithExtension(f func(request IOSSystemExtensionRequest, existing IOSSystemExtensionProperties, ext IOSSystemExtensionProperties) OSSystemExtensionReplacementAction) {
	d._RequestActionForReplacingExtensionWithExtension = f
}

// SetRequestDidFailWithError sets the handler for the RequestDidFailWithError delegate method.
//
// Tells the delegate the manager failed to complete the request.
func (d *OSSystemExtensionRequestDelegate) SetRequestDidFailWithError(f func(request IOSSystemExtensionRequest, error_ objc.IObject /* cross-framework: Error */)) {
	d._RequestDidFailWithError = f
}

// SetRequestDidFinishWithResult sets the handler for the RequestDidFinishWithResult delegate method.
//
// Tells the delegate that the manager completed the request.
func (d *OSSystemExtensionRequestDelegate) SetRequestDidFinishWithResult(f func(request IOSSystemExtensionRequest, result OSSystemExtensionRequestResult)) {
	d._RequestDidFinishWithResult = f
}

// SetRequestNeedsUserApproval sets the handler for the RequestNeedsUserApproval delegate method.
//
// Tells the delegate that the user must grant approval before the manager can activate the extension.
func (d *OSSystemExtensionRequestDelegate) SetRequestNeedsUserApproval(f func(request IOSSystemExtensionRequest)) {
	d._RequestNeedsUserApproval = f
}

// RequestFoundProperties implements the POSSystemExtensionRequestDelegate interface.
func (d *OSSystemExtensionRequestDelegate) RequestFoundProperties(request IOSSystemExtensionRequest, properties []OSSystemExtensionProperties) {
	if d._RequestFoundProperties != nil {
		d._RequestFoundProperties(request, properties)
	}
}

// HasRequestFoundProperties returns true if a handler for RequestFoundProperties has been set.
func (d *OSSystemExtensionRequestDelegate) HasRequestFoundProperties() bool {
	return d._RequestFoundProperties != nil
}

// RequestActionForReplacingExtensionWithExtension implements the POSSystemExtensionRequestDelegate interface.
func (d *OSSystemExtensionRequestDelegate) RequestActionForReplacingExtensionWithExtension(request IOSSystemExtensionRequest, existing IOSSystemExtensionProperties, ext IOSSystemExtensionProperties) OSSystemExtensionReplacementAction {
	if d._RequestActionForReplacingExtensionWithExtension != nil {
		return d._RequestActionForReplacingExtensionWithExtension(request, existing, ext)
	}
	var zero OSSystemExtensionReplacementAction
	return zero
}

// HasRequestActionForReplacingExtensionWithExtension returns true if a handler for RequestActionForReplacingExtensionWithExtension has been set.
func (d *OSSystemExtensionRequestDelegate) HasRequestActionForReplacingExtensionWithExtension() bool {
	return d._RequestActionForReplacingExtensionWithExtension != nil
}

// RequestDidFailWithError implements the POSSystemExtensionRequestDelegate interface.
func (d *OSSystemExtensionRequestDelegate) RequestDidFailWithError(request IOSSystemExtensionRequest, error_ objc.IObject /* cross-framework: Error */) {
	if d._RequestDidFailWithError != nil {
		d._RequestDidFailWithError(request, error_)
	}
}

// HasRequestDidFailWithError returns true if a handler for RequestDidFailWithError has been set.
func (d *OSSystemExtensionRequestDelegate) HasRequestDidFailWithError() bool {
	return d._RequestDidFailWithError != nil
}

// RequestDidFinishWithResult implements the POSSystemExtensionRequestDelegate interface.
func (d *OSSystemExtensionRequestDelegate) RequestDidFinishWithResult(request IOSSystemExtensionRequest, result OSSystemExtensionRequestResult) {
	if d._RequestDidFinishWithResult != nil {
		d._RequestDidFinishWithResult(request, result)
	}
}

// HasRequestDidFinishWithResult returns true if a handler for RequestDidFinishWithResult has been set.
func (d *OSSystemExtensionRequestDelegate) HasRequestDidFinishWithResult() bool {
	return d._RequestDidFinishWithResult != nil
}

// RequestNeedsUserApproval implements the POSSystemExtensionRequestDelegate interface.
func (d *OSSystemExtensionRequestDelegate) RequestNeedsUserApproval(request IOSSystemExtensionRequest) {
	if d._RequestNeedsUserApproval != nil {
		d._RequestNeedsUserApproval(request)
	}
}

// HasRequestNeedsUserApproval returns true if a handler for RequestNeedsUserApproval has been set.
func (d *OSSystemExtensionRequestDelegate) HasRequestNeedsUserApproval() bool {
	return d._RequestNeedsUserApproval != nil
}
