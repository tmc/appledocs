// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import "github.com/ebitengine/purego/objc"

// sessionProtocol is the session: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//
// Use this protocol when registering custom classes that conform to session:.
var sessionProtocol *objc.Protocol

func init() {
	sessionProtocol = objc.GetProtocol("session:")
}
