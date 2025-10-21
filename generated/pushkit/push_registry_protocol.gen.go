// Code generated from Apple documentation for PushKit. DO NOT EDIT.

package pushkit

import "github.com/ebitengine/purego/objc"

// pushRegistryProtocol is the pushRegistry: protocol.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.15+
//   - visionOS 1.0+
//   - watchOS 6.0+
//
// Use this protocol when registering custom classes that conform to pushRegistry:.
var pushRegistryProtocol *objc.Protocol

func init() {
	pushRegistryProtocol = objc.GetProtocol("pushRegistry:")
}
