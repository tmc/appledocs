// Code generated from Apple documentation for CoreMotion. DO NOT EDIT.

package coremotion

import "github.com/ebitengine/purego/objc"

// managerProtocol is the manager: protocol.
//
// Availability:
//   - iOS 16.0+
//   - iPadOS 16.0+
//   - visionOS 1.0+
//   - watchOS 9.0+
//
// Use this protocol when registering custom classes that conform to manager:.
var managerProtocol *objc.Protocol

func init() {
	managerProtocol = objc.GetProtocol("manager:")
}
