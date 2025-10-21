// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import "github.com/ebitengine/purego/objc"

// isMaximumSizeReachedProtocol is the isMaximumSizeReached protocol.
//
// Availability:
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to isMaximumSizeReached.
var isMaximumSizeReachedProtocol *objc.Protocol

func init() {
	isMaximumSizeReachedProtocol = objc.GetProtocol("isMaximumSizeReached")
}
