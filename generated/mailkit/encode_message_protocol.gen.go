// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import "github.com/ebitengine/purego/objc"

// encodeMessageProtocol is the encodeMessage: protocol.
//
// Availability:
//   - macOS 12.0+
//
// Use this protocol when registering custom classes that conform to encodeMessage:.
var encodeMessageProtocol *objc.Protocol

func init() {
	encodeMessageProtocol = objc.GetProtocol("encodeMessage:")
}

