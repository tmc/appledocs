// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec

import "github.com/ebitengine/purego/objc"

// autoreleaseProtocol is the autorelease protocol.
//
// Availability:
//   - Mac Catalyst +
//   - iOS +
//   - iPadOS +
//   - macOS +
//   - tvOS +
//   - visionOS +
//   - watchOS +
//
// Use this protocol when registering custom classes that conform to autorelease.
var autoreleaseProtocol *objc.Protocol

func init() {
	autoreleaseProtocol = objc.GetProtocol("autorelease")
}
