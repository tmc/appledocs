// Code generated from Apple documentation for ModelIO. DO NOT EDIT.

package modelio

import "github.com/ebitengine/purego/objc"

// canResolveAssetNamedProtocol is the canResolveAssetNamed: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 11.0+
//   - iPadOS 11.0+
//   - macOS 10.13+
//   - tvOS 11.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to canResolveAssetNamed:.
var canResolveAssetNamedProtocol *objc.Protocol

func init() {
	canResolveAssetNamedProtocol = objc.GetProtocol("canResolveAssetNamed:")
}
