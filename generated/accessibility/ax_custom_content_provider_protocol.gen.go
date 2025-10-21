// Code generated from Apple documentation for Accessibility. DO NOT EDIT.

package accessibility

import "github.com/ebitengine/purego/objc"

// AXCustomContentProviderProtocol is the AXCustomContentProvider protocol.
//
// Availability:
//   - Mac Catalyst 14.0+
//   - iOS 14.0+
//   - iPadOS 14.0+
//   - macOS 11.0+
//   - tvOS 14.0+
//   - visionOS 1.0+
//   - watchOS 7.0+
//
// Use this protocol when registering custom classes that conform to AXCustomContentProvider.
var AXCustomContentProviderProtocol *objc.Protocol

func init() {
	AXCustomContentProviderProtocol = objc.GetProtocol("AXCustomContentProvider")
}
