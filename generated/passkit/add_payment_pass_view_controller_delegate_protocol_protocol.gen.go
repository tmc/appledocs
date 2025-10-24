// Code generated from Apple documentation for PassKit. DO NOT EDIT.

package passkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/objc"

	"github.com/tmc/appledocs/generated/coretelephony"

	"github.com/tmc/appledocs/generated/foundation"
)

// PAddPaymentPassViewControllerDelegate is the PKAddPaymentPassViewControllerDelegate protocol interface.
//
// Methods that let the system prompt you for an add payment request, and inform you when a request has succeeded or failed.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - visionOS 1.0+
//
// See: doc://com.apple.passkit/documentation/PassKit/PKAddPaymentPassViewControllerDelegate
type PAddPaymentPassViewControllerDelegate interface {
	// Required methods
	AddPaymentPassViewControllerDidFinishAddingPaymentPassError(controller AddPaymentPassViewController /* not a class type */, pass PaymentPass /* not a class type */, error_ objc.IObject /* cross-framework: Error */)/* debug [protocol_interface/required_method]: AddPaymentPassViewControllerDidFinishAddingPaymentPassError */
	AddPaymentPassViewControllerGenerateRequestWithCertificateChainNonceNonceSignatureCompletionHandler(controller AddPaymentPassViewController /* not a class type */, certificates []foundation.Data, nonce objc.IObject /* cross-framework: NSData */, nonceSignature objc.IObject /* cross-framework: NSData */, handler unsafe.Pointer)/* debug [protocol_interface/required_method]: AddPaymentPassViewControllerGenerateRequestWithCertificateChainNonceNonceSignatureCompletionHandler */
}

// AddPaymentPassViewControllerDelegate is a delegate implementation builder for the PAddPaymentPassViewControllerDelegate protocol.
//
// Use this struct to create a custom delegate by setting handler functions for the methods you want to implement.
type AddPaymentPassViewControllerDelegate struct {
	_AddPaymentPassViewControllerDidFinishAddingPaymentPassError func(controller AddPaymentPassViewController /* not a class type */, pass PaymentPass /* not a class type */, error_ objc.IObject /* cross-framework: Error */)
	_AddPaymentPassViewControllerGenerateRequestWithCertificateChainNonceNonceSignatureCompletionHandler func(controller AddPaymentPassViewController /* not a class type */, certificates []foundation.Data, nonce objc.IObject /* cross-framework: NSData */, nonceSignature objc.IObject /* cross-framework: NSData */, handler unsafe.Pointer)
}

// SetAddPaymentPassViewControllerDidFinishAddingPaymentPassError sets the handler for the AddPaymentPassViewControllerDidFinishAddingPaymentPassError delegate method.
func (d *AddPaymentPassViewControllerDelegate) SetAddPaymentPassViewControllerDidFinishAddingPaymentPassError(f func(controller AddPaymentPassViewController /* not a class type */, pass PaymentPass /* not a class type */, error_ objc.IObject /* cross-framework: Error */)) {
	d._AddPaymentPassViewControllerDidFinishAddingPaymentPassError = f
}

// SetAddPaymentPassViewControllerGenerateRequestWithCertificateChainNonceNonceSignatureCompletionHandler sets the handler for the AddPaymentPassViewControllerGenerateRequestWithCertificateChainNonceNonceSignatureCompletionHandler delegate method.
//
// Asks the delegate to create an add payment request.
func (d *AddPaymentPassViewControllerDelegate) SetAddPaymentPassViewControllerGenerateRequestWithCertificateChainNonceNonceSignatureCompletionHandler(f func(controller AddPaymentPassViewController /* not a class type */, certificates []foundation.Data, nonce objc.IObject /* cross-framework: NSData */, nonceSignature objc.IObject /* cross-framework: NSData */, handler unsafe.Pointer)) {
	d._AddPaymentPassViewControllerGenerateRequestWithCertificateChainNonceNonceSignatureCompletionHandler = f
}

// AddPaymentPassViewControllerDidFinishAddingPaymentPassError implements the PAddPaymentPassViewControllerDelegate interface.
func (d *AddPaymentPassViewControllerDelegate) AddPaymentPassViewControllerDidFinishAddingPaymentPassError(controller AddPaymentPassViewController /* not a class type */, pass PaymentPass /* not a class type */, error_ objc.IObject /* cross-framework: Error */) {
	if d._AddPaymentPassViewControllerDidFinishAddingPaymentPassError != nil {
		d._AddPaymentPassViewControllerDidFinishAddingPaymentPassError(controller, pass, error_)
	}
}

// HasAddPaymentPassViewControllerDidFinishAddingPaymentPassError returns true if a handler for AddPaymentPassViewControllerDidFinishAddingPaymentPassError has been set.
func (d *AddPaymentPassViewControllerDelegate) HasAddPaymentPassViewControllerDidFinishAddingPaymentPassError() bool {
	return d._AddPaymentPassViewControllerDidFinishAddingPaymentPassError != nil
}

// AddPaymentPassViewControllerGenerateRequestWithCertificateChainNonceNonceSignatureCompletionHandler implements the PAddPaymentPassViewControllerDelegate interface.
func (d *AddPaymentPassViewControllerDelegate) AddPaymentPassViewControllerGenerateRequestWithCertificateChainNonceNonceSignatureCompletionHandler(controller AddPaymentPassViewController /* not a class type */, certificates []foundation.Data, nonce objc.IObject /* cross-framework: NSData */, nonceSignature objc.IObject /* cross-framework: NSData */, handler unsafe.Pointer) {
	if d._AddPaymentPassViewControllerGenerateRequestWithCertificateChainNonceNonceSignatureCompletionHandler != nil {
		d._AddPaymentPassViewControllerGenerateRequestWithCertificateChainNonceNonceSignatureCompletionHandler(controller, certificates, nonce, nonceSignature, handler)
	}
}

// HasAddPaymentPassViewControllerGenerateRequestWithCertificateChainNonceNonceSignatureCompletionHandler returns true if a handler for AddPaymentPassViewControllerGenerateRequestWithCertificateChainNonceNonceSignatureCompletionHandler has been set.
func (d *AddPaymentPassViewControllerDelegate) HasAddPaymentPassViewControllerGenerateRequestWithCertificateChainNonceNonceSignatureCompletionHandler() bool {
	return d._AddPaymentPassViewControllerGenerateRequestWithCertificateChainNonceNonceSignatureCompletionHandler != nil
}
