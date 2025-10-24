// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import "github.com/ebitengine/purego/objc"

// copyWithZoneProtocol is the copyWithZone: protocol.
//
// Availability:
//   - Mac Catalyst 1.0+
//   - iOS 1.0+
//   - iPadOS 1.0+
//   - macOS 10.0+
//   - tvOS 1.0+
//   - visionOS 1.0+
//   - watchOS 1.0+
//
// Use this protocol when registering custom classes that conform to copyWithZone:.
var copyWithZoneProtocol *objc.Protocol

func init() {
	copyWithZoneProtocol = objc.GetProtocol("copyWithZone:")
}
