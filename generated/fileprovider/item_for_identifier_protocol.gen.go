// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import "github.com/ebitengine/purego/objc"

// itemForIdentifierProtocol is the itemForIdentifier: protocol.
//
// Availability:
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 11.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to itemForIdentifier:.
var itemForIdentifierProtocol *objc.Protocol

func init() {
	itemForIdentifierProtocol = objc.GetProtocol("itemForIdentifier:")
}
