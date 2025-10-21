// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// IndirectComputeCommandProtocol is the MTLIndirectComputeCommand protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - macOS 11.0+
//   - tvOS 13.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to MTLIndirectComputeCommand.
var IndirectComputeCommandProtocol *objc.Protocol

func init() {
	IndirectComputeCommandProtocol = objc.GetProtocol("MTLIndirectComputeCommand")
}
