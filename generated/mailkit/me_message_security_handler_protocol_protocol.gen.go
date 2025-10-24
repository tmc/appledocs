// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (
	"unsafe"

	"github.com/tmc/appledocs/generated/foundation"
)

// PMEMessageSecurityHandler is the MEMessageSecurityHandler protocol interface.
//
// An object that digitally signs or encrypts messages the user sends and receives.
//
// Availability:
//   - macOS 12.0+
//
// See: doc://com.apple.mailkit/documentation/MailKit/MEMessageSecurityHandler
type PMEMessageSecurityHandler interface {
	// Required methods
	ExtensionViewControllerForMessageContext(context objc.IObject /* cross-framework: NSData */) MEExtensionViewController/* debug [protocol_interface/required_method]: ExtensionViewControllerForMessageContext */
	ExtensionViewControllerForMessageSigners(messageSigners []MEMessageSigner) MEExtensionViewController/* debug [protocol_interface/required_method]: ExtensionViewControllerForMessageSigners */
	PrimaryActionClickedForMessageContextCompletionHandler(context objc.IObject /* cross-framework: NSData */, completionHandler unsafe.Pointer)/* debug [protocol_interface/required_method]: PrimaryActionClickedForMessageContextCompletionHandler */
}
