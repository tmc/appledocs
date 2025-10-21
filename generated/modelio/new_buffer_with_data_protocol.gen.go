// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio

import "github.com/ebitengine/purego/objc"

// newBufferWithDataProtocol is the newBufferWithData: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to newBufferWithData:.
var newBufferWithDataProtocol *objc.Protocol

func init() {
	newBufferWithDataProtocol = objc.GetProtocol("newBufferWithData:")
}
