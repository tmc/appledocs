// Code generated from Apple documentation for Photos. DO NOT EDIT.

package photos

import "github.com/ebitengine/purego/objc"

// photoLibraryDidChangeProtocol is the photoLibraryDidChange: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.13+
//   - tvOS 10.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to photoLibraryDidChange:.
var photoLibraryDidChangeProtocol *objc.Protocol

func init() {
	photoLibraryDidChangeProtocol = objc.GetProtocol("photoLibraryDidChange:")
}
