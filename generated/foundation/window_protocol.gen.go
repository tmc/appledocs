// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import "github.com/ebitengine/purego/objc"

// windowProtocol is the window protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 5.0+
//   - iPadOS 5.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to window.
var windowProtocol *objc.Protocol

func init() {
	windowProtocol = objc.GetProtocol("window")
}
