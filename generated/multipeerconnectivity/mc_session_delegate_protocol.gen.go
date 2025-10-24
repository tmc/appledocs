// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity

import "github.com/ebitengine/purego/objc"

// MCSessionDelegateProtocol is the MCSessionDelegate protocol.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 7.0+
//   - iPadOS 7.0+
//   - macOS 10.10+
//   - tvOS 10.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to MCSessionDelegate.
var MCSessionDelegateProtocol *objc.Protocol

func init() {
	MCSessionDelegateProtocol = objc.GetProtocol("MCSessionDelegate")
}

