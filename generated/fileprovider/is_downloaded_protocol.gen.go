// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import "github.com/ebitengine/purego/objc"

// isDownloadedProtocol is the isDownloaded protocol.
//
// Availability:
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 11.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to isDownloaded.
var isDownloadedProtocol *objc.Protocol

func init() {
	isDownloadedProtocol = objc.GetProtocol("isDownloaded")
}

