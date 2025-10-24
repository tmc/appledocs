// Code generated from Apple documentation for Metal. DO NOT EDIT.

package metal

import "github.com/ebitengine/purego/objc"

// recommendedMaxWorkingSetSizeProtocol is the recommendedMaxWorkingSetSize protocol.
//
// Availability:
//   - Mac Catalyst 13.0+
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - macOS 10.12+
//   - tvOS 16.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to recommendedMaxWorkingSetSize.
var recommendedMaxWorkingSetSizeProtocol *objc.Protocol

func init() {
	recommendedMaxWorkingSetSizeProtocol = objc.GetProtocol("recommendedMaxWorkingSetSize")
}

