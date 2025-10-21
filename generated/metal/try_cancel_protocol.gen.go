// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// tryCancelProtocol is the tryCancel protocol.
//
// Availability:
//   - Mac Catalyst 16.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - tvOS 16.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to tryCancel.
var tryCancelProtocol *objc.Protocol

func init() {
	tryCancelProtocol = objc.GetProtocol("tryCancel")
}
