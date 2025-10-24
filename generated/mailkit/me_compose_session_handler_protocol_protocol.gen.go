// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PMEComposeSessionHandler is the MEComposeSessionHandler protocol interface.
//
// An object that participates in the composition of mail messages, and annotates recipient tokens.
//
// Availability:
//   - macOS 12.0+
//
// See: doc://com.apple.mailkit/documentation/MailKit/MEComposeSessionHandler
type PMEComposeSessionHandler interface {
	// Required methods
	MailComposeSessionDidBegin(session IMEComposeSession)/* debug [protocol_interface/required_method]: MailComposeSessionDidBegin */
	MailComposeSessionDidEnd(session IMEComposeSession)/* debug [protocol_interface/required_method]: MailComposeSessionDidEnd */
	ViewControllerForSession(session IMEComposeSession) MEExtensionViewController/* debug [protocol_interface/required_method]: ViewControllerForSession */
	// Optional methods
	AdditionalHeadersForSession(session IMEComposeSession) foundation.IDictionary
	HasAdditionalHeadersForSession() bool
	SessionCanSendMessageWithCompletionHandler(session IMEComposeSession, completion unsafe.Pointer)
	HasSessionCanSendMessageWithCompletionHandler() bool
	SessionAnnotateAddressesWithCompletionHandler(session IMEComposeSession, completionHandler unsafe.Pointer)
	HasSessionAnnotateAddressesWithCompletionHandler() bool
}
