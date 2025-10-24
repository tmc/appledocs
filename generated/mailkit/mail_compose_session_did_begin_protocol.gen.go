// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import "github.com/ebitengine/purego/objc"

// mailComposeSessionDidBeginProtocol is the mailComposeSessionDidBegin: protocol.
//
// Availability:
//   - macOS 12.0+
//
// Use this protocol when registering custom classes that conform to mailComposeSessionDidBegin:.
var mailComposeSessionDidBeginProtocol *objc.Protocol

func init() {
	mailComposeSessionDidBeginProtocol = objc.GetProtocol("mailComposeSessionDidBegin:")
}

