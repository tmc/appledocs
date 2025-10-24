// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// isShareableProtocol is the isShareable protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.14+
//   - tvOS 13.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to isShareable.
var isShareableProtocol *objc.Protocol

func init() {
	isShareableProtocol = objc.GetProtocol("isShareable")
}

