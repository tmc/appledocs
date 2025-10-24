// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import "github.com/ebitengine/purego/objc"

// URLProtocolClientProtocol is the NSURLProtocolClient protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.2+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// Use this protocol when registering custom classes that conform to NSURLProtocolClient.
var URLProtocolClientProtocol *objc.Protocol

func init() {
	URLProtocolClientProtocol = objc.GetProtocol("NSURLProtocolClient")
}
