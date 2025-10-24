// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec

import "github.com/ebitengine/purego/objc"

// selfProtocol is the self protocol.
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
// Use this protocol when registering custom classes that conform to self.
var selfProtocol *objc.Protocol

func init() {
	selfProtocol = objc.GetProtocol("self")
}

