// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"unsafe"
)

// PMEMessageEncoder is the MEMessageEncoder protocol interface.
//
// An object that encrypts or digitally signs outgoing messages.
//
// Availability:
//   - macOS 12.0+
//
// See: doc://com.apple.mailkit/documentation/MailKit/MEMessageEncoder
type PMEMessageEncoder interface {
	// Required methods
	EncodeMessageComposeContextCompletionHandler(message IMEMessage, composeContext IMEComposeContext, completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: EncodeMessageComposeContextCompletionHandler */
	GetEncodingStatusForMessageComposeContextCompletionHandler(message IMEMessage, composeContext IMEComposeContext, completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: GetEncodingStatusForMessageComposeContextCompletionHandler */
}
