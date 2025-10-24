// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import "github.com/ebitengine/purego/objc"

// SecureCodingProtocol is the NSSecureCoding protocol.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 2.0+
//   - iPadOS 2.0+
//   - macOS 10.0+
//   - tvOS 9.0+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// Use this protocol when registering custom classes that conform to NSSecureCoding.
var SecureCodingProtocol *objc.Protocol

func init() {
	SecureCodingProtocol = objc.GetProtocol("NSSecureCoding")
}

