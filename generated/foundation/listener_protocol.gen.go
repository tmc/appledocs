// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import "github.com/ebitengine/purego/objc"

// listenerProtocol is the listener: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 10.8+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// Use this protocol when registering custom classes that conform to listener:.
var listenerProtocol *objc.Protocol

func init() {
	listenerProtocol = objc.GetProtocol("listener:")
}
