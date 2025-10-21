// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import "github.com/ebitengine/purego/objc"

// productViewControllerDidFinishProtocol is the productViewControllerDidFinish: protocol.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 6.0+
//   - iPadOS 6.0+
//   - macOS 11.0+
//
// Use this protocol when registering custom classes that conform to productViewControllerDidFinish:.
var productViewControllerDidFinishProtocol *objc.Protocol

func init() {
	productViewControllerDidFinishProtocol = objc.GetProtocol("productViewControllerDidFinish:")
}
