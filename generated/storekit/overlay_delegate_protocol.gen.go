// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import "github.com/ebitengine/purego/objc"

// OverlayDelegateProtocol is the SKOverlayDelegate protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to SKOverlayDelegate.
var OverlayDelegateProtocol *objc.Protocol

func init() {
	OverlayDelegateProtocol = objc.GetProtocol("SKOverlayDelegate")
}
