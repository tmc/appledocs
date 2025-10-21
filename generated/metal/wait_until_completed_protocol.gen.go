// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// waitUntilCompletedProtocol is the waitUntilCompleted protocol.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to waitUntilCompleted.
var waitUntilCompletedProtocol *objc.Protocol

func init() {
	waitUntilCompletedProtocol = objc.GetProtocol("waitUntilCompleted")
}
