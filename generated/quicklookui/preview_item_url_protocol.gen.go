// Code generated from Apple documentation for QuickLookUI. DO NOT EDIT.

package quicklookui

import "github.com/ebitengine/purego/objc"

// previewItemURLProtocol is the previewItemURL protocol.
//
// Availability:
//   - macOS 10.6+
//
// Use this protocol when registering custom classes that conform to previewItemURL.
var previewItemURLProtocol *objc.Protocol

func init() {
	previewItemURLProtocol = objc.GetProtocol("previewItemURL")
}

