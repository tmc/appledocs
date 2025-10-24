// Code generated from Apple documentation for PassKit. DO NOT EDIT.

package passkit

import (
	"unsafe"
)

// PPaymentInformationRequestHandling is the PKPaymentInformationRequestHandling protocol interface.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - visionOS +
//   - watchOS +
//
// See: doc://com.apple.passkit/documentation/PassKit/PKPaymentInformationRequestHandling
type PPaymentInformationRequestHandling interface {
	// Required methods
	HandleSignatureRequestCompletion(signatureRequest BarcodeEventSignatureRequest /* not a class type */, completion SignatureRequestCompletionBlock /* not a class type */)/* debug [protocol_interface/required_method]: HandleSignatureRequestCompletion */
	HandleConfigurationRequestCompletion(configurationRequest BarcodeEventConfigurationRequest /* not a class type */, completion unsafe.Pointer)/* debug [protocol_interface/required_method]: HandleConfigurationRequestCompletion */
	HandleInformationRequestCompletion(infoRequest BarcodeEventMetadataRequest /* not a class type */, completion InformationRequestCompletionBlock /* not a class type */)/* debug [protocol_interface/required_method]: HandleInformationRequestCompletion */
}
