// Code generated from Apple documentation for QuickLookUI. DO NOT EDIT.

package quicklookui

import "github.com/ebitengine/purego/objc"

// providePreviewForFileRequestProtocol is the providePreviewForFileRequest: protocol.
//
// Availability:
//   - macOS 12.0+
//
// Use this protocol when registering custom classes that conform to providePreviewForFileRequest:.
var providePreviewForFileRequestProtocol *objc.Protocol

func init() {
	providePreviewForFileRequestProtocol = objc.GetProtocol("providePreviewForFileRequest:")
}

