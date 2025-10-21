// Code generated from Apple documentation for MultipeerConnectivity. DO NOT EDIT.

package multipeerconnectivity

import "github.com/ebitengine/purego/objc"

// advertiserProtocol is the advertiser: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 8.0+
//   - iPadOS 8.0+
//   - macOS 10.10+
//   - tvOS 9.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to advertiser:.
var advertiserProtocol *objc.Protocol

func init() {
	advertiserProtocol = objc.GetProtocol("advertiser:")
}
