// Code generated from Apple documentation for StoreKit. DO NOT EDIT.

package storekit

import "github.com/ebitengine/purego/objc"

// productsRequestProtocol is the productsRequest: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 18.0)
//   - iOS 3.0+ (Deprecated in 18.0)
//   - iPadOS 3.0+ (Deprecated in 18.0)
//   - macOS 10.7+ (Deprecated in 15.0)
//   - tvOS + (Deprecated in 18.0)
//   - visionOS 1.0+ (Deprecated in 2.0)
//   - watchOS 6.2+ (Deprecated in 11.0)
//
// Use this protocol when registering custom classes that conform to productsRequest:.
var productsRequestProtocol *objc.Protocol

func init() {
	productsRequestProtocol = objc.GetProtocol("productsRequest:")
}

