// Code generated from Apple documentation for FileProvider. DO NOT EDIT.

package fileprovider

import "github.com/ebitengine/purego/objc"

// filenameProtocol is the filename protocol.
//
// Availability:
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 11.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to filename.
var filenameProtocol *objc.Protocol

func init() {
	filenameProtocol = objc.GetProtocol("filename")
}
