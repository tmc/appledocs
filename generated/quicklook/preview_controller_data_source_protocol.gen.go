// Code generated from Apple documentation for QuickLook. DO NOT EDIT.

package quicklook

import "github.com/ebitengine/purego/objc"

// PreviewControllerDataSourceProtocol is the QLPreviewControllerDataSource protocol.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 4.0+
//   - iPadOS 4.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to QLPreviewControllerDataSource.
var PreviewControllerDataSourceProtocol *objc.Protocol

func init() {
	PreviewControllerDataSourceProtocol = objc.GetProtocol("QLPreviewControllerDataSource")
}
