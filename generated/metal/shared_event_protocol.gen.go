// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// SharedEventProtocol is the MTLSharedEvent protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 12.0+
//   - iPadOS 12.0+
//   - macOS 10.14+
//   - tvOS 12.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to MTLSharedEvent.
var SharedEventProtocol *objc.Protocol

func init() {
	SharedEventProtocol = objc.GetProtocol("MTLSharedEvent")
}

