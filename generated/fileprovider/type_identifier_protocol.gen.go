// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import "github.com/ebitengine/purego/objc"

// typeIdentifierProtocol is the typeIdentifier protocol.
//
// Availability:
//   - iOS 11.0+ (Deprecated in 15.0)
//   - iPadOS 11.0+ (Deprecated in 15.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//
// Use this protocol when registering custom classes that conform to typeIdentifier.
var typeIdentifierProtocol *objc.Protocol

func init() {
	typeIdentifierProtocol = objc.GetProtocol("typeIdentifier")
}

