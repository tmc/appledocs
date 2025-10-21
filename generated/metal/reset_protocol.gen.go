// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// resetProtocol is the reset protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to reset.
var resetProtocol *objc.Protocol

func init() {
	resetProtocol = objc.GetProtocol("reset")
}
