// Code generated from Apple documentation for ObjectiveC. DO NOT EDIT.

package objectivec

import "github.com/ebitengine/purego/objc"

// performSelectorProtocol is the performSelector: protocol.
//
// Availability:
//   - Mac Catalyst 1.0+
//   - iOS 1.0+
//   - iPadOS 1.0+
//   - macOS 10.0+
//   - tvOS 1.0+
//   - visionOS 1.0+
//   - watchOS 1.0+
//
// Use this protocol when registering custom classes that conform to performSelector:.
var performSelectorProtocol *objc.Protocol

func init() {
	performSelectorProtocol = objc.GetProtocol("performSelector:")
}
