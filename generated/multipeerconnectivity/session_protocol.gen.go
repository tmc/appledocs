// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity

import "github.com/ebitengine/purego/objc"

// sessionProtocol is the session: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to session:.
var sessionProtocol *objc.Protocol

func init() {
	sessionProtocol = objc.GetProtocol("session:")
}
