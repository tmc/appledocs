// Code generated from Apple documentation for CallKit. DO NOT EDIT.

package callkit

import "github.com/ebitengine/purego/objc"

// requestFailedForExtensionContextProtocol is the requestFailedForExtensionContext: protocol.
//
// Availability:
//   - Mac Catalyst 10.0+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to requestFailedForExtensionContext:.
var requestFailedForExtensionContextProtocol *objc.Protocol

func init() {
	requestFailedForExtensionContextProtocol = objc.GetProtocol("requestFailedForExtensionContext:")
}

