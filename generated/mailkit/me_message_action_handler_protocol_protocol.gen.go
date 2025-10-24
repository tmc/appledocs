// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"unsafe"
)

// PMEMessageActionHandler is the MEMessageActionHandler protocol interface.
//
// An object that performs actions on messages as the system downloads them.
//
// Availability:
//   - macOS 12.0+
//
// See: doc://com.apple.mailkit/documentation/MailKit/MEMessageActionHandler
type PMEMessageActionHandler interface {
	// Required methods
	DecideActionForMessageCompletionHandler(message IMEMessage, completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: DecideActionForMessageCompletionHandler */
}
