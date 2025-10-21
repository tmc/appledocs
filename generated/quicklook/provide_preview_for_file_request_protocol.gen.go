// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import "github.com/ebitengine/purego/objc"

// providePreviewForFileRequestProtocol is the providePreviewForFileRequest: protocol.
//
// Availability:
//   - Mac Catalyst 15.0+
//   - iOS 15.0+
//   - iPadOS 15.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to providePreviewForFileRequest:.
var providePreviewForFileRequestProtocol *objc.Protocol

func init() {
	providePreviewForFileRequestProtocol = objc.GetProtocol("providePreviewForFileRequest:")
}
