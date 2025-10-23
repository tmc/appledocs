// Code generated from Apple documentation for Foundation. DO NOT EDIT.

package foundation

import "github.com/ebitengine/purego/objc"

// sceneProtocol is the scene: protocol.
//
// Availability:
//   - Mac Catalyst 13.1+
//   - iOS 13.0+
//   - iPadOS 13.0+
//   - tvOS 13.0+
//   - visionOS 1.0+
//
// Use this protocol when registering custom classes that conform to scene:.
var sceneProtocol *objc.Protocol

func init() {
	sceneProtocol = objc.GetProtocol("scene:")
}
