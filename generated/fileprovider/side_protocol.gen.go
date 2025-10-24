// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import "github.com/ebitengine/purego/objc"

// sideProtocol is the side protocol.
//
// Availability:
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 11.3+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to side.
var sideProtocol *objc.Protocol

func init() {
	sideProtocol = objc.GetProtocol("side")
}

