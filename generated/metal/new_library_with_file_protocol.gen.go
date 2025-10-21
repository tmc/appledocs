// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// newLibraryWithFileProtocol is the newLibraryWithFile: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 16.0)
//   - iOS 8.0+ (Deprecated in 16.0)
//   - iPadOS 8.0+ (Deprecated in 16.0)
//   - macOS 10.11+ (Deprecated in 13.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//
// Use this protocol when registering custom classes that conform to newLibraryWithFile:.
var newLibraryWithFileProtocol *objc.Protocol

func init() {
	newLibraryWithFileProtocol = objc.GetProtocol("newLibraryWithFile:")
}
