// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio

import "github.com/ebitengine/purego/objc"

// allocatorProtocol is the allocator protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to allocator.
var allocatorProtocol *objc.Protocol

func init() {
	allocatorProtocol = objc.GetProtocol("allocator")
}
