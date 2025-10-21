// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import "github.com/ebitengine/purego/objc"

// modifyItemProtocol is the modifyItem: protocol.
//
// Availability:
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 11.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to modifyItem:.
var modifyItemProtocol *objc.Protocol

func init() {
	modifyItemProtocol = objc.GetProtocol("modifyItem:")
}
