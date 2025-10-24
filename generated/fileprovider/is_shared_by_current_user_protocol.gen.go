// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import "github.com/ebitengine/purego/objc"

// isSharedByCurrentUserProtocol is the isSharedByCurrentUser protocol.
//
// Availability:
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 11.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to isSharedByCurrentUser.
var isSharedByCurrentUserProtocol *objc.Protocol

func init() {
	isSharedByCurrentUserProtocol = objc.GetProtocol("isSharedByCurrentUser")
}

