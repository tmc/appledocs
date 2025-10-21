// Code generated from Apple documentation for MetalKit. DO NOT EDIT.

package metalkit

import "github.com/ebitengine/purego/objc"

// ViewDelegateProtocol is the MTKViewDelegate protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 9.0+
//   - iPadOS 9.0+
//   - macOS 10.11+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to MTKViewDelegate.
var ViewDelegateProtocol *objc.Protocol

func init() {
	ViewDelegateProtocol = objc.GetProtocol("MTKViewDelegate")
}
