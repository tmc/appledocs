// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import "github.com/ebitengine/purego/objc"

// additionalHeadersForSessionProtocol is the additionalHeadersForSession: protocol.
//
// Availability:
//   - macOS 12.0+
//
// Use this protocol when registering custom classes that conform to additionalHeadersForSession:.
var additionalHeadersForSessionProtocol *objc.Protocol

func init() {
	additionalHeadersForSessionProtocol = objc.GetProtocol("additionalHeadersForSession:")
}

