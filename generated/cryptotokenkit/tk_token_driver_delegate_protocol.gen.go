// Code generated from Apple documentation for CryptoTokenKit. DO NOT EDIT.

package cryptotokenkit

import "github.com/ebitengine/purego/objc"

// TKTokenDriverDelegateProtocol is the TKTokenDriverDelegate protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 10.0+
//   - iPadOS 10.0+
//   - macOS 10.12+
//   - tvOS 11.0+
//   - visionOS 1.0+
//   - watchOS 4.0+
//
// Use this protocol when registering custom classes that conform to TKTokenDriverDelegate.
var TKTokenDriverDelegateProtocol *objc.Protocol

func init() {
	TKTokenDriverDelegateProtocol = objc.GetProtocol("TKTokenDriverDelegate")
}
