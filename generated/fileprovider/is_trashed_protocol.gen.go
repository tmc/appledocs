// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import "github.com/ebitengine/purego/objc"

// isTrashedProtocol is the isTrashed protocol.
//
// Availability:
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to isTrashed.
var isTrashedProtocol *objc.Protocol

func init() {
	isTrashedProtocol = objc.GetProtocol("isTrashed")
}
