// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// areBarycentricCoordsSupportedProtocol is the areBarycentricCoordsSupported protocol.
//
// Availability:
//   - Mac Catalyst 14.0+ (Deprecated in 16.0)
//   - iOS 14.0+ (Deprecated in 16.0)
//   - iPadOS 14.0+ (Deprecated in 16.0)
//   - macOS 10.15+ (Deprecated in 13.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//
// Use this protocol when registering custom classes that conform to areBarycentricCoordsSupported.
var areBarycentricCoordsSupportedProtocol *objc.Protocol

func init() {
	areBarycentricCoordsSupportedProtocol = objc.GetProtocol("areBarycentricCoordsSupported")
}
