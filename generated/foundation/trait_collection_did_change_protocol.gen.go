// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import "github.com/ebitengine/purego/objc"

// traitCollectionDidChangeProtocol is the traitCollectionDidChange: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+ (Deprecated in 17.0)
//   - iOS 8.0+ (Deprecated in 17.0)
//   - iPadOS 8.0+ (Deprecated in 17.0)
//   - visionOS 1.0+ (Deprecated in 1.0)
//
// Use this protocol when registering custom classes that conform to traitCollectionDidChange:.
var traitCollectionDidChangeProtocol *objc.Protocol

func init() {
	traitCollectionDidChangeProtocol = objc.GetProtocol("traitCollectionDidChange:")
}
