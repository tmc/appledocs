// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import "github.com/ebitengine/purego/objc"

// isRestrictedProtocol is the isRestricted protocol.
//
// Availability:
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 13.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to isRestricted.
var isRestrictedProtocol *objc.Protocol

func init() {
	isRestrictedProtocol = objc.GetProtocol("isRestricted")
}

