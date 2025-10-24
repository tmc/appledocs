// Code generated from Apple documentation for WebKit. DO NOT EDIT.

package webkit

import "github.com/ebitengine/purego/objc"

// identifierProtocol is the identifier protocol.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 13.1)
//   - iOS 10.0+ (Deprecated in 13.0)
//   - iPadOS 10.0+ (Deprecated in 13.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//
// Use this protocol when registering custom classes that conform to identifier.
var identifierProtocol *objc.Protocol

func init() {
	identifierProtocol = objc.GetProtocol("identifier")
}

