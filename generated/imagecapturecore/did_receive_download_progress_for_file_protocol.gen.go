// Code generated from Apple documentation for ImageCaptureCore. DO NOT EDIT.

package imagecapturecore

import "github.com/ebitengine/purego/objc"

// didReceiveDownloadProgressForFileProtocol is the didReceiveDownloadProgressForFile: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 10.4+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to didReceiveDownloadProgressForFile:.
var didReceiveDownloadProgressForFileProtocol *objc.Protocol

func init() {
	didReceiveDownloadProgressForFileProtocol = objc.GetProtocol("didReceiveDownloadProgressForFile:")
}

