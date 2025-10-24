// Code generated from Apple documentation for PushKit. DO NOT EDIT.

package pushkit

import "github.com/ebitengine/purego/objc"

// pushRegistryProtocol is the pushRegistry: protocol.
//
// Availability:
//   - Mac Catalyst 8.0+ (Deprecated in 11.0)
//   - iOS 8.0+ (Deprecated in 11.0)
//   - iPadOS 8.0+ (Deprecated in 11.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//
// Use this protocol when registering custom classes that conform to pushRegistry:.
var pushRegistryProtocol *objc.Protocol

func init() {
	pushRegistryProtocol = objc.GetProtocol("pushRegistry:")
}
