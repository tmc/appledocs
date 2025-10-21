// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// typeProtocol is the type protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to type.
var typeProtocol *objc.Protocol

func init() {
	typeProtocol = objc.GetProtocol("type")
}
