// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import "github.com/ebitengine/purego/objc"

// preparePreviewOfFileAtURLProtocol is the preparePreviewOfFileAtURL: protocol.
//
// Availability:
//   - Mac Catalyst 8.0+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to preparePreviewOfFileAtURL:.
var preparePreviewOfFileAtURLProtocol *objc.Protocol

func init() {
	preparePreviewOfFileAtURLProtocol = objc.GetProtocol("preparePreviewOfFileAtURL:")
}
