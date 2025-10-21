// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// iosurfaceProtocol is the iosurface protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.11+
//   - tvOS 11.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to iosurface.
var iosurfaceProtocol *objc.Protocol

func init() {
	iosurfaceProtocol = objc.GetProtocol("iosurface")
}
