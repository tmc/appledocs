// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import (

	"github.com/tmc/appledocs/generated/foundation"
)

// PMEMessageDecoder is the MEMessageDecoder protocol interface.
//
// An object that decrypts messages and provides details about digital signatures.
//
// Availability:
//   - macOS 12.0+
//
// See: doc://com.apple.mailkit/documentation/MailKit/MEMessageDecoder
type PMEMessageDecoder interface {
	// Required methods
	DecodedMessageForMessageData(data objc.IObject /* cross-framework: NSData */) MEDecodedMessage/* debug [protocol_interface/required_method]: DecodedMessageForMessageData */
}
