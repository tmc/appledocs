// Code generated from Apple documentation for Matter. DO NOT EDIT.

package matter

import "github.com/ebitengine/purego/objc"

// publicKeyProtocol is the publicKey protocol.
//
// Availability:
//   - Mac Catalyst 16.1+ (Deprecated in 18.4)
//   - iOS 16.1+ (Deprecated in 18.4)
//   - iPadOS 16.1+ (Deprecated in 18.4)
//   - macOS 13.0+ (Deprecated in 15.4)
//   - tvOS 16.1+ (Deprecated in 18.4)
//   - visionOS 1.0+ (Deprecated in 2.4)
//   - watchOS 9.1+ (Deprecated in 11.4)
//
// Use this protocol when registering custom classes that conform to publicKey.
var publicKeyProtocol *objc.Protocol

func init() {
	publicKeyProtocol = objc.GetProtocol("publicKey")
}
