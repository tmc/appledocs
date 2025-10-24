// Code generated from Apple documentation for PassKit. DO NOT EDIT.

package passkit

import "github.com/ebitengine/purego/objc"

// sessionDidChangeConnectionStateProtocol is the sessionDidChangeConnectionState: protocol.
//
// Availability:
//   - Mac Catalyst 15.4+
//   - iOS 15.4+
//   - iPadOS 15.4+
//   - macOS +
//   - visionOS 1.0+
//   - watchOS 8.5+
//
// Use this protocol when registering custom classes that conform to sessionDidChangeConnectionState:.
var sessionDidChangeConnectionStateProtocol *objc.Protocol

func init() {
	sessionDidChangeConnectionStateProtocol = objc.GetProtocol("sessionDidChangeConnectionState:")
}

