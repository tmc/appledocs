// Code generated from Apple documentation for MailKit. DO NOT EDIT.

package mailkit

import "github.com/ebitengine/purego/objc"

// sessionProtocol is the session: protocol.
//
// Availability:
//   - macOS 12.0+
//
// Use this protocol when registering custom classes that conform to session:.
var sessionProtocol *objc.Protocol

func init() {
	sessionProtocol = objc.GetProtocol("session:")
}
