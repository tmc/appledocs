// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import "github.com/ebitengine/purego/objc"

// previewControllerWillDismissProtocol is the previewControllerWillDismiss: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to previewControllerWillDismiss:.
var previewControllerWillDismissProtocol *objc.Protocol

func init() {
	previewControllerWillDismissProtocol = objc.GetProtocol("previewControllerWillDismiss:")
}

