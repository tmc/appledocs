// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import "github.com/ebitengine/purego/objc"

// setInteractionSuppressedProtocol is the setInteractionSuppressed: protocol.
//
// Availability:
//   - macOS 12.0+
//
// Use this protocol when registering custom classes that conform to setInteractionSuppressed:.
var setInteractionSuppressedProtocol *objc.Protocol

func init() {
	setInteractionSuppressedProtocol = objc.GetProtocol("setInteractionSuppressed:")
}

