// Code generated from Apple documentation for CoreData. DO NOT EDIT.

package coredata

import "github.com/ebitengine/purego/objc"

// controllerWillChangeContentProtocol is the controllerWillChangeContent: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 3.0+
//   - iPadOS 3.0+
//   - macOS 10.12+
//   - visionOS 1.0+
//   - watchOS 2.0+
//
// Use this protocol when registering custom classes that conform to controllerWillChangeContent:.
var controllerWillChangeContentProtocol *objc.Protocol

func init() {
	controllerWillChangeContentProtocol = objc.GetProtocol("controllerWillChangeContent:")
}
