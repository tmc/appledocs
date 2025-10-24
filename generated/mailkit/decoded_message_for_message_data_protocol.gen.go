// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import "github.com/ebitengine/purego/objc"

// decodedMessageForMessageDataProtocol is the decodedMessageForMessageData: protocol.
//
// Availability:
//   - macOS 12.0+
//
// Use this protocol when registering custom classes that conform to decodedMessageForMessageData:.
var decodedMessageForMessageDataProtocol *objc.Protocol

func init() {
	decodedMessageForMessageDataProtocol = objc.GetProtocol("decodedMessageForMessageData:")
}

